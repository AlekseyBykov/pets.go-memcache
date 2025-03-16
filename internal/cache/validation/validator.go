package validation

import (
	"errors"
	"fmt"
	"github.com/AlekseyBykov/pets.go-memcache/internal/cache/consts"
	"github.com/AlekseyBykov/pets.go-memcache/internal/utils"
	"time"
)

const (
	minTTL = time.Second * 1
	maxTTL = time.Minute * 5
)

func ValidateKey(key string) error {
	if len(key) == 0 {
		return errors.New(consts.ErrKeyIsEmpty)
	}
	return nil
}

func ValidateValue(value any) error {
	if value == nil {
		return errors.New(consts.ErrValueIsEmpty)
	}
	return nil
}

func ValidateItemTtl(ttl time.Duration) error {
	if ttl < minTTL {
		return fmt.Errorf(consts.ErrTtlTooShortFmt,
			utils.FormatDuration(ttl, utils.TimeFormatMinutesSeconds),
			utils.FormatDuration(minTTL, utils.TimeFormatMinutesSeconds))
	}

	if ttl > maxTTL {
		return fmt.Errorf(consts.ErrTtlTooHighFmt,
			utils.FormatDuration(ttl, utils.TimeFormatMinutesSeconds),
			utils.FormatDuration(maxTTL, utils.TimeFormatMinutesSeconds))
	}

	return nil
}

func GetKeyNotFoundError(key string) error {
	return fmt.Errorf(consts.ErrKeyNotFoundFmt, key)
}

func GetItemExpiredError(key string) error {
	return fmt.Errorf(consts.ErrItemExpiredFmt, key)
}
