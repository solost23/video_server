package response

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type JSONTime time.Time

func (j JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(j).Format("2006/01/02 15:04:05"))), nil
}

type JSONDate time.Time

func (j JSONDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(j).Format("2006/01/02"))), nil
}

type Response struct {
	ErrorCode int         `json:"code"`
	ErrorMsg  string      `json:"message"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
}

func Error(c *gin.Context, code int, err error) {
	resp := &Response{ErrorCode: code, Success: false, ErrorMsg: err.Error(), Data: ""}
	c.JSON(http.StatusOK, resp)
	// 记录日志
	_ = c.AbortWithError(http.StatusOK, err)
}

func Success(c *gin.Context, data interface{}) {
	resp := &Response{ErrorCode: 0, Success: true, ErrorMsg: "", Data: data}
	c.JSON(http.StatusOK, resp)
	// 记录日志
}

func MessageSuccess(c *gin.Context, errorMsg string, data interface{}) {
	resp := &Response{ErrorCode: 0, Success: true, ErrorMsg: errorMsg, Data: data}
	c.JSON(http.StatusOK, resp)
	// 记录日志
}

func CSV(c *gin.Context, filename string, data [][]string) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+time.Now().Format("2006/01/02")+filename)

	bytesBuffer := &bytes.Buffer{}
	bytesBuffer.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(bytesBuffer)
	err := w.Write(data[0])
	if err != nil {
		Error(c, 1000, err)
		return
	}
	w.Flush()

	err = w.WriteAll(data[1:])
	if err != nil {
		Error(c, 1000, err)
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", bytesBuffer.Bytes())
}

func Binary(c *gin.Context, filename string, content []byte) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+time.Now().Format("2006/01/02")+filename)

	bytesBuffer := &bytes.Buffer{}
	bytesBuffer.Write(content)

	c.Data(http.StatusOK, "application/octet-stream", bytesBuffer.Bytes())
}
