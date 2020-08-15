package kcg

type Namespaces struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []Namespace       `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

type Namespace struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   NamespaceMetadata `json:"metadata"`
	Spec       interface{}       `json:"spec"`
	Status     interface{}       `json:"status"`
}

type NamespaceMetadata struct {
	CreationTimestamp string `json:"creationTimestamp"`
	Name              string `json:"name"`
	ResourceVersion   string `json:"resourceVersion"`
	SelfLink          string `json:"selfLink"`
	UID               string `json:"uid"`
}
