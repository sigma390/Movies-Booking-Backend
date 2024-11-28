package main

import (
	"encoding/json"
	"net/http"
)

// JSONResponse defines the structure of the JSON object sent to the frontend
type JSONResponse struct {
	Error   bool        `json:"error"`   // Indicates if the response contains an error
	Message string      `json:"message"` // A message describing the response or error
	Data    interface{} `json:"data,omitempty"` // Optional field to hold additional data
}

// writeJSON is a utility function to send a JSON response to the client.
// Parameters:
// - w: http.ResponseWriter to send the response.
// - statusCode: HTTP status code to send (e.g., 200, 400, 500).
// - data: The data to be serialized and sent in the JSON response.
// - headers: Optional HTTP headers to be included in the response.
func (app *application) writeJSON(w http.ResponseWriter, statusCode int, data interface{}, headers ...http.Header) error {
	// Marshal the data into a JSON byte slice
	out, err := json.Marshal(data)
	if err != nil {
		// Return an error if JSON marshalling fails
		return err
	}

	// Add custom headers if any are provided
	if len(headers) > 0 {
		// Loop through the first header map in the variadic headers argument
		for key, value := range headers[0] {
			// Set each header key and value
			w.Header()[key] = value
		}
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the HTTP status code to the response
	w.WriteHeader(statusCode)

	// Write the marshalled JSON to the response body
	_, err = w.Write(out)
	if err != nil {
		// Return an error if writing the response fails
		return err
	}

	// Return nil to indicate success
	return nil
}
