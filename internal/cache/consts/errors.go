package consts

const (
	ErrKeyIsEmpty     = "key is null or empty"
	ErrKeyNotFoundFmt = "key not found, key = %q"

	ErrValueIsEmpty = "value is null or empty"

	ErrItemExpiredFmt = "item expired, key = %q"

	ErrTtlTooShortFmt = "ttl %q is too short, min = %q"
	ErrTtlTooHighFmt  = "ttl %q is too high, max = %q"
)
