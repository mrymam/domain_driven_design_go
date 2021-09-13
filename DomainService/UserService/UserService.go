package UserService

import (
	"github.com/onyanko-pon/domain_driven_design_go/Entities/User"
)

type UserService struct{}

func (u UserService) Exists(user User.User) bool {
	// TODO DBアクセスなどする
	return false
}
