package proxy

// User 用户
// @proxy IUser
type User struct {
}
type IUser interface {
	Login(username, password string) error
}
