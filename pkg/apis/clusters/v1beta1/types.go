/*
Copyright 2021 The Clusternet Authors.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// Important: Run "make generated" to regenerate code after modifying this file

type ClusterType string

// These are the valid values for ClusterType
const (
	// self provisioned edge cluster
	EdgeClusterSelfProvisioned ClusterType = "EdgeClusterSelfProvisioned"

	// todo: add more types
)

// ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest
type ClusterRegistrationRequestSpec struct {
	// ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster.
	// It is typically generated by the clusternet agent on the successful creation of a "self-cluster" Lease
	// in the child cluster.
	// Also it is not allowed to change on PUT operations.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	ClusterID types.UID `json:"clusterId"`

	// ClusterType denotes the type of the child cluster.
	//
	// +optional
	// +kubebuilder:validation:Type=string
	ClusterType ClusterType `json:"clusterType,omitempty"`

	// ClusterName is the cluster name.
	// a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character
	//
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:MaxLength=30
	// +kubebuilder:validation:Pattern="[a-z0-9]([-a-z0-9]*[a-z0-9])?([a-z0-9]([-a-z0-9]*[a-z0-9]))*"
	ClusterName string `json:"clusterName,omitempty"`
}

// ClusterRegistrationRequestStatus defines the observed state of ClusterRegistrationRequest
type ClusterRegistrationRequestStatus struct {
	// DedicatedNamespace is a dedicated namespace for the edge cluster, which is created in the parent cluster.
	//
	// +optional
	DedicatedNamespace string `json:"dedicatedNamespace,omitempty"`

	// DedicatedToken is populated by clusternet-hub when Result is RequestApproved.
	// With this token, the client could have full access on the resources created in DedicatedNamespace.
	//
	// +optional
	DedicatedToken []byte `json:"token,omitempty"`

	// CACertificate is the public certificate that is the root of trust for parent cluster
	// The certificate is encoded in PEM format.
	//
	// +optional
	CACertificate []byte `json:"caCertificate,omitempty"`

	// Result indicates whether this request has been approved.
	// When all necessary objects have been created and ready for child cluster registration,
	// this field will be set to "Approved". If any illegal updates on this object, "Illegal" will be set to this filed.
	//
	// +optional
	Result *ApprovedResult `json:"result,omitempty"`

	// ErrorMessage tells the reason why the request is not approved successfully.
	//
	// +optional
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ManagedClusterName is the name of ManagedCluster object in the parent cluster corresponding to the child cluster
	//
	// +optional
	ManagedClusterName string `json:"managedClusterName,omitempty"`
}

type ApprovedResult string

// These are the possible results for a cluster registration request.
const (
	RequestDenied   ApprovedResult = "Denied"
	RequestApproved ApprovedResult = "Approved"
	RequestFailed   ApprovedResult = "Failed"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Cluster",shortName=clsrr,categories=clusternet
// +kubebuilder:printcolumn:name="Cluster-ID",type=string,JSONPath=`.spec.clusterId`,description="The unique id for the cluster"
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.result`,description="The status of current cluster registration request"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API
type ClusterRegistrationRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRegistrationRequestSpec   `json:"spec,omitempty"`
	Status ClusterRegistrationRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRegistrationRequestList contains a list of ClusterRegistrationRequest
type ClusterRegistrationRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRegistrationRequest `json:"items"`
}

// ManagedClusterSpec defines the desired state of ManagedCluster
type ManagedClusterSpec struct {
	// ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster.
	// It is typically generated by the clusternet agent on the successful creation of a "self-cluster" Lease
	// in the child cluster.
	// Also it is not allowed to change on PUT operations.
	//
	// +required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	ClusterID types.UID `json:"clusterId"`

	// ClusterType denotes the type of the child cluster.
	//
	// +optional
	// +kubebuilder:validation:Type=string
	ClusterType ClusterType `json:"clusterType,omitempty"`
}

// ManagedClusterStatus defines the observed state of ManagedCluster
type ManagedClusterStatus struct {
	// lastObservedTime is the time when last status from the series was seen before last heartbeat.
	// RFC 3339 date and time at which the object was acknowledged by the Clusternet Agent.
	// +optional
	LastObservedTime metav1.Time `json:"lastObservedTime,omitempty"`

	// k8sVersion is the Kubernetes version of the cluster
	// +optional
	KubernetesVersion string `json:"k8sVersion,omitempty"`

	// platform indicates the running platform of the cluster
	// +optional
	Platform string `json:"platform,omitempty"`

	// Healthz indicates the healthz status of the cluster
	// which is deprecated since Kubernetes v1.16. Please use Livez and Readyz instead.
	// Leave it here only for compatibility.
	// +optional
	Healthz bool `json:"healthz,omitempty"`

	// Livez indicates the livez status of the cluster
	// +optional
	Livez bool `json:"livez,omitempty"`

	// Readyz indicates the readyz status of the cluster
	// +optional
	Readyz bool `json:"readyz,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope="Namespaced",shortName=mcls,categories=clusternet
// +kubebuilder:printcolumn:name="Cluster-ID",type=string,JSONPath=`.spec.clusterId`,description="The unique id for the cluster"
// +kubebuilder:printcolumn:name="Cluster-Type",type=string,JSONPath=`.spec.clusterType`,description="The type of the cluster"
// +kubebuilder:printcolumn:name="Kubernetes",type=string,JSONPath=".status.k8sVersion"
// +kubebuilder:printcolumn:name="Readyz",type=string,JSONPath=".status.readyz"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ManagedCluster is the Schema for the managedclusters API
type ManagedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedClusterSpec   `json:"spec,omitempty"`
	Status ManagedClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedClusterList contains a list of ManagedCluster
type ManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedCluster `json:"items"`
}
