/*
Copyright 2022 ArayLee.

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

package v1

import (
	"context"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var vpclog = logf.Log.WithName("vpc-resource")

func (r *Vpc) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		WithDefaulter(&Vpc{}).For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
//+kubebuilder:webhook:path=/mutate-sdn-github-com-v1-vpc,mutating=true,failurePolicy=fail,sideEffects=None,groups=sdn.github.com,resources=vpcs,verbs=create;update,versions=v1,name=mvpc.kb.io,admissionReviewVersions=v1

func (r *Vpc) Default(ctx context.Context, obj runtime.Object) error {
	vpclog.Info("mutate", "default", obj)
	req, err := admission.RequestFromContext(ctx)
	if err != nil {
		return err
	}

	vpc := obj.(*Vpc)
	if req.Operation == admissionv1.Create {
		vpc.Status.VNI = 303
	}
	return nil
}

//+kubebuilder:webhook:path=/validate-sdn-github-com-v1-vpc,mutating=false,failurePolicy=fail,sideEffects=None,groups=sdn.github.com,resources=vpcs,verbs=create;update;delete,versions=v1,name=vvpc.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Vpc{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Vpc) ValidateCreate() error {
	vpclog.Info("validate create", "name", r.Name)

	if r.VNI != 0 {
		return errors.NewBadRequest("vni is forbidden to create.")
	}
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Vpc) ValidateUpdate(old runtime.Object) error {
	vpclog.Info("validate update", "name", r.Name)

	oldVpc := old.(*Vpc)
	if oldVpc.VNI != 0 && oldVpc.VNI != r.VNI {
		return errors.NewBadRequest("vni is forbidden to update.")
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Vpc) ValidateDelete() error {
	vpclog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
