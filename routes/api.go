package routes

import (
	api "faceapi/app/handlers"
	// "faceapi/app/repositories/face"
	userRepo "faceapi/app/repositories/user"
	faceRepo "faceapi/app/repositories/face"
	authSrv "faceapi/app/services/auth"

	userSrv "faceapi/app/services/user"

	faceSrv "faceapi/app/services/face"
	"faceapi/config"
	"net/http"

	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
)

var (
	BaseRoute = "/api/v1"
)

func InitializeRoutes(router *mux.Router, dbSession *mgo.Session, conf *config.Configuration) {
	userRepository := userRepo.New(dbSession, conf)
	userService := userSrv.New(userRepository)
	authService := authSrv.New(userRepository)
	authAPI := api.NewAuthAPI(authService, conf)
	userAPI := api.NewUserAPI(userService)

	faceRepository := faceRepo.New(dbSession, conf)
	faceGroupService := faceSrv.New(faceRepository)
	faceGroupAPI := api.NewFaceGroupAPI(faceGroupService)

	// Routes

	//  -------------------------- Auth APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/auth/register", authAPI.Create).Methods(http.MethodPost)
	router.HandleFunc(BaseRoute+"/auth/login", authAPI.Login).Methods(http.MethodPost)

	// -------------------------- User APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/users/me", userAPI.Get).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/users", userAPI.Update).Methods(http.MethodPut)

	// -------------------------- FaceGroup APIs ------------------------------------
	router.HandleFunc(BaseRoute+"/facegroups", faceGroupAPI.GetAll).Methods(http.MethodGet)
	router.HandleFunc(BaseRoute+"/facegroup", faceGroupAPI.Create).Methods(http.MethodPost)
}
