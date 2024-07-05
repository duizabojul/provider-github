/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConfigurationInitParameters struct {

	// The content type for the payload. Valid values are either form or json.
	// The content type for the payload. Valid values are either 'form' or 'json'.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Insecure SSL boolean toggle. Defaults to false.
	// Insecure SSL boolean toggle. Defaults to 'false'.
	InsecureSSL *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`

	// The shared secret for the webhook. See API documentation.
	// The shared secret for the webhook
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// The URL of the webhook.
	// The URL of the webhook.
	URLSecretRef v1.SecretKeySelector `json:"urlSecretRef" tf:"-"`
}

type ConfigurationObservation struct {

	// The content type for the payload. Valid values are either form or json.
	// The content type for the payload. Valid values are either 'form' or 'json'.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Insecure SSL boolean toggle. Defaults to false.
	// Insecure SSL boolean toggle. Defaults to 'false'.
	InsecureSSL *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`
}

type ConfigurationParameters struct {

	// The content type for the payload. Valid values are either form or json.
	// The content type for the payload. Valid values are either 'form' or 'json'.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Insecure SSL boolean toggle. Defaults to false.
	// Insecure SSL boolean toggle. Defaults to 'false'.
	// +kubebuilder:validation:Optional
	InsecureSSL *bool `json:"insecureSsl,omitempty" tf:"insecure_ssl,omitempty"`

	// The shared secret for the webhook. See API documentation.
	// The shared secret for the webhook
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// The URL of the webhook.
	// The URL of the webhook.
	// +kubebuilder:validation:Optional
	URLSecretRef v1.SecretKeySelector `json:"urlSecretRef" tf:"-"`
}

type RepositoryWebhookInitParameters struct {

	// Indicate if the webhook should receive events. Defaults to true.
	// Indicate if the webhook should receive events. Defaults to 'true'.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Configuration block for the webhook. Detailed below.
	// Configuration for the webhook.
	Configuration []ConfigurationInitParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of events which should trigger the webhook. See a list of available events.
	// A list of events which should trigger the webhook
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The repository of the webhook.
	// The repository of the webhook.
	// +crossplane:generate:reference:type=Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

type RepositoryWebhookObservation struct {

	// Indicate if the webhook should receive events. Defaults to true.
	// Indicate if the webhook should receive events. Defaults to 'true'.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Configuration block for the webhook. Detailed below.
	// Configuration for the webhook.
	Configuration []ConfigurationObservation `json:"configuration,omitempty" tf:"configuration,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A list of events which should trigger the webhook. See a list of available events.
	// A list of events which should trigger the webhook
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The repository of the webhook.
	// The repository of the webhook.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// URL of the webhook.  This is a sensitive attribute because it may include basic auth credentials.
	// Configuration block for the webhook
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type RepositoryWebhookParameters struct {

	// Indicate if the webhook should receive events. Defaults to true.
	// Indicate if the webhook should receive events. Defaults to 'true'.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Configuration block for the webhook. Detailed below.
	// Configuration for the webhook.
	// +kubebuilder:validation:Optional
	Configuration []ConfigurationParameters `json:"configuration,omitempty" tf:"configuration,omitempty"`

	// A list of events which should trigger the webhook. See a list of available events.
	// A list of events which should trigger the webhook
	// +kubebuilder:validation:Optional
	// +listType=set
	Events []*string `json:"events,omitempty" tf:"events,omitempty"`

	// The repository of the webhook.
	// The repository of the webhook.
	// +crossplane:generate:reference:type=Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

// RepositoryWebhookSpec defines the desired state of RepositoryWebhook
type RepositoryWebhookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryWebhookParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RepositoryWebhookInitParameters `json:"initProvider,omitempty"`
}

// RepositoryWebhookStatus defines the observed state of RepositoryWebhook.
type RepositoryWebhookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryWebhook is the Schema for the RepositoryWebhooks API. Creates and manages repository webhooks within GitHub organizations or personal accounts
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type RepositoryWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.events) || (has(self.initProvider) && has(self.initProvider.events))",message="spec.forProvider.events is a required parameter"
	Spec   RepositoryWebhookSpec   `json:"spec"`
	Status RepositoryWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryWebhookList contains a list of RepositoryWebhooks
type RepositoryWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryWebhook `json:"items"`
}

// Repository type metadata.
var (
	RepositoryWebhook_Kind             = "RepositoryWebhook"
	RepositoryWebhook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryWebhook_Kind}.String()
	RepositoryWebhook_KindAPIVersion   = RepositoryWebhook_Kind + "." + CRDGroupVersion.String()
	RepositoryWebhook_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryWebhook_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryWebhook{}, &RepositoryWebhookList{})
}