// Copyright 2025- The sacloud/iam-api-go authors
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

package user2fa

import (
	"context"

	iam "github.com/sacloud/iam-api-go"
	v1 "github.com/sacloud/iam-api-go/apis/v1"
)

type User2FAAPI interface {
	ActivateOTP(ctx context.Context) (*v1.CompatUsersUserIDActivateOtpPostOK, error)
	DeactivateOTP(ctx context.Context) error

	CreateRecoveryCode(ctx context.Context) (string, error)
	ListTrustedDevices(ctx context.Context) (*v1.CompatUsersUserIDTrustedDevicesGetOK, error)
	DeleteTrustedDevice(ctx context.Context, trustedDeviceID int) error
	ClearTrustedDevices(ctx context.Context) error
	StartSecurityKeyRegistration(ctx context.Context) (string, error)
	ValidateSecurityKeyRegistration(ctx context.Context, credential string) error

	ListSecurityKeys(ctx context.Context) (*v1.CompatUsersUserIDSecurityKeysGetOK, error)
	GetSecurityKey(ctx context.Context, securityKeyID int) (*v1.UserSecurityKey, error)
	UpdateSecurityKey(ctx context.Context, securityKeyID int, name string) (*v1.UserSecurityKey, error)
	DeleteSecurityKey(ctx context.Context, securityKeyID int) error
}

type user2faOp struct {
	client *v1.Client
	user   *v1.User
}

func NewUser2FAOp(client *v1.Client, user *v1.User) User2FAAPI {
	return &user2faOp{
		client: client,
		user:   user,
	}
}

func (u *user2faOp) getUserID() int { return u.user.GetID() }

func (u *user2faOp) ActivateOTP(ctx context.Context) (*v1.CompatUsersUserIDActivateOtpPostOK, error) {
	return iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDActivateOtpPostOK]("User2FA.ActivateOTP", func() (any, error) {
		return u.client.CompatUsersUserIDActivateOtpPost(ctx, v1.CompatUsersUserIDActivateOtpPostParams{UserID: u.getUserID()})
	})
}

func (u *user2faOp) DeactivateOTP(ctx context.Context) error {
	_, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDDeactivateOtpPostNoContent]("User2FA.DeactivateOTP", func() (any, error) {
		return u.client.CompatUsersUserIDDeactivateOtpPost(ctx, v1.CompatUsersUserIDDeactivateOtpPostParams{UserID: u.getUserID()})
	})
	return err
}

func (u *user2faOp) CreateRecoveryCode(ctx context.Context) (string, error) {
	if res, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDRecoveryCodePostOK]("User2FA.CreateRecoveryCode", func() (any, error) {
		return u.client.CompatUsersUserIDRecoveryCodePost(ctx, v1.CompatUsersUserIDRecoveryCodePostParams{UserID: u.getUserID()})
	}); err != nil {
		return "", err
	} else {
		return res.GetCode(), nil
	}
}

func (u *user2faOp) ListTrustedDevices(ctx context.Context) (*v1.CompatUsersUserIDTrustedDevicesGetOK, error) {
	return iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDTrustedDevicesGetOK]("User2FA.ListTrustedDevices", func() (any, error) {
		return u.client.CompatUsersUserIDTrustedDevicesGet(ctx, v1.CompatUsersUserIDTrustedDevicesGetParams{UserID: u.getUserID()})
	})
}

func (u *user2faOp) DeleteTrustedDevice(ctx context.Context, trustedDeviceID int) error {
	_, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDTrustedDevicesTrustedDeviceIDDeleteNoContent]("User2FA.DeleteTrustedDevice", func() (any, error) {
		return u.client.CompatUsersUserIDTrustedDevicesTrustedDeviceIDDelete(ctx, v1.CompatUsersUserIDTrustedDevicesTrustedDeviceIDDeleteParams{
			UserID:          u.getUserID(),
			TrustedDeviceID: trustedDeviceID,
		})
	})
	return err
}

func (u *user2faOp) ClearTrustedDevices(ctx context.Context) error {
	_, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDClearTrustedDevicesPostNoContent]("User2FA.ClearTrustedDevices", func() (any, error) {
		return u.client.CompatUsersUserIDClearTrustedDevicesPost(ctx, v1.CompatUsersUserIDClearTrustedDevicesPostParams{UserID: u.getUserID()})
	})
	return err
}

func (u *user2faOp) StartSecurityKeyRegistration(ctx context.Context) (string, error) {
	if res, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDStartSecurityKeyRegistrationPostOK]("User2FA.StartSecurityKeyRegistration", func() (any, error) {
		return u.client.CompatUsersUserIDStartSecurityKeyRegistrationPost(ctx, v1.CompatUsersUserIDStartSecurityKeyRegistrationPostParams{UserID: u.getUserID()})
	}); err != nil {
		return "", err
	} else {
		return res.GetPublicKeyCredentialCreationOptions(), nil
	}
}

func (u *user2faOp) ValidateSecurityKeyRegistration(ctx context.Context, credential string) error {
	_, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDValidateSecurityKeyRegistrationPostNoContent]("User2FA.ValidateSecurityKeyRegistration", func() (any, error) {
		req := v1.CompatUsersUserIDValidateSecurityKeyRegistrationPostReq{Credential: credential}
		params := v1.CompatUsersUserIDValidateSecurityKeyRegistrationPostParams{UserID: u.getUserID()}
		return u.client.CompatUsersUserIDValidateSecurityKeyRegistrationPost(ctx, &req, params)
	})
	return err
}

func (u *user2faOp) ListSecurityKeys(ctx context.Context) (*v1.CompatUsersUserIDSecurityKeysGetOK, error) {
	return iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDSecurityKeysGetOK]("User2FA.ListSecurityKeys", func() (any, error) {
		return u.client.CompatUsersUserIDSecurityKeysGet(ctx, v1.CompatUsersUserIDSecurityKeysGetParams{UserID: u.getUserID()})
	})
}

func (u *user2faOp) GetSecurityKey(ctx context.Context, securityKeyID int) (*v1.UserSecurityKey, error) {
	return iam.ErrorFromDecodedResponse[v1.UserSecurityKey]("User2FA.GetSecurityKey", func() (any, error) {
		return u.client.CompatUsersUserIDSecurityKeysSecurityKeyIDGet(ctx, v1.CompatUsersUserIDSecurityKeysSecurityKeyIDGetParams{
			UserID:        u.getUserID(),
			SecurityKeyID: securityKeyID,
		})
	})
}

func (u *user2faOp) UpdateSecurityKey(ctx context.Context, securityKeyID int, name string) (*v1.UserSecurityKey, error) {
	return iam.ErrorFromDecodedResponse[v1.UserSecurityKey]("User2FA.UpdateSecurityKey", func() (any, error) {
		req := v1.NewOptCompatUsersUserIDSecurityKeysSecurityKeyIDPutReq(v1.CompatUsersUserIDSecurityKeysSecurityKeyIDPutReq{Name: name})
		params := v1.CompatUsersUserIDSecurityKeysSecurityKeyIDPutParams{
			UserID:        u.getUserID(),
			SecurityKeyID: securityKeyID,
		}
		return u.client.CompatUsersUserIDSecurityKeysSecurityKeyIDPut(ctx, req, params)
	})
}

func (u *user2faOp) DeleteSecurityKey(ctx context.Context, securityKeyID int) error {
	_, err := iam.ErrorFromDecodedResponse[v1.CompatUsersUserIDSecurityKeysSecurityKeyIDDeleteNoContent]("User2FA.DeleteSecurityKey", func() (any, error) {
		return u.client.CompatUsersUserIDSecurityKeysSecurityKeyIDDelete(ctx, v1.CompatUsersUserIDSecurityKeysSecurityKeyIDDeleteParams{
			UserID:        u.getUserID(),
			SecurityKeyID: securityKeyID,
		})
	})
	return err
}
