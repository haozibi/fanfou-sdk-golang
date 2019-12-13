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

// IsDebug check debug mode
func IsDebug() bool {
	return strings.ToLower(os.Getenv("FANFOU_SDK_DEBUG")) == "true"
}
