package schema

import (
	"errors"
	"fmt"
	"log"
)

// 错误码规则 17***
var (
	// 通用错误码
	SUCCESS = errors.New("00000:ok")
	FAIL    = errors.New("99999:系统错误")

	// 业务错误码
)

func Panic(userDefinedErr error, args ...interface{}) {
	if userDefinedErr == nil {
		Panic(FAIL, "Panic is nil")
	}
	e := userDefinedErr.Error()
	errMsg := fmt.Sprintf(e, args...)
	log.Printf("Panic Exception： %s", errMsg)
	panic(errors.New(errMsg))
}
