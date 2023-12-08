package helper

import (
	"os/exec"
)

func UUIDGen() string {
	uuid, _ := exec.Command("uuidgen").Output()
	return string(uuid)
}
