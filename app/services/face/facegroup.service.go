package user

import (
	"context"
	// "faceapi/utility"

	model "faceapi/app/models"
	"faceapi/config"

	repository "faceapi/app/repositories/face"

	mgo "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type FaceGroupServiceInterface interface {
	// Get(context.Context, string) (*model.FaceGroup, error)
	GetAll(context.Context) ([]*model.FaceGroup, error)
	Create(context.Context, *model.FaceGroup) error
	IsFaceGroupAlreadyExists(context.Context, string) bool
}

// UserService , implements UserService
// and perform user related business logics
type FaceGroupService struct {
	db         *mgo.Session
	repository repository.FaceGroupRepository
	config     *config.Configuration
}

// New function will initialize UserService
func New(faceGroupRepo repository.FaceGroupRepository) FaceGroupServiceInterface {
	return &FaceGroupService{repository: faceGroupRepo}
}


func (service *FaceGroupService) Create(ctx context.Context, faceGroup *model.FaceGroup) error {

	return service.repository.Create(ctx, faceGroup)
}
// Get function will find user by id
// return user and error if any
func (service *FaceGroupService) GetAll(ctx context.Context) ([]*model.FaceGroup, error) {
	return service.repository.FindAll(ctx)
}


// IsUserAlreadyExists , checks if user already exists in DB
func (service *FaceGroupService) IsFaceGroupAlreadyExists(ctx context.Context, email string) bool {

	return service.repository.IsFaceGroupAlreadyExists(ctx, email)

}