package json

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []string  `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}

func unmarshal() (person Person) {
	file, err := os.Open("skywalker.json")
	if err != nil {
		log.Println("Error opening json file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading json data:", err)
	}

	err = json.Unmarshal(data, &person)
	if err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func unmarshalAPI() (person Person) {
	r, err := http.Get("https://swapi.dev/api/people/1")
	if err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading json data:", err)
	}

	err = json.Unmarshal(data, &person)
	if err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func unstructured() (output map[string]interface{}) {
	file, err := os.Open("unstructured.json")
	if err != nil {
		log.Println("Error opening json file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading json data:", err)
	}

	err = json.Unmarshal(data, &output)
	if err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func unmarshalStructArray() (people []Person) {
	file, err := os.Open("people.json")
	if err != nil {
		log.Println("Error opening json file:", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading json data:", err)
	}

	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func decode(p chan Person) {
	file, err := os.Open("people_stream.json")
	if err != nil {
		log.Println("Error opening json file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for {
		var person Person
		err = decoder.Decode(&person)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error decoding json data:", err)
			break
		}
		p <- person
	}
	close(p)
}

func marshal() {
	person := unmarshal()
	data, err := json.Marshal(&person)
	if err != nil {
		log.Println("Cannot marshal person:", err)
	}
	err = os.WriteFile("skywalker_marshalled.json", data, 0644)
	if err != nil {
		log.Println("Cannot write to file", err)
	}
}
