package kcg

type Services struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []Service         `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

type Service struct {
	ApiVersion string          `json:"apiVersion"`
	Kind       string          `json:"kind"`
	Metadata   ServiceMetadata `json:"metadata"`
	Spec       ServiceSpec     `json:"spec"`
	Status     interface{}     `json:"status"`
}

type ServiceMetadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	ResourceVersion   string            `json:"resourceVersion"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences"`
}

type ServiceSpec struct {
	ClusterIP             string            `json:"clusterIP"`
	ExternalTrafficPolicy string            `json:"externalTrafficPolicy"`
	Type                  string            `json:"type"`
	SessionAffinity       string            `json:"sessionAffinity"`
	Selector              map[string]string `json:"selector"`
	Ports                 []ServicePort     `json:"ports,omitempty"`
}

type ServicePort struct {
	Name       string      `json:"name"`
	NodePort   int         `json:"nodePort"`
	Port       int         `json:"port"`
	Protocol   string      `json:"protocol"`
	TargetPort interface{} `json:"targetPort"` // TODO: This should be interface{} because it is sometimes string, sometimes int
}
