package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func main() {
	dir, _ := os.ReadDir("/home/oaitbenh/Desktop/")
	Total(dir, "/home/oaitbenh/Desktop/")
	fmt.Println()
}

func Total(Dir []fs.DirEntry, Path string) string {
	fmt.Println(Path, ":")
	var Dirs []fs.DirEntry
	for _, file := range Dir {
		if file.IsDir() {
			if strings.HasPrefix(file.Name(), ".") {
				continue
			}
			fmt.Print(Blue + file.Name() + Reset + "  ")
			Dirs = append(Dirs, file)
		} else {
			fmt.Print(file.Name(), "  ")
		}
	}
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(Dirs); i++ {
		dir, _ := os.ReadDir(Path + Dirs[i].Name())
		Total(dir, Path+Dirs[i].Name()+"/")
	}
	return Path
}
