package delete

type Request struct {
	RoleName string `json:"roleName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}
