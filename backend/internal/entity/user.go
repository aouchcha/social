package entity

type UserProfile struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	BirthDate string
	Avatar    string
	About     string
	Status    bool
}
