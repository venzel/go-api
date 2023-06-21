package user

type CreateUserDto struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type ResponseUserDto struct {
	ID    string
	Name  string
	Email string
}
