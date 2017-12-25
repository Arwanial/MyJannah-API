package request

type RegisterRequest struct {
  Email                 string        `json:"email"`
  Passwords             string        `json:"password"`
  PasswordConfirmation  string        `json:"confirmationPassword"`
  TravelName            string        `json:"travelName"`
  Phone                 string        `json:"phone"`
  Mobile                string        `json:"mobile"`
  Fax                   string        `json:"fax"`
  Website               string        `json:"website"`
  OfficeAddress         string        `json:"officeAddress"`
  OfficeCity            string        `json:"city"`
  OfficeProvince        string        `json:"province"`
  KemenagHajiNo         string        `json:"kemenagHajiNo"`
  KemenagHajiPath       string        `json:"kemenagHajiPath"`
  KemenagUmrohNo        string        `json:"kemenangUmrohNo"`
  KemenangUmrohPath     string        `json:"kemenangUmrohPath"`
}
