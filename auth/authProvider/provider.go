package authProvider

type AuthProvider interface {
	LoginUser(string, string) (string, error)
	VerifyToken(string) bool
}
