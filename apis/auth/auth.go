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

package auth

import (
	"context"

	iam "github.com/sacloud/iam-api-go"
	v1 "github.com/sacloud/iam-api-go/apis/v1"
)

type AuthAPI interface {
	GetPasswordPolicy(ctx context.Context) (*v1.PasswordPolicy, error)
	PutPasswordPolicy(ctx context.Context) (*v1.PasswordPolicy, error)

	GetAuthConditions(ctx context.Context) (*v1.AuthConditions, error)
	PutAuthConditions(ctx context.Context, req *v1.AuthConditions) (*v1.AuthConditions, error)
}

type authOp struct {
	client *v1.Client
}

var _ AuthAPI = (*authOp)(nil)

func NewAuthOp(client *v1.Client) AuthAPI { return &authOp{client} }

func (a *authOp) GetPasswordPolicy(ctx context.Context) (*v1.PasswordPolicy, error) {
	return iam.ErrorFromDecodedResponse[v1.PasswordPolicy]("Auth.GetPasswordPolicy", func() (any, error) {
		return a.client.OrganizationPasswordPolicyGet(ctx)
	})
}

func (a *authOp) PutPasswordPolicy(ctx context.Context) (*v1.PasswordPolicy, error) {
	return iam.ErrorFromDecodedResponse[v1.PasswordPolicy]("Auth.PutPasswordPolicy", func() (any, error) {
		return a.client.OrganizationPasswordPolicyPut(ctx)
	})
}

func (a *authOp) GetAuthConditions(ctx context.Context) (*v1.AuthConditions, error) {
	return iam.ErrorFromDecodedResponse[v1.AuthConditions]("Auth.GetAuthConditions", func() (any, error) {
		return a.client.OrganizationAuthConditionsGet(ctx)
	})
}

func (a *authOp) PutAuthConditions(ctx context.Context, req *v1.AuthConditions) (*v1.AuthConditions, error) {
	return iam.ErrorFromDecodedResponse[v1.AuthConditions]("Auth.PutAuthConditions", func() (any, error) {
		return a.client.OrganizationAuthConditionsPut(ctx, req)
	})
}
