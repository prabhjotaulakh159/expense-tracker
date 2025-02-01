package services


type UserService struct {}

func NewUserService() *UserService {
	return &UserService {}
}

func (userService *UserService) Register(username string, password string) error {
	return nil
}