package controllers

import (

<<<<<<< HEAD
"github.com/Arwanial/MyJannah-API/models"
=======
"MyJannah-API/models"
"MyJannah-API/request"
"MyJannah-API/response"
"MyJannah-API/utils"
>>>>>>> 4c21f4f1ebf6b2f23d2f23ee2e9e9d26428c4658
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
// RegistrationRequest
// @Param	body		body 	models.Travelagent	true		"body for Travelagent content"
// @Success 201 {int} models.Travelagent
// @Failure 403 body is empty
// @router / reg [post]
func(r *RegistrationController) RegisterTravel(){

  var registrationReq request.RegistrationRequest

  beego.Debug("test ")

  if err := json.Unmarshal(l.Ctx.Input.RequestBody, &registrationReq); err != nil {
     l.Ctx.ResponseWriter.WriteHeader(410)
      l.Data["json"] = ErrResponse{410000, "Invalid Request"}
      l.ServeJSON()
      return
  }

  valid := validation.Validation{}

  valid.Email(registrationReq.Email, "email").Message("insert valid email, please")
  valid.Required(registrationReq.Email, "email").Message("email must be insert")
	valid.Required(registrationReq.Password, "password").Message("password must be insert")
	valid.Required(registrationReq.TravelName, "travelName").Message("Travel Name must be insert")
  valid.MinSize(registrationReq.TravelName, 2, "travelName").Message("Minimum Travel Name is 2 character")
	valid.Tel(registrationReq.Phone, "phone").Message("Insert valid phone, please")
  valid.Required(registrationReq.Phone, "phone").Message("password must be insert")
	valid.Tel(registrationReq.Mobile, "mobile").Message("Insert valid mobile, please")
  valid.Required(registrationReq.Mobile, "mobile").Message("mobile must be insert")
  valid.Tel(registrationReq.Fax, "fax").Message("Insert valid fax, please")
  valid.MinSize(registrationReq.OfficeAddress, 10, "officeAddress").Message("Minimum Character of Address is 10")
  valid.Required(registrationReq.officeAddress, "officeAddress").Message("Office Adddress must be insert")
  valid.Required(registrationReq.OfficeCity, "city").Message("City must be insert")
  valid.Required(registrationReq.OfficeProvince, "province").Message("Province must be insert")
  valid.Required(registrationReq.KemenagHajiNo, "kemenagHajiNo").Message("Kemenang Haji must be insert")
  valid.Required(registrationReq.KemenagUmrohNo, "kemenangUmrohNo").Message("Kemenang Umroh must be insert")
  valid.Length(registrationReq.Password, 8, "password").Message("Password Mus be 8 character")
  valid.Required(registrationReq.PasswordConfirmation, "confirmationPassword").Message("Confirmation Password cant be empty")
  //
  if valid.HasErrors() {
      for _, err := range valid.Errors {
  			r.Ctx.ResponseWriter.WriteHeader(403)
  			r.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
  			r.ServeJSON()
  			return
  		}
  }else{
    if password != passwordConfirmation {
      r.Ctx.ResponseWriter.WriteHeader(422)
      r.Data["json"] = ErrResponse{422001, "Your Confirmation Password doesnt match"}
      r.ServeJSON()
      return
    }else{
      // Encrypt Passwords
      encryptPassword := sha1.New()
      encryptPassword.Write([]byte(password))
      password = hex.EncodeToString(encryptPassword.Sum(nil))

      registration_uid := uuid.NewV4().String()

      //setter
      travelagent := models.Travelagent{
        Email             : registrationReq.Email,
      	NamaTravel        : registrationReq.TravelName,
      	Phone             : registrationReq.Phone,
      	Fax               : registrationReq.Fax,
      	Mobile            : registrationReq.Mobile,
      	Website           : registrationReq.Website,
      	AlamatKantor      : registrationReq.OfficeAddress,
      	KotaKantor        : registrationReq.OfficeCity,
      	Provinsi          : registrationReq.OfficeProvince,
      	NoKemenagUmroh    : registrationReq.KemenagUmrohNo,
      	KemenangUmrohPath : registrationReq.KemenangUmrohPath,
      	NoKemenangHaji    : registrationReq.KemenagHajiNo,
      	KemenagHajiPath   : registrationReq.KemenagHajiPath,
      	Status            : globalConstantsSTATUS_PENDING,
      	Password          : registrationReq.Password,
    		RegisterNumber		: registration_uid,
      	RoleId            : 3,
      }

      if _, err := models.AddTravelagent(&travelagent); err != nil {
    		r.Data["json"] = ErrResponse{403, err.Error()}
    		r.ServeJSON()
    		return
    	}else{

        //create token
        et := utils.EasyToken{
          RegistrationId : registration_uid,
          Expires:  time.Now().Unix() + 3600,
        }
        token, err := et.GetToken()
        if token == "" || err != nil {
          r.Data["json"] = ErrResponse{-0, err.Error()}
        } else {
          link := "http://localhost:8080/v1/verify/" + token

          mail := utils.NewEMail(email_config)
          mail.To = []string{registrationReq.Email}
          mail.From = "ceniebettygreditasari@gmail.com"
          mail.Subject = "Beego-Ureg - Account Activation"
          mail.HTML = "To verify your account, please click on the following link.<br><br><a href=\""+link+
          "\">"+link+"</a><br><br>Best Regards,<br>Awesome's team"

           if err := mail.Send(); err != nil {
           r.Data["json"] = ErrResponse{403, err.Error()}
           r.ServeJSON()
           return
           }
           r.Data["json"] = Response{200, "success.", response.RegistrationResponse{registrationReq.Email, registrationReq.TravelName, token}}
        }
      	r.ServeJSON()
      }
    }
  }
}
