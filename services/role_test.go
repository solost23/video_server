package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/http"
	"testing"
	"video_server/forms"
)

func TestService_ListRole(t *testing.T) {
	ginCtx, _ := gin.CreateTestContext(&http.TestResponseWriter{})
	type arg struct {
		ctx    *gin.Context
		params *forms.RoleListForm
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
			},
		},
	}

	for _, test := range tests {
		results, err := (&Service{}).ListRole(test.arg.ctx, test.arg.params)
		if err != nil {
			t.Errorf("err: %+v \n", err)
		}
		fmt.Printf("results: %+v \n", results)
	}
}
