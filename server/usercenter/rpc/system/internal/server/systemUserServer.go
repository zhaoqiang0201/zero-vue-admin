// Code generated by goctl. DO NOT EDIT!
// Source: sysuser.proto

package server

import (
	"context"

	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/logic"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/rpc/system/pb"
)

type SystemUserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSystemUserServer
}

func NewSystemUserServer(svcCtx *svc.ServiceContext) *SystemUserServer {
	return &SystemUserServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemUserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
