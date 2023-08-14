package user

import (
	"context"
	"errors"
	"testing"

	"github.com/MuhAndriJP/dealls.git/action/user/mocks"
	"github.com/MuhAndriJP/dealls.git/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin_Handle(t *testing.T) {
	type fields struct {
		uRepo *mocks.IUser
	}
	type args struct {
		ctx context.Context
		req *entity.Users
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		wantRes entity.Users
		wantErr error
	}{
		{
			name: "Found Email",
			prepare: func(f *fields) {
				f.uRepo.On("GetUserByEmail", context.Background(), "registered@gmail.com").Return(
					entity.Users{
						Email: "registered@gmail.com",
					}, nil,
				)
			},
			args: args{
				ctx: context.Background(),
				req: &entity.Users{
					Email: "registered@gmail.com",
				},
			},
			wantRes: entity.Users{
				Email: "registered@gmail.com",
			},
			wantErr: nil,
		},
		{
			name: "Not Found Email",
			prepare: func(f *fields) {
				f.uRepo.On("GetUserByEmail", context.Background(), "unregistered@gmail.com").Return(
					entity.Users{}, errors.New("not found"),
				)
			},
			args: args{
				ctx: context.Background(),
				req: &entity.Users{
					Email: "unregistered@gmail.com",
				},
			},
			wantRes: entity.Users{},
			wantErr: errors.New("not found"),
		},
	}

	tests2 := []struct {
		name    string
		prepare func(f *fields)
		args    args
		wantRes entity.Users
		wantErr error
	}{
		{
			name: "Upsert Token",
			prepare: func(f *fields) {
				f.uRepo.On("Upsert", context.Background(), &entity.Users{
					ID:       1,
					Name:     "registered user",
					Email:    "registered@gmail.com",
					Password: "existingPassword",
					Token:    "new_token_generated",
				}).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				req: &entity.Users{
					ID:       1,
					Name:     "registered user",
					Email:    "registered@gmail.com",
					Password: "existingPassword",
					Token:    "new_token_generated",
				},
			},
			wantErr: nil,
		},
		{
			name: "Upsert Token Fail",
			prepare: func(f *fields) {
				f.uRepo.On("Upsert", context.Background(), &entity.Users{
					ID:       1,
					Name:     "registered user",
					Email:    "registered@gmail.com",
					Password: "existingPassword",
					Token:    "new_token_generated",
				}).Return(errors.New("internal server error"))
			},
			args: args{
				ctx: context.Background(),
				req: &entity.Users{
					ID:       1,
					Name:     "registered user",
					Email:    "registered@gmail.com",
					Password: "existingPassword",
					Token:    "new_token_generated",
				},
			},
			wantErr: errors.New("internal server error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				uRepo: new(mocks.IUser),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			gotRes, gotErr := f.uRepo.GetUserByEmail(tt.args.ctx, tt.args.req.Email)
			if gotErr != nil {
				assert.Equal(t, tt.wantErr, gotErr)
			} else {
				assert.NoError(t, gotErr)
				assert.Equal(t, tt.wantRes, gotRes)
			}
		})
	}

	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			f := fields{
				uRepo: new(mocks.IUser),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			gotErr := f.uRepo.Upsert(tt.args.ctx, tt.args.req)
			if gotErr != nil {
				assert.Equal(t, tt.wantErr, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}
