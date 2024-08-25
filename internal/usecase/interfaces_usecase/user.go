package interfaces_usecase

type Hash interface {
	Encrypt(pass string) (string, error)
	Compare(pass, hash string) bool
}
