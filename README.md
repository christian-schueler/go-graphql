# go-graphql

This project is an example how to use [GraphQL](https://graphql.org/) with [Go](https://go.dev/).
It implements ths project [graphql-go](https://github.com/graphql-go/graphql).

## Currently defined Queries

There are, at the moment, only read queries defined. The mutation queries are yet to be implemented.

### List all locations

With the query __locations__ you'll get a list of all locations, defined in the example data.

### Query by Id

This query gives you the location by Id, often used by machines, because of the GUID.

### Query by Name

This query gives you the location by name. It's more convenient than the query by Id.
