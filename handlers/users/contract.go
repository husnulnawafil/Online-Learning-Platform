package handlers

type RequestBody struct {
	Name         string `file:"name"`
	Email        string `file:"email"`
	Password     string `file:"password"`
	ProfileImage string `file:"profile_image,omitempty"`
}

type ResponseData struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image,omitempty"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
