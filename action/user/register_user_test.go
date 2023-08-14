package user

import (
	"context"
	"errors"
	"testing"

	"github.com/MuhAndriJP/dealls.git/action/user/mocks"
	"github.com/MuhAndriJP/dealls.git/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister_Handle(t *testing.T) {
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
			name: "Upsert Token",
			prepare: func(f *fields) {
				f.uRepo.On("Insert", context.Background(), &entity.Users{
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
				f.uRepo.On("Insert", context.Background(), &entity.Users{
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
			gotErr := f.uRepo.Insert(tt.args.ctx, tt.args.req)
			if gotErr != nil {
				assert.Equal(t, tt.wantErr, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
		})
	}
}
