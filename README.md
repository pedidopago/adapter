Adapter
=======

```toml
# atob.toml
[[spec]]
    name = "UserAtoB"
    src = "UserA"
    dst = "UserB"
    [[spec.field]]
        from = "Name"
    [[spec.field]]
        from = "Age"
        use = "int8"
    [[spec.field]]
        from = "Score"
        to = "MaxScore"
        use = "f64tostr"
    [[spec.field]]
        from = "Props"
        use = "PropAToBSlice"
[[spec]]
    name = "PropAToB"
    src = "PropA"
    dst = "PropB"
    [[spec.field]]
        from = "Name"
    [[spec.field]]
        from = "Value"
        use = "atoim"
```

```go
// atob.go

package atob

import (
	"strconv"
)

//go:generate go run github.com/pedidopago/adapter/cmd/adapter -f atob.toml

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
```

Generates:

```go
func UserAtoB(src UserA) (dst UserB)
func UserAtoBSlice(src []UserA) (dst []UserB)
func PropAToB(src PropA) (dst PropB)
func PropAToBSlice(src []PropA) (dst []PropB)
```