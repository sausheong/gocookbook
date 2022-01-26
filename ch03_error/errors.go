package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

func anotherUnmarshal() (person Person) {
	var r *http.Response
	var data []byte
	var err error

	if r, err = http.Get("https://swapi.dev/api/people/1"); err != nil {
		log.Println("Cannot get from URL", err)
	}
	defer r.Body.Close()

	if data, err = io.ReadAll(r.Body); err != nil {
		log.Println("Error reading json data:", err)
	}

	if err = json.Unmarshal(data, &person); err != nil {
		log.Println("Error unmarshalling json data:", err)
	}
	return
}

func helperUnmarshal() (person Person) {
	r, err := http.Get("https://swapi.dev/api/people/1")
	check(err, func() {})
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	check(err, func() {})

	err = json.Unmarshal(data, &person)
	check(err, func() {})
	return
}

func check(err error, handler func()) {
	if err != nil {
		handler()
	}
}

func must(arg interface{}, err error) interface{} {
	if err != nil {
		fmt.Println("error:", err)
	}
	return arg
}

func mustUnmarshal() (person Person) {
	r := must(http.Get("https://swapi.dev/api/people/1")).(*http.Response)
	defer r.Body.Close()
	data := must(io.ReadAll(r.Body)).([]byte)
	must(nil, json.Unmarshal(data, &person))
	return
}

type unmarshaller struct {
	response *http.Response
	data     []byte
	err      error
	person   Person
}

func (u *unmarshaller) get(url string) {
	u.response, u.err = http.Get(url)
}

func (u *unmarshaller) read() {
	if u.err != nil {
		return
	}
	u.data, u.err = io.ReadAll(u.response.Body)
}

func (u *unmarshaller) unmarshal() {
	if u.err != nil {
		return
	}
	u.err = json.Unmarshal(u.data, &u.person)
}

type CommsError struct{}

func (m CommsError) Error() string {
	return "An error happened during data transfer"
}

func send(data []byte) error {
	return &CommsError{}
}

type SyntaxError struct {
	Line int
	Col  int
}

func (err *SyntaxError) Error() string {
	return fmt.Sprintf("Error at line %d, column %d", err.Line, err.Col)
}

func run() error {
	return &SyntaxError{
		Line: 119,
		Col:  45,
	}
}

type ConnectionError struct {
	Host string
	Port int
	Err  error
}

func (err *ConnectionError) Error() string {
	return fmt.Sprintf("Error connecting to %s at port %d", err.Host, err.Port)
}

func (err *ConnectionError) Unwrap() error {
	return err.Err
}

func connectAPI() error {
	return ApiErr
}

func connect() error {
	return &ConnectionError{
		Host: "localhost",
		Port: 8080,
		Err:  ApiErr,
	}
}

var ApiErr error = errors.New("Error trying to get data from API")

func A() {
	defer fmt.Println("defer on A")
	fmt.Println("A")
	B()
	fmt.Println("Unreachable in A")
}

func B() {
	defer fmt.Println("defer on B")
	fmt.Println("B")
	C()
	fmt.Println("Unreachable in B")
}

func C() {
	defer fmt.Println("defer on C")
	fmt.Println("C")
	panic("panicked in C")
	fmt.Println("Unreachable in C")
}
