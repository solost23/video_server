package register

type Request struct {
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}
