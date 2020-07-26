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

type Services struct {
	ApiVersion string            `Json:"apiVersion"`
	Kind       string            `Json:"kind"`
	Items      []Service         `Json:"item"`
	Metadata   map[string]string `Json:"metadata"`
}

type Service struct {
	ApiVersion string          `Json:"apiVersion"`
	Kind       string          `Json:"kind"`
	Metadata   ServiceMetadata `Json:"metadata"`
	Spec       ServiceSpec     `Json:"spec"`
}

type ServiceMetadata struct {
	Name              string            `Json:"name"`
	Namespace         string            `Json:"namespace"`
	ResourceVersion   string            `Json:"resourceVersion"`
	SelfLink          string            `Json:"selfLink"`
	UID               string            `Json:"uid"`
	CreationTimestamp string            `Json:"creationTimestamp"`
	Labels            map[string]string `Json:"labels"`
	Annotations       map[string]string `Json:"annotations"`
}

type ServiceSpec struct {
	ClusterIP             string            `Json:"clusterIP"`
	ExternalTrafficPolicy string            `Json:"externalTrafficPolicy"`
	Type                  string            `Json:"type"`
	SessionAffinity       string            `Json:"sessionAffinity"`
	Selector              map[string]string `Json:"selector"`
	Ports                 []ServicePort     `Json:"ports,omitempty"`
	Status                interface{}       `Json:"status"`
}

type ServicePort struct {
	Name       string `Json:"name"`
	NodePort   int    `Json:"nodePort"`
	Port       int    `Json:"port"`
	Protocol   string `Json:"protocol"`
	TargetPort int    `Json:"targetPort"` // TODO: This should be interface{} because it is sometimes string, sometimes int
}
