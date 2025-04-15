package user

import "github.com/jeissoni/EventLine/internal/ports"

type UserHandler struct {
	UserService ports.UserService
}
