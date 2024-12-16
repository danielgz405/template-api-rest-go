package structures

type CreateRequest struct {
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Name     string   `json:"name"`
	Roles    []string `json:"roles"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type ProfileRequest struct {
	Id    string   `bson:"_id" json:"_id"`
	Name  string   `bson:"name" json:"name"`
	Email string   `bson:"email" json:"email"`
	Roles []string `bson:"roles" json:"roles"`
}
