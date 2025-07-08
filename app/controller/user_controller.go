package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/skncvo/Scalable_API/app/service"
)

// 사용자 관련 컨트롤러 함수 목록 정의
type UserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// 내부에 userservice 포함, 모든 요청 서비스 계층에 위임
type UserControllerImpl struct {
	svc service.UserService
}

// 요청 처리 -> 서비스 호출
// GetAllUserData -> svc.GetAllUser ...
func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
	u.svc.GetAllUser(c)
}

func (u UserControllerImpl) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImpl) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

func (u UserControllerImpl) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}

// 의존성 주입을 위한 생성자
// 외부에서 userService 구현체를 주입받아 userControllerImpl 초기화
func UserControllerInit(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		svc: userService,
	}
}
