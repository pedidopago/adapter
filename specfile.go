package adapter

type Spec struct {
	Name      string `json:"name" toml:"name"`
	SliceName string `json:"slice_name" toml:"slice_name"`

	Src      string      `json:"src" toml:"src"`
	Dst      string      `json:"dst" toml:"dst"`
	Fields   []SpecField `json:"fields" toml:"field"`
	DstIsPtr bool        `json:"-" toml:"-"`
}

type SpecField struct {
	From string `json:"from" toml:"from"`
	To   string `json:"to" toml:"to"`
	Use  string `json:"use" toml:"use"`
}

type Specs struct {
	Items    []Spec   `json:"specs" toml:"spec"`
	Packages []string `json:"packages" toml:"packages"`
}
