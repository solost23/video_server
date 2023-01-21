package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/http"
	"testing"
	"video_server/forms"
	"video_server/pkg/utils"
)

func TestService_VideoList(t *testing.T) {
	ginCtx, _ := gin.CreateTestContext(&http.TestResponseWriter{})
	type arg struct {
		ctx    *gin.Context
		params *forms.VideoListForm
	}
	type want struct {
		results []*forms.VideoListResponse
		err     error
	}
	tests := []struct {
		arg  arg
		want want
	}{
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.VideoListForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 20,
					},
					VideoTitle: "",
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, test := range tests {
		results, err := (&Service{}).VideoList(test.arg.ctx, test.arg.params)
		if err != nil {
			t.Errorf("err: %+v \n", err)
		}
		fmt.Printf("results: %+v \n", results)
	}
}
