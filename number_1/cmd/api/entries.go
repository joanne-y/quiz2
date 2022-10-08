//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"

	"quiz2.joanneyong.net/internal/data"
	"quiz2.joanneyong.net/internal/validator"
)

// createEntryHandler for the "POST /v1/entries" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination
	var input struct {
		Name    string   `json:"name"`
		Level   string   `json:"level"`
		Contact string   `json:"contact"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Website string   `json:"website"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}
	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Entry struct

	entry := &data.Entry{
		Name:    input.Name,
		Level:   input.Level,
		Contact: input.Contact,
		Phone:   input.Phone,
		Email:   input.Email,
		Website: input.Website,
		Address: input.Address,
		Mode:    input.Mode,
	}
	// Initialize a new Validator instance
	v := validator.New()

	//Check the map to determine if there were any validation errors
	if data.ValidateEntry(v, entry); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Display the request
	fmt.Fprintf(w, "%+v\n", input)
}

// showRandomStringHandler for the "GET /v1/random/:id" endpoint
func (app *application) showRandomStringHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Display the random string
	//Converts the int64 id to an int
	num := int(id)
	fmt.Fprintf(w, "show random string for %d\n", num)
	fmt.Fprintln(w, app.Tools.generateRandomString(num))
}
