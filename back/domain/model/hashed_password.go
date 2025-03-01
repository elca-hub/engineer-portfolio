package model

type HashedPassword struct {
	hashedPassword string
}

func NewHashedPassword(hashedPassword string) *HashedPassword {
	return &HashedPassword{hashedPassword: hashedPassword}
}

func (h *HashedPassword) HashedPassword() string {
	return h.hashedPassword
}
