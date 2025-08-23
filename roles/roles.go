package roles

const (
	Guest = "guest" // guest
	User  = "user"
	Admin = "admin"
)

func IsValid(role string) bool {
	switch role {
	case Guest, User, Admin:
		return true
	default:
		return false
	}
}
