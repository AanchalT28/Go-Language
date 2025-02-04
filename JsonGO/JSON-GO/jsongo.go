package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Defining the structure of our weather data in Go is straightforward.
// Note that the json package only encodes struct fields that are public
// (and hence start with an uppercase letter).
// The JSON fields are all lowercase, so we need to map the struct field
// names to the corresponding JSON field names.
// Luckily, [Go structs come with a string tag feature](https://golang.org/ref/spec#Struct_types).
// This way we can tag every struct field with the corresponding JSON field name.
type weatherData struct {
	LocationName string   `json: locationName`
	Weather      string   `json: weather`
	Temperature  int      `json: temperature`
	Celsius      bool     `json: celsius`
	TempForecast []int    `json: temp_forecast`
	Wind         windData `json: wind`
}

type windData struct {
	Direction string `json: direction`
	Speed     int    `json: speed`
}

// Let's implement a tiny server application. The client sends its location, and
// the server responds by sending weather data.
//
// Location data is just a latitude and a longitude.
type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

// For the server, we need a function for handling the request
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// First, we need a location struct to receive the decoded data.
	location := loc{}

	// The location data is inside the request body which is an io.ReadCloser,
	// but we need a byte slice for unmarshalling.
	// ReadAll from ioutil just comes in handy.
	// Note we use ReadAll here for simplicity. Be careful when using ReadAll in larger
	// projects, as reading large files can consume a lot of memory.
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}
	// Now we can decode the request data using the Unmarshal function.
	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	// To see if the request was correctly received, let's print it to the console.
	log.Printf("Received: %v\n", location)

	// Now it's time to prepare our response by setting up a weatherData structure.
	// We could try fetching the data from a weather service, but for the purpose of
	// demonstrating JSON handling, let's just use some mock-up data.
	weather := weatherData{
		LocationName: "Zzyzx",
		Weather:      "cloudy",
		Temperature:  31,
		Celsius:      true,
		TempForecast: []int{30, 32, 29},
		Wind: windData{
			Direction: "S",
			Speed:     20,
		},
	}

	// For encoding the Go struct as JSON, we use the Marshal function from `encoding/json`.
	weatherJson, err := json.Marshal(weather)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	// We send a JSON response, so we need to set the Content-Type header accordingly.
	w.Header().Set("Content-Type", "application/json")

	// Sending the response is as easy as writing to the ResponseWriter object.
	w.Write(weatherJson)

}

// Thanks to Go's http package, starting the server is a piece of cake.
func server() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8088", nil)
}

// Our mock client is almost as simple as the server.
func client() {
	// Again we create JSON by marshalling a struct; in this case a loc struct literal.
	locJson, err := json.Marshal(loc{Lat: 35.14326, Lon: -116.104})
	// Then we set up a new HTTP request for posting the JSON data to local port 8080.
	req, err := http.NewRequest("POST", "http://localhost:8088", bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type", "application/json")

	// An HTTP client will send our HTTP request to the server and collect the response.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	// Finally, we print the received response and close the response body.
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

// The main function is as easy as it can get. We start the server in a goroutine
// and then run the client.
func main() {
	go server()
	client()
}
