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
// gen_go_data -package compute -var YAML_forwarding_rule blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/forwarding_rule.yaml

package compute

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/compute/forwarding_rule.yaml
var YAML_forwarding_rule = []byte("info:\n  title: Compute/ForwardingRule\n  description: The Compute ForwardingRule resource\n  x-dcl-struct-name: ForwardingRule\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a ForwardingRule\n    parameters:\n    - name: ForwardingRule\n      required: true\n      description: A full instance of a ForwardingRule\n  apply:\n    description: The function used to apply information about a ForwardingRule\n    parameters:\n    - name: ForwardingRule\n      required: true\n      description: A full instance of a ForwardingRule\n  delete:\n    description: The function used to delete a ForwardingRule\n    parameters:\n    - name: ForwardingRule\n      required: true\n      description: A full instance of a ForwardingRule\n  deleteAll:\n    description: The function used to delete all ForwardingRule\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many ForwardingRule\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    ForwardingRule:\n      title: ForwardingRule\n      x-dcl-id: projects/{{project}}/global/forwardingRules/{{name}}\n      x-dcl-locations:\n      - region\n      - global\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - project\n      properties:\n        allPorts:\n          type: boolean\n          x-dcl-go-name: AllPorts\n          description: This field is used along with the `backend_service` field for\n            internal load balancing or with the `target` field for internal TargetInstance.\n            This field cannot be used with `port` or `portRange` fields. When the\n            load balancing scheme is `INTERNAL` and protocol is TCP/UDP, specify this\n            field to allow packets addressed to any ports will be forwarded to the\n            backends configured with this forwarding rule.\n          x-kubernetes-immutable: true\n        allowGlobalAccess:\n          type: boolean\n          x-dcl-go-name: AllowGlobalAccess\n          description: This field is used along with the `backend_service` field for\n            internal load balancing or with the `target` field for internal TargetInstance.\n            If the field is set to `TRUE`, clients can access ILB from all regions.\n            Otherwise only allows access from clients in the same region as the internal\n            load balancer.\n        backendService:\n          type: string\n          x-dcl-go-name: BackendService\n          description: This field is only used for `INTERNAL` load balancing. For\n            internal load balancing, this field identifies the BackendService resource\n            to receive the matched traffic.\n          x-kubernetes-immutable: true\n        creationTimestamp:\n          type: string\n          x-dcl-go-name: CreationTimestamp\n          readOnly: true\n          description: '[Output Only] Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt)\n            text format.'\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: An optional description of this resource. Provide this property\n            when you create the resource.\n          x-kubernetes-immutable: true\n        ipAddress:\n          type: string\n          x-dcl-go-name: IPAddress\n          description: 'IP address that this forwarding rule serves. When a client\n            sends traffic to this IP address, the forwarding rule directs the traffic\n            to the target that you specify in the forwarding rule. If you don''t specify\n            a reserved IP address, an ephemeral IP address is assigned. Methods for\n            specifying an IP address: * IPv4 dotted decimal, as in `100.1.2.3` * Full\n            URL, as in `https://www.googleapis.com/compute/v1/projects/project_id/regions/region/addresses/address-name`\n            * Partial URL or by name, as in: * `projects/project_id/regions/region/addresses/address-name`\n            * `regions/region/addresses/address-name` * `global/addresses/address-name`\n            * `address-name` The loadBalancingScheme and the forwarding rule''s target\n            determine the type of IP address that you can use. For detailed information,\n            refer to [IP address specifications](/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).'\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        ipProtocol:\n          type: string\n          x-dcl-go-name: IPProtocol\n          x-dcl-go-type: ForwardingRuleIPProtocolEnum\n          description: The IP protocol to which this rule applies. For protocol forwarding,\n            valid options are `TCP`, `UDP`, `ESP`, `AH`, `SCTP` or `ICMP`. For Internal\n            TCP/UDP Load Balancing, the load balancing scheme is `INTERNAL`, and one\n            of `TCP` or `UDP` are valid. For Traffic Director, the load balancing\n            scheme is `INTERNAL_SELF_MANAGED`, and only `TCP`is valid. For Internal\n            HTTP(S) Load Balancing, the load balancing scheme is `INTERNAL_MANAGED`,\n            and only `TCP` is valid. For HTTP(S), SSL Proxy, and TCP Proxy Load Balancing,\n            the load balancing scheme is `EXTERNAL` and only `TCP` is valid. For Network\n            TCP/UDP Load Balancing, the load balancing scheme is `EXTERNAL`, and one\n            of `TCP` or `UDP` is valid.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          enum:\n          - TCP\n          - UDP\n          - ESP\n          - AH\n          - SCTP\n          - ICMP\n          - L3_DEFAULT\n        ipVersion:\n          type: string\n          x-dcl-go-name: IPVersion\n          x-dcl-go-type: ForwardingRuleIPVersionEnum\n          description: 'The IP Version that will be used by this forwarding rule.\n            Valid options are `IPV4` or `IPV6`. This can only be specified for an\n            external global forwarding rule. Possible values: UNSPECIFIED_VERSION,\n            IPV4, IPV6'\n          x-kubernetes-immutable: true\n          enum:\n          - UNSPECIFIED_VERSION\n          - IPV4\n          - IPV6\n        isMirroringCollector:\n          type: boolean\n          x-dcl-go-name: IsMirroringCollector\n          description: Indicates whether or not this load balancer can be used as\n            a collector for packet mirroring. To prevent mirroring loops, instances\n            behind this load balancer will not have their traffic mirrored even if\n            a `PacketMirroring` rule applies to them. This can only be set to true\n            for load balancers that have their `loadBalancingScheme` set to `INTERNAL`.\n          x-kubernetes-immutable: true\n        labelFingerprint:\n          type: string\n          x-dcl-go-name: LabelFingerprint\n          readOnly: true\n          description: Used internally during label updates.\n          x-kubernetes-immutable: true\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: Labels to apply to this rule.\n        loadBalancingScheme:\n          type: string\n          x-dcl-go-name: LoadBalancingScheme\n          x-dcl-go-type: ForwardingRuleLoadBalancingSchemeEnum\n          description: \"Specifies the forwarding rule type.\\n\\n*   `EXTERNAL` is used\n            for:\\n    *   Classic Cloud VPN gateways\\n    *   Protocol forwarding\n            to VMs from an external IP address\\n    *   The following load balancers:\n            HTTP(S), SSL Proxy, TCP Proxy, and Network TCP/UDP\\n*   `INTERNAL` is\n            used for:\\n    *   Protocol forwarding to VMs from an internal IP address\\n\n            \\   *   Internal TCP/UDP load balancers\\n*   `INTERNAL_MANAGED` is used\n            for:\\n    *   Internal HTTP(S) load balancers\\n*   `INTERNAL_SELF_MANAGED`\n            is used for:\\n    *   Traffic Director\\n*   `EXTERNAL_MANAGED` is used\n            for:\\n    *   Global external HTTP(S) load balancers \\n\\nFor more information\n            about forwarding rules, refer to [Forwarding rule concepts](/load-balancing/docs/forwarding-rule-concepts).\n            Possible values: INVALID, INTERNAL, INTERNAL_MANAGED, INTERNAL_SELF_MANAGED,\n            EXTERNAL, EXTERNAL_MANAGED\"\n          x-kubernetes-immutable: true\n          enum:\n          - INVALID\n          - INTERNAL\n          - INTERNAL_MANAGED\n          - INTERNAL_SELF_MANAGED\n          - EXTERNAL\n          - EXTERNAL_MANAGED\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location of this resource.\n          x-kubernetes-immutable: true\n        metadataFilter:\n          type: array\n          x-dcl-go-name: MetadataFilter\n          description: |-\n            Opaque filter criteria used by Loadbalancer to restrict routing configuration to a limited set of [xDS](https://github.com/envoyproxy/data-plane-api/blob/master/XDS_PROTOCOL.md) compliant clients. In their xDS requests to Loadbalancer, xDS clients present [node metadata](https://github.com/envoyproxy/data-plane-api/search?q=%22message+Node%22+in%3A%2Fenvoy%2Fapi%2Fv2%2Fcore%2Fbase.proto&). If a match takes place, the relevant configuration is made available to those proxies. Otherwise, all the resources (e.g. `TargetHttpProxy`, `UrlMap`) referenced by the `ForwardingRule` will not be visible to those proxies.\n\n            For each `metadataFilter` in this list, if its `filterMatchCriteria` is set to MATCH_ANY, at least one of the `filterLabel`s must match the corresponding label provided in the metadata. If its `filterMatchCriteria` is set to MATCH_ALL, then all of its `filterLabel`s must match with corresponding labels provided in the metadata.\n\n            `metadataFilters` specified here will be applifed before those specified in the `UrlMap` that this `ForwardingRule` references.\n\n            `metadataFilters` only applies to Loadbalancers that have their loadBalancingScheme set to `INTERNAL_SELF_MANAGED`.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ForwardingRuleMetadataFilter\n            required:\n            - filterMatchCriteria\n            - filterLabel\n            properties:\n              filterLabel:\n                type: array\n                x-dcl-go-name: FilterLabel\n                description: |-\n                  The list of label value pairs that must match labels in the provided metadata based on `filterMatchCriteria`\n\n                  This list must not be empty and can have at the most 64 entries.\n                x-kubernetes-immutable: true\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: object\n                  x-dcl-go-type: ForwardingRuleMetadataFilterFilterLabel\n                  required:\n                  - name\n                  - value\n                  properties:\n                    name:\n                      type: string\n                      x-dcl-go-name: Name\n                      description: |-\n                        Name of metadata label.\n\n                        The name can have a maximum length of 1024 characters and must be at least 1 character long.\n                      x-kubernetes-immutable: true\n                    value:\n                      type: string\n                      x-dcl-go-name: Value\n                      description: |-\n                        The value of the label must match the specified value.\n\n                        value can have a maximum length of 1024 characters.\n                      x-kubernetes-immutable: true\n              filterMatchCriteria:\n                type: string\n                x-dcl-go-name: FilterMatchCriteria\n                x-dcl-go-type: ForwardingRuleMetadataFilterFilterMatchCriteriaEnum\n                description: |-\n                  Specifies how individual `filterLabel` matches within the list of `filterLabel`s contribute towards the overall `metadataFilter` match.\n\n                  Supported values are:\n\n                  *   MATCH_ANY: At least one of the `filterLabels` must have a matching label in the provided metadata.\n                  *   MATCH_ALL: All `filterLabels` must have matching labels in the provided metadata. Possible values: NOT_SET, MATCH_ALL, MATCH_ANY\n                x-kubernetes-immutable: true\n                enum:\n                - NOT_SET\n                - MATCH_ALL\n                - MATCH_ANY\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: Name of the resource; provided by the client when the resource\n            is created. The name must be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).\n            Specifically, the name must be 1-63 characters long and match the regular\n            expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character\n            must be a lowercase letter, and all following characters must be a dash,\n            lowercase letter, or digit, except the last character, which cannot be\n            a dash.\n          x-kubernetes-immutable: true\n        network:\n          type: string\n          x-dcl-go-name: Network\n          description: This field is not used for external load balancing. For `INTERNAL`\n            and `INTERNAL_SELF_MANAGED` load balancing, this field identifies the\n            network that the load balanced IP should belong to for this Forwarding\n            Rule. If this field is not specified, the default network will be used.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        networkTier:\n          type: string\n          x-dcl-go-name: NetworkTier\n          x-dcl-go-type: ForwardingRuleNetworkTierEnum\n          description: 'This signifies the networking tier used for configuring this\n            load balancer and can only take the following values: `PREMIUM`, `STANDARD`.\n            For regional ForwardingRule, the valid values are `PREMIUM` and `STANDARD`.\n            For GlobalForwardingRule, the valid value is `PREMIUM`. If this field\n            is not specified, it is assumed to be `PREMIUM`. If `IPAddress` is specified,\n            this value must be equal to the networkTier of the Address.'\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n          enum:\n          - PREMIUM\n          - STANDARD\n        portRange:\n          type: string\n          x-dcl-go-name: PortRange\n          description: |-\n            When the load balancing scheme is `EXTERNAL`, `INTERNAL_SELF_MANAGED` and `INTERNAL_MANAGED`, you can specify a `port_range`. Use with a forwarding rule that points to a target proxy or a target pool. Do not use with a forwarding rule that points to a backend service. This field is used along with the `target` field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when `IPProtocol` is `TCP`, `UDP`, or `SCTP`, only packets addressed to ports in the specified range will be forwarded to `target`. Forwarding rules with the same `[IPAddress, IPProtocol]` pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:\n\n            *   TargetHttpProxy: 80, 8080\n            *   TargetHttpsProxy: 443\n            *   TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222\n            *   TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995, 1688, 1883, 5222\n            *   TargetVpnGateway: 500, 4500\n\n            @pattern: d+(?:-d+)?\n          x-kubernetes-immutable: true\n        ports:\n          type: array\n          x-dcl-go-name: Ports\n          description: 'This field is used along with the `backend_service` field\n            for internal load balancing. When the load balancing scheme is `INTERNAL`,\n            a list of ports can be configured, for example, [''80''], [''8000'',''9000''].\n            Only packets addressed to these ports are forwarded to the backends configured\n            with the forwarding rule. If the forwarding rule''s loadBalancingScheme\n            is INTERNAL, you can specify ports in one of the following ways: * A list\n            of up to five ports, which can be non-contiguous * Keyword `ALL`, which\n            causes the forwarding rule to forward traffic on any port of the forwarding\n            rule''s protocol. @pattern: d+(?:-d+)? For more information, refer to\n            [Port specifications](/load-balancing/docs/forwarding-rule-concepts#port_specifications).'\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: set\n          items:\n            type: string\n            x-dcl-go-type: string\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project this resource belongs in.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        region:\n          type: string\n          x-dcl-go-name: Region\n          description: '[Output Only] URL of the region where the regional forwarding\n            rule resides. This field is not applicable to global forwarding rules.\n            You must specify this field as part of the HTTP request URL. It is not\n            settable as a field in the request body.'\n          x-kubernetes-immutable: true\n        selfLink:\n          type: string\n          x-dcl-go-name: SelfLink\n          readOnly: true\n          description: '[Output Only] Server-defined URL for the resource.'\n          x-kubernetes-immutable: true\n        serviceDirectoryRegistrations:\n          type: array\n          x-dcl-go-name: ServiceDirectoryRegistrations\n          description: Service Directory resources to register this forwarding rule\n            with. Currently, only supports a single Service Directory resource.\n          x-kubernetes-immutable: true\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: ForwardingRuleServiceDirectoryRegistrations\n            properties:\n              namespace:\n                type: string\n                x-dcl-go-name: Namespace\n                description: Service Directory namespace to register the forwarding\n                  rule under.\n                x-kubernetes-immutable: true\n              service:\n                type: string\n                x-dcl-go-name: Service\n                description: Service Directory service to register the forwarding\n                  rule under.\n                x-kubernetes-immutable: true\n        serviceLabel:\n          type: string\n          x-dcl-go-name: ServiceLabel\n          description: An optional prefix to the service name for this Forwarding\n            Rule. If specified, the prefix is the first label of the fully qualified\n            service name. The label must be 1-63 characters long, and comply with\n            [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Specifically, the label\n            must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`\n            which means the first character must be a lowercase letter, and all following\n            characters must be a dash, lowercase letter, or digit, except the last\n            character, which cannot be a dash. This field is only used for internal\n            load balancing.\n          x-kubernetes-immutable: true\n        serviceName:\n          type: string\n          x-dcl-go-name: ServiceName\n          readOnly: true\n          description: '[Output Only] The internal fully qualified service name for\n            this Forwarding Rule. This field is only used for internal load balancing.'\n          x-kubernetes-immutable: true\n        subnetwork:\n          type: string\n          x-dcl-go-name: Subnetwork\n          description: This field is only used for `INTERNAL` load balancing. For\n            internal load balancing, this field identifies the subnetwork that the\n            load balanced IP should belong to for this Forwarding Rule. If the network\n            specified is in auto subnet mode, this field is optional. However, if\n            the network is in custom subnet mode, a subnetwork must be specified.\n          x-kubernetes-immutable: true\n          x-dcl-server-default: true\n        target:\n          type: string\n          x-dcl-go-name: Target\n          description: The URL of the target resource to receive the matched traffic.\n            For regional forwarding rules, this target must live in the same region\n            as the forwarding rule. For global forwarding rules, this target must\n            be a global load balancing resource. The forwarded traffic must be of\n            a type appropriate to the target object. For `INTERNAL_SELF_MANAGED` load\n            balancing, only `targetHttpProxy` is valid, not `targetHttpsProxy`.\n")

// 21483 bytes
// MD5: 1c98dc9b2143e3fcb5ead8d127e10efd
