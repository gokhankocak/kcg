package kcg

// DaemonSets holds data about daemonsets
type DaemonSets struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []DaemonSet       `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

// DaemonSet holds data about a daemonset
type DaemonSet struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   DaemonSetMetadata `json:"metadata"`
	Spec       DaemonSetSpec     `json:"spec"`
	Status     DaemonSetStatus   `json:"status"`
}

// DaemonSetMetadata holds metadata for a daemonset
type DaemonSetMetadata struct {
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

// DaemonSetSpec holds specifications for a daemonset
type DaemonSetSpec struct {
	RevisionHistoryLimit int64                   `json:"revisionHistoryLimit"`
	Selector             map[string]interface{}  `json:"selector"`
	UpdateStrategy       DaemonSetUpdateStrategy `json:"updateStrategy"`
	Template             DaemonSetTemplate       `json:"template"`
}

// DaemonSetStatus holds status for a daemonset
type DaemonSetStatus struct {
	CurrentNumberScheduled int64 `json:"currentNumberScheduled"`
	DesiredNumberScheduled int64 `json:"desiredNumberScheduled"`
	NumberAvailable        int64 `json:"numberAvailable"`
	NumberMisscheduled     int64 `json:"numberMisscheduled"`
	NumberReady            int64 `json:"numberReady"`
	ObservedGeneration     int64 `json:"observedGeneration"`
	UpdatedNumberScheduled int64 `json:"updatedNumberScheduled"`
}

// DaemonSetUpdateStrategy holds update strategy for a daemonset
type DaemonSetUpdateStrategy struct {
	Type           string      `json:"type"`
	UpdateStrategy interface{} `json:"updateStrategy"`
}

// DaemonSetTemplate holds data for a daemonset spec template
type DaemonSetTemplate struct {
	Metadata DaemonSetTemplateMetadata `json:"metadata"`
	Spec     DaemonSetTemplateSpec     `json:"spec"`
}

// DaemonSetTemplateMetadata holds data about template metedata
type DaemonSetTemplateMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

// DaemonSetTemplateSpec holds data about template spec
type DaemonSetTemplateSpec struct {
	DNSPolicy                     string                 `json:"dnsPolicy"`
	RestartPolicy                 string                 `json:"restartPolicy"`
	SchedulerName                 string                 `json:"schedulerName"`
	SecurityContext               map[string]interface{} `json:"securityContext"`
	ServiceAccount                string                 `json:"serviceAccount"`
	ServiceAccountName            string                 `json:"serviceAccountName"`
	TerminationGracePeriodSeconds int64                  `json:"terminationGracePeriodSeconds"`
	Containers                    []Container            `json:"containers"`
}
