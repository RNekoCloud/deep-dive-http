package datasource

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/RNekoCloud/deep-dive-http/helper"
)

func TestModify(t *testing.T) {
	data1 := City{
		Name:     "Solo",
		Province: "Jawa Tengah",
	}

	id := helper.UUIDGen()

	Store[id] = data1

	for _, val := range Store {
		fmt.Println("City:", val.Name)
	}
}

func TestJSONMap(t *testing.T) {
	b, err := json.Marshal(Store)

	if err != nil {
		t.Error(err)
	}

	data := []City{}

	for _, val := range Store {
		data = append(data, val)
	}

	fmt.Println(string(b))
	fmt.Println("All:", data)
}
