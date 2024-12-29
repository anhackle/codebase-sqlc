package impl

import (
	"context"

	"github.com/anle/codebase/internal/database"
	"github.com/anle/codebase/internal/service"
	"github.com/anle/codebase/model"
)

type sUserAuthen struct {
	r *database.Queries
}

func (s *sUserAuthen) Login(ctx context.Context) error {
	return nil
}

func (s *sUserAuthen) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	panic("error")
}

func (s *sUserAuthen) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out *model.VerifyOTPOutput, err error) {
	panic("error")
}

func (s *sUserAuthen) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}

func NewUserAuthenImpl(r *database.Queries) service.IUserAuthen {
	return &sUserAuthen{
		r: r,
	}
}
