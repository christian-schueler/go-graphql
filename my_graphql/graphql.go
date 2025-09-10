package my_graphql

import (
	"fmt"
	"log"

	gq "github.com/graphql-go/graphql"
)

type Location struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

func GetExampleData() []Location {
	var locations = []Location{
		{
			Id:     "d3f78c69-04b8-4821-a2e5-92ed7f6d419a",
			Name:   "Planet Erde",
			Parent: "",
		},
		{
			Id:     "1fd351c6-ed6b-4946-a099-f45ae79ea320",
			Name:   "Europa",
			Parent: "d3f78c69-04b8-4821-a2e5-92ed7f6d419a",
		},
		{
			Id:     "87de2ea0-3952-4df1-afb1-9eb7e17ca99a",
			Name:   "Asien",
			Parent: "d3f78c69-04b8-4821-a2e5-92ed7f6d419a",
		},
		{
			Id:     "f16e4ae2-1e2a-4cbb-8004-cc7b207d0ce0",
			Name:   "Deutschland",
			Parent: "1fd351c6-ed6b-4946-a099-f45ae79ea320",
		},
		{
			Id:     "a815bf06-75ca-490e-8244-e0cf9e22482d",
			Name:   "Thüringen",
			Parent: "f16e4ae2-1e2a-4cbb-8004-cc7b207d0ce0",
		},
		{
			Id:     "2da7d16f-c603-4626-b94a-c2aa0ceafc95",
			Name:   "Bayern",
			Parent: "f16e4ae2-1e2a-4cbb-8004-cc7b207d0ce0",
		},
		{
			Id:     "9fc6f442-1735-405e-846a-ae237884024a",
			Name:   "Landkreis Hildburghausen",
			Parent: "a815bf06-75ca-490e-8244-e0cf9e22482d",
		},
		{
			Id:     "a5583da3-c658-46cf-8d74-2d91e0a45cb1",
			Name:   "Landkreis Coburg",
			Parent: "2da7d16f-c603-4626-b94a-c2aa0ceafc95",
		},
		{
			Id:     "12a04f4f-6f41-46a5-937b-78175e8c4885",
			Name:   "Stadt Hildburghausen",
			Parent: "9fc6f442-1735-405e-846a-ae237884024a",
		},
		{
			Id:     "15148d92-c58c-42d9-8866-03a3954dfbf4",
			Name:   "Birkenfelder Straße",
			Parent: "12a04f4f-6f41-46a5-937b-78175e8c4885",
		},
		{
			Id:     "cfa36886-6d0a-40cf-99d5-795d6a5fa1e4",
			Name:   "50",
			Parent: "15148d92-c58c-42d9-8866-03a3954dfbf4",
		},
	}

	return locations
}

func GetSchema(dataPool []Location) gq.Schema {
	var LocationGqlType = gq.NewObject(
		gq.ObjectConfig{
			Name: "Location",
			Fields: gq.Fields{
				"id": &gq.Field{
					Type: gq.String,
				},
				"name": &gq.Field{
					Type: gq.String,
				},
				"parent": &gq.Field{
					Type: gq.String,
				},
			},
		},
	)

	var queries = gq.NewObject(
		gq.ObjectConfig{
			Name: "Query",
			Fields: gq.Fields{
				/* Get (read) a single location by Id
				   https://localhost:8080/location?query={location(id:"xxxx"){id,name,parent}}
				*/
				"location": &gq.Field{
					Type:        LocationGqlType,
					Description: "Get a location by Id (GUID)",
					Args: gq.FieldConfigArgument{
						"id": &gq.ArgumentConfig{
							Type: gq.String,
						},
					},
					Resolve: func(rp gq.ResolveParams) (interface{}, error) {
						var id string = rp.Args["id"].(string)
						if id != "" {
							for _, location := range dataPool {
								if location.Id == id {
									return location, nil
								}
							}
						}

						return nil, nil
					},
				},
				/* Get (read) a single location by Name
				   https://localhost:8080/location?query={locationByName(name:"xxxx"){id,name,parent}}
				*/
				"locationByName": &gq.Field{
					Type:        LocationGqlType,
					Description: "Get a location by name",
					Args: gq.FieldConfigArgument{
						"name": &gq.ArgumentConfig{
							Type: gq.String,
						},
					},
					Resolve: func(rp gq.ResolveParams) (interface{}, error) {
						var name string = rp.Args["name"].(string)
						if name != "" {
							for _, location := range dataPool {
								if location.Name == name {
									return location, nil
								}
							}
						}

						return nil, nil
					},
				},
				/* Get (read) location list
				   https://localhost:8080/location?query={locations{id,name,parent}}
				*/
				"locations": &gq.Field{
					Type:        gq.NewList(LocationGqlType),
					Description: "Get the location list",
					Resolve: func(rp gq.ResolveParams) (interface{}, error) {
						return dataPool, nil
					},
				},
			},
		},
	)
	schemaConfig := gq.SchemaConfig{
		Query: queries,
	}
	schema, err := gq.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	return schema
}

func ExecuteQuery(query string, dataPool []Location) *gq.Result {
	var schema gq.Schema = GetSchema(dataPool)
	result := gq.Do(gq.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("errors occurred: %v", result.Errors)
	}

	return result
}
