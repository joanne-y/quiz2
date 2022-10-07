//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"quiz2.joanneyong.net/internal/data"
	"quiz2.joanneyong.net/internal/validator"
)

// createSchoolHandler for the "POST /v1/entries" endpoint
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

// showSchoolHnadler for the "GET v1/entries/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Use the "ParamsFromContext()" function to get the request context as a slice
	id, err := app.readIDParam(r)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	// Create a new instance of the Entry struct containing the ID we extracted
	// from our URL and some sample data
	entry := data.Entry{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Apple Tree",
		Level:     "High School",
		Contact:   "Anna Smith",
		Phone:     "601-4411",
		Address:   "14 Apple street",
		Mode:      []string{"blended", "online"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"entry": entry}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
