package config

import (
	"github.com/skncvo/Scalable_API/app/controller"
	"github.com/skncvo/Scalable_API/app/repository"
	"github.com/skncvo/Scalable_API/app/service"
)

// 의존성 구조체
type Initialization struct {
	// 사용자 DB 처리
	userRepo repository.UserRepository
	// 사용자 비즈니스 로직
	userSvc service.UserService
	// 사용자 HTTP 핸들러
	UserCtrl controller.UserController
	// role 관련 DB 처리
	RoleRepo repository.RoleRepository
}

// 외부에서 구성요소 주입 받아 의존성 구조체 Initialization 생성
func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
		RoleRepo: roleRepo,
	}
}
