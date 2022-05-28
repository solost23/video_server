package detail

type Response struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	ClassID      string `json:"classId"`
	Title        string `json:"title"`
	Introduce    string `json:"introduce"`
	ImageUrl     string `json:"imageUrl"`
	VideoUrl     string `json:"videoUrl"`
	ThumbCount   int64  `json:"thumbCount"`
	CommentCount int64  `json:"commentCount"`
	DeleteStatus string `json:"deleteStatus"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}
