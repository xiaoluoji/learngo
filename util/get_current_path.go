package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

//noinspection GoUnusedFunction
func GetCurrentPath() {
	fmt.Println("getTmpDir（当前系统临时目录） = ", os.TempDir())
	fmt.Println("getCurrentAbPathByExecutable（仅支持go build） = ", getCurrentAbPathByExecutable())
	fmt.Println("getCurrentAbPathByCaller（仅支持go run） = ", getCurrentAbPathByCaller())
	fmt.Println("getCurrentAbPath（最终方案-全兼容） = ", getCurrentAbPath())
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

//获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
