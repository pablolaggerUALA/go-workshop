package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"uala/go-workshop/internal/processor"
	"uala/go-workshop/internal/repository"
	"uala/go-workshop/pkg/handler"
)

// Punto de entrada de cualquier programa hecho de Go
func main() {
	r := repository.New()
	p := processor.New(r)
	h := handler.New(p)

	lambda.Start(h.DeleteItem)
}
