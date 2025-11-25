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

package serviceprincipal_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/sacloud/iam-api-go/apis/serviceprincipal"
	v1 "github.com/sacloud/iam-api-go/apis/v1"
	iam_test "github.com/sacloud/iam-api-go/testutil"
	"github.com/sacloud/packages-go/testutil"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T, v any, s ...int) (*require.Assertions, ServicePrincipalAPI) {
	assert := require.New(t)
	client := iam_test.NewTestClient(v, s...)
	api := NewServicePrincipalOp(client)
	return assert, api
}

func TestList(t *testing.T) {
	var expected v1.ServicePrincipalsGetOK
	expected.SetFake()
	expected.SetItems(make([]v1.ServicePrincipal, 1))
	expected.Items[0].SetFake()
	assert, api := setup(t, &expected)

	actual, err := api.List(t.Context(), ListParams{})
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestList_Fail(t *testing.T) {
	var res v1.Http403Forbidden
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusForbidden)
	res.SetDetail(expected)
	assert, api := setup(t, &res, res.Status)

	actual, err := api.List(t.Context(), ListParams{})
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestCreate(t *testing.T) {
	var expected v1.ServicePrincipal
	expected.SetFake()
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &expected, http.StatusCreated)

	actual, err := api.Create(t.Context(), req)
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestCreate_Fail(t *testing.T) {
	var res v1.Http400BadRequest
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusBadRequest)
	res.SetDetail(expected)
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	actual, err := api.Create(t.Context(), req)
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestGet(t *testing.T) {
	var expected v1.ServicePrincipal
	expected.SetFake()
	assert, api := setup(t, &expected)

	actual, err := api.Get(t.Context(), 123)
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestGet_Fail(t *testing.T) {
	var res v1.Http404NotFound
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusNotFound)
	res.SetDetail(expected)
	assert, api := setup(t, &res, res.Status)

	actual, err := api.Get(t.Context(), 123)
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestUpdate(t *testing.T) {
	var expected v1.ServicePrincipal
	expected.SetFake()
	var req v1.ServicePrincipalsServicePrincipalIDPutReq
	req.SetFake()
	assert, api := setup(t, &expected)

	actual, err := api.Update(t.Context(), 123, req)
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestUpdate_Fail(t *testing.T) {
	var res v1.Http403Forbidden
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusForbidden)
	res.SetDetail(expected)
	var req v1.ServicePrincipalsServicePrincipalIDPutReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	actual, err := api.Update(t.Context(), 123, req)
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestDelete(t *testing.T) {
	assert, api := setup(t, &v1.ServicePrincipalsServicePrincipalIDDeleteNoContent{}, http.StatusNoContent)

	err := api.Delete(t.Context(), 123)
	assert.NoError(err)
}

func TestDelete_Fail(t *testing.T) {
	var res v1.Http404NotFound
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusNotFound)
	res.SetDetail(expected)
	assert, api := setup(t, &res, res.Status)

	err := api.Delete(t.Context(), 123)
	assert.Error(err)
	assert.Contains(err.Error(), expected)
}

func TestListKeys(t *testing.T) {
	var expected v1.ServicePrincipalsServicePrincipalIDKeysGetOK
	expected.SetFake()
	expected.SetItems(make([]v1.ServicePrincipalKey, 1))
	expected.Items[0].SetFake()
	assert, api := setup(t, &expected)

	actual, err := api.ListKeys(t.Context(), 123, ListKeysParams{})
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestListKeys_Fail(t *testing.T) {
	var res v1.Http403Forbidden
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusForbidden)
	res.SetDetail(expected)
	assert, api := setup(t, &res, res.Status)

	actual, err := api.ListKeys(t.Context(), 123, ListKeysParams{})
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestUploadKey(t *testing.T) {
	var expected v1.ServicePrincipalKey
	expected.SetFake()
	assert, api := setup(t, &expected, http.StatusCreated)

	actual, err := api.UploadKey(t.Context(), 123, "test key")
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestUploadKey_Fail(t *testing.T) {
	var res v1.Http400BadRequest
	res.SetFake()
	res.SetStatus(http.StatusBadRequest)
	res.SetDetail("bad request")
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	actual, err := api.Create(t.Context(), req)
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), "bad request")
}

func TestEnableKey(t *testing.T) {
	var expected v1.ServicePrincipalKey
	expected.SetFake()
	assert, api := setup(t, &expected)

	actual, err := api.EnableKey(t.Context(), 123, uuid.New())
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestEnableKey_Fail(t *testing.T) {
	var res v1.Http403Forbidden
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusForbidden)
	res.SetDetail(expected)
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	actual, err := api.EnableKey(t.Context(), 123, uuid.New())
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestDisableKey(t *testing.T) {
	var expected v1.ServicePrincipalKey
	expected.SetFake()
	assert, api := setup(t, &expected)

	actual, err := api.DisableKey(t.Context(), 123, uuid.New())
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestDisableKey_Fail(t *testing.T) {
	var res v1.Http403Forbidden
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusForbidden)
	res.SetDetail(expected)
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	actual, err := api.DisableKey(t.Context(), 123, uuid.New())
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}

func TestDeleteKey(t *testing.T) {
	assert, api := setup(t, nil, http.StatusNoContent)

	err := api.DeleteKey(t.Context(), 123, uuid.New())
	assert.NoError(err)
}

func TestDeleteKey_Fail(t *testing.T) {
	var res v1.Http401Unauthorized
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusUnauthorized)
	res.SetDetail(expected)
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	err := api.DeleteKey(t.Context(), 123, uuid.New())
	assert.Error(err)
	assert.Contains(err.Error(), expected)
}

func TestIssueToken(t *testing.T) {
	var expected v1.ServicePrincipalOAuth2AccessToken
	expected.SetFake()
	expected.SetTokenExpiredAt(time.UnixMicro(0).UTC())
	assert, api := setup(t, &expected)

	actual, err := api.IssueToken(t.Context(), "test test")
	assert.NoError(err)
	assert.NotNil(actual)
	assert.Equal(&expected, actual)
}

func TestIssueToken_Fail(t *testing.T) {
	var res v1.Http400BadRequest
	expected := testutil.Random(64, testutil.CharSetAlphaNum)
	res.SetFake()
	res.SetStatus(http.StatusBadRequest)
	res.SetDetail(expected)
	var req v1.ServicePrincipalsPostReq
	req.SetFake()
	assert, api := setup(t, &res, res.Status)

	actual, err := api.IssueToken(t.Context(), "test test")
	assert.Error(err)
	assert.Nil(actual)
	assert.Contains(err.Error(), expected)
}
