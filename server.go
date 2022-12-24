package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
      http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
<<<<<<< HEAD
        fmt.Fprintf(w, "Hello! test 2")
=======
        fmt.Fprintf(w, "Hello! test 1)
>>>>>>> a772803e744dea50e0e59d6681bda5774a1eda67
    })

    fmt.Printf("Starting server at port 8081\n")
        if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}
