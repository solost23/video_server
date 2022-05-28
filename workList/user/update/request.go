package update

type Request struct {
	ID        string `json:"id"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"`
}
