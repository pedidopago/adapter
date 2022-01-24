package atob

import (
	"strconv"
)

// for production, use "go:generate go run github.com/pedidopago/adapter/cmd/adapter -f atob.toml"
//go:generate go run ./../../cmd/adapter/main.go -f atob.toml

type UserA struct {
	Name  string
	Age   int
	Score float64
	Props []PropA
}

type PropA struct {
	Name  string
	Value string
}

type UserB struct {
	Name     string
	Age      int8
	MaxScore string
	Props    []PropB
}

type PropB struct {
	Name  string
	Value int
}

func f64tostr(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', -1, 64)
}

func atoim(i string) int {
	v, _ := strconv.Atoi(i)
	return v
}
