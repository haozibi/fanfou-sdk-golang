package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// ShowInfomation show defail information
func ShowInfomation(name string, info ...interface{}) {

	fmt.Printf("=== %s ===\n", name)
	spew.Dump(info...)
	fmt.Println("")
}

const (
	debugFlag = "FANFOU_SDK_DEBUG"
)

// IsDebug check debug mode
func IsDebug() bool {
	return strings.ToLower(os.Getenv(debugFlag)) == "true"
}
