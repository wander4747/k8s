package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/configmap-volume", ConfigMapVolume)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":8000", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func ConfigMapVolume(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/src/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Password: %s", user, password)
}
