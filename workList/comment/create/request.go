package create

type Request struct {
	VideoID  string `json:"videoId"`
	Content  string `json:"content"`
	ParentID string `json:"parentId"`
	ISThumb  string `json:"ISThumb"`
}
