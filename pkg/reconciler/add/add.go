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
	logger := logging.FromContext(ctx)

	expression := ""
	result := 0

	for _, op := range o.Spec.Operands {
		if op.Ref != nil {
			// Make sure we are tracking the Operand.
			if err := r.tracker.TrackReference(tracker.Reference{
				APIVersion: op.Ref.APIVersion,
				Kind:       op.Ref.Kind,
				Name:       op.Ref.Name,
				Namespace:  op.Ref.Namespace,
			}, o); err != nil {
				ref := fmt.Sprintf("%s.%s %s/%s", op.Ref.Kind, op.Ref.APIVersion, op.Ref.Namespace, op.Ref.Name)
				logger.Errorf("Error tracking operand %s: %v", ref, err)
				o.Status.MarkResultsMissing(ref)
				return err
			}
			gvr, _ := meta.UnsafeGuessKindToResource(
				schema.FromAPIVersionAndKind(op.Ref.APIVersion, op.Ref.Kind))

			// Get a cached informer.
			_, lister, err := r.informerFactory.Get(ctx, gvr)
			if err != nil {
				return err
			}

			// Get result.
			obj, err := lister.ByNamespace(op.Ref.Namespace).Get(op.Ref.Name)
			if err != nil {
				return err
			}

			rt, ok := obj.(*duckv1.ResultsType)

			logger.Infof("Got Results ducktype: %#v\n", rt)

			if !ok {
				return apierrs.NewBadRequest(fmt.Sprintf("%+v (%T) is not an ResultsType", op.Ref, obj))
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
