/*

Copyright 2020 Gokhan Kocak (gokhan.kocak@mail.ru)
https://github.com/gokhankocak/VisualKubernetes

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package visualkubernetes

type Container struct {
	Args            []string               `Json:"args"`
	Command         []string               `Json:"command"`
	Env             []ContainerEnv         `Json:"env"`
	Image           string                 `Json:"image"`
	ImagePullPolicy string                 `Json:"imagePullPolicy"`
	Name            string                 `Json:"name"`
	ContainerPorts  []ContainerPort        `Json:"ports"`
	VolumeMounts    []ContainerVolumeMount `Json:"volumeMounts"`
	LivenessProbe   interface{}            `Json:"livenessProbe"`
	ReadinessProbe  interface{}            `Json:"readinessProbe"`
}

type ContainerEnv struct {
	Name  string `Json:"name"`
	Value string `Json:"value"`
}

type ContainerPort struct {
	ContainerPort int    `Json:"containerPort"`
	Protocol      string `Json:"protocol"`
}

type ContainerResource struct {
	Limits   Resource `Json:"limits"`
	Requests Resource `Json:"requests"`
}

type Resource struct {
	CPU    string `Json:"cpu"`
	Memory string `Json:"memory"`
}

type ContainerVolumeMount struct {
	MountPath string `Json:"mountPath"`
	Name      string `Json:"name"`
	SubPath   string `Json:"subPath"`
	ReadOnly  bool   `Json:"readOnly"`
}
