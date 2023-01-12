package validation

import (
	"errors"

	"github.com/eliassebastian/r6index-api/pkg/utils"
)

// check request body platform, name, and uuid validity
func All(platform, name, uuid string) error {
	if name != "" && uuid != "" {
		return errors.New("name and uid params both used")
	}

	if name == "" && uuid == "" {
		return errors.New("both name and uid params are empty")
	}

	if uuid != "" && !utils.IsValidUUID(uuid) {
		return errors.New("invalid player uuid provided")
	}

	if !utils.IsValidPlatform(platform) {
		return errors.New("invalid platform provided")
	}

	return nil
}
