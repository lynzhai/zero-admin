// Code generated by goctl. DO NOT EDIT!
// Source: us.proto

//go:generate mockgen -destination ./us_mock.go -package usclient -source $GOFILE

package usclient

import (
	"context"

	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	PersionLoginReq     = us.PersionLoginReq
	PersionAddReq       = us.PersionAddReq
	PersionListResp     = us.PersionListResp
	RoleListReq         = us.RoleListReq
	RoleAddReq          = us.RoleAddReq
	RoleListResp        = us.RoleListResp
	RoleDeleteReq       = us.RoleDeleteReq
	RoleDeleteResp      = us.RoleDeleteResp
	PersionLoginResp    = us.PersionLoginResp
	PersionUpdateResp   = us.PersionUpdateResp
	PersionDeleteResp   = us.PersionDeleteResp
	RoleData            = us.RoleData
	PersionDeleteReq    = us.PersionDeleteReq
	RoleAddResp         = us.RoleAddResp
	PersionUpdateReq    = us.PersionUpdateReq
	RoleUpdateReq       = us.RoleUpdateReq
	RoleUpdateResp      = us.RoleUpdateResp
	PersionRegisterReq  = us.PersionRegisterReq
	PersionRegisterResp = us.PersionRegisterResp
	PersionData         = us.PersionData
	PersionAddResp      = us.PersionAddResp
	PersionListReq      = us.PersionListReq

	Us interface {
		PersionAdd(ctx context.Context, in *PersionAddReq) (*PersionAddResp, error)
		PersionList(ctx context.Context, in *PersionListReq) (*PersionListResp, error)
		PersionUpdate(ctx context.Context, in *PersionUpdateReq) (*PersionUpdateResp, error)
		PersionDelete(ctx context.Context, in *PersionDeleteReq) (*PersionDeleteResp, error)
		RoleAdd(ctx context.Context, in *RoleAddReq) (*RoleAddResp, error)
		RoleList(ctx context.Context, in *RoleListReq) (*RoleListResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq) (*RoleUpdateResp, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq) (*RoleDeleteResp, error)
		PersionLogin(ctx context.Context, in *PersionLoginReq) (*PersionLoginResp, error)
		PersionRegister(ctx context.Context, in *PersionRegisterReq) (*PersionRegisterResp, error)
	}

	defaultUs struct {
		cli zrpc.Client
	}
)

func NewUs(cli zrpc.Client) Us {
	return &defaultUs{
		cli: cli,
	}
}

func (m *defaultUs) PersionAdd(ctx context.Context, in *PersionAddReq) (*PersionAddResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.PersionAdd(ctx, in)
}

func (m *defaultUs) PersionList(ctx context.Context, in *PersionListReq) (*PersionListResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.PersionList(ctx, in)
}

func (m *defaultUs) PersionUpdate(ctx context.Context, in *PersionUpdateReq) (*PersionUpdateResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.PersionUpdate(ctx, in)
}

func (m *defaultUs) PersionDelete(ctx context.Context, in *PersionDeleteReq) (*PersionDeleteResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.PersionDelete(ctx, in)
}

func (m *defaultUs) RoleAdd(ctx context.Context, in *RoleAddReq) (*RoleAddResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.RoleAdd(ctx, in)
}

func (m *defaultUs) RoleList(ctx context.Context, in *RoleListReq) (*RoleListResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.RoleList(ctx, in)
}

func (m *defaultUs) RoleUpdate(ctx context.Context, in *RoleUpdateReq) (*RoleUpdateResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in)
}

func (m *defaultUs) RoleDelete(ctx context.Context, in *RoleDeleteReq) (*RoleDeleteResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.RoleDelete(ctx, in)
}

func (m *defaultUs) PersionLogin(ctx context.Context, in *PersionLoginReq) (*PersionLoginResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.PersionLogin(ctx, in)
}

func (m *defaultUs) PersionRegister(ctx context.Context, in *PersionRegisterReq) (*PersionRegisterResp, error) {
	client := us.NewUsClient(m.cli.Conn())
	return client.PersionRegister(ctx, in)
}
