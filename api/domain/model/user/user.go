package user

type User struct {
	ID    ID
	Email Email
	Name  Name
}

func NewUser(id ID, email Email, name Name) *User {
	return &User{
		ID:    id,
		Email: email,
		Name:  name,
	}
}
