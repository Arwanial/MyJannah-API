package request

type RegisterRequest struct {
  Email                 String        `json:"email"`
  Passwords             String        `json:"password"`
  PasswordConfirmation  String        `json:"confirmationPassword"`
  TravelName            String        `json:"travelName"`
  Phone                 String        `json:"phone"`
  Mobile                String        `json:"mobile"`
  Fax                   String        `json:"fax"`
  Website               String        `json:"website"`
  OfficeAddress         String        `json:"officeAddress"`
  OfficeCity            String        `json:"city"`
  OfficeProvince        String        `json:"province"`
  KemenagHajiNo         String        `json:"kemenagHajiNo"`
  KemenagHajiPath       String        `json:"kemenagHajiPath"`
  KemenagUmrohNo        String        `json:"kemenangUmrohNo"`
  KemenangUmrohPath     String        `json:"kemenangUmrohPath"`
}
