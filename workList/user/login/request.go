package login

type Request struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
