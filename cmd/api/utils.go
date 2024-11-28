package main

import (
	"encoding/json"
	"errors"
	"io"
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


//==============> Read JSON function <==============

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
    // Step 1: Limit the size of the JSON body
    maxBytes := 1024 * 1024 // Allow a maximum JSON size of 1MB
    r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

    // Step 2: Create a JSON decoder for the request body
    dec := json.NewDecoder(r.Body)

    // Step 3: Disallow unknown fields
    // This ensures that if the incoming JSON contains fields that are not mapped to the target struct,
    // an error is thrown. This helps enforce strict validation.
    dec.DisallowUnknownFields()

    // Step 4: Decode the JSON body into the target `data` struct
    err := dec.Decode(data)
    if err != nil {
        // Return the error if decoding fails (e.g., invalid JSON format, type mismatch)
        return err
    }

    // Step 5: Ensure the JSON body contains only one object
    // Attempt to decode again. If additional JSON data exists in the body,
    // this indicates that the body contains multiple JSON objects, which is not allowed.
    err = dec.Decode(&struct{}{}) // Decode into an empty struct to check for extra data
    if err != io.EOF {
        // If the error is not `io.EOF`, it means there is extra data beyond the first JSON object
        return errors.New("body must contain a single JSON object")
    }

    // Step 6: Return nil to indicate successful decoding
    return nil
}

//==================> Error JSON <=====================
