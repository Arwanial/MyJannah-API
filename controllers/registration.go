package controllers

import (

"MyJannah-API/models"
"crypto/sha1"
"encoding/hex"
"github.com/astaxie/beego"
"github.com/astaxie/beego/validation"
  "github.com/twinj/uuid"
	"github.com/astaxie/beego/utils"
)


var(
	base_url string = beego.AppConfig.String("base_url")
	gmail_account string = beego.AppConfig.String("gmail_account")
	gmail_account_password string = beego.AppConfig.String("gmail_account_password")
	mail_host string = beego.AppConfig.String("mail_host")
	mail_host_port, err = beego.AppConfig.Int("mail_host_port")
  email_config string = `{"username":"` + gmail_account + `","password":"` + gmail_account_password + `","host":"smtp.gmail.com","port":587}`
)

// RegistrationController operations for Registration
type RegistrationController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegistrationController) URLMapping() {
	c.Mapping("RegisterTravel", c.RegisterTravel)
}

// Post ...
// @Email
// @Password
// @Nama Travel
// @Phone
// @Mobile
// @Fax
// @AlamatKantor
// @Kota
// @Provinsi
// @No Kemenag Haji
// @No Kemenag Umroh
// @path Kemenag Haji
// @path Kemenag Umroh
// @Website
// @Param	body		body 	models.Travelagent	true		"body for Travelagent content"
// @Success 201 {int} models.Travelagent
// @Failure 403 body is empty
// @router / reg [post]
func(r *RegistrationController) RegisterTravel(){
  email := r.GetString("email")
  password := r.GetString("password")
  namaTravel := r.GetString("namatravel")
  phone := r.GetString("phone")
  mobile := r.GetString("mobile")
  fax := r.GetString("fax")
  alamatKantor := r.GetString("alamatkantor")
  kota := r.GetString("kota")
  provinsi := r.GetString("provinsi")
  noKemenagHaji := r.GetString("nokemenaghaji")
  noKemenagUmroh := r.GetString("nokemenagumroh")
  kemenagHajiPath := r.GetString("kemenanghajipath")
  kemenangUmrohPath := r.GetString("kemenangumrohpath")
  website := r.GetString("website")
  passwordConfirmation := r.GetString("passwordconfirmation")

  valid := validation.Validation{}

  valid.Email(email, "email").Message("insert valid email, please")
  valid.Required(email, "email").Message("email must be insert")
	valid.Required(password, "password").Message("password must be insert")
	valid.Required(namaTravel, "namaTravel").Message("Travel Name must be insert")
  valid.MinSize(namaTravel, 2, "namaTravel").Message("Minimum Travel Name is 2 character")
	valid.Tel(phone, "phone").Message("Insert valid phone, please")
  valid.Required(phone, "phone").Message("password must be insert")
	valid.Tel(mobile, "mobile").Message("Insert valid mobile, please")
  valid.Required(mobile, "mobile").Message("mobile must be insert")
  valid.Tel(fax, "fax").Message("Insert valid fax, please")
	// valid.MaxSize(alamatKantor, 100, "alamatKantor").Message("Maximum Character of Address is 100")
  valid.MinSize(alamatKantor, 10, "alamatKantor").Message("Minimum Character of Address is 10")
  valid.Required(alamatKantor, "alamatKantor").Message("Office Adddress must be insert")
  valid.Required(kota, "kota").Message("City must be insert")
  valid.Required(provinsi, "provinsi").Message("Province must be insert")
  valid.Required(noKemenagHaji, "noKemenagHaji").Message("Kemenang Haji must be insert")
  valid.Required(noKemenagUmroh, "noKemenagUmroh").Message("Kemenang Umroh must be insert")
  // valid.Required(kemenanghajipath, "kemenagHajiPath").Message("Kemenang Haji Path must be insert")
  // valid.Required(kemenangumrohpath, "kemenangumrohpath").Message("Kemenang Umroh Path must be insert")
  valid.Length(password, 8, "password").Message("Password Mus be 8 character")
  valid.Required(passwordConfirmation, "passwordConfirmation").Message("Confirmation Password cant be empty")
  //
  if valid.HasErrors() {
      for _, err := range valid.Errors {
  			r.Ctx.ResponseWriter.WriteHeader(403)
  			r.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
  			r.ServeJSON()
  			return
  		}
  }

  if password != passwordConfirmation {
    r.Ctx.ResponseWriter.WriteHeader(422)
    r.Data["json"] = ErrResponse{422001, "Your Confirmation Password doesnt match"}
    r.ServeJSON()
    return
  }


  // Encrypt Passwords
	encryptPassword := sha1.New()
	encryptPassword.Write([]byte(password))
  password = hex.EncodeToString(encryptPassword.Sum(nil))

	registration_uid := uuid.NewV4().String()

  travelagent := models.Travelagent{
    Email         : email,
  	NamaTravel    : namaTravel,
  	Phone         : phone,
  	Fax           : fax,
  	Mobile        : mobile,
  	Website       : website,
  	AlamatKantor  : alamatKantor,
  	KotaKantor    : kota,
  	Provinsi      : provinsi,
  	NoKemenagUmroh   : noKemenagUmroh,
  	KemenangUmrohPath : kemenangUmrohPath,
  	NoKemenangHaji    : noKemenagHaji,
  	KemenagHajiPath   : kemenagHajiPath,
  	Status            : "PENDING",
  	Password          : password,
		RegisterNumber		: registration_uid,
  	RoleId            : 3,
  }

	if _, err := models.AddTravelagent(&travelagent); err != nil {
		r.Data["json"] = ErrResponse{403, err.Error()}
		r.ServeJSON()
		return
	}else{
			//sending
		link := "http://localhost:8080/v1/verify/" + registration_uid

		 mail := utils.NewEMail(email_config)
		 mail.To = []string{email}
		 mail.From = "ceniebettygreditasari@gmail.com"
		 mail.Subject = "Beego-Ureg - Account Activation"
		 mail.HTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+
		 "\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"

			if err := mail.Send(); err != nil {
			r.Data["json"] = ErrResponse{403, err.Error()}
			r.ServeJSON()
			return
			}
	}
  r.Data["json"] = ErrResponse{200, "Succes"}
	r.ServeJSON()
}
