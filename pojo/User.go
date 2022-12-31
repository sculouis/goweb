package pojo

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	PassWord string `json:"UserPassWord"`
	Email    string `json:"UserEmail"`
}
