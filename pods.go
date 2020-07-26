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

type Pods struct {
	ApiVersion string            `Json:"apiVersion"`
	Kind       string            `Json:"kind"`
	Items      []Pod             `Json:"items"`
	Metadata   map[string]string `Json:"metadata"`
}

type Pod struct {
	ApiVersion string      `Json:"apiVersion"`
	Kind       string      `Json:"kind"`
	Metadata   PodMetadata `Json:"metadata"`
	Spec       PodSpec     `Json:"spec"`
}

type PodMetadata struct {
	CreationTimestamp string            `Json:"creationTimestamp"`
	GenerateName      string            `Json:"generateName"`
	Name              string            `Json:"name"`
	Namespace         string            `Json:"namespace"`
	ResourceVersion   string            `Json:"resourceVersion"`
	SelfLink          string            `Json:"selfLink"`
	UID               string            `Json:"uid"`
	Labels            map[string]string `Json:"labels"`
	Annotations       map[string]string `Json:"annotations"`
	OwnerReferences   []OwnerReference  `Json:"ownerReferences"`
}

type PodSpec struct {
	DnsPolicy                     string               `Json:"dnsPolicy"`
	NodeName                      string               `Json:"nodeName"`
	HostName                      string               `Json:"hostName"`
	RestartPolicy                 string               `Json:"restartPolicy"`
	TerminationGracePeriodSeconds int                  `Json:"terminationGracePeriodSeconds"`
	Containers                    []Container          `Json:"containers"`
	InitContainers                interface{}          `Json:"initContainers"`
	ImagePullSecrets              []PodImagePullSecret `Json:"imagePullSecrets"`
	Tolerations                   []PodToleration      `Json:"tolerations"`
	Volumes                       []PodVolume          `Json:"volumes"`
	SecurityContext               interface{}          `Json:"securityContext"`
	ServiceAccount                string               `Json:"serviceAccount"`
	ServiceAccountName            string               `Json:"serviceAccountName"`
	Subdomain                     string               `Json:"subdomain"`
}

type PodImagePullSecret struct {
	Name string `Json:"name"`
}

type PodToleration struct {
	Effect            string `Json:"effect"`
	Key               string `Json:"key"`
	Operator          string `Json:"operator"`
	TolerationSeconds int    `Json:"tolerationSeconds"`
}

type PodVolume struct {
	Name                  string            `Json:"name"`
	PersistentVolumeClaim map[string]string `Json:"persistentVolumeClaim"`
	EmptyDir              interface{}       `Json:"emptyDir"`
	Secret                map[string]string `Json:"secret"`
}
