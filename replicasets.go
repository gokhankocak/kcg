package kcg

// ReplicaSets holds data about Replicasets
type ReplicaSets struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []ReplicaSet      `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

// ReplicaSet holds data about a Replicaset
type ReplicaSet struct {
	APIVersion string             `json:"apiVersion"`
	Kind       string             `json:"kind"`
	Metadata   ReplicaSetMetadata `json:"metadata"`
	Spec       ReplicaSetSpec     `json:"spec"`
	Status     ReplicaSetStatus   `json:"status"`
}

// ReplicaSetMetadata holds metadata for a Replicaset
type ReplicaSetMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	ResourceVersion   string            `json:"resourceVersion"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	Generation        int64             `json:"generation"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences"`
}

// ReplicaSetSpec holds specifications for a Replicaset
type ReplicaSetSpec struct {
	Replicas int64                  `json:"replicas"`
	Selector map[string]interface{} `json:"selector"`
	Template ReplicaSetTemplate     `json:"template"`
}

// ReplicaSetStatus holds status for a Replicaset
type ReplicaSetStatus struct {
	AvailableReplicas    int64 `json:"availableReplicas"`
	FullyLabeledReplicas int64 `json:"fullyLabeledReplicas"`
	ObservedGeneration   int64 `json:"observedGeneration"`
	ReadyReplicas        int64 `json:"readyReplicas"`
	Replicas             int64 `json:"replicas"`
}

// ReplicaSetTemplate holds data for a Replicaset spec template
type ReplicaSetTemplate struct {
	Metadata ReplicaSetTemplateMetadata `json:"metadata"`
	Spec     ReplicaSetTemplateSpec     `json:"spec"`
}

// ReplicaSetTemplateMetadata holds data about template metedata
type ReplicaSetTemplateMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

// ReplicaSetTemplateSpec holds data about template spec
type ReplicaSetTemplateSpec struct {
	DNSPolicy                     string                 `json:"dnsPolicy"`
	RestartPolicy                 string                 `json:"restartPolicy"`
	SchedulerName                 string                 `json:"schedulerName"`
	SecurityContext               map[string]interface{} `json:"securityContext"`
	ServiceAccount                string                 `json:"serviceAccount"`
	ServiceAccountName            string                 `json:"serviceAccountName"`
	TerminationGracePeriodSeconds int64                  `json:"terminationGracePeriodSeconds"`
	Containers                    []Container            `json:"containers"`
}
