package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Klusterlet represents controllers on the managed cluster. When configured,
// the Klusterlet requires a secret named of bootstrap-hub-kubeconfig in the
// same namespace to allow API requests to the hub for the registration protocol.
type Klusterlet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents the desired deployment configuration of Klusterlet agent.
	Spec KlusterletSpec `json:"spec,omitempty"`

	// Status represents the current status of Klusterlet agent.
	Status KlusterletStatus `json:"status,omitempty"`
}

// KlusterletSpec represents the desired deployment configuration of Klusterlet agent.
type KlusterletSpec struct {
	// Namespace is the namespace to deploy the agent.
	// The namespace must have a prefix of "open-cluster-management-", and if it is not set,
	// the namespace of "open-cluster-management-agent" is used to deploy agent.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// RegistrationImagePullSpec represents the desired image configuration of registration agent.
	// +required
	RegistrationImagePullSpec string `json:"registrationImagePullSpec"`

	// WorkImagePullSpec represents the desired image configuration of work agent.
	// +required
	WorkImagePullSpec string `json:"workImagePullSpec,omitempty"`

	// ClusterName is the name of the managed cluster to be created on hub.
	// The Klusterlet agent generates a random name if it is not set, or discovers the appropriate cluster name on OpenShift.
	// +optional
	ClusterName string `json:"clusterName,omitempty"`

	// ExternalServerURLs represents the a list of apiserver urls and ca bundles that is accessible externally
	// If it is set empty, managed cluster has no externally accessible url that hub cluster can visit.
	// +optional
	ExternalServerURLs []ServerURL `json:"externalServerURLs,omitempty"`
}

// ServerURL represents the apiserver url and ca bundle that is accessible externally
type ServerURL struct {
	// URL is the url of apiserver endpoint of the managed cluster.
	// +required
	URL string `json:"url"`

	// CABundle is the ca bundle to connect to apiserver of the managed cluster.
	// System certs are used if it is not set.
	// +optional
	CABundle []byte `json:"caBundle,omitempty"`
}

// KlusterletStatus represents the current status of Klusterlet agent.
type KlusterletStatus struct {
	// ObservedGeneration is the last generation change you've dealt with
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions contain the different condition statuses for this Klusterlet.
	// Valid condition types are:
	// Applied: Components have been applied in the managed cluster.
	// Available: Components in the managed cluster are available and ready to serve.
	// Progressing: Components in the managed cluster are in a transitioning state.
	// Degraded: Components in the managed cluster do not match the desired configuration and only provide
	// degraded service.
	Conditions []metav1.Condition `json:"conditions"`

	// Generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.
	// +optional
	Generations []GenerationStatus `json:"generations,omitempty"`

	// RelatedResources are used to track the resources that are related to this Klusterlet.
	// +optional
	RelatedResources []RelatedResourceMeta `json:"relatedResources,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KlusterletList is a collection of Klusterlet agents.
type KlusterletList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of Klusterlet agents.
	Items []Klusterlet `json:"items"`
}
