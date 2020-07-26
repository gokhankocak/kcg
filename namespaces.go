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

type Namespaces struct {
	ApiVersion string            `Json:"apiVersion"`
	Kind       string            `Json:"kind"`
	Items      []Namespace       `Json:"items"`
	Metadata   map[string]string `Json:"metadata"`
}

type Namespace struct {
	ApiVersion string            `Json:"apiVersion"`
	Kind       string            `Json:"kind"`
	Metadata   NamespaceMetadata `Json:"metadata"`
	Spec       interface{}       `Json:"spec"`
	Status     interface{}       `Json:"status"`
}

type NamespaceMetadata struct {
	CreationTimestamp string `Json:"creationTimestamp"`
	Name              string `Json:"name"`
	ResourceVersion   string `Json:"resourceVersion"`
	SelfLink          string `Json:"selfLink"`
	UID               string `Json:"uid"`
}