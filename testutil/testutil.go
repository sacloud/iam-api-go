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

package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/sacloud/iam-api-go"
	v1 "github.com/sacloud/iam-api-go/apis/v1"
	"github.com/sacloud/saclient-go"
)

var theCLient saclient.Client

func NewTestClient(v any, s ...int) *v1.Client {
	s = append(s, http.StatusOK)
	j, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		st := s[0]

		w.WriteHeader(st)
		if st == http.StatusNoContent {
			return
		}
		if _, e = w.Write(j); e != nil {
			panic(e)
		}
	})
	sv := httptest.NewServer(h)
	sa, e := theCLient.DupWith(saclient.WithTestServer(sv), saclient.WithTraceMode("all"))
	if e != nil {
		panic(e)
	}
	c, e := iam.NewClientWithAPIRootURL(sa, sv.URL)
	if e != nil {
		panic(e)
	}
	if e := sa.Populate(); e != nil {
		panic(e)
	}
	return c
}
