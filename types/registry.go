package types

type Registry struct {
	URL      string `json:"registry"`
	Name     string `json:"registry_name"`
	Username string `json:"registry_user"`
	Password string `json:"registry_pass"`
	Type     string `json:"registry_type"`
}
