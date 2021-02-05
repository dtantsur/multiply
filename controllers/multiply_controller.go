/*


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

package controllers

import (
	"context"
	"math"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mathopsv1alpha1 "github.com/dtantsur/multiply/api/v1alpha1"
)

// MultiplyReconciler reconciles a Multiply object
type MultiplyReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=mathops.owlet.today,resources=multiplies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mathops.owlet.today,resources=multiplies/status,verbs=get;update;patch
func (r *MultiplyReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("multiply", req.NamespacedName)

	multiply := &mathopsv1alpha1.Multiply{}
	err := r.Get(ctx, req.NamespacedName, multiply)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Resource no longer found")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Could not get the resource")
		return ctrl.Result{}, err
	}

	overflow := false
	result := int64(multiply.Spec.Operand1) * int64(multiply.Spec.Operand2)
	if result > math.MaxInt32 {
		overflow = true
	}
	trimmedResult := int32(result)

	if trimmedResult != multiply.Status.Result || overflow != multiply.Status.Overflow {
		multiply.Status.Result = trimmedResult
		multiply.Status.Overflow = overflow
		err = r.Status().Update(ctx, multiply)
		if err != nil {
			log.Error(err, "Cound not update the result")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *MultiplyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mathopsv1alpha1.Multiply{}).
		Complete(r)
}
