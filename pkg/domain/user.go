package domain

type User struct {
	Id     ItemId     `json:"id,omitempty"`
	Name   string     `json:"name,omitempty"`
	Email  string     `json:"email,omitempty"`
	Gender UserGender `json:"gender,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

type ItemId uint16
type UserStatus string
type UserGender string

const (
	ActiveStatus   UserStatus = "active"
	InactiveStatus UserStatus = "inactive"
)

const (
	FemaleGender UserGender = "female"
	MaleGender   UserGender = "male"
)

type UserApiResponse interface {
	User | []User
}
