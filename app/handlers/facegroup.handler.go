package handlers

import (
	"encoding/json"
	"faceapi/app/models"
	faceGroupSvc "faceapi/app/services/face"
	"faceapi/utility"
	"net/http"
)

// FaceGroupHandler - handles face group request
type FaceGroupHandler struct {
	us faceGroupSvc.FaceGroupServiceInterface
}

func NewFaceGroupAPI(faceGroupService faceGroupSvc.FaceGroupServiceInterface) *FaceGroupHandler {
	return &FaceGroupHandler{
		us: faceGroupService,
	}
}

// Create godoc
// @Summary Register user
// @Description Register user api if not exists
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   payload     body    signupReq     true        "User Data"
// @Success 200 {object} basicResponse
// @Success 200 {object} errorRes
// @Router /auth/register [post]
func (h *FaceGroupHandler) Create(w http.ResponseWriter, r *http.Request) {
	payload := new(faceGroupReq)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&payload)
	requestFaceGroup := &models.FaceGroup{Name: payload.Name, Description: payload.Description}
	result := make(map[string]interface{})

	requestFaceGroup.Initialize()

	if h.us.IsFaceGroupAlreadyExists(r.Context(), requestFaceGroup.Name) {
		result = utility.NewHTTPError(utility.FaceGroupAlreadyExists, http.StatusBadRequest)
		utility.Response(w, result)
		return
	}
	err := h.us.Create(r.Context(), requestFaceGroup)
	if err != nil {
		result = utility.NewHTTPError(utility.EntityCreationError, http.StatusBadRequest)
	} else {

		result = utility.SuccessPayload(nil, "Create FaceGoup Successfully ", 201)
	}
	utility.Response(w, result)
}

// Get godoc
// @Summary Get Profile
// @Description Get user profile info
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Success 200 {object} errorRes
// @Security ApiKeyAuth
// @Router /users/me [get]
func (h *FaceGroupHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	facegroups, err := h.us.GetAll(r.Context())

	if err != nil {
		utility.Response(w, utility.NewHTTPError(utility.InternalError, 500))
	} else {
		utility.Response(w, utility.SuccessPayload(facegroups, ""))
	}
}

