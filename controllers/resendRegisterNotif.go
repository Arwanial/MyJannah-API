package controllers

import (
  "github.com/Arwanial/MyJannah-API/request"
  "github.com/Arwanial/MyJannah-API/models"
  "github.com/Arwanial/MyJannah-API/response"
  "github.com/astaxie/beego"
  "encoding/json"
  "github.com/astaxie/beego/utils"
)

var(

	// gmail_account string = beego.AppConfig.String("gmail_account")
	// gmail_account_password string = beego.AppConfig.String("gmail_account_password")
	// mail_host string = beego.AppConfig.String("mail_host")
	// mail_host_port, err = beego.AppConfig.Int("mail_host_port")
  // email_config string = `{"username":"` + gmail_account + `","password":"` + gmail_account_password + `","host":"smtp.gmail.com","port":587}`

)

/**
 ** ResendController operations for Resend Register Notification
 */
type ResendRegisterNotifController struct {
	beego.Controller
}

/**
 ** URL Mapping
 */
 func (r *ResendRegisterNotifController) URLMapping() {
 	r.Mapping("Post", r.ResendNotifReg)
 }

 /**
  ** Function Resend Notif Register Method
  ** AllowMethod : POST
  */
  func(r *ResendRegisterNotifController) ResendNotifReg(){
    var resendRegNotifReq request.ResendRegisterRequest


  /**
   ** Parshing Json
   */
   if err := json.Unmarshal(r.Ctx.Input.RequestBody, &resendRegNotifReq); err != nil {
      r.Ctx.ResponseWriter.WriteHeader(403)
      r.Data["json"] = ErrResponse{410000, "Invalid Request"}
      r.ServeJSON()
       return
   }

   /**
    ** Check Account
    */
    if account, err := models.CheckAccountByEmailOrPhone(resendRegNotifReq.Account); err != nil {
      r.Data["json"] = ErrResponse{403, err.Error()}
      r.ServeJSON()
      return
    } else{
      /**
       **  Check If Account Already Exist
       */
       if(account == 1){
         if resendRegNotifReq.TypeAccount == "EMAIL" {

           /**
            ** Resend Email
            */
            mail := utils.NewEMail(email_config)
            mail.To = []string{resendRegNotifReq.Account}
            mail.From = "ceniebettygreditasari@gmail.com"
            mail.Subject = "My Jannah  -  Travel Agent Account Activation"
            mail.HTML = "Your Account Has Been Registered. Admin will be verify at office hours.<br><br>Best Regards,<br>Awesome's team"

             if err := mail.Send(); err != nil {
               r.Data["json"] = ErrResponse{403, err.Error()}
               r.ServeJSON()
               return
             }
         }else{

           /**
            ** Mobile Notification
            */
         }
         r.Ctx.ResponseWriter.WriteHeader(200)
         r.Data["json"] = response.ResendRegisterResponse{"200", "Notification Has Been Sent"}
         r.ServeJSON()
       }
    }

}
