package user

type User struct {
	NameFull string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// func (u *User) NewUser(name, email, phone string) *User {
// 	return &User{
// 		NameFull: name,
// 		Email:    email,
// 		Phone:    phone,
// 	}
// }
