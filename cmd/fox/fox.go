package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/StevenZack/openurl"

	"github.com/StevenZack/tools/fileToolkit"

	"github.com/StevenZack/tools/strToolkit"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("not enough args")
		fmt.Println("examples:")
		fmt.Println("1. fox build_386 // build fox module for x86 platform")
		fmt.Println("2. fox build_arm64")
		fmt.Println("3. fox build_arm")
		fmt.Println("4. fox build_arm*")
		fmt.Println("\n5. fox new // create new fox project")
		fmt.Println("6. fox update //")
		fmt.Println("7. fox build_apk  // build debug apk")
		return
	}
	gopath := fileToolkit.Getrpath(os.Getenv("GOPATH")) + "src"
	w, e := os.Getwd()
	if e != nil {
		fmt.Println("somewhere got error:", e)
		return
	}
	if !strToolkit.StartsWith(w, gopath) {
		fmt.Println("current path is not in GOPATH")
		return
	}
	curPath := w[len(fileToolkit.Getrpath(gopath)):]
	var c *exec.Cmd
	switch args[0] {
	case "build_386":
		c = exec.Command("gomobile", "bind", "--target=android/386", "-o", "android/fox/fox.aar", curPath)
	case "build_arm64":
		c = exec.Command("gomobile", "bind", "--target=android/arm64", "-o", "android/fox/fox.aar", curPath)
	case "build_arm":
		c = exec.Command("gomobile", "bind", "--target=android/arm", "-o", "android/fox/fox.aar", curPath)
	case "build_arm*":
		c = exec.Command("gomobile", "bind", "--target=android/arm,android/arm64", "-o", "android/fox/fox.aar", curPath)
	case "build_apk":
		e = os.Chdir("android")
		if e != nil {
			fmt.Println(`chdir error :`, e)
			return
		}
		build := exec.Command("./gradlew", "assembleDebug")
		build.Stdin = os.Stdin
		build.Stderr = os.Stderr
		build.Stdout = os.Stdout
		e = build.Run()
		if e != nil {
			fmt.Println(`./gradlew error :`, e)
			return
		}
		os.Chdir("..")
		fmt.Println("OK")
		openurl.Open("./android/app/build/outputs/apk/debug")
		return
	case "new":
		c = exec.Command("cp", "-r", fileToolkit.Getrpath(gopath)+"github.com/gofaith/fox", ".")
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		e = c.Run()
		if e != nil {
			fmt.Println(`new error :`, e)
			return
		}
		return
	case "update":
		fs, e := listDir(fileToolkit.Getrpath(gopath) + "github.com/gofaith/fox/")
		if e != nil {
			fmt.Println(`ls error :`, e)
			return
		}
		for _, v := range fs {
			c = exec.Command("cp", "-r", fileToolkit.Getrpath(gopath)+"github.com/gofaith/fox/"+v, ".")
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			e = c.Run()
			if e != nil {
				fmt.Println(`command.Run error :`, e)
				continue
			}
		}
		return
	}
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	e = c.Run()
	if e != nil {
		fmt.Println(`command.Run() error :`, e)
		return
	}
}
func listDir(dir string) ([]string, error) {
	c := exec.Command("ls", dir)
	buf := bytes.NewBufferString("")
	c.Stdout = buf
	e := c.Run()
	if e != nil {
		return nil, e
	}
	ss := strings.Split(buf.String(), "\n")
	return ss[:len(ss)-2], nil
}
