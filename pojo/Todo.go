package pojo

type Todo struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Done     bool   `json:"Done"`
	Selected bool   `json:"Selected"`
}
