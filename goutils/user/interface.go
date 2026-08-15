package user

// MinimalisticUserI is general interface for user data
type MinimalisticUserI interface {
	GetLogin() *string
	GetPassword() *string
}
