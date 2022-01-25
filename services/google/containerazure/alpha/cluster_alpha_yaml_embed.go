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
// gen_go_data -package alpha -var YAML_cluster blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containerazure/alpha/cluster.yaml

package alpha

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containerazure/alpha/cluster.yaml
var YAML_cluster = []byte("info:\n  title: ContainerAzure/Cluster\n  description: An Anthos cluster running on Azure.\n  x-dcl-struct-name: Cluster\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: API reference\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest/v1/projects.locations.azureClusters\n  x-dcl-guides:\n  - text: Multicloud overview\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud\npaths:\n  get:\n    description: The function used to get information about a Cluster\n    parameters:\n    - name: Cluster\n      required: true\n      description: A full instance of a Cluster\n  apply:\n    description: The function used to apply information about a Cluster\n    parameters:\n    - name: Cluster\n      required: true\n      description: A full instance of a Cluster\n  delete:\n    description: The function used to delete a Cluster\n    parameters:\n    - name: Cluster\n      required: true\n      description: A full instance of a Cluster\n  deleteAll:\n    description: The function used to delete all Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Cluster\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Cluster:\n      title: Cluster\n      x-dcl-id: projects/{{project}}/locations/{{location}}/azureClusters/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - azureRegion\n      - resourceGroupId\n      - client\n      - networking\n      - controlPlane\n      - authorization\n      - project\n      - location\n      - fleet\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: 'Optional. Annotations on the cluster. This field has the same\n            restrictions as Kubernetes annotations. The total size of all keys and\n            values combined is limited to 256k. Keys can have 2 segments: prefix (optional)\n            and name (required), separated by a slash (/). Prefix must be a DNS subdomain.\n            Name must be 63 characters or less, begin and end with alphanumerics,\n            with dashes (-), underscores (_), dots (.), and alphanumerics between.'\n          x-kubernetes-immutable: true\n        authorization:\n          type: object\n          x-dcl-go-name: Authorization\n          x-dcl-go-type: ClusterAuthorization\n          description: Required. Configuration related to the cluster RBAC settings.\n          x-kubernetes-immutable: true\n          required:\n          - adminUsers\n          properties:\n            adminUsers:\n              type: array\n              x-dcl-go-name: AdminUsers\n              description: Required. Users that can perform operations as a cluster\n                admin. A new ClusterRoleBinding will be created to grant the cluster-admin\n                ClusterRole to the users. At most one user can be specified. For more\n                info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ClusterAuthorizationAdminUsers\n                required:\n                - username\n                properties:\n                  username:\n                    type: string\n                    x-dcl-go-name: Username\n                    description: Required. The name of the user, e.g. `my-gcp-id@gmail.com`.\n                    x-kubernetes-immutable: true\n        azureRegion:\n          type: string\n          x-dcl-go-name: AzureRegion\n          description: Required. The Azure region where the cluster runs. Each Google\n            Cloud region supports a subset of nearby Azure regions. You can call to\n            list all supported Azure regions within a given Google Cloud region.\n          x-kubernetes-immutable: true\n        client:\n          type: string\n          x-dcl-go-name: Client\n          description: Required. Name of the AzureClient. The `AzureClient` resource\n            must reside on the same GCP project and region as the `AzureCluster`.\n            `AzureClient` names are formatted as `projects/<project-number>/locations/<region>/azureClients/<client-id>`.\n            See Resource Names (https:cloud.google.com/apis/design/resource_names)\n            for more details on Google Cloud resource names.\n          x-dcl-references:\n          - resource: ContainerAzure/AzureClient\n            field: name\n        controlPlane:\n          type: object\n          x-dcl-go-name: ControlPlane\n          x-dcl-go-type: ClusterControlPlane\n          description: Required. Configuration related to the cluster control plane.\n          required:\n          - version\n          - subnetId\n          - sshConfig\n          properties:\n            databaseEncryption:\n              type: object\n              x-dcl-go-name: DatabaseEncryption\n              x-dcl-go-type: ClusterControlPlaneDatabaseEncryption\n              description: Optional. Configuration related to application-layer secrets\n                encryption.\n              x-kubernetes-immutable: true\n              required:\n              - keyId\n              properties:\n                keyId:\n                  type: string\n                  x-dcl-go-name: KeyId\n                  description: 'The ARM ID of the Azure Key Vault key to encrypt /\n                    decrypt data. For example: `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name>`\n                    Encryption will always take the latest version of the key and\n                    hence specific version is not supported.'\n                  x-kubernetes-immutable: true\n            mainVolume:\n              type: object\n              x-dcl-go-name: MainVolume\n              x-dcl-go-type: ClusterControlPlaneMainVolume\n              description: Optional. Configuration related to the main volume provisioned\n                for each control plane replica. The main volume is in charge of storing\n                all of the cluster's etcd state. When unspecified, it defaults to\n                a 8-GiB Azure Disk.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              properties:\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the disk, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n            proxyConfig:\n              type: object\n              x-dcl-go-name: ProxyConfig\n              x-dcl-go-type: ClusterControlPlaneProxyConfig\n              description: Proxy configuration for outbound HTTP(S) traffic.\n              x-kubernetes-immutable: true\n              required:\n              - resourceGroupId\n              - secretId\n              properties:\n                resourceGroupId:\n                  type: string\n                  x-dcl-go-name: ResourceGroupId\n                  description: The ARM ID the of the resource group containing proxy\n                    keyvault. Resource group ids are formatted as `/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>`\n                  x-kubernetes-immutable: true\n                secretId:\n                  type: string\n                  x-dcl-go-name: SecretId\n                  description: The URL the of the proxy setting secret with its version.\n                    Secret ids are formatted as `https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>`.\n                  x-kubernetes-immutable: true\n            replicaPlacements:\n              type: array\n              x-dcl-go-name: ReplicaPlacements\n              description: Configuration for where to place the control plane replicas.\n                Up to three replica placement instances can be specified. If replica_placements\n                is set, the replica placement instances will be applied to the three\n                control plane replicas as evenly as possible.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: ClusterControlPlaneReplicaPlacements\n                required:\n                - subnetId\n                - azureAvailabilityZone\n                properties:\n                  azureAvailabilityZone:\n                    type: string\n                    x-dcl-go-name: AzureAvailabilityZone\n                    description: For a given replica, the Azure availability zone\n                      where to provision the control plane VM and the ETCD disk.\n                    x-kubernetes-immutable: true\n                  subnetId:\n                    type: string\n                    x-dcl-go-name: SubnetId\n                    description: For a given replica, the ARM ID of the subnet where\n                      the control plane VM is deployed. Make sure it's a subnet under\n                      the virtual network in the cluster configuration.\n                    x-kubernetes-immutable: true\n            rootVolume:\n              type: object\n              x-dcl-go-name: RootVolume\n              x-dcl-go-type: ClusterControlPlaneRootVolume\n              description: Optional. Configuration related to the root volume provisioned\n                for each control plane replica. When unspecified, it defaults to 32-GiB\n                Azure Disk.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              properties:\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the disk, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n            sshConfig:\n              type: object\n              x-dcl-go-name: SshConfig\n              x-dcl-go-type: ClusterControlPlaneSshConfig\n              description: Required. SSH configuration for how to access the underlying\n                control plane machines.\n              x-kubernetes-immutable: true\n              required:\n              - authorizedKey\n              properties:\n                authorizedKey:\n                  type: string\n                  x-dcl-go-name: AuthorizedKey\n                  description: Required. The SSH public key data for VMs managed by\n                    Anthos. This accepts the authorized_keys file format used in OpenSSH\n                    according to the sshd(8) manual page.\n                  x-kubernetes-immutable: true\n            subnetId:\n              type: string\n              x-dcl-go-name: SubnetId\n              description: 'Required. The ARM ID of the subnet where the control plane\n                VMs are deployed. Example: `/subscriptions//resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/default`.'\n              x-kubernetes-immutable: true\n            tags:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Tags\n              description: Optional. A set of tags to apply to all underlying control\n                plane Azure resources.\n              x-kubernetes-immutable: true\n            version:\n              type: string\n              x-dcl-go-name: Version\n              description: Required. The Kubernetes version to run on control plane\n                replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions\n                on a given Google Cloud region by calling GetAzureServerConfig.\n            vmSize:\n              type: string\n              x-dcl-go-name: VmSize\n              description: 'Optional. The Azure VM size name. Example: `Standard_DS2_v2`.\n                For available VM sizes, see https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions.\n                When unspecified, it defaults to `Standard_DS2_v2`.'\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this cluster was created.\n          x-kubernetes-immutable: true\n        description:\n          type: string\n          x-dcl-go-name: Description\n          description: Optional. A human readable description of this cluster. Cannot\n            be longer than 255 UTF-8 encoded bytes.\n        endpoint:\n          type: string\n          x-dcl-go-name: Endpoint\n          readOnly: true\n          description: Output only. The endpoint of the cluster's API server.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Allows clients to perform consistent read-modify-writes through\n            optimistic concurrency control. May be sent on update and delete requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        fleet:\n          type: object\n          x-dcl-go-name: Fleet\n          x-dcl-go-type: ClusterFleet\n          description: Fleet configuration.\n          x-kubernetes-immutable: true\n          required:\n          - project\n          properties:\n            membership:\n              type: string\n              x-dcl-go-name: Membership\n              readOnly: true\n              description: The name of the managed Hub Membership resource associated\n                to this cluster. Membership names are formatted as projects/<project-number>/locations/global/membership/<cluster-id>.\n              x-kubernetes-immutable: true\n            project:\n              type: string\n              x-dcl-go-name: Project\n              description: The number of the Fleet host project where this cluster\n                will be registered.\n              x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of this resource.\n          x-kubernetes-immutable: true\n        networking:\n          type: object\n          x-dcl-go-name: Networking\n          x-dcl-go-type: ClusterNetworking\n          description: Required. Cluster-wide networking configuration.\n          x-kubernetes-immutable: true\n          required:\n          - virtualNetworkId\n          - podAddressCidrBlocks\n          - serviceAddressCidrBlocks\n          properties:\n            podAddressCidrBlocks:\n              type: array\n              x-dcl-go-name: PodAddressCidrBlocks\n              description: Required. The IP address range of the pods in this cluster,\n                in CIDR notation (e.g. `10.96.0.0/14`). All pods in the cluster get\n                assigned a unique RFC1918 IPv4 address from these ranges. Only a single\n                range is supported. This field cannot be changed after creation.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            serviceAddressCidrBlocks:\n              type: array\n              x-dcl-go-name: ServiceAddressCidrBlocks\n              description: Required. The IP address range for services in this cluster,\n                in CIDR notation (e.g. `10.96.0.0/14`). All services in the cluster\n                get assigned a unique RFC1918 IPv4 address from these ranges. Only\n                a single range is supported. This field cannot be changed after creating\n                a cluster.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            virtualNetworkId:\n              type: string\n              x-dcl-go-name: VirtualNetworkId\n              description: 'Required. The Azure Resource Manager (ARM) ID of the VNet\n                associated with your cluster. All components in the cluster (i.e.\n                control plane and node pools) run on a single VNet. Example: `/subscriptions/*/resourceGroups/*/providers/Microsoft.Network/virtualNetworks/*`\n                This field cannot be changed after creation.'\n              x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        reconciling:\n          type: boolean\n          x-dcl-go-name: Reconciling\n          readOnly: true\n          description: Output only. If set, there are currently changes in flight\n            to the cluster.\n          x-kubernetes-immutable: true\n        resourceGroupId:\n          type: string\n          x-dcl-go-name: ResourceGroupId\n          description: 'Required. The ARM ID of the resource group where the cluster\n            resources are deployed. For example: `/subscriptions/*/resourceGroups/*`'\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: ClusterStateEnum\n          readOnly: true\n          description: 'Output only. The current state of the cluster. Possible values:\n            STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,\n            DEGRADED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PROVISIONING\n          - RUNNING\n          - RECONCILING\n          - STOPPING\n          - ERROR\n          - DEGRADED\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. A globally unique identifier for the cluster.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time at which this cluster was last updated.\n          x-kubernetes-immutable: true\n        workloadIdentityConfig:\n          type: object\n          x-dcl-go-name: WorkloadIdentityConfig\n          x-dcl-go-type: ClusterWorkloadIdentityConfig\n          readOnly: true\n          description: Output only. Workload Identity settings.\n          x-kubernetes-immutable: true\n          properties:\n            identityProvider:\n              type: string\n              x-dcl-go-name: IdentityProvider\n              description: The ID of the OIDC Identity Provider (IdP) associated to\n                the Workload Identity Pool.\n              x-kubernetes-immutable: true\n            issuerUri:\n              type: string\n              x-dcl-go-name: IssuerUri\n              description: The OIDC issuer URL for this cluster.\n              x-kubernetes-immutable: true\n            workloadPool:\n              type: string\n              x-dcl-go-name: WorkloadPool\n              description: The Workload Identity Pool associated to the cluster.\n              x-kubernetes-immutable: true\n")

// 20337 bytes
// MD5: ad25f4df175be0272ffae17fb1960feb
