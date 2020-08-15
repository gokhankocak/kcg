package kcg

// StatefulSets holds data about Statefulsets
type StatefulSets struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []StatefulSet     `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

// StatefulSet holds data about a Statefulset
type StatefulSet struct {
	APIVersion string              `json:"apiVersion"`
	Kind       string              `json:"kind"`
	Metadata   StatefulSetMetadata `json:"metadata"`
	Spec       StatefulSetSpec     `json:"spec"`
	Status     StatefulSetStatus   `json:"status"`
}

// StatefulSetMetadata holds metadata for a Statefulset
type StatefulSetMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	ResourceVersion   string            `json:"resourceVersion"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	Generation        int64             `json:"generation"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

// StatefulSetSpec holds specifications for a Statefulset
type StatefulSetSpec struct {
	PodManagementPolicy  string                            `json:"podManagementPolicy"`
	Replicas             int64                             `json:"replicas"`
	RevisionHistoryLimit int64                             `json:"revisionHistoryLimit"`
	ServiceName          string                            `json:"serviceName"`
	Selector             map[string]interface{}            `json:"selector"`
	UpdateStrategy       map[string]interface{}            `json:"updateStrategy"`
	Template             StatefulSetTemplate               `json:"template"`
	VolumeClaimTemplates []StatefulSetsVolumeClaimTemplate `json:"volumeClaimTemplates"`
}

// StatefulSetStatus holds status for a Statefulset
type StatefulSetStatus struct {
	CollisionCount     int64  `json:"collisionCount"`
	CurrentReplicas    int64  `json:"currentReplicas"`
	CurrentRevision    string `json:"currentRevision"`
	ObservedGeneration int64  `json:"observedGeneration"`
	ReadyReplicas      int64  `json:"readyReplicas"`
	Replicas           int64  `json:"replicas"`
	UpdateRevision     string `json:"updateRevision"`
	UpdatedReplicas    int64  `json:"updatedReplicas"`
}

// StatefulSetTemplate holds data for a Statefulset spec template
type StatefulSetTemplate struct {
	Metadata StatefulSetTemplateMetadata `json:"metadata"`
	Spec     StatefulSetTemplateSpec     `json:"spec"`
}

// StatefulSetTemplateMetadata holds data about template metedata
type StatefulSetTemplateMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

// StatefulSetTemplateSpec holds data about template spec
type StatefulSetTemplateSpec struct {
	DNSPolicy                     string                 `json:"dnsPolicy"`
	RestartPolicy                 string                 `json:"restartPolicy"`
	SchedulerName                 string                 `json:"schedulerName"`
	SecurityContext               map[string]interface{} `json:"securityContext"`
	ServiceAccount                string                 `json:"serviceAccount"`
	ServiceAccountName            string                 `json:"serviceAccountName"`
	TerminationGracePeriodSeconds int64                  `json:"terminationGracePeriodSeconds"`
	Containers                    []Container            `json:"containers"`
	Volumes                       []StatefulSetVolume    `json:"volumes"`
}

// StatefulSetsVolumeClaimTemplate holds data for volumeClaimTemplate
type StatefulSetsVolumeClaimTemplate struct {
}

// StatefulSetVolume holds data for a volume
type StatefulSetVolume struct {
	Name      string                 `json:"name"`
	Secret    map[string]interface{} `json:"secret"`
	ConfigMap map[string]interface{} `json:"configMap"`
	EmptyDir  map[string]interface{} `json:"emptyDir"`
}
