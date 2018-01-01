package controllers

import (
  "github.com/Arwanial/MyJannah-API/models"
  "github.com/Arwanial/MyJannah-API/request"
  "github.com/Arwanial/MyJannah-API/response"
  "github.com/astaxie/beego/validation"
  "github.com/astaxie/beego"
  "github.com/Arwanial/MyJannah-API/util"
  "time"
  "crypto/sha1"
  "encoding/json"
  "encoding/hex"
  //"github.com/astaxie/beego/utils"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Login)
}

// Login ...
// @Username
// @Password
// @TypeLogin
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {int} models.Travelagent
// @Failure 403 body is empty
// @router / login [post]
func(l *LoginController) Login(){

  var loginReq request.LoginRequest

  if err := json.Unmarshal(l.Ctx.Input.RequestBody, &loginReq); err != nil {
     l.Ctx.ResponseWriter.WriteHeader(403)
      l.Data["json"] = ErrResponse{410000, "Invalid Request"}
      l.ServeJSON()
      return
  }

     valid := validation.Validation{}

    //validation
    valid.Required(loginReq.Username, "username").Message("username must be insert")
    valid.Required(loginReq.Password, "password").Message("password must be insert")

   //Encrypt Passwords
   encryptPassword := sha1.New()
   encryptPassword.Write([]byte(loginReq.Password))
   loginReq.Password = hex.EncodeToString(encryptPassword.Sum(nil))

   //If Has Error From validation
       if valid.HasErrors() {
         for _, err := range valid.Errors {
             l.Ctx.ResponseWriter.WriteHeader(403)
             l.Data["json"] = ErrResponse{403001, map[string]string{err.Key: err.Message}}
             l.ServeJSON()
             return
            }
       }else{
           //Check Available Account
            w, err := models.CheckUserByEmailAndPassword(loginReq.Username, loginReq.Password);
            if err != nil {
              l.Ctx.ResponseWriter.WriteHeader(403)
              l.Data["json"] = ErrResponse{428001, "Account Not Be Found"}
              l.ServeJSON()
              return
            }else{

              //create token
              et := util.EasyToken{
                Username : loginReq.Username,
                Expires:  time.Now().Unix() + 3600,
              }
              token, err := et.GetToken()

              if token == "" || err != nil {
                l.Ctx.ResponseWriter.WriteHeader(403)
                l.Data["json"] = ErrResponse{-0, err.Error()}
              } else{

             //Check Status Account
                if w.Status != "ACTIVE" {

                if  w.NamaTravel.Valid == false && w.Phone == "" && w.Fax == "" && w.Mobile.Valid == false && w.AlamatKantor == "" && w.KotaKantor == "" && w.Provinsi == "" && w.NoKemenagUmroh.Valid == false && w.KemenangUmrohPath == "" && w.NoKemenangHaji.Valid == false && w.KemenagHajiPath == "" {
                     l.Data["json"] = Response{200, "success.", response.LoginResponse{"INACTIVE", "INCOMPLETE", token}}
                } else {
                    l.Data["json"] = Response{200, "success.", response.LoginResponse{"INACTIVE", "COMPLETE", "null"}}
                }
                  l.ServeJSON()
                } else if w.Status == "ACTIVE" {
                if  w.NamaTravel.Valid == false && w.Phone == "" && w.Fax == "" && w.Mobile.Valid == false && w.AlamatKantor == "" && w.KotaKantor == "" && w.Provinsi == "" && w.NoKemenagUmroh.Valid == false && w.KemenangUmrohPath == "" && w.NoKemenangHaji.Valid == false && w.KemenagHajiPath == "" {
                    l.Data["json"] = Response{200, "success.", response.LoginResponse{"ACTIVE", "INCOMPLETE", token}}
                 } else {
                    l.Data["json"] = Response{200, "success.", response.LoginResponse{"ACTIVE", "COMPLETE", token}}
                  }
                  l.ServeJSON()
                }
             }
          }
        }
  }
