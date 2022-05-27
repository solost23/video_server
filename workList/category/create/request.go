package create

type Request struct {
	UserId    string `json:"userId"`
	Title     string `json:"title"`
	Introduce string `json:"introduce"`
}
