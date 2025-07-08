package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/skncvo/Scalable_API/app/domain/dao"
	"gorm.io/gorm"
)

// 서비스 계층에서 사용할 수 있는 메서드 시그니처 정의
type UserRepository interface {
	FindAllUser() ([]dao.User, error)
	FindUserById(id int) (dao.User, error)
	Save(user *dao.User) (dao.User, error)
	DeleteUserById(id int) error
}

// GORM의 DB 핸들러를 내부에 갖고 있고, 모든 동작은 이 인스턴스로 수행됨
type UserRepositoryImpl struct {
	db *gorm.DB
}

// 모든 사용자 조회
func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
	var users []dao.User

	// user 구조체에 관련된 Role에 조인
	var err = u.db.Preload("Role").Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return users, nil
}

// 특정 사용자 ID로 조회
func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
	user := dao.User{
		ID: id,
	}
	// Role 정보도 함께 조회
	err := u.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return dao.User{}, err
	}
	return user, nil
}

// 유저 정보 저장
func (u UserRepositoryImpl) Save(user *dao.User) (dao.User, error) {
	// DB에 user정보 save
	var err = u.db.Save(user).Error
	if err != nil {
		log.Error("Got an error when save user. Error: ", err)
		return dao.User{}, err
	}
	// 성공 시 저장된 유저 구조체 반환
	return *user, nil
}

// ID 기반 삭제
func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&dao.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete user. Error: ", err)
		return err
	}
	return nil
}

// 생성자
// 외부에서 *gorm.DB 주입
func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	// AutoMigrate : user테이블이 없으면 생성
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
