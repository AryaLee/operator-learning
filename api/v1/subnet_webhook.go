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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var subnetlog = logf.Log.WithName("subnet-resource")

func (r *Subnet) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).WithValidator(&SubnetValidator{
		Client: mgr.GetClient(),
	}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-sdn-github-com-v1-subnet,mutating=true,failurePolicy=fail,sideEffects=None,groups=sdn.github.com,resources=subnets,verbs=create;update,versions=v1,name=msubnet.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Subnet{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Subnet) Default() {
	subnetlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-sdn-github-com-v1-subnet,mutating=false,failurePolicy=fail,sideEffects=None,groups=sdn.github.com,resources=subnets,verbs=create;update,versions=v1,name=vsubnet.kb.io,admissionReviewVersions=v1

// +kubebuilder:object:generate=false
type SubnetValidator struct {
	client.Client
}

func (v *SubnetValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	subnet := obj.(*Subnet)
	var vpc Vpc
	vpcKey := client.ObjectKey{
		Namespace: subnet.Namespace,
		Name:      subnet.Spec.VpcID,
	}
	if err := v.Get(ctx, vpcKey, &vpc); err != nil {
		subnetlog.Error(err, "get vpc error")
		return err
	}
	subnetlog.Info("validate create", "vpc", vpc)
	return nil
}

func (v *SubnetValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

func (v *SubnetValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}

var _ webhook.Validator = &Subnet{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateCreate() error {
	subnetlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateUpdate(old runtime.Object) error {
	subnetlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Subnet) ValidateDelete() error {
	subnetlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
