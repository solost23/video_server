package deleteVideo

import (
	"context"
	"testing"
	"video/mysql"
)

func TestTask_Deal(t *testing.T) {
	var dbConn = mysql.GetMysqlConn()
	task := &Task{
		Ctx:  context.Background(),
		Conn: dbConn,
	}
	task.DeleteVideo()
}
