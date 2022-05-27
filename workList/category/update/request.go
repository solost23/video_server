package update

type Request struct {
	UserId     string `json:"userId"`
	CategoryId string `json:"categoryId"`

	Title     string `json:"title"`
	Introduce string `json:"introduce"`
}
