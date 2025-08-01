package Usecase

type UserUsecase struct {
	UserRepo IUserRepo
	UserService IUserServices
}

func NewUserUscase(userRepo IUserRepo,userService IUserServices) *UserUsecase{
	return &UserUsecase{userRepo,userService}
}


//User usecase methodes