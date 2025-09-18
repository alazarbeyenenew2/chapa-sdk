package service

import "fmt"

type Chapa struct {
	secretKey string
}

func NewClient(secretKey string) *Chapa {
	return &Chapa{
		secretKey: secretKey,
	}
}

func (c *Chapa) Test() string {
	return fmt.Sprintf("test")
}
