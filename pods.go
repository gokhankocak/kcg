package kcg

type Pods struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []Pod             `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

type Pod struct {
	ApiVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   PodMetadata `json:"metadata"`
	Spec       PodSpec     `json:"spec"`
}

type PodMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	GenerateName      string            `json:"generateName"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	ResourceVersion   string            `json:"resourceVersion"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences"`
}

type PodSpec struct {
	DnsPolicy                     string               `json:"dnsPolicy"`
	NodeName                      string               `json:"nodeName"`
	HostName                      string               `json:"hostName"`
	RestartPolicy                 string               `json:"restartPolicy"`
	TerminationGracePeriodSeconds int                  `json:"terminationGracePeriodSeconds"`
	Containers                    []Container          `json:"containers"`
	InitContainers                []InitContainer      `json:"initContainers"`
	ImagePullSecrets              []PodImagePullSecret `json:"imagePullSecrets"`
	Tolerations                   []PodToleration      `json:"tolerations"`
	Volumes                       []PodVolume          `json:"volumes"`
	SecurityContext               interface{}          `json:"securityContext"`
	ServiceAccount                string               `json:"serviceAccount"`
	ServiceAccountName            string               `json:"serviceAccountName"`
	Subdomain                     string               `json:"subdomain"`
}

type PodImagePullSecret struct {
	Name string `json:"name"`
}

type PodToleration struct {
	Effect            string `json:"effect"`
	Key               string `json:"key"`
	Operator          string `json:"operator"`
	TolerationSeconds int    `json:"tolerationSeconds"`
}

type PodVolume struct {
	Name                  string            `json:"name"`
	PersistentVolumeClaim map[string]string `json:"persistentVolumeClaim"`
	EmptyDir              interface{}       `json:"emptyDir"`
	Secret                map[string]string `json:"secret"`
}
