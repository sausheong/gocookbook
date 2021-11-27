package json

import (
	"bytes"
	"encoding/json"
	"testing"
)

var luke []byte = []byte(
	`{
	"name": "Luke Skywalker",
	"height": "172",
	"mass": "77",
	"hair_color": "blond",
	"skin_color": "fair",
	"eye_color": "blue",
	"birth_year": "19BBY",
	"gender": "male"
}`)

func BenchmarkUnmarshal(b *testing.B) {
	var person Person
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(luke, &person)
	}
}

func BenchmarkDecode(b *testing.B) {
	var person Person
	data := bytes.NewReader(luke)
	decoder := json.NewDecoder(data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(&person)
		data.Seek(0, 0)
	}
}
