package model

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	MinPasswordLen = 8
	MaxPasswordLen = 64
)

type RawPassword struct {
	rawPassword string
}

func NewRawPassword(rawPassword string) (*RawPassword, error) {
	if len(rawPassword) < MinPasswordLen {
		return nil, errors.New(fmt.Sprintf("The password must be at least %d characters long.", MinPasswordLen))
	}

	if len(rawPassword) > MaxPasswordLen {
		return nil, errors.New(fmt.Sprintf("The password must be at most %d characters long.", MaxPasswordLen))
	}

	// パスワードの正規表現
	// 8文字以上の半角英数字と. + - [ ] * ~ _ # : ?を必ず含む
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(rawPassword)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(rawPassword)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(rawPassword)
	hasSymbol := regexp.MustCompile(`[.+\-[\]*~_#:?]`).MatchString(rawPassword)

	if !hasLower || !hasUpper || !hasNumber || !hasSymbol {
		return nil, errors.New("the password must contain at least one uppercase letter, one lowercase letter, one number, and one symbol")
	}

	return &RawPassword{rawPassword: rawPassword}, nil
}

func (r *RawPassword) RawPassword() string {
	return r.rawPassword
}
