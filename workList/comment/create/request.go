package create

type Request struct {
	VideoID  string `json:"videoID"`
	Content  string `json:"content"`
	ParentID string `json:"parentId"`
	ISThumb  string `json:"ISThumb"`
}
