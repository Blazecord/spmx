package utils

import (
	"encoding/json"
	"net/http"
)

type PackageJson struct {
	Version string `json:"version"`
}

func GetVersion() (string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/Blazecord/blaze/master/package.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var packageInfo PackageJson
	err = json.NewDecoder(resp.Body).Decode(&packageInfo)
	if err != nil {
		return "", err
	}

	version := packageInfo.Version

	return version, err
}