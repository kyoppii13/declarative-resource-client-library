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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/gkehub/alpha/gkehub_alpha_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/alpha"
)

// FeatureMembershipServer implements the gRPC interface for FeatureMembership.
type FeatureMembershipServer struct{}

// ProtoToFeatureMembershipMeshControlPlaneEnum converts a FeatureMembershipMeshControlPlaneEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipMeshControlPlaneEnum(e alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum) *alpha.FeatureMembershipMeshControlPlaneEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipMeshControlPlaneEnum(n[len("GkehubAlphaFeatureMembershipMeshControlPlaneEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum converts a FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum enum from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(e alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum) *alpha.FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum {
	if e == 0 {
		return nil
	}
	if n, ok := alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum_name[int32(e)]; ok {
		e := alpha.FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(n[len("GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum"):])
		return &e
	}
	return nil
}

// ProtoToFeatureMembershipMesh converts a FeatureMembershipMesh object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipMesh(p *alphapb.GkehubAlphaFeatureMembershipMesh) *alpha.FeatureMembershipMesh {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipMesh{
		ControlPlane: ProtoToGkehubAlphaFeatureMembershipMeshControlPlaneEnum(p.GetControlPlane()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagement converts a FeatureMembershipConfigmanagement object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagement(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagement) *alpha.FeatureMembershipConfigmanagement {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagement{
		ConfigSync:          ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSync(p.GetConfigSync()),
		PolicyController:    ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyController(p.GetPolicyController()),
		Binauthz:            ProtoToGkehubAlphaFeatureMembershipConfigmanagementBinauthz(p.GetBinauthz()),
		HierarchyController: ProtoToGkehubAlphaFeatureMembershipConfigmanagementHierarchyController(p.GetHierarchyController()),
		Version:             dcl.StringOrNil(p.GetVersion()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSync converts a FeatureMembershipConfigmanagementConfigSync object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSync(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSync) *alpha.FeatureMembershipConfigmanagementConfigSync {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementConfigSync{
		Git:          ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit(p.GetGit()),
		SourceFormat: dcl.StringOrNil(p.GetSourceFormat()),
		PreventDrift: dcl.Bool(p.GetPreventDrift()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementConfigSyncGit converts a FeatureMembershipConfigmanagementConfigSyncGit object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit) *alpha.FeatureMembershipConfigmanagementConfigSyncGit {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementConfigSyncGit{
		SyncRepo:               dcl.StringOrNil(p.GetSyncRepo()),
		SyncBranch:             dcl.StringOrNil(p.GetSyncBranch()),
		PolicyDir:              dcl.StringOrNil(p.GetPolicyDir()),
		SyncWaitSecs:           dcl.StringOrNil(p.GetSyncWaitSecs()),
		SyncRev:                dcl.StringOrNil(p.GetSyncRev()),
		SecretType:             dcl.StringOrNil(p.GetSecretType()),
		HttpsProxy:             dcl.StringOrNil(p.GetHttpsProxy()),
		GcpServiceAccountEmail: dcl.StringOrNil(p.GetGcpServiceAccountEmail()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyController converts a FeatureMembershipConfigmanagementPolicyController object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyController(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyController) *alpha.FeatureMembershipConfigmanagementPolicyController {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementPolicyController{
		Enabled:                    dcl.Bool(p.GetEnabled()),
		ReferentialRulesEnabled:    dcl.Bool(p.GetReferentialRulesEnabled()),
		LogDeniesEnabled:           dcl.Bool(p.GetLogDeniesEnabled()),
		MutationEnabled:            dcl.Bool(p.GetMutationEnabled()),
		PolicyControllerMonitoring: ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring(p.GetPolicyControllerMonitoring()),
		TemplateLibraryInstalled:   dcl.Bool(p.GetTemplateLibraryInstalled()),
		AuditIntervalSeconds:       dcl.StringOrNil(p.GetAuditIntervalSeconds()),
	}
	for _, r := range p.GetExemptableNamespaces() {
		obj.ExemptableNamespaces = append(obj.ExemptableNamespaces, r)
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring converts a FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring) *alpha.FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring{}
	for _, r := range p.GetBackends() {
		obj.Backends = append(obj.Backends, *ProtoToGkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(r))
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementBinauthz converts a FeatureMembershipConfigmanagementBinauthz object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementBinauthz(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementBinauthz) *alpha.FeatureMembershipConfigmanagementBinauthz {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.Bool(p.GetEnabled()),
	}
	return obj
}

// ProtoToFeatureMembershipConfigmanagementHierarchyController converts a FeatureMembershipConfigmanagementHierarchyController object from its proto representation.
func ProtoToGkehubAlphaFeatureMembershipConfigmanagementHierarchyController(p *alphapb.GkehubAlphaFeatureMembershipConfigmanagementHierarchyController) *alpha.FeatureMembershipConfigmanagementHierarchyController {
	if p == nil {
		return nil
	}
	obj := &alpha.FeatureMembershipConfigmanagementHierarchyController{
		Enabled:                         dcl.Bool(p.GetEnabled()),
		EnablePodTreeLabels:             dcl.Bool(p.GetEnablePodTreeLabels()),
		EnableHierarchicalResourceQuota: dcl.Bool(p.GetEnableHierarchicalResourceQuota()),
	}
	return obj
}

// ProtoToFeatureMembership converts a FeatureMembership resource from its proto representation.
func ProtoToFeatureMembership(p *alphapb.GkehubAlphaFeatureMembership) *alpha.FeatureMembership {
	obj := &alpha.FeatureMembership{
		Mesh:             ProtoToGkehubAlphaFeatureMembershipMesh(p.GetMesh()),
		Configmanagement: ProtoToGkehubAlphaFeatureMembershipConfigmanagement(p.GetConfigmanagement()),
		Project:          dcl.StringOrNil(p.GetProject()),
		Location:         dcl.StringOrNil(p.GetLocation()),
		Feature:          dcl.StringOrNil(p.GetFeature()),
		Membership:       dcl.StringOrNil(p.GetMembership()),
	}
	return obj
}

// FeatureMembershipMeshControlPlaneEnumToProto converts a FeatureMembershipMeshControlPlaneEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipMeshControlPlaneEnumToProto(e *alpha.FeatureMembershipMeshControlPlaneEnum) alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum_value["FeatureMembershipMeshControlPlaneEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipMeshControlPlaneEnum(0)
}

// FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnumToProto converts a FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum enum to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnumToProto(e *alpha.FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum) alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum {
	if e == nil {
		return alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(0)
	}
	if v, ok := alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum_value["FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum"+string(*e)]; ok {
		return alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(v)
	}
	return alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(0)
}

// FeatureMembershipMeshToProto converts a FeatureMembershipMesh object to its proto representation.
func GkehubAlphaFeatureMembershipMeshToProto(o *alpha.FeatureMembershipMesh) *alphapb.GkehubAlphaFeatureMembershipMesh {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipMesh{}
	p.SetControlPlane(GkehubAlphaFeatureMembershipMeshControlPlaneEnumToProto(o.ControlPlane))
	return p
}

// FeatureMembershipConfigmanagementToProto converts a FeatureMembershipConfigmanagement object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementToProto(o *alpha.FeatureMembershipConfigmanagement) *alphapb.GkehubAlphaFeatureMembershipConfigmanagement {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagement{}
	p.SetConfigSync(GkehubAlphaFeatureMembershipConfigmanagementConfigSyncToProto(o.ConfigSync))
	p.SetPolicyController(GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerToProto(o.PolicyController))
	p.SetBinauthz(GkehubAlphaFeatureMembershipConfigmanagementBinauthzToProto(o.Binauthz))
	p.SetHierarchyController(GkehubAlphaFeatureMembershipConfigmanagementHierarchyControllerToProto(o.HierarchyController))
	p.SetVersion(dcl.ValueOrEmptyString(o.Version))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncToProto converts a FeatureMembershipConfigmanagementConfigSync object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementConfigSyncToProto(o *alpha.FeatureMembershipConfigmanagementConfigSync) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSync {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSync{}
	p.SetGit(GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGitToProto(o.Git))
	p.SetSourceFormat(dcl.ValueOrEmptyString(o.SourceFormat))
	p.SetPreventDrift(dcl.ValueOrEmptyBool(o.PreventDrift))
	return p
}

// FeatureMembershipConfigmanagementConfigSyncGitToProto converts a FeatureMembershipConfigmanagementConfigSyncGit object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGitToProto(o *alpha.FeatureMembershipConfigmanagementConfigSyncGit) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementConfigSyncGit{}
	p.SetSyncRepo(dcl.ValueOrEmptyString(o.SyncRepo))
	p.SetSyncBranch(dcl.ValueOrEmptyString(o.SyncBranch))
	p.SetPolicyDir(dcl.ValueOrEmptyString(o.PolicyDir))
	p.SetSyncWaitSecs(dcl.ValueOrEmptyString(o.SyncWaitSecs))
	p.SetSyncRev(dcl.ValueOrEmptyString(o.SyncRev))
	p.SetSecretType(dcl.ValueOrEmptyString(o.SecretType))
	p.SetHttpsProxy(dcl.ValueOrEmptyString(o.HttpsProxy))
	p.SetGcpServiceAccountEmail(dcl.ValueOrEmptyString(o.GcpServiceAccountEmail))
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerToProto converts a FeatureMembershipConfigmanagementPolicyController object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerToProto(o *alpha.FeatureMembershipConfigmanagementPolicyController) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetReferentialRulesEnabled(dcl.ValueOrEmptyBool(o.ReferentialRulesEnabled))
	p.SetLogDeniesEnabled(dcl.ValueOrEmptyBool(o.LogDeniesEnabled))
	p.SetMutationEnabled(dcl.ValueOrEmptyBool(o.MutationEnabled))
	p.SetPolicyControllerMonitoring(GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringToProto(o.PolicyControllerMonitoring))
	p.SetTemplateLibraryInstalled(dcl.ValueOrEmptyBool(o.TemplateLibraryInstalled))
	p.SetAuditIntervalSeconds(dcl.ValueOrEmptyString(o.AuditIntervalSeconds))
	sExemptableNamespaces := make([]string, len(o.ExemptableNamespaces))
	for i, r := range o.ExemptableNamespaces {
		sExemptableNamespaces[i] = r
	}
	p.SetExemptableNamespaces(sExemptableNamespaces)
	return p
}

// FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringToProto converts a FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringToProto(o *alpha.FeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoring{}
	sBackends := make([]alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum, len(o.Backends))
	for i, r := range o.Backends {
		sBackends[i] = alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum(alphapb.GkehubAlphaFeatureMembershipConfigmanagementPolicyControllerPolicyControllerMonitoringBackendsEnum_value[string(r)])
	}
	p.SetBackends(sBackends)
	return p
}

// FeatureMembershipConfigmanagementBinauthzToProto converts a FeatureMembershipConfigmanagementBinauthz object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementBinauthzToProto(o *alpha.FeatureMembershipConfigmanagementBinauthz) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementBinauthz {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementBinauthz{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	return p
}

// FeatureMembershipConfigmanagementHierarchyControllerToProto converts a FeatureMembershipConfigmanagementHierarchyController object to its proto representation.
func GkehubAlphaFeatureMembershipConfigmanagementHierarchyControllerToProto(o *alpha.FeatureMembershipConfigmanagementHierarchyController) *alphapb.GkehubAlphaFeatureMembershipConfigmanagementHierarchyController {
	if o == nil {
		return nil
	}
	p := &alphapb.GkehubAlphaFeatureMembershipConfigmanagementHierarchyController{}
	p.SetEnabled(dcl.ValueOrEmptyBool(o.Enabled))
	p.SetEnablePodTreeLabels(dcl.ValueOrEmptyBool(o.EnablePodTreeLabels))
	p.SetEnableHierarchicalResourceQuota(dcl.ValueOrEmptyBool(o.EnableHierarchicalResourceQuota))
	return p
}

// FeatureMembershipToProto converts a FeatureMembership resource to its proto representation.
func FeatureMembershipToProto(resource *alpha.FeatureMembership) *alphapb.GkehubAlphaFeatureMembership {
	p := &alphapb.GkehubAlphaFeatureMembership{}
	p.SetMesh(GkehubAlphaFeatureMembershipMeshToProto(resource.Mesh))
	p.SetConfigmanagement(GkehubAlphaFeatureMembershipConfigmanagementToProto(resource.Configmanagement))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetFeature(dcl.ValueOrEmptyString(resource.Feature))
	p.SetMembership(dcl.ValueOrEmptyString(resource.Membership))

	return p
}

// applyFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) applyFeatureMembership(ctx context.Context, c *alpha.Client, request *alphapb.ApplyGkehubAlphaFeatureMembershipRequest) (*alphapb.GkehubAlphaFeatureMembership, error) {
	p := ProtoToFeatureMembership(request.GetResource())
	res, err := c.ApplyFeatureMembership(ctx, p)
	if err != nil {
		return nil, err
	}
	r := FeatureMembershipToProto(res)
	return r, nil
}

// applyGkehubAlphaFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Apply() method.
func (s *FeatureMembershipServer) ApplyGkehubAlphaFeatureMembership(ctx context.Context, request *alphapb.ApplyGkehubAlphaFeatureMembershipRequest) (*alphapb.GkehubAlphaFeatureMembership, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyFeatureMembership(ctx, cl, request)
}

// DeleteFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembership Delete() method.
func (s *FeatureMembershipServer) DeleteGkehubAlphaFeatureMembership(ctx context.Context, request *alphapb.DeleteGkehubAlphaFeatureMembershipRequest) (*emptypb.Empty, error) {

	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteFeatureMembership(ctx, ProtoToFeatureMembership(request.GetResource()))

}

// ListGkehubAlphaFeatureMembership handles the gRPC request by passing it to the underlying FeatureMembershipList() method.
func (s *FeatureMembershipServer) ListGkehubAlphaFeatureMembership(ctx context.Context, request *alphapb.ListGkehubAlphaFeatureMembershipRequest) (*alphapb.ListGkehubAlphaFeatureMembershipResponse, error) {
	cl, err := createConfigFeatureMembership(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListFeatureMembership(ctx, request.GetProject(), request.GetLocation(), request.GetFeature())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.GkehubAlphaFeatureMembership
	for _, r := range resources.Items {
		rp := FeatureMembershipToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListGkehubAlphaFeatureMembershipResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigFeatureMembership(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
