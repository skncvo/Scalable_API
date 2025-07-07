package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/skncvo/Scalable_API/app/constant"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, NULL()))
			c.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, NULL()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, NULL()))
			c.Abort()
		}
	}
}
