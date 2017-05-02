package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/demizer/go-logs/src/logs"
)

var envs = make(map[string]string)

//goExec execute
func goExec(cmd string) bool {
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		logs.Println("unable to execute command: " + cmd)
		return false
	}
	return true
}

func init() {
	for _, env := range os.Environ() {
		option := strings.Split(env, "=")
		envs[option[0]] = option[1]
	}
}

func main() {

	if envs["idekey"] != "" && envs["remoteport"] != "" {
		remotePort := "xdebug.remote_port = \"" + envs["remoteport"] + "\""
		cmd := "echo '" + remotePort + "' >> /opt/xdebug.ini"
		if !goExec(cmd) {
			os.Exit(1)
		}

		idekey := "xdebug.idekey = \"" + envs["idekey"] + "\""
		cmd = "echo '" + idekey + "' >> /opt/xdebug.ini"
		if !goExec(cmd) {
			os.Exit(1)
		}

		if envs["path"] != "" {
			cmd = "cp /opt/newrelic.ini " + envs["path"] + "/.user.ini"
		} else {
			cmd = "cp /opt/newrelic.ini /var/www/web/.user.ini"
		}

		if !goExec(cmd) {
			os.Exit(1)
		}

	}

	os.Exit(0)
}
