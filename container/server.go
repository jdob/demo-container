package main

import (
    "fmt"
    "net/http"
)

const asciiArt = `
    ________  ___  ___  ________  _______      
   |\   ____\|\  \|\  \|\   ____\|\  ___ \     
   \ \  \___|\ \  \\\  \ \  \___|\ \   __/|    
    \ \_____  \ \  \\\  \ \_____  \ \  \_|/__  
     \|____|\  \ \  \\\  \|____|\  \ \  \_|\ \ 
       ____\_\  \ \_______\____\_\  \ \_______\
      |\_________\|_______|\_________\|_______|
      \|_________|        \|_________|         
`

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.NotFound(w, r)
        return
    }
    
    message := fmt.Sprintf("Welcome to SUSECON!\n%s", asciiArt)
    fmt.Fprint(w, message)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    
    port := ":9977"
    fmt.Printf("Server starting on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}