package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("无法获取可执行文件路径:", err)
		return
	}

	exeDir := filepath.Dir(exePath)
	go func() {
		fs := getFileSystem(useOS)
		r := gin.Default()
		r.GET("/", func(ctx *gin.Context) {
			ctx.FileFromFS("/", fs)
		})
		r.GET("/video/*any", func(ctx *gin.Context) {
			ctx.FileFromFS("/", fs)
		})
		r.GET("/rating/*any", func(ctx *gin.Context) {
			ctx.FileFromFS("/", fs)
		})
		r.GET("/favicon.ico", func(ctx *gin.Context) {
			ctx.FileFromFS("favicon.ico", fs)
		})
		r.GET("/assets/*any", func(ctx *gin.Context) {
			filename := ctx.Param("any")
			ctx.FileFromFS("assets/"+filename, fs)
		})
		r.GET("/videos/*any", func(ctx *gin.Context) {
			filename := ctx.Param("any")
			ctx.FileFromFS(filename, http.FS(os.DirFS(exeDir+"/videos")))
		})
		r.Run()
	}()

	openBrowser("http://127.0.0.1:8080")

	// 阻止主 goroutine 退出
	select {}
}

//go:embed static
var staticFiles embed.FS

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}

	log.Print("using embed mode")

	fsys, err := fs.Sub(staticFiles, "static")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}
