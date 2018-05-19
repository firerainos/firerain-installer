package installer

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
)

type IpApi struct {
	Query    string `json:"query"`
	Timezone string `json:"timezone"`
}

func AutoSetTimezone() error {
	timezone, err := getTimezone()
	if err != nil {
		return err
	}

	if err := os.Symlink("/usr/share/zoneinfo/"+timezone, "/tmp/localtime"); err != nil {
		return err
	}

	cmd := exec.Command("hwclock", "--systohc")
	return cmd.Run()
}

func getTimezone() (string, error) {
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipApi := IpApi{}

	if err := json.NewDecoder(resp.Body).Decode(&ipApi); err != nil {
		return "", err
	}

	return ipApi.Timezone, nil
}
