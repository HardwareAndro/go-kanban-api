package services

import (
	"log"
	"os"

	model "github.com/HardwareAndro/go-kanban-api/app/models"
	"github.com/HardwareAndro/go-kanban-api/app/shared/constants"
	repository "github.com/HardwareAndro/go-kanban-api/app/shared/repositories"
	"github.com/HardwareAndro/go-kanban-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.GenericRepository[model.User]
	App            config.GoAppTools
}

func NewUserService(userRepository *repository.GenericRepository[model.User]) *UserService {
	InfoLogger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)
	var app config.GoAppTools
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger
	return &UserService{
		userRepository: userRepository,
		App:            app,
	}
}

// HashPassword hashes the user's password
func (us *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password matches the hashed password
func (us *UserService) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (us *UserService) RegisterUser(user *model.User) (*model.User, error) {
	// Hash the password before storing
	hashedPassword, err := us.HashPassword(user.Password)
	if err != nil {
		us.App.ErrorLogger.Fatalln(constants.ERR_REGISTER_USER, err)
		return nil, err
	}
	user.Password = hashedPassword // Store the hashed password

	_, err = us.userRepository.Create(user)
	if err != nil {
		us.App.ErrorLogger.Fatalln(constants.ERR_REGISTER_USER, err)
		return nil, err
	}
	us.App.InfoLogger.Println(constants.SUCCESS_REGISTER_USER, user.ID)
	return user, nil
}

func (us *UserService) LoginUser(credentials model.User) (*model.User, error) {
	// Use FindAsync with a function literal to filter by email
	user, err := us.userRepository.FindAsync(func(filter bson.M) bson.M {
		return bson.M{"email": credentials.Email}
	})
	if err != nil || !us.CheckPassword(user.Password, credentials.Password) {
		us.App.ErrorLogger.Println(constants.ERR_INVALID_CREDENTIALS, err)
		return nil, err
	}
	us.App.InfoLogger.Println(constants.SUCCESS_LOGIN_USER, user.ID)
	return user, nil
}

func (us *UserService) GetUserById(id string) (*model.User, error) {
	user, err := us.userRepository.FindById(id)
	if err != nil {
		us.App.ErrorLogger.Println(constants.ERR_USER_NOT_FOUND, err)
		return nil, err
	}
	return user, nil
}

func (us *UserService) LogoutUser() error {
	// Logic for logout (e.g., invalidating session or token)
	us.App.InfoLogger.Println(constants.SUCCESS_LOGOUT_USER)
	return nil
}
