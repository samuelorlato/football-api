package entities

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

func (r *RegisterRequest) ToUserEntity(registerRequest RegisterRequest, ID string, passwordHash string) User {
	return User{
		ID:           ID,
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		PasswordHash: passwordHash,
	}
}
