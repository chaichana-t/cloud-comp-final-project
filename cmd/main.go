package main

import (
	"cloud-final-project/cmd/endpoint"
	"cloud-final-project/cmd/runner"
)

func main() {
	go runner.SaveSnapshot()
	endpoint.ListenAndServe()
}
