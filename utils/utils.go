package utils

import (
	"os"
	"encoding/json"
)

type PackageJson struct {
	Version string `json:"version"`
}

func GetVersion() string {
	packageJsonPath := "~/blaze/package.json"

	packageJsonPath, err := expandTilde(packageJsonPath)

	if err != nil {
		return ""
	}

    packageJsonData, err := os.ReadFile(packageJsonPath)

    if err != nil {
        return ""
    }

	var packageInfo PackageJson

	if err := json.Unmarshal(packageJsonData, &packageInfo); err != nil {
		return ""
	}

	version := packageInfo.Version

	return version
}

func expandTilde(path string) (string, error) {
    if path[:2] == "~/" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            return "", err
        }
        return homeDir + path[1:], nil
    }
    return path, nil
}