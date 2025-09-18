package service

import "fmt"

type Chapa struct {
	Username string
	Password string
}

func NewClient(username, password string) *Chapa {
	// test response
	return &Chapa{
		Username: username,
		Password: password,
	}

}

func (c *Chapa) Test() string {
	return fmt.Sprintf("username %s password %s", c.Username, c.Password)
}
