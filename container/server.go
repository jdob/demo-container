package main

import (
    "fmt"
    "net/http"
    "time"
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

const suseconAsciiArt = `
   _______  __    __  _______  _______  _______  _______  __    _ 
  |       ||  |  |  ||       ||       ||       ||       ||  |  | |
  |  _____||  |  |  ||  _____||    ___||     __||   _   ||   |_| |
  | |_____ |  |  |  || |_____ |   |___ |    |   |  | |  ||       |
  |_____  ||  |__|  ||_____  ||    ___||    |__ |  |_|  ||  _    |
   _____| ||        | _____| ||   |___ |       ||       || | |   |
  |_______||________||_______||_______||_______||_______||_|  |__|
`

// Animation frames for SUSECON
var animationFrames = []string{
    // Frame 1: S
    `
   _______  
  |       | 
  |  _____| 
  | |_____ 
  |_____  |
   _____| |
  |_______|
`,
    // Frame 2: SU
    `
   _______  __    __ 
  |       ||  |  |  |
  |  _____||  |  |  |
  | |_____ |  |  |  |
  |_____  ||  |__|  |
   _____| ||        |
  |_______||________|
`,
    // Frame 3: SUS
    `
   _______  __    __  _______ 
  |       ||  |  |  ||       |
  |  _____||  |  |  ||  _____|
  | |_____ |  |  |  || |_____ 
  |_____  ||  |__|  ||_____  |
   _____| ||        | _____| |
  |_______||________||_______|
`,
    // Frame 4: SUSE
    `
   _______  __    __  _______  _______ 
  |       ||  |  |  ||       ||       |
  |  _____||  |  |  ||  _____||    ___|
  | |_____ |  |  |  || |_____ |   |___ 
  |_____  ||  |__|  ||_____  ||    ___|
   _____| ||        | _____| ||   |___ 
  |_______||________||_______||_______|
`,
    // Frame 5: SUSEC
    `
   _______  __    __  _______  _______  _______ 
  |       ||  |  |  ||       ||       ||       |
  |  _____||  |  |  ||  _____||    ___||     __|
  | |_____ |  |  |  || |_____ |   |___ |    |
  |_____  ||  |__|  ||_____  ||    ___||    |__
   _____| ||        | _____| ||   |___ |       | 
  |_______||________||_______||_______||_______|
`,
    // Frame 6: SUSECO
    `
   _______  __    __  _______  _______  _______  _______ 
  |       ||  |  |  ||       ||       ||       ||       |
  |  _____||  |  |  ||  _____||    ___||     __||   _   |
  | |_____ |  |  |  || |_____ |   |___ |    |   |  | |  |
  |_____  ||  |__|  ||_____  ||    ___||    |__ |  |_|  |
   _____| ||        | _____| ||   |___ |       ||       |
  |_______||________||_______||_______||_______||_______|
`,
    // Frame 7: SUSECON
    suseconAsciiArt,
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.NotFound(w, r)
        return
    }
    
    message := fmt.Sprintf("Welcome to SUSECON!\n%s", asciiArt)
    fmt.Fprint(w, message)
}

func animationHandler(w http.ResponseWriter, r *http.Request) {
    // Set headers for streaming response
    w.Header().Set("Content-Type", "text/plain")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    
    // Get flusher to force sending chunks
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming not supported", http.StatusInternalServerError)
        return
    }
    
    // Control sequences for terminal
    clearScreen := "\033[2J"   // Clear the entire screen
    homeCursor := "\033[H"     // Move cursor to home position (0,0)
    
    // Number of animation cycles
    cycles := 100
    
    // Play the animation
    for cycle := 0; cycle < cycles; cycle++ {
        for _, frame := range animationFrames {
            // Clear screen and move cursor to top-left
            fmt.Fprint(w, clearScreen+homeCursor)
            
            // Draw the current frame
            fmt.Fprintf(w, "Welcome to:\n%s", frame)
            
            // Force sending this chunk
            flusher.Flush()
            
            // Wait before showing next frame
            time.Sleep(300 * time.Millisecond)
        }
        
        // Pause for a moment on the complete SUSECON text
        time.Sleep(500 * time.Millisecond)
    }
    
    // Add a color flourish at the end (green for SUSE)
    greenColor := "\033[32m"
    resetColor := "\033[0m"
    
    fmt.Fprint(w, clearScreen+homeCursor)
    fmt.Fprintf(w, "Welcome to:\n%s%s%s\n\nAnimation complete! Enjoy SUSECON!\n", 
                greenColor, suseconAsciiArt, resetColor)
    flusher.Flush()
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/ani", animationHandler)
    
    port := ":9977"
    fmt.Printf("Server starting on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}