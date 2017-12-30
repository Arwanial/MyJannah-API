package request

type RegisterRequest struct {
  Username              string        `json:"username"`
  Email                 string        `json:"email"`
  Passwords             string        `json:"password"`
  Mobile                string        `json:"mobile"`

}
