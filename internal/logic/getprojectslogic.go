package logic

import (
	"context"

	"projects/api/internal/svc"
	"projects/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectsLogic {
	return &GetProjectsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectsLogic) GetProjects(req *types.GetReq) (resp *types.GetResp, err error) {
	type Project struct {
		ID          int
		Name        string
		Description string
	}

	projects := []Project{
		{ID: 1, Name: "Project 1", Description: "Description Project 1"},
		{ID: 2, Name: "Project 2", Description: "Description Project 2"},
	}

	resp = &types.GetResp{
		Id:          uint(projects[0].ID), // Example for the first project
		Name:        projects[0].Name,
		Description: projects[0].Description,
		Created_at:  "2023-01-01T00:00:00Z", // Example value
		Updated_at:  "2023-01-01T00:00:00Z", // Example value
	}

	return resp, nil
}
