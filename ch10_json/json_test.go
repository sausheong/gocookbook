package json

import (
	"testing"
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
