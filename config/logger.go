package config

import (
	// 구조화된 로그 출력을 위한 nested formatter
	nested "github.com/antonfisher/nested-logrus-formatter"

	// .env 파일을 로딩하여 환경 변수로 등록하는 패키지
	"github.com/joho/godotenv"

	// logrus: Go용 고급 로깅 라이브러리
	log "github.com/sirupsen/logrus"

	// 환경 변수 및 OS 관련 함수 제공
	"os"
)

// InitLog은 logrus를 초기화하는 함수로, 로그 레벨, 포맷, 호출자 정보 등을 설정함
func InitLog() {
	// .env 파일을 로드하여 os.Getenv에서 사용할 수 있도록 함
	// 실패해도 무시 (예: .env 파일이 없을 수도 있음)
	_ = godotenv.Load()

	// 환경 변수 LOG_LEVEL 값을 기반으로 로그 레벨 설정
	log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))

	// 로그 호출자 정보 (파일명과 라인 넘버 등)를 포함
	log.SetReportCaller(true)

	// 로그 출력 포맷 설정 (nested-logrus-formatter 사용)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,                              // 필드 이름을 숨김 (component=value → value)
		FieldsOrder:     []string{"component", "category"}, // 출력할 필드 순서 지정
		TimestampFormat: "2006-01-02 15:04:05",             // 타임스탬프 포맷 (Go 고유 형식)
		ShowFullLevel:   true,                              // 로그 레벨을 대문자로 출력 (INFO, DEBUG 등)
		CallerFirst:     true,                              // 호출자 정보를 로그의 가장 앞에 표시
	})
}

// getLoggerLevel은 문자열로 받은 로그 레벨 값을 logrus의 log.Level로 변환
func getLoggerLevel(value string) log.Level {
	switch value {
	case "DEBUG":
		return log.DebugLevel
	case "TRACE":
		return log.TraceLevel
	}
	// 기본값은 InfoLevel
	return log.InfoLevel
}
