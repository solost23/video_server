package delete

type Request struct {
	UserID  string `json:"userId"`
	ClassID string `json:"classId"`
	VideoID string `json:"videoId"`
}
