package datasource

type City struct {
	Name     string `json:"name,omitempty"`
	Province string `json:"province,omitempty"`
}

var Store map[string]City
