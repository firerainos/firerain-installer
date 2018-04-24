package config

type Account struct {
	Username string
	Password string
}

func (a *Account) SetUsername(username string) {
	a.Username = username
}

func (a *Account) SetPassword(password string) {
	a.Password = password
}
