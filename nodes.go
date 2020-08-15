package kcg

// Nodes holds data for an array of nodes
type Nodes struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Items      []Node            `json:"items"`
	Metadata   map[string]string `json:"metadata"`
}

// Node holds data for a node
type Node struct {
	APIVersion string       `json:"apiVersion"`
	Kind       string       `json:"kind"`
	Metadata   NodeMetadata `json:"metadata"`
	Spec       NodeSpec     `json:"spec"`
	Status     NodeStatus   `json:"status"`
}

// NodeMetadata holds metadata for a node
type NodeMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp"`
	Name              string            `json:"name"`
	ResourceVersion   string            `json:"resourceVersion"`
	SelfLink          string            `json:"selfLink"`
	UID               string            `json:"uid"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

// NodeSpec holds spcification for a node
type NodeSpec struct {
	PodCIDR    string   `json:"podCIDR"`
	PodCIDRs   []string `json:"podCIDRs"`
	ProviderID string   `json:"providerID"`
}

// NodeAddress holds a node address
type NodeAddress struct {
	Address string
	Type    string
}

// NodeCPUMemPod holds CPU, Memory data
type NodeCPUMemPod struct {
	CPU              string `json:"cpu"`
	EphemeralStorage string `json:"ephemeral-storage"`
	HugePages1Gi     bool   `json:"hugepages-1Gi"`
	HugePages2Mi     bool   `json:"hugepages-2Mi"`
	Memory           string `json:"memory"`
	Pods             string `json:"pods"`
}

// NodeImage holds the image data
type NodeImage struct {
	Names     []string `json:"names"`
	SizeBytes int64    `json:"sizeBytes"`
}

// NodeInfo holds node OS and Architecture data
type NodeInfo struct {
	Architecture            string `json:"architecture"`
	BootID                  string `json:"bootID"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KernelVersion           string `json:"kernelVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	MachineID               string `json:"machineID"`
	OperatingSystem         string `json:"operatingSystem"`
	OsImage                 string `json:"osImage"`
	SystemUUID              string `json:"systemUUID"`
}

// NodeStatus holds status data for a node
type NodeStatus struct {
	Addresses       []NodeAddress `json:"addresses"`
	Allocatable     NodeCPUMemPod `json:"allocatable"`
	Capacity        NodeCPUMemPod `json:"capacity"`
	Conditions      []Condition   `json:"conditions"`
	DaemonEndPoints interface{}   `json:"daemonEndpoints"`
	Images          []NodeImage   `json:"images"`
	Info            NodeInfo      `json:"nodeInfo"`
}
