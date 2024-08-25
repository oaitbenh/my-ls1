package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func main() {
	Options := map[string]bool{}
	Options["l"] = true
	Files, _ := os.ReadDir("./")
	for i, File := range Files {
		info, _ := File.Info()
		if Options["l"] {
			fmt.Print(info.Mode(), " ")
			Sys := info.Sys().(*syscall.Stat_t)
			fmt.Print(Sys.Nlink, " ")
			User, _ := user.LookupId(strconv.Itoa(int(Sys.Uid)))
			Group, _ := user.LookupGroupId(strconv.Itoa(int(Sys.Gid)))
			fmt.Print(User.Username, " ")
			fmt.Print(Group.Name, " ")
			fmt.Print(info.Size(), " ")
			fmt.Print(info.ModTime().Month(), " ")
			fmt.Print(info.ModTime().Day(), " ")
			h, m, _ := info.ModTime().Clock()
			if h < 10 {
				fmt.Print("0")
			}
			fmt.Print(h, ":")
			if m < 10 {
				fmt.Print("0")
			}
			fmt.Print(m, " ")
		}
		fmt.Print(File.Name())
		if Options["l"] || i == len(Files)-1 {
			fmt.Println()
		} else {
			fmt.Print("  ")
		}
	}
}

