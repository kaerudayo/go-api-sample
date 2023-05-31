//go:build parallel
// +build parallel

package user_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/api-sample/app/domain/model"
	mock_query "github.com/api-sample/app/infra/reader/mock"
	"github.com/api-sample/app/pkg/result"
	"github.com/api-sample/app/usecase/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser_FindById(t *testing.T) {
	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) user.U
		in    user.FindByIDInput
		out   user.FindByIDOutput
		res   result.Response
	}{
		"error_return_error": {
			setup: func(ctrl *gomock.Controller) user.U {
				userMock := mock_query.NewMockUserQuery(ctrl)
				userMock.EXPECT().FindByID("hoge").Return(
					model.User{},
					errors.New("error"),
				)
				return user.U{
					Q: user.Q{
						UserQuery: userMock,
					},
				}
			},
			in:  user.FindByIDInput{UserID: "hoge"},
			out: user.FindByIDOutput{ID: "", Name: ""},
			res: result.NewResponse(
				http.StatusInternalServerError,
				"Internal Server Error",
			),
		},
		"error_user_empty": {
			setup: func(ctrl *gomock.Controller) user.U {
				userMock := mock_query.NewMockUserQuery(ctrl)
				userMock.EXPECT().FindByID("a").Return(model.User{}, nil)
				return user.U{
					Q: user.Q{
						UserQuery: userMock,
					},
				}
			},
			in:  user.FindByIDInput{UserID: "a"},
			out: user.FindByIDOutput{ID: "", Name: ""},
			res: result.NewResponse(http.StatusNotFound, "User not found"),
		},
		"success": {
			setup: func(ctrl *gomock.Controller) user.U {
				userMock := mock_query.NewMockUserQuery(ctrl)
				userMock.EXPECT().FindByID("hoge_id").Return(model.User{
					ID:   "hoge_id",
					Name: "hoge_name",
				}, nil)
				return user.U{
					Q: user.Q{
						UserQuery: userMock,
					},
				}
			},
			in:  user.FindByIDInput{UserID: "hoge_id"},
			out: user.FindByIDOutput{ID: "hoge_id", Name: "hoge_name"},
			res: result.NewResponse(http.StatusOK, ""),
		},
	}
	for name, c := range cases {
		c := c
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		r := c.setup(ctrl)
		t.Run(name, func(t *testing.T) {
			u := user.NewTestUsecase(r)
			out, res := u.FindByID(c.in, nil)
			assert.Equal(t, c.res, res)
			assert.Equal(t, c.out, out)
		})
	}
}
