package model

import (
	"fmt"
	"testing"
)

func TestGetMysqlConn(t *testing.T) {
	db, err := NewMysqlClient(false)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(db)
}
