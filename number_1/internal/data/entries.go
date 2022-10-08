// Filename: internal/data/entries.go

package data

import (
	"time"

	"quiz2.joanneyong.net/internal/validator"
)

type Entry struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Website   string    `json:"website,omitempty"`
	Address   string    `json:"address"`
	Mode      []string  `json:"mode"`
	Version   int32     `json:"version"`
}

func ValidateEntry(v *validator.Validator, entry *Entry) {
	// Use the Check() method to execute our validation checks
	v.Check(entry.Name != "", "name", "must be provided")
	v.Check(len(entry.Name) <= 200, "name", "must not be more than 200 bytes long")

	v.Check(entry.Level != "", "level", "must be provided")
	v.Check(len(entry.Level) <= 200, "level", "must not be more than 200 bytes long")

	v.Check(entry.Contact != "", "contact", "must be provided")
	v.Check(len(entry.Contact) <= 200, "contact", "must not be more than 200 bytes long")

	v.Check(entry.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(entry.Phone, validator.PhoneRX), "phone", "must be a valid phone number")

	v.Check(entry.Email != "", "email", "must be provided")
	v.Check(validator.Matches(entry.Email, validator.EmailRX), "email", "must be a valid email address")

	v.Check(entry.Website != "", "website", "must be provided")
	v.Check(validator.ValidWebsite(entry.Website), "website", "must be a valid URL")

	v.Check(entry.Address != "", "address", "must be provided")
	v.Check(len(entry.Address) <= 500, "address", "must not be more than 500 bytes long")

	v.Check(entry.Mode != nil, "mode", "must be provided")
	v.Check(len(entry.Mode) >= 1, "mode", "must contain at least 1 entry")
	v.Check(len(entry.Mode) <= 5, "mode", "must contain at most 5 entries")
	v.Check(validator.Unique(entry.Mode), "mode", "must not contain duplicate entries")
}
