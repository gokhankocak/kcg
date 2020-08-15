package kcg

type Container struct {
	Args            []string               `json:"args"`
	Command         []string               `json:"command"`
	Env             []ContainerEnv         `json:"env"`
	Image           string                 `json:"image"`
	ImagePullPolicy string                 `json:"imagePullPolicy"`
	Name            string                 `json:"name"`
	Ports           []ContainerPort        `json:"ports"`
	VolumeMounts    []ContainerVolumeMount `json:"volumeMounts"`
	LivenessProbe   interface{}            `json:"livenessProbe"`
	ReadinessProbe  interface{}            `json:"readinessProbe"`
}

type ContainerEnv struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ContainerPort struct {
	CntrPort int    `json:"containerPort"`
	Protocol string `json:"protocol"`
	Name     string `json:"name"`
}

type ContainerResource struct {
	Limits   Resource `json:"limits"`
	Requests Resource `json:"requests"`
}

type Resource struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

type ContainerVolumeMount struct {
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
	SubPath   string `json:"subPath"`
	ReadOnly  bool   `json:"readOnly"`
}

type InitContainer struct {
	Name                     string                 `json:"name"`
	Image                    string                 `json:"image"`
	ImagePullPolicy          string                 `json:"imagePullPolicy"`
	Resources                interface{}            `json:"resources"`
	TerminationMessagePath   string                 `json:"terminationMessagePath"`
	TerminationMessagePolicy string                 `json:"terminationMessagePolicy"`
	VolumeMounts             []ContainerVolumeMount `json:"volumeMounts"`
}
