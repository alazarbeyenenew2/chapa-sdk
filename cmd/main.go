package main

import "fmt"

type chapa struct {
}

func NewClient() *chapa {
	return &chapa{}
}

func (c *chapa) Test() string {
	return fmt.Sprintf("test")
}
