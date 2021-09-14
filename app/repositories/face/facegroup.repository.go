package face

import (
	"context"
	model "faceapi/app/models"
	"faceapi/config"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
// So that, db opration can be perform easily
type FaceGroupRepository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *model.FaceGroup) error

	// FildAll, returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]*model.FaceGroup, error)

	// FindOneById, find the user by the provided id
	// return matched user and error if any
	FindOneById(context.Context, string) (*model.FaceGroup, error)

	// Update, will update user data by id
	// return error if any
	Update(context.Context, interface{}, interface{}) error

	// Delete, will remove user entry from DB
	// Return error if any
	Delete(context.Context, *model.FaceGroup) error

	// FindOne, will find one entry of user matched by the query
	// Query object is an interface type that can accept any object
	// return matched user and error if any
	FindOne(context.Context, interface{}) (*model.FaceGroup, error)

	IsFaceGroupAlreadyExists(context.Context, string) bool
}

type FaceGroupRepositoryImp struct {
	db     *mgo.Session
	config *config.Configuration
	
}


func New(db *mgo.Session, c *config.Configuration) FaceGroupRepository {
	return &FaceGroupRepositoryImp{db: db, config: c}
}

func (service *FaceGroupRepositoryImp) Create(ctx context.Context, facegroup *model.FaceGroup) error {
	return service.collection().Insert(facegroup)
}

func (service *FaceGroupRepositoryImp) FindAll(ctx context.Context) ([]*model.FaceGroup, error) {
	var listFacegroup []*model.FaceGroup
	err := service.collection().Find(nil).All(&listFacegroup)
	return listFacegroup , err
}

func (service *FaceGroupRepositoryImp) Update(ctx context.Context, query, change interface{}) error {

	return service.collection().Update(query, change)
}

func (service *FaceGroupRepositoryImp) FindOneById(ctx context.Context, id string) (*model.FaceGroup, error) {
	var facegroup model.FaceGroup
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	e := service.collection().Find(query).One(&facegroup)
	return &facegroup, e
}

func (service *FaceGroupRepositoryImp) Delete(ctx context.Context, facegroup *model.FaceGroup) error {
	return nil
}

func (service *FaceGroupRepositoryImp) FindOne(ctx context.Context, query interface{}) (*model.FaceGroup, error) {
	var facegroup model.FaceGroup
	e := service.collection().Find(query).One(&facegroup)
	return &facegroup, e
}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *FaceGroupRepositoryImp) IsFaceGroupAlreadyExists(ctx context.Context, name string) bool {
	query := bson.M{"name": name}
	_, e := service.FindOne(ctx, query)
	if e != nil {
		return false
	}
	return true
}

func (service *FaceGroupRepositoryImp) collection() *mgo.Collection {
	return service.db.DB(service.config.DataBaseName).C("face_group")
}
