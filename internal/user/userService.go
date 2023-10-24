package user

import (
	"Github.com/Synoptic2023/internal/util"
)

type UserService interface {
	CreateUser(username string, password string, role string) (User, error)
	Login(username string, password string) (User, error)
	GetProfile(userId int64) (Profile, error)
	ListUsers() ([]User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{repo: userRepository}
}

// returns a user with no id set since this is set once the user is entered into the database
func (s *userService) CreateUser(username string, password string, role string) (User, error) {

	//hashes password using the standard bcrypt hashing alogrithm included in go
	passwordHash, err := util.HashPassword(password)

	if err != nil {
		return User{}, err
	}

	user := User{Username: username, Password: passwordHash, Role: role}

	newUser, err := s.repo.CreateUser(user)

	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (s *userService) Login(username string, password string) (User, error) {

	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return User{}, err
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return User{}, err
	}

	user.Password = ""

	return user, nil
}

func (s *userService) GetProfile(userId int64) (Profile, error) {

	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return Profile{}, err
	}

	profile, err := s.repo.GetUserProfile(user)
	if err != nil {
		return Profile{}, err
	}

	return profile, nil

}

func (s *userService) ListUsers() ([]User, error) {

	users, err := s.repo.ListUsers()
	if err != nil {
		return []User{}, err
	}

	return users, nil
}
