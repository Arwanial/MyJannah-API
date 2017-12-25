package response

type RegisterResponse struct {
  Email                 String        `json:"email"`
  TravelName            String        `json:"travelName"`
  Token                 string        `json:"token"`
}
