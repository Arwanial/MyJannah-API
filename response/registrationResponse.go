package response

type RegisterResponse struct {
  Email                 string        `json:"username"`
  Token                 string        `json:"token"`
}
