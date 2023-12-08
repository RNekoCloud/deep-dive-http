package datasource

import (
	"fmt"
	"testing"
)

func TestModify(t *testing.T) {
	data1 := City{
		Name:     "Solo",
		Province: "Sukoharjo",
	}

	Store["1"] = data1
	fmt.Println(Store)
}
