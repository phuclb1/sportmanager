package models

import (
	"faceapi/utility"

	// "github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// User , definds user model
type FaceGroup struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name,omitempty" bson:"name,omitempty"`
	Description     string  `json:"description,omitempty" bson:"email,omitempty"`
	CreatedAT int64         `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAT int64         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

// Initialize , will set the hashed password, createdAt and updatedAt
// date in milliseconds
func (u *FaceGroup) Initialize() error {
	u.CreatedAT = utility.CurrentTimeInMilli()
	u.UpdatedAT = utility.CurrentTimeInMilli()
	return nil
}
