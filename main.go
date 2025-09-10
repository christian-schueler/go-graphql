package main

import (
	"encoding/json"
	"fmt"
	"log"

	mgq "github.com/christian-schueler/go-graphql/my_graphql"
	gq "github.com/graphql-go/graphql"
)

func main() {
	var datapool []mgq.LocationType = mgq.GetExampleData()
	if datapool == nil {
		log.Fatalf("failed to load example data")
	}
	var schema gq.Schema = mgq.GetSchema(datapool)
	// Query
	query := `
		{
			id,
			name,
			parent
		}
	`
	params := gq.Params{Schema: schema, RequestString: query}
	r := gq.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON) // {"data":{"hello":"world"}}
}
