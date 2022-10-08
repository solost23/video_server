package constants

const (
	RedisPrefix = "video_server_"

	WebRedisPrefix    = RedisPrefix + "web_"
	MobileRedisPrefix = RedisPrefix + "mobile_"
)

const (
	// 管理员
	ROLEADMIN = "ADMIN"
	ROLEUSER  = "USER"

	// 删除状态
	DELETENORMAL = "DELETE_STATUS_NORMAL"
	DELETEDEL    = "DELETE_STATUS_DEL"

	// 评论类型
	ISTHUMB   = "ISTHUMB"
	ISCOMMENT = "ISCOMMENT"
)
