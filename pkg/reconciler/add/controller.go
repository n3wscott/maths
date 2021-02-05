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
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/apis/duck"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/tracker"

	resultsinformer "tableflip.dev/maths/pkg/client/injection/ducks/duck/v1/results"
	addinformer "tableflip.dev/maths/pkg/client/injection/informers/maths/v1alpha1/add"
	addreconciler "tableflip.dev/maths/pkg/client/injection/reconciler/maths/v1alpha1/add"
)

// NewController creates a Reconciler and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	addInformer := addinformer.Get(ctx)
	resultsInformer := resultsinformer.Get(ctx)

	r := &Reconciler{}
	impl := addreconciler.NewImpl(ctx, r)
	r.tracker = tracker.New(impl.EnqueueKey, controller.GetTrackerLease(ctx))

	cif := &duck.CachedInformerFactory{
		Delegate: &duck.EnqueueInformerFactory{
			Delegate: resultsInformer,
			EventHandler: cache.ResourceEventHandlerFuncs{
				AddFunc:    r.tracker.OnChanged,
				UpdateFunc: controller.PassNew(r.tracker.OnChanged),
			},
		},
	}

	r.informerFactory = cif

	logger.Info("Setting up event handlers.")

	addInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	return impl
}
