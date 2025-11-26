// Copyright 2025- The sacloud/iam-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package organization

import (
	"context"

	iam "github.com/sacloud/iam-api-go"
	v1 "github.com/sacloud/iam-api-go/apis/v1"
)

type OrganizationAPI interface {
	Get(ctx context.Context) (*v1.Organization, error)
	Put(ctx context.Context, name string) (*v1.Organization, error)

	GetServicePolicy(ctx context.Context, params GetServicePolicyParams) ([]v1.RuleResponse, error)
	PutServicePolicy(ctx context.Context, rules []v1.Rule) ([]v1.RuleResponse, error)
}

type organizationOp struct {
	client *v1.Client
}

var _ OrganizationAPI = (*organizationOp)(nil)

func NewOrganizationOp(client *v1.Client) OrganizationAPI { return &organizationOp{client} }

func (o *organizationOp) Get(ctx context.Context) (*v1.Organization, error) {
	return iam.ErrorFromDecodedResponse[v1.Organization]("Organization.Get", func() (any, error) {
		return o.client.OrganizationGet(ctx)
	})
}

func (o *organizationOp) Put(ctx context.Context, name string) (*v1.Organization, error) {
	return iam.ErrorFromDecodedResponse[v1.Organization]("Organization.Put", func() (any, error) {
		return o.client.OrganizationPut(ctx, &v1.OrganizationPutReq{Name: name})
	})
}

type GetServicePolicyParams struct {
	IsActive *bool
	IsDryRun *bool
	Name     *string
	Code     *string
	Type     *v1.OrganizationServicePolicyGetType
}

func (o *organizationOp) GetServicePolicy(ctx context.Context, params GetServicePolicyParams) ([]v1.RuleResponse, error) {
	if ret, err := iam.ErrorFromDecodedResponse[v1.OrganizationServicePolicyGetOK]("Organization.GetServicePolicy", func() (any, error) {
		return o.client.OrganizationServicePolicyGet(ctx, v1.OrganizationServicePolicyGetParams{
			IsActive: iam.IntoOpt[v1.OptBool](params.IsActive),
			IsDryRun: iam.IntoOpt[v1.OptBool](params.IsDryRun),
			Name:     iam.IntoOpt[v1.OptString](params.Name),
			Code:     iam.IntoOpt[v1.OptString](params.Code),
			Type:     iam.IntoOpt[v1.OptOrganizationServicePolicyGetType](params.Type),
		})
	}); err != nil {
		return nil, err
	} else {
		return ret.Rules, nil
	}
}

func (o *organizationOp) PutServicePolicy(ctx context.Context, rules []v1.Rule) ([]v1.RuleResponse, error) {
	if ret, err := iam.ErrorFromDecodedResponse[v1.OrganizationServicePolicyPutOK]("Organization.PutServicePolicy", func() (any, error) {
		return o.client.OrganizationServicePolicyPut(ctx, &v1.OrganizationServicePolicyPutReq{Rules: rules})
	}); err != nil {
		return nil, err
	} else {
		return ret.Rules, nil
	}
}
