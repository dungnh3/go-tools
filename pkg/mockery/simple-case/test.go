package main

type Stringer interface {
	String(string, ...interface{}) string
}

type SendFunc func(data string) (int, error)
