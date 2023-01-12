package ubisoft

import "fmt"

func profileUri(name, uuid, platform string) string {
	if name != "" {
		return fmt.Sprintf("https://public-ubiservices.ubi.com/v3/profiles?namesOnPlatform=%s&platformType=%s", name, platform)
	}

	return fmt.Sprintf("https://public-ubiservices.ubi.com/v3/users/%s/profiles?platformType=%s}", uuid, platform)
}
