package ports

type EncryptionService interface {
	HashPassword(password string) (*string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}
