package json

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestUnmarshal(t *testing.T) {
	luke := unmarshal()
	if luke.Name != "Luke Skywalker" || luke.Starships[0] != "https://swapi.dev/api/starships/12/" {
		t.Error("Cannot unmarshal")
	}
}

func TestUnmarshalAPI(t *testing.T) {
	luke := unmarshalAPI()
	if luke.Name != "Luke Skywalker" || luke.Starships[0] != "https://swapi.dev/api/starships/12/" {
		t.Error("Cannot unmarshal with API")
	}
}

func TestUnstructured(t *testing.T) {
	unstruct := unstructured()
	vader := unstruct["Darth Vader"].([]interface{})
	if len(vader) != 4 {
		t.Error("Wrong number of elements")
	}
	if vader[0] != "https://swapi.dev/api/films/1/" {
		t.Error("Not parsed correctly")
	}
}

func TestUnmarshalStructArray(t *testing.T) {
	people := unmarshalStructArray()
	if len(people) != 3 {
		t.Error("Wrong number of people parsed")
	}
	if people[0].Name != "Luke Skywalker" {
		t.Error("Cannot unmarshal")
	}
}

func TestDecode(t *testing.T) {
	count := 0
	ch := make(chan Person)
	go decode(ch)
	for {
		p, ok := <-ch
		if ok {
			if !(p.Name == "Luke Skywalker" || p.Name == "C-3PO" || p.Name == "R2-D2") {
				t.Error("Did not decode person")
			}
			count++
			t.Log(fmt.Sprintf("%# v\n", pretty.Formatter(p)))
		} else {
			break
		}
	}
	if count != 3 {
		t.Error("Did not decode 3 elements")
	}
}
