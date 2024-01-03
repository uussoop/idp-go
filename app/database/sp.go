package database

type ServiceProviders struct {
	ID          int    `json:"id"          yaml:"id"`
	Name        string `json:"name"        yaml:"name"`
	Description string `json:"description" yaml:"description"`
	URL         string `json:"url"         yaml:"url"`
	Token       string `json:"token"       yaml:"token"`
}
