package rest

type CreateProfileRequest struct {
	Name string
}

type CreateProfileResponse struct {
	Name       string
	AccessCode string
}
