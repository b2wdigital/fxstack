package util

import (
	"fmt"

	giconfig "github.com/b2wdigital/goignite/config"
)

func GetStringConfigOrPanic(configPath string) string {
	s := giconfig.String(configPath)
	if s == "" {
		panic(fmt.Sprintf("Config %s is not set", configPath))
	}
	return s
}
