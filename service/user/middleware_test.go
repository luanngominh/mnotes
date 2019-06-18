package user

import (
	"context"
	"testing"

	"github.com/luanngominh/mnotes/model"
)

func Test_validationMiddleware_Active(t *testing.T) {
	serviceMock2User := &ServiceMock{
		ActiveFunc: func(ctx context.Context, userID string, verifyCode string) (*model.User, error) {
			return nil, nil
		},
		GetFunc: func(ctx context.Context, query *UserQuery) ([]*model.User, error) {
			users := []*model.User{
				&model.User{},
				&model.User{},
			}
			return users, nil
		},
	}
	serviceMock := &ServiceMock{
		ActiveFunc: func(ctx context.Context, userID string, verifyCode string) (*model.User, error) {
			return nil, nil
		},
		GetFunc: func(ctx context.Context, query *UserQuery) ([]*model.User, error) {
			users := []*model.User{}
			return users, nil
		},
	}

	type args struct {
		ctx        context.Context
		userID     string
		verifyCode string
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		checkUserIDExist bool
	}{
		{
			name: "Valid test case",
			args: args{
				userID:     "7521bd92-91ab-4948-b4f1-b16b2dd9f47b",
				verifyCode: "123456",
			},
			wantErr:          false,
			checkUserIDExist: false,
		},
		{
			name: "Empty user id",
			args: args{
				userID:     "",
				verifyCode: "123455",
			},
			wantErr:          true,
			checkUserIDExist: false,
		},
		{
			name: "Empty user verify code",
			args: args{
				userID:     "7521bd92-91ab-4948-b4f1-b16b2dd9f47b",
				verifyCode: "",
			},
			wantErr:          true,
			checkUserIDExist: false,
		},
		{
			name: "Empty user verify code and userID",
			args: args{
				userID:     "",
				verifyCode: "",
			},
			wantErr:          true,
			checkUserIDExist: false,
		},
		{
			name: "User id not existed",
			args: args{
				userID:     "adfdf",
				verifyCode: "223333",
			},
			wantErr:          true,
			checkUserIDExist: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{}
			if !tt.checkUserIDExist {
				mw = validationMiddleware{
					Service: serviceMock2User,
				}
			} else {
				mw = validationMiddleware{
					Service: serviceMock,
				}
			}

			_, err := mw.Active(tt.args.ctx, tt.args.userID, tt.args.verifyCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Active() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
