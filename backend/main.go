package main

import (
   "socialai/backend"
   "socialai/handler"
   "fmt"
   "log"
   "net/http"
)

func main() {

   fmt.Println("started-service")

   backend.InitElasticsearchBackend()
   backend.InitGCSBackend()

   log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))

}