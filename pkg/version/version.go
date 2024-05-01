package version

import (
	"fmt"
	"os"
	"runtime"
)

var (
	GoVersion  = runtime.Version()
	CommitId   string
	BranchName string
	BuildTime  string
	AppVersion string
)

func PrintVersion() {
	fmt.Printf("go version: %s\r\n", GoVersion)
	fmt.Printf("git commit ID: %s\r\n", CommitId)
	fmt.Printf("git branch name: %s\r\n", BranchName)
	fmt.Printf("app build xtime: %s\r\n", BuildTime)
	fmt.Printf("app version: %s\r\n", AppVersion)
	// 打印完退出
	os.Exit(0)
}
