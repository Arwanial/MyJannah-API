package response

type RegisterResponse struct {
  Email                 string        `json:"email"`
  TravelName            string        `json:"travelName"`
  Token                 string        `json:"token"`
}
