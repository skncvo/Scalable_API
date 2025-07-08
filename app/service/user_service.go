package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/skncvo/Scalable_API/app/constant"
	"github.com/skncvo/Scalable_API/app/domain/dao"
	"github.com/skncvo/Scalable_API/app/pkg"
	"github.com/skncvo/Scalable_API/app/repository"
	"golang.org/x/crypto/bcrypt"
)

// 컨트롤러가 이 인터페이스만 보고 의존하도록 (DI구조)
type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

// 실제 서비스 구현체
// 비즈니스 로직에서 DB에 접근할 때 이를 통해 접근
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// 모든 유저 조회
// 실패 시 UnknownError
// 성공 시 JSON 응답 반환
func (u UserServiceImpl) GetAllUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get all data user")

	data, err := u.userRepository.FindAllUser()
	if err != nil {
		log.Error("Happened Error when find all user data. Error: ", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

// URL에서 ID 파싱 -> 조회
// 없으면 DataNotFound
// 성공 시, JSON 반환
func (u UserServiceImpl) GetUserById(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program get user by id")
	// URL에서 ID 파싱
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

// 요청 body를 user 구조체로 바인딩
// bcrypt로 비밀번호 암호화
// 저장 후 성공 응답
func (u UserServiceImpl) AddUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program add data user")
	var request dao.User
	// shoudbind : http의 요청의 데이터를 지정된 Go 구조체에 바인딩하고 유효성 검사
	// 바인딩 : 데이터를 특정위치에 동적으로 연결하는 것
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	// bcrypt : 암호 해싱 기능, salt값을 넣어 hash값 생성
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	// 저장
	data, err := u.userRepository.Save(&request)
	if err != nil {
		log.Error("Happened error when saving data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

// ID로 기존 유저 조회
// 필드 덮어쓰기 후 저장
// 성공 응답
func (u UserServiceImpl) UpdateUserData(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute program update user data by id")
	// URL에서 ID 파싱
	userID, _ := strconv.Atoi(c.Param("userID"))

	// 요청 body를 구조체 user로 바인딩
	var request dao.User
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request from FE. Error", err)
		pkg.PanicException(constant.InvalidRequest)
	}

	// userID로 조회, 없으면 패닉
	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		log.Error("Happened error when get data from database. Error", err)
		pkg.PanicException(constant.DataNotFound)
	}

	// 필드 덮어쓰기 후 저장
	data.RoleID = request.RoleID
	data.Email = request.Email
	data.Name = request.Name
	data.Status = request.Status
	u.userRepository.Save(&data)

	if err != nil {
		log.Error("Happened error when updating data to database. Error", err)
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

// URL에서 ID 파싱
// ID를 기반으로 삭제
// 성공 응답
func (u UserServiceImpl) DeleteUser(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute delete data user by id")
	// ID 파싱
	userID, _ := strconv.Atoi(c.Param("userID"))

	// ID기반 삭제
	err := u.userRepository.DeleteUserById(userID)
	if err != nil {
		log.Error("Happened Error when try delete data user from DB. Error:", err)
		pkg.PanicException(constant.UnknownError)
	}

	// 성공 응답
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.NULL()))
}

// 생성자
// DI 방식으로 repository를 주입 받아 서비스 생성
func UserServiceInit(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}
