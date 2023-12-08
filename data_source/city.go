package datasource

type City struct {
	Name     string `json:"name,omitempty"`
	Province string `json:"province,omitempty"`
}

var baseData = City{
	Name:     "Sukoharjo",
	Province: "Jawa Tengah",
}

var Store map[string]City = map[string]City{
	"01": baseData,
}
