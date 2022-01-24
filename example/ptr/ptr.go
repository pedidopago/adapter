package ptr

//go:generate go run ./../../cmd/adapter/main.go -f convert.toml

type A struct {
	X int
}

type B struct {
	X2 int
}
