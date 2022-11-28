// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package networkservices -var YAML_gateway blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/gateway.yaml

package networkservices

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/networkservices/gateway.yaml
var YAML_gateway = []byte("info:\n  title: NetworkServices/Gateway\n  description: The NetworkServices Gateway resource\n  x-dcl-struct-name: Gateway\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Gateway\n    parameters:\n    - name: gateway\n      required: true\n      description: A full instance of a Gateway\n  apply:\n    description: The function used to apply information about a Gateway\n    parameters:\n    - name: gateway\n      required: true\n      description: A full instance of a Gateway\n  delete:\n    description: The function used to delete a Gateway\n    parameters:\n    - name: gateway\n      required: true\n      description: A full instance of a Gateway\n  deleteAll:\n    description: The function used to delete all Gateway\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Gateway\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Gateway:\n      title: Gateway\n      x-dcl-id: projects/{{project}}/locations/{{location}}/gateways/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - name\n      - ports\n      - scope\n      - project\n      - location\n      properties:\n        addresses:\n          type: array\n          x-dcl-go-name: Addresses\n          description: One or more addresses with ports in format of \":\" that the\n            Gateway must receive traffic on. The proxy binds to the ports specified.\n            IP address can be anything that is allowed by the underlying infrastructure\n            (auto-allocation, static IP, BYOIP).\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: string\n            x-dcl-go-type: string\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A free-text description of the resource. Max length\n            1024 characters.\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Optional. Set of label tags associated with the Gateway resource.\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Required. Name of the Gateway resource. It matches pattern\n            `projects/*/locations/global/gateways/`.\n        ports:\n          type: array\n          x-dcl-go-name: Ports\n          description: Required. One or more ports that the Gateway must receive traffic\n            on. The proxy binds to the ports specified. Gateway listen on 0.0.0.0\n            on the ports specified below.\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: integer\n            format: int64\n            x-dcl-go-type: int64\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        scope:\n          type: string\n          x-dcl-go-name: Scope\n          description: Required. Immutable. Scope determines how configuration across\n            multiple Gateway instances are merged. The configuration for multiple\n            Gateway instances with the same scope will be merged as presented as a\n            single coniguration to the proxy/load balancer. Max length 64 characters.\n            Scope should start with a letter and can only have letters, numbers, hyphens.\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: Output only. Server-defined URL of this resource\n          x-kubernetes-immutable: true\n        serverTlsPolicy:\n          type: string\n          x-dcl-go-name: ServerTlsPolicy\n          description: Optional. A fully-qualified ServerTLSPolicy URL reference.\n            Specifies how TLS traffic is terminated. If empty, TLS termination is\n            disabled.\n          x-dcl-references:\n          - resource: Networksecurity/ServerTlsPolicy\n            field: name\n            format: projects/{{project}}/locations/global/serverTlsPolicies/{{name}}\n        type:\n          type: string\n          x-dcl-go-name: Type\n          x-dcl-go-type: GatewayTypeEnum\n          description: 'Immutable. The type of the customer managed gateway. Possible\n            values: TYPE_UNSPECIFIED, OPEN_MESH, SECURE_WEB_GATEWAY'\n          x-kubernetes-immutable: true\n          enum:\n          - TYPE_UNSPECIFIED\n          - OPEN_MESH\n          - SECURE_WEB_GATEWAY\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The timestamp when the resource was updated.\n          x-kubernetes-immutable: true\n")

// 5806 bytes
// MD5: 6ff42fa64dd77ca8b9b258fe0875b9c8
