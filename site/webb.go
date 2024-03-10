package site

import (
	"os/exec"
	"runtime"
)

func Open(website string) error {
	var cmd *exec.Cmd
	var siteLink = "https://" + website + ".com"
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", siteLink)
	case "linux":
		cmd = exec.Command("xdg-open", siteLink)
	case "darwin":
		cmd = exec.Command("open", siteLink)
	}
	return cmd.Run()

}
