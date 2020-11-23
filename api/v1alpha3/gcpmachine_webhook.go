/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	"reflect"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var _ = logf.Log.WithName("gcpmachine-resource")

func (r *GCPMachine) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// +kubebuilder:webhook:verbs=create;update,path=/validate-infrastructure-cluster-x-k8s-io-v1alpha3-gcpmachine,mutating=false,failurePolicy=fail,matchPolicy=Equivalent,groups=infrastructure.cluster.x-k8s.io,resources=gcpmachines,versions=v1alpha3,name=validation.gcpmachine.infrastructure.cluster.x-k8s.io

var _ webhook.Validator = &GCPMachine{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPMachine) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPMachine) ValidateUpdate(old runtime.Object) error {
	newGCPMachine, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		return apierrors.NewInvalid(GroupVersion.WithKind("GCPMachine").GroupKind(), r.Name, field.ErrorList{
			field.InternalError(nil, errors.Wrap(err, "failed to convert new GCPMachine to unstructured object")),
		})
	}
	oldGCPMachine, err := runtime.DefaultUnstructuredConverter.ToUnstructured(old)
	if err != nil {
		return apierrors.NewInvalid(GroupVersion.WithKind("GCPMachine").GroupKind(), r.Name, field.ErrorList{
			field.InternalError(nil, errors.Wrap(err, "failed to convert old GCPMachine to unstructured object")),
		})
	}

	newGCPMachineSpec := newGCPMachine["spec"].(map[string]interface{})
	oldGCPMachineSpec := oldGCPMachine["spec"].(map[string]interface{})

	// allow changes to providerID
	delete(oldGCPMachineSpec, "providerID")
	delete(newGCPMachineSpec, "providerID")

	// allow changes to additionalLabels
	delete(oldGCPMachineSpec, "additionalLabels")
	delete(newGCPMachineSpec, "additionalLabels")

	// allow changes to additionalNetworkTags
	delete(oldGCPMachineSpec, "additionalNetworkTags")
	delete(newGCPMachineSpec, "additionalNetworkTags")

	if !reflect.DeepEqual(oldGCPMachineSpec, newGCPMachineSpec) {
		return apierrors.NewInvalid(GroupVersion.WithKind("GCPMachine").GroupKind(), r.Name, field.ErrorList{
			field.Forbidden(field.NewPath("spec"), "cannot be modified"),
		})
	}

	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *GCPMachine) ValidateDelete() error {
	return nil
}
