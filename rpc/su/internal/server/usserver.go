// Code generated by goctl. DO NOT EDIT!
// Source: su.proto

package server

import (
	"context"

	"go-zero-admin/rpc/su/internal/logic"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"
)

type UsServer struct {
	svcCtx *svc.ServiceContext
}

func NewUsServer(svcCtx *svc.ServiceContext) *UsServer {
	return &UsServer{
		svcCtx: svcCtx,
	}
}

func (s *UsServer) SubjectAdd(ctx context.Context, in *su.SubjectAddReq) (*su.SubjectAddResp, error) {
	l := logic.NewSubjectAddLogic(ctx, s.svcCtx)
	return l.SubjectAdd(in)
}

func (s *UsServer) SubjectList(ctx context.Context, in *su.MainTeacherSubjectListReq) (*su.MainTeacherSubjectListReap, error) {
	l := logic.NewSubjectListLogic(ctx, s.svcCtx)
	return l.SubjectList(in)
}

func (s *UsServer) SubjectUpdate(ctx context.Context, in *su.SubjectUpdateReq) (*su.SubjectUpdateResp, error) {
	l := logic.NewSubjectUpdateLogic(ctx, s.svcCtx)
	return l.SubjectUpdate(in)
}

func (s *UsServer) SubjectDelete(ctx context.Context, in *su.SubjectDeleteReq) (*su.SubjectDeleteResp, error) {
	l := logic.NewSubjectDeleteLogic(ctx, s.svcCtx)
	return l.SubjectDelete(in)
}

func (s *UsServer) SubjectFind(ctx context.Context, in *su.SubjectFindReq) (*su.SubjectFindResp, error) {
	l := logic.NewSubjectFindLogic(ctx, s.svcCtx)
	return l.SubjectFind(in)
}
