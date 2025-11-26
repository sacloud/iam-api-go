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

package iam

import (
	"context"
	"fmt"
	"runtime"

	v1 "github.com/sacloud/iam-api-go/apis/v1"
	"github.com/sacloud/saclient-go"
)

// DefaultAPIRootURL デフォルトのAPIルートURL
const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/api/iam/1.0/"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"iam-api-go/%s (%s/%s; +https://github.com/sacloud/iam-api-go)",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
)

// voidSecuritySource is a placeholder to satisfy the SecuritySource interface.
// saclientにて処理するためここにはロジック不要だが何か渡さないといけないので空の構造体を用意する
type voidSecuritySource struct{}

func (voidSecuritySource) CompatAccessTokenAuth(context.Context, v1.OperationName) (v1.CompatAccessTokenAuth, error) {
	return v1.CompatAccessTokenAuth{}, nil
}

func NewClient(client saclient.ClientAPI) (*v1.Client, error) {
	return v1.NewClient(DefaultAPIRootURL, voidSecuritySource{}, v1.WithClient(client))
}

func NewClientWithAPIRootURL(client saclient.ClientAPI, apiRootURL string) (*v1.Client, error) {
	return v1.NewClient(apiRootURL, voidSecuritySource{}, v1.WithClient(client))
}
