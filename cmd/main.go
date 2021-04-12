package main

import (
	"cloud-final-project/cmd/endpoint"
	_ "cloud-final-project/cmd/endpoint"
)

func main() {
	endpoint.ListenAndServe()
}
