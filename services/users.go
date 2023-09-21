package services

import (
	"restaurant/errors"
	"restaurant/repositorys"

	"github.com/gofiber/fiber/v2"
)

func NewUserService(userRepo repositorys.UsersRepository) UserService {
	return userService{userRepo}
}

func (s userService) GetUsers(c *fiber.Ctx) (*ResponseUsers, error) {

	// Repository
	users, err := s.userRepo.GetUsers(c)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}


	var data []DataResponseUsers
	for _, item := range users.Data {
		responseEvent := DataResponseUsers{
			ID:         item.ID,
			FirstName:  item.FirstName,
			LastName:   item.LastName,
		}
		data = append(data, responseEvent)
	}

	user := &ResponseUsers{
		Data:       data,
		Messages:   users.Messages,
	}



	
	return user, nil
}

func (s userService) PutUsers(c *fiber.Ctx) (*repositorys.ResponseUsers,error) {
	// Repository
	users, err := s.userRepo.PutUsers(c)
	if err != nil {

		return nil, errors.NewUnexpectedError(err.Error())
	}
	return users, nil
}

func (s userService) DeleteUsers(c *fiber.Ctx) (*repositorys.ResponseUsers,  error) {
	// Repository
	users, err := s.userRepo.DeleteUsers(c)
	if err != nil {
		return nil, errors.NewUnexpectedError(err.Error())
	}
	return users, nil
}
