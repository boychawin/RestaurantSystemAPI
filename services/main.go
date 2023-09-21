package services

import (

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"restaurant/errors"
	"restaurant/models"
	"restaurant/repositorys"
)

type userService struct {
	userRepo    repositorys.UsersRepository
}

func NewMainService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) Session(c *fiber.Ctx) (users *UserResponse, err error) {


	// Repository
	usersDB, err := s.userRepo.Session(c)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError(err.Error())
		}

		// logs.Error(err)
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &UserResponse{
		Data: &models.Users{
			ID:        usersDB.Data.ID,
			FirstName: usersDB.Data.FirstName,
			LastName:  usersDB.Data.LastName,
			UserName:  usersDB.Data.UserName,
		},
		Messages: usersDB.Messages,
	}



	return user, nil
}

func (s userService) Refresh(c *fiber.Ctx) (users *UserResponse, err error) {

	// Repository
	usersRepo, err := s.userRepo.Refresh(c)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("user not found")
		}

		// logs.Error(err)
		return nil, errors.NewUnexpectedError("unexpected error")
	}

	user := &UserResponse{
		Messages:     usersRepo.Messages,
		AccessToken:  usersRepo.AccessToken,
		RefreshToken: usersRepo.RefreshToken,
	}

	return user, nil
}

func (s userService) Login(c *fiber.Ctx) (users *UserResponse, err error) {
	usersRepo, err := s.userRepo.Login(c)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("user not found")
		}
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &UserResponse{
		Messages:     usersRepo.Messages,
		AccessToken:  usersRepo.AccessToken,
		RefreshToken: usersRepo.RefreshToken,
	}

	return user, nil
}

func (s userService) Register(c *fiber.Ctx) (users *UserResponse, err error) {

	usersRepo, err := s.userRepo.Register(c)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("ไม่พบข้อมูล")
		}
		return nil, errors.NewUnexpectedError(err.Error())
	}

	user := &UserResponse{
		Messages: usersRepo.Messages,
	}

	return user, nil
}