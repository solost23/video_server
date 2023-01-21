package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/http"
	"os"
	"testing"
	"video_server/forms"
	"video_server/global/initialize"
	"video_server/pkg/utils"
)

func TestMain(m *testing.M) {
	initialize.Initialize("../configs/configs.yml")
	os.Exit(m.Run())
}

func TestService_SearchCategory(t *testing.T) {
	ginCtx, _ := gin.CreateTestContext(&http.TestResponseWriter{})
	type arg struct {
		ctx    *gin.Context
		params *forms.SearchForm
	}
	type want struct {
		results []*forms.CategoryListResponse
		err     error
	}
	tests := []struct {
		arg  arg
		want want
	}{
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.SearchForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 10,
					},
					Keyword: "",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.SearchForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 10,
					},
					Keyword: "",
				},
			},
			want: want{
				err: nil,
			},
		},
		{
			arg: arg{
				ctx: ginCtx,
				params: &forms.SearchForm{
					PageForm: utils.PageForm{
						Page: 1,
						Size: 10,
					},
					Keyword: "",
				},
			},
			want: want{
				err: nil,
			},
		},
	}

	for _, test := range tests {
		results, err := (&Service{}).SearchCategory(test.arg.ctx, test.arg.params)
		if err != nil {
			t.Errorf("error: %+v \n", err)
		}
		fmt.Printf("results: %+v \n", results)
	}
}
