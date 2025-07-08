// wire 조건 빌드 태그, wire가 코드 생성 시에만 사용, 실행시 사용 x

package config

// Google wire 사용으로 의존성 주입 자동화
import (
	"github.com/google/wire"
	"github.com/skncvo/Scalable_API/app/controller"
	"github.com/skncvo/Scalable_API/app/repository"
	"github.com/skncvo/Scalable_API/app/service"
)

// wire.NewSet : 구성 요소의 그래프를 정의
var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet)
	return nil
}
