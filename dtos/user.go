package dtos

type UserData struct {
	UserId    string `json:"user_id" valid:"required"`
	FirstName string `json:"first_name" valid:"required"`
	LastName  string `json:"last_name" valid:"required"`
	Nickname  string `json:"nickname" valid:"required"`
	Password  string `json:"password" valid:"required"`
	Email     string `json:"email" valid:"required,email"`
	Country   string `json:"country" valid:"required"`
}
