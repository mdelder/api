package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// ClusterManager configures the controllers on the hub that govern registration and work distribution for attached Klusterlets.
// ClusterManager will only be deployed in open-cluster-management-hub namespace.
type ClusterManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.
	Spec ClusterManagerSpec `json:"spec"`

	// Status represents the current status of controllers that govern the lifecycle of managed clusters.
	// +optional
	Status ClusterManagerStatus `json:"status,omitempty"`
}

// ClusterManagerSpec represents a desired deployment configuration of controllers that govern registration and work distribution for attached Klusterlets.
type ClusterManagerSpec struct {
	// RegistrationImagePullSpec represents the desired image of registration controller/webhook installed on hub.
	// +required
	RegistrationImagePullSpec string `json:"registrationImagePullSpec"`

	// WorkImagePullSpec represents the desired image configuration of work controller/webhook installed on hub.
	// +required
	WorkImagePullSpec string `json:"workImagePullSpec,omitempty"`
}

// ClusterManagerStatus represents the current status of the registration and work distribution controllers running on the hub.
type ClusterManagerStatus struct {
	// ObservedGeneration is the last generation change you've dealt with
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions contain the different condition statuses for this ClusterManager.
	// Valid condition types are:
	// Applied: Components in hub are applied.
	// Available: Components in hub are available and ready to serve.
	// Progressing: Components in hub are in a transitioning state.
	// Degraded: Components in hub do not match the desired configuration and only provide
	// degraded service.
	Conditions []metav1.Condition `json:"conditions"`

	// Generations are used to determine when an item needs to be reconciled or has changed in a way that needs a reaction.
	// +optional
	Generations []GenerationStatus `json:"generations,omitempty"`

	// RelatedResources are used to track the resources that are related to this ClusterManager.
	// +optional
	RelatedResources []RelatedResourceMeta `json:"relatedResources,omitempty"`
}

// RelatedResourceMeta represents the resource that is managed by an operator
type RelatedResourceMeta struct {
	// group is the group of the resource that you're tracking
	// +required
	Group string `json:"group"`

	// version is the version of the thing you're tracking
	// +required
	Version string `json:"version"`

	// resource is the resource type of the resource that you're tracking
	// +required
	Resource string `json:"resource"`

	// namespace is where the thing you're tracking is
	// +optional
	Namespace string `json:"namespace"`

	// name is the name of the resource that you're tracking
	// +required
	Name string `json:"name"`
}

// GenerationStatus keeps track of the generation for a given resource so that decisions about forced updates can be made.
// The definition matches the GenerationStatus defined in github.com/openshift/api/v1
type GenerationStatus struct {
	// group is the group of the resource that you're tracking
	// +required
	Group string `json:"group"`

	// version is the version of the resource that you're tracking
	// +required
	Version string `json:"version"`

	// resource is the resource type of the resource that you're tracking
	// +required
	Resource string `json:"resource"`

	// namespace is where the resource that you're tracking is
	// +optional
	Namespace string `json:"namespace"`

	// name is the name of the resource that you're tracking
	// +required
	Name string `json:"name"`

	// lastGeneration is the last generation of the resource that controller applies
	// +required
	LastGeneration int64 `json:"lastGeneration" protobuf:"varint,5,opt,name=lastGeneration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterManagerList is a collection of deployment configurations for registration and work distribution controllers.
type ClusterManagerList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	// Items is a list of deployment configurations for registration and work distribution controllers.
	Items []ClusterManager `json:"items"`
}
