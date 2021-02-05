/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package add

import (
	"context"
	"fmt"

	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	pkgapisduck "knative.dev/pkg/apis/duck"
	v1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"
	"knative.dev/pkg/tracker"
	duckv1 "tableflip.dev/maths/pkg/apis/duck/v1"
	mathsv1alpha1 "tableflip.dev/maths/pkg/apis/maths/v1alpha1"
	addreconciler "tableflip.dev/maths/pkg/client/injection/reconciler/maths/v1alpha1/add"
)

// Reconciler implements addressableservicereconciler.Interface for
// Add resources.
type Reconciler struct {
	// Tracker builds an index of what resources are watching other resources
	// so that we can immediately react to changes tracked resources.
	tracker         tracker.Interface
	informerFactory pkgapisduck.InformerFactory
}

// Check that our Reconciler implements Interface
var _ addreconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, o *mathsv1alpha1.Add) reconciler.Event {
	expression := ""
	result := 0

	for _, op := range o.Spec.Operands {
		if op.Ref != nil {
			rt, err := r.getResults(ctx, o, op.Ref)
			if err != nil {
				return err
			}

			if len(expression) == 0 {
				expression = fmt.Sprintf("(%s)", rt.Status.Expression)
			} else {
				expression = fmt.Sprintf("%s + (%s)", expression, rt.Status.Expression)
			}
			result = result + rt.Status.Result
		} else if op.Value != nil {
			if len(expression) == 0 {
				expression = fmt.Sprintf("%d", *op.Value)
			} else {
				expression = fmt.Sprintf("%s + %d", expression, *op.Value)
			}
			result = result + *op.Value
		}

	}

	o.Status.Expression = expression
	o.Status.Result = result

	o.Status.MarkComputed()

	return nil
}

func (r *Reconciler) getResults(ctx context.Context, o *mathsv1alpha1.Add, ref *v1.KReference) (*duckv1.ResultsType, error) {
	l := logging.FromContext(ctx)

	refDesc := fmt.Sprintf("%s.%s %s/%s", ref.Kind, ref.APIVersion, ref.Namespace, ref.Name)

	// Make sure we are tracking the Operand.
	if err := r.tracker.TrackReference(tracker.Reference{
		APIVersion: ref.APIVersion,
		Kind:       ref.Kind,
		Name:       ref.Name,
		Namespace:  ref.Namespace,
	}, o); err != nil {
		l.Errorf("Error tracking operand %s: %v", refDesc, err)
		return nil, err
	}
	gvk := schema.FromAPIVersionAndKind(ref.APIVersion, ref.Kind)
	gvr, _ := meta.UnsafeGuessKindToResource(gvk)

	l.Infof("===> Looking for ducktype Results with gvr: %+v\n", gvr)

	// Get a cached informer.
	_, lister, err := r.informerFactory.Get(ctx, gvr)
	if err != nil {
		return nil, err
	}
	// Get result.
	obj, err := lister.ByNamespace(ref.Namespace).Get(ref.Name)
	if err != nil {
		o.Status.MarkResultsMissing(refDesc)
		return nil, err
	}

	rt, ok := obj.(*duckv1.ResultsType)
	if !ok {
		return nil, apierrs.NewBadRequest(fmt.Sprintf("%+v (%T) is not an ResultsType", ref, obj))
	}
	return rt, nil
}
