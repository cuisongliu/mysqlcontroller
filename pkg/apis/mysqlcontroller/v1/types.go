package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroup describes a PodGroup resource
type PodGroup struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodGroupSpec `json:"spec"`
}

// PodGroupSpec is the spec for a MyResource resource
type PodGroupSpec struct {
	DeploymentName  string `json:"deploymentName"`
	Replicas        *int32 `json:"replicas"`
	DeploymentImage string `json:"deploymentImage"`

	MysqlService  string `json:"mysqlService"`
	MysqlPort     string `json:"mysqlPort"`
	MysqlUsername string `json:"mysqlUsername"`
	MysqlPassword string `json:"mysqlPassword"`
	MysqlDb       string `json:"mysqlDb"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroupList is a list of PodGroup resources
type PodGroupList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []PodGroup `json:"items"`
}
