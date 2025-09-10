package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mgq "github.com/christian-schueler/go-graphql/my_graphql"
)

func main() {
	var datapool []mgq.Location = mgq.GetExampleData()
	if datapool == nil {
		log.Fatalf("failed to load example data")
	}

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		result := mgq.ExecuteQuery(r.URL.Query().Get("query"), datapool)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
