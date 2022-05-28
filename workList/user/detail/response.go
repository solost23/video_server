package detail

type Response struct {
	ID           string `json:"id"`
	UserName     string `json:"user_name"`
	Nickname     string `json:"nickname"`
	Role         string `json:"role"`
	Avatar       string `json:"avatar"`
	Introduce    string `json:"introduce"`
	FansCount    int64  `json:"fansCount"`
	CommentCount int64  `json:"commentCount"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}
