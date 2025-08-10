package core

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Banners() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()

	fmt.Print(FR + `                                                                                                                                                                                                                                                                                                                                                                                 
█████      █████      ████     ████    ████    ███     
██████    ██████     ███████    ███   ██████  ████     
███ ███  ███ ███   ████  ████    ███ ███ ███ ████      
███  █████   ███  ████████████    █████   ██████       
███   ███    ███ ████       ███   ████     ████ 3six     
                                                                                                                                                                                                  
==== [Shell Finder Advance With Big List Path] ====
`)
	fmt.Println(FC + "[Maw - Scanner] - " + FG + "Nyari Webshell")
}
