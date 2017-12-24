package models

import (

  "regexp"

	"github.com/astaxie/beego/validation"
)

type TravelAgentRegister struct {
  Email                  string    `form:"email"`
	NamaTravel             string    `form:"namatravel"`
	Phone                  string    `form:"phone`
	Fax                    string    `form:"fax"`
	Mobile                 string    `form:"mobile"`
	Website                string    `form:"website" `
	AlamatKantor           string    `form:"alamatkantor" `
	KotaKantor             string    `form:"kotakantor`
	Provinsi               string    `form:"provinsi"`
	NoKemenagUmroh         string    `form:"nokemenagumroh"`
	KemenangUmrohPath      string    `form:"kemenangumrohpath"`
	NoKemenangHaji         string    `form:"nokemenanghaji"`
	KemenagHajiPath        string    `form:"kemenaghajipath"`
  Password               string    `form:"password"`
  PasswordConfirmation   string    `form:"passwordconfirmation"`
  Status                 string    `form:"status"`
}

func (taRegister *TravelAgentRegister) Valid(v *validation.Validation) {

  // Check passwords match
  if taRegister.Password != taRegister.PasswordConfirmation {
    v.SetError("Password", "Passwords don't match")
    v.SetError("ConfirmPassword", "Passwords don't match")
  }
  // Check password contain capital letter
  r, _ := regexp.Compile(`[a-zA-Z0-9&.-!]`)

  if !r.MatchString(taRegister.Password) {
    v.SetError("Password", "Password must contain at least Capital Letter, Numeric, and symbol")
  }
}
