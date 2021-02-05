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

package v1alpha1

import (
	"context"
	"fmt"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (a *Subtract) Validate(ctx context.Context) *apis.FieldError {
	return a.Spec.Validate(ctx).ViaField("spec")
}

// Validate implements apis.Validatable
func (as *SubtractSpec) Validate(ctx context.Context) *apis.FieldError {
	for k, v := range as.Operands {
		if v.Ref != nil && v.Value != nil {
			return apis.ErrMultipleOneOf("sub", fmt.Sprintf("[%d]", k))
		} else if v.Ref != nil {
			if err := v.Ref.Validate(ctx); err != nil {
				return err.ViaFieldIndex("sub", k)
			}
		}
	}
	return nil
}
