package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/http"
	"testing"
	"video_server/forms"
	"video_server/pkg/utils"
)

func TestService_CommentList(t *testing.T) {
	ginCtx, _ := gin.CreateTestContext(&http.TestResponseWriter{})
	type arg struct {
		ctx    *gin.Context
		params *forms.CommentListForm
	}
	type want struct {
		err error
	}
	tests := []struct {
		arg  arg
		want want
	}{
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.CommentListForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 10,
					},
					VideoId: 1,
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.CommentListForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 10,
					},
					VideoId: 2,
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.CommentListForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 10,
					},
					VideoId: 3,
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, test := range tests {
		results, err := (&Service{}).CommentList(test.arg.ctx, test.arg.params)
		if err != nil {
			t.Errorf("err: %+v \n", err)
		}
		fmt.Printf("results: %+v \n", results)
	}
}
