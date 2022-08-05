package utils

type PageForm struct {
	Page int `form:"page" comment:"当前页码"`
	Size int `form:"size" comment:"每页显示记录数"`
}

type IdForm struct {
	Id int `uri:"id" comment:"id" binding:"min=1"`
}

type UIdForm struct {
	Id uint `uri:"id" comment:"id" binding:"min=1"`
}

type IdsForm struct {
	Ids string `uri:"ids" comment:"ids"`
}

type All struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PageList struct {
	Size    int   `json:"size"`
	Pages   int64 `json:"pages"`
	Total   int64 `json:"total"`
	Current int   `json:"current"`
}
