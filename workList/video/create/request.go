package create

type Request struct {
	UserID    string `json:"userId"`
	ClassID   string `json:"classId"`
	Title     string `json:"title"`
	Introduce string `json:"introduce"`
	ImageUrl  string `json:"imageUrl"`
	VideoUrl  string `json:"videoUrl"`
}
