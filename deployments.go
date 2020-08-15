package kcg

// Deployments holds the deployments data for an array of deployments
type Deployments struct {
	APIVersion string            `json:"APIVersion"`
	Kind       string            `json:"kind"`
	Items      []Deployment      `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

// Deployment hold deployment data for each deployment object
type Deployment struct {
	APIVersion string             `json:"APIVersion"`
	Kind       string             `json:"kind"`
	Metadata   DeploymentMetadata `json:"metadata"`
	Spec       DeploymentSpec     `json:"spec"`
	Status     DeploymentStatus   `json:"status"`
}

// DeploymentMetadata holds metadata for a deployment
type DeploymentMetadata struct {
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

// DeploymentSpec holds specification for a deployment
type DeploymentSpec struct {
	ProgressDeadlineSeconds int64                  `json:"progressDeadlineSeconds"`
	Replicas                int64                  `json:"replicas"`
	RevisionHistoryLimit    int64                  `json:"revisionHistoryLimit"`
	Selector                map[string]interface{} `json:"selector"`
	Strategy                DeploymentStrategy     `json:"strategy"`
	Template                DeploymentTemplate     `json:"template"`
}

// DeploymentStatus holds status of a deployment
type DeploymentStatus struct {
	AvailableReplicas  int64       `json:"availableReplicas"`
	Conditions         []Condition `json:"conditions"`
	ObservedGeneration int64       `json:"observedGeneration"`
	ReadyReplicas      int64       `json:"readyReplicas"`
	Replicas           int64       `json:"replicas"`
	UpdatedReplicas    int64       `json:"updatedReplicas"`
}

// DeploymentStrategy holds data about the strategy
type DeploymentStrategy struct {
	Type          string                  `json:"type"`
	RollingUpdate DeploymentRollingUpdate `json:"rollingUpdate"`
}

// DeploymentRollingUpdate holds data about rolling update
type DeploymentRollingUpdate struct {
	MaxSurge       interface{} `json:"maxSurge"`
	MaxUnavailable interface{} `json:"maxUnavailable"`
}

// DeploymentTemplate holds data about a deployment template
type DeploymentTemplate struct {
	Metadata DeploymentTemplateMetadata `json:"metadata"`
	Spec     DeploymentTemplateSpec     `json:"spec"`
}

// DeploymentTemplateMetadata holds data about template metedata
type DeploymentTemplateMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

// DeploymentTemplateSpec holds data about template spec
type DeploymentTemplateSpec struct {
	DNSPolicy                     string                 `json:"dnsPolicy"`
	RestartPolicy                 string                 `json:"restartPolicy"`
	SchedulerName                 string                 `json:"schedulerName"`
	SecurityContext               map[string]interface{} `json:"securityContext"`
	ServiceAccount                string                 `json:"serviceAccount"`
	ServiceAccountName            string                 `json:"serviceAccountName"`
	TerminationGracePeriodSeconds int64                  `json:"terminationGracePeriodSeconds"`
	Containers                    []Container            `json:"containers"`
}
