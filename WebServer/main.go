package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pluralsight/gobigpicture/module"
)

func main() {

	user := module.User{
		Id:        0,
		FirstName: "tris",
		LastName:  "mack",
	}

	fmt.Println(user)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string

		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)

		enc.Encode(m)
		//w.Write([]byte("Hello " + name))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
