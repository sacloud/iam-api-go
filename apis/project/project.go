package project

import (
	"context"

	iam "github.com/sacloud/iam-api-go"
	v1 "github.com/sacloud/iam-api-go/apis/v1"
	"github.com/sacloud/iam-api-go/common"
)

type ProjectAPI interface {
	List(ctx context.Context, params ListParams) (*v1.ProjectsGetOK, error)

	Create(ctx context.Context, params CreateParams) (*v1.Project, error)
	Read(ctx context.Context, id int) (*v1.Project, error)
	Update(ctx context.Context, id int, name string, description string) (*v1.Project, error)
	Delete(ctx context.Context, id int) error

	Move(ctx context.Context, ids []int, parentFolderID *int) error
}

type projectOp struct {
	client *v1.Client
}

var _ ProjectAPI = (*projectOp)(nil)

func NewProjectOp(client *v1.Client) ProjectAPI { return &projectOp{client} }

type ListParams struct {
	Page           *int
	PerPage        *int
	Ordering       *v1.ProjectsGetOrdering
	IamRole        *string
	ParentFolderID *int
}

func (p *projectOp) List(ctx context.Context, params ListParams) (*v1.ProjectsGetOK, error) {
	return common.ErrorFromDecodedResponse[v1.ProjectsGetOK]("Project.List", func() (any, error) {
		return p.client.ProjectsGet(ctx, v1.ProjectsGetParams{
			Page:           iam.IntoOpt[v1.OptInt](params.Page),
			PerPage:        iam.IntoOpt[v1.OptInt](params.PerPage),
			Ordering:       iam.IntoOpt[v1.OptProjectsGetOrdering](params.Ordering),
			IamRole:        iam.IntoOpt[v1.OptString](params.IamRole),
			ParentFolderID: iam.IntoOpt[v1.OptInt](params.ParentFolderID),
		})
	})
}

type CreateParams struct {
	Code           string
	Name           string
	Description    string
	ParentFolderID *int
}

func (p *projectOp) Create(ctx context.Context, params CreateParams) (*v1.Project, error) {
	return common.ErrorFromDecodedResponse[v1.Project]("Project.Create", func() (any, error) {
		return p.client.ProjectsPost(ctx, &v1.ProjectsPostReq{
			Code:           params.Code,
			Name:           params.Name,
			Description:    params.Description,
			ParentFolderID: iam.IntoOpt[v1.OptInt](params.ParentFolderID),
		})
	})
}

func (p *projectOp) Read(ctx context.Context, id int) (*v1.Project, error) {
	return common.ErrorFromDecodedResponse[v1.Project]("Project.Read", func() (any, error) {
		return p.client.ProjectsProjectIDGet(ctx, v1.ProjectsProjectIDGetParams{ProjectID: id})
	})
}

func (p *projectOp) Update(ctx context.Context, id int, name string, description string) (*v1.Project, error) {
	return common.ErrorFromDecodedResponse[v1.Project]("Project.Update", func() (any, error) {
		params := v1.ProjectsProjectIDPutParams{
			ProjectID: id,
		}
		request := v1.ProjectsProjectIDPutReq{
			Name:        name,
			Description: description,
		}
		return p.client.ProjectsProjectIDPut(ctx, &request, params)
	})
}

func (p *projectOp) Delete(ctx context.Context, projectID int) error {
	_, err := common.ErrorFromDecodedResponse[v1.ProjectsProjectIDDeleteNoContent]("Project.Delete", func() (any, error) {
		return p.client.ProjectsProjectIDDelete(ctx, v1.ProjectsProjectIDDeleteParams{ProjectID: projectID})
	})

	return err
}

type MoveProjectsParams struct {
	ProjectIDs     []int
	ParentFolderID *int
}

func (p *projectOp) Move(ctx context.Context, ids []int, parentFolderID *int) error {
	_, err := common.ErrorFromDecodedResponse[v1.MoveProjectsPostNoContent]("Project.Move", func() (any, error) {
		return p.client.MoveProjectsPost(ctx, &v1.MoveProjects{
			ProjectIds:     ids,
			ParentFolderID: iam.IntoNullable[v1.NilInt](parentFolderID),
		})
	})

	return err
}
