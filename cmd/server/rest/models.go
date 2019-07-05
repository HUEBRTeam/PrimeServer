package rest

type CreateProfileRequest struct {
	Name      string
	CountryID int
	Avatar    int
	Modifiers int
}

type CreateProfileResponse struct {
	Name       string
	AccessCode string
}

type CreateProfileChange struct {
	AccessCode string
	Nickname   string
	CountryID  int
	Avatar     int
	Modifiers  int
}

type CreateStatus struct {
	Status string
}
