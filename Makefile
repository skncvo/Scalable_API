init-dependency:
	# 로그 출력을 트리 구조로 보기 좋게 해주는 Logrus 포맷터
	go get -u github.com/antonfisher/nested-logrus-formatter
	
	# Gin 웹 프레임워크
	go get -u github.com/gin-gonic/gin
	
	# 보안 및 암호화 기능 제공
	go get -u golang.org/x/crypto
	
	# ORM(Object Relational Mapping) 라이브러리
	go get -u gorm.io/gorm
	
	# gorm에서 postgresql을 쓰기위한 드라이버 
	go get -u gorm.io/driver/postgres
	
	# 로깅 라이브러리
	go get -u github.com/sirupsen/logrus
	
	# .env 파일에서 환경변수를 불러와 사용할 수 있게 해주는 도구
	go get -u github.com/joho/godotenv