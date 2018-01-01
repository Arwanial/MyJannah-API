package controllers

import (

"github.com/Arwanial/MyJannah-API/models"
"github.com/Arwanial/MyJannah-API/util"
"github.com/Arwanial/MyJannah-API/request"
"github.com/Arwanial/MyJannah-API/response"
"crypto/sha1"
"encoding/hex"
"github.com/astaxie/beego"
"github.com/astaxie/beego/validation"
  "github.com/twinj/uuid"
"encoding/json"
"time"
"github.com/astaxie/beego/utils"
"database/sql"
)


var(
	base_url string = beego.AppConfig.String("base_url")
	gmail_account string = beego.AppConfig.String("gmail_account")
	gmail_account_password string = beego.AppConfig.String("gmail_account_password")
	mail_host string = beego.AppConfig.String("mail_host")
	mail_host_port, err = beego.AppConfig.Int("mail_host_port")
  email_config string = `{"username":"` + gmail_account + `","password":"` + gmail_account_password + `","host":"smtp.gmail.com","port":587}`
  sqlEmail sql.NullString
  sqlMobile sql.NullString
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

  var registrationReq request.RegisterRequest

 /**
  ** Parshing To Json
  */
  if err := json.Unmarshal(r.Ctx.Input.RequestBody, &registrationReq); err != nil {
     r.Ctx.ResponseWriter.WriteHeader(403)
      r.Data["json"] = ErrResponse{410000, "Invalid Request"}
      r.ServeJSON()
      return
  }
  // End of Parsing JSON

  /**
  ** Validation
  */
  valid := validation.Validation{}

  valid.Required(registrationReq.Username, "username").Message("username must be insert")
	valid.Required(registrationReq.Passwords, "password").Message("password must be insert")
  valid.MinSize(registrationReq.Passwords, 8, "password").Message("Password Must be 8 character")

  if registrationReq.Email != "" {
    sqlEmail =  sql.NullString { String : registrationReq.Email, Valid: true}
    valid.Required(registrationReq.Email, "email").Message("email must be insert")
    valid.Email(registrationReq.Email, "email").Message("invalid email format")
  }

  if registrationReq.Mobile != "" {
    sqlMobile =  sql.NullString { String : registrationReq.Mobile, Valid: true}
    valid.Required(registrationReq.Mobile, "mobile").Message("mobile must be insert")
    // valid.Mobile(registrationReq.Mobile, "mobile").Message("invalid mobile format")
  }

  if valid.HasErrors() {
      for _, err := range valid.Errors {
  			r.Ctx.ResponseWriter.WriteHeader(403)
  			r.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
  			r.ServeJSON()
  			return
  		}
  // End of validation


  }else{
      /**
      ** Created ENcryption
      */
      encryptPassword := sha1.New()
      encryptPassword.Write([]byte(registrationReq.Passwords))
      registrationReq.Passwords = hex.EncodeToString(encryptPassword.Sum(nil))

      registration_uid := uuid.NewV4().String()
      // End of  Encryption

      /** Insert To TravelAgent Model
      */
      travelagent := models.Travelagent{
        Email             : sqlEmail,
        Username          : registrationReq.Username,
        Mobile            : sqlMobile,
      	Status            : util.STATUS_PENDING,
      	Password          : registrationReq.Passwords,
    		RegisterNumber		: registration_uid,
      	RoleId            : 3,
        NamaTravel        : sql.NullString { String : "", Valid: false},
        NoKemenagUmroh    : sql.NullString { String : "", Valid: false},
        NoKemenangHaji    : sql.NullString { String : "", Valid: false},
      }
      // End of TravelAgent Model

        /** Created token
        */
        et := util.EasyToken{
          Username : registrationReq.Username,
          Expires:  time.Now().Unix() + 3600,
        }
        token, err := et.GetToken()

        if token == "" || err != nil {
          r.Data["json"] = ErrResponse{-0, err.Error()}

        } else {

          /** Insert To Tabel TravelAgent*/
          if _, err := models.AddTravelagent(&travelagent); err != nil {
           r.Data["json"] = ErrResponse{403, err.Error()}
           r.ServeJSON()
           return
         }
         // End of Insert To Tabel Travelagent

         /** Sending Notification
          ** By Email Notification
         */


          if(registrationReq.Email != ""){
          // link := "http://localhost:8080/v1/verify/" + token
            mail := utils.NewEMail(email_config)
            mail.To = []string{registrationReq.Email}
            mail.From = "ceniebettygreditasari@gmail.com"
            mail.Subject = "My Jannah  -  Travel Agent Account Activation"
            mail.HTML = "Your Account Has Been Registered. Admin will be verify at office hours.<br><br>Best Regards,<br>Awesome's team"

             if err := mail.Send(); err != nil {
               r.Data["json"] = ErrResponse{403, err.Error()}
               r.ServeJSON()
               return
             }
             // End of Email Notification
         } else {

          /** By OTP Notification
          */
         }
           r.Data["json"] = Response{200, "success.", response.RegisterResponse{registrationReq.Username, token}}
        }
      	r.ServeJSON()
  }
}
