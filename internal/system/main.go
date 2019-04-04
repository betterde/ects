package system

type Information struct {
	Version    string `json:"version"`
	Installed  bool   `json:"installed"`
	Permission bool   `json:"permission"`
}

var (
	Info *Information
)