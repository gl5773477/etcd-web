package config

import (
	"errors"
)

const (
	ERR_CODE_INTERNEL = 0x0001
)

var (
	ERR_MSG_INTERNEL = errors.New("内部错误")
)
