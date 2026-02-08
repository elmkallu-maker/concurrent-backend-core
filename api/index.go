package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // JSON-reitit
    if r.URL.Path == "/api/status" || r.URL.Path == "/api/health" {
        w.Header().Set("Content-Type", "application/json")
        response := map[string]string{
            "status": "ok",
            "path":   r.URL.Path,
        }
        json.NewEncoder(w).Encode(response)
        return
    }
    
    // Muut reitit
    switch r.URL.Path {
    case "/":
        fmt.Fprintf(w, "Etusivu - Server toimii!")
    case "/hello":
        fmt.Fprintf(w, "Terve! 👋")
    default:
        w.WriteHeader(404)
        fmt.Fprintf(w, "404 - Polkua ei löydy: %s", r.URL.Path)
    }
}
