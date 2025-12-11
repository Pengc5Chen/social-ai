package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pengc5Chen/social-ai/backend/backend"
	"github.com/Pengc5Chen/social-ai/backend/handler"
)

func main() {

	fmt.Println("started-service")

	backend.InitElasticsearchBackend()
	backend.InitGCSBackend()

	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))

}
