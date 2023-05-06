package user_command_controller_dto

type CreateUserRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	LoginMethod string `json:"loginMethod"`
}
