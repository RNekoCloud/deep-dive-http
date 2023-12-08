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
	"553f3b9d-1189-4e22-80af-1f0d0d946a2f": baseData,
}
