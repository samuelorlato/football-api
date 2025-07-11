package entities

type RegisterFanRequest struct {
	Name  string
	Email string
	Team  string
}

func (r *RegisterFanRequest) ToEntity(ID string) Fan {
	return Fan{
		ID:    ID,
		Name:  r.Name,
		Email: r.Email,
		Team:  r.Team,
	}
}
