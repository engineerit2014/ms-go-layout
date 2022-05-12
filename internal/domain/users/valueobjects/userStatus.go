package vo

// UserStatus is a value object that represents only a concept with no identity
type UserStatus string

const (
	ActiveUserStatus UserStatus = "ACTIVE"
)
