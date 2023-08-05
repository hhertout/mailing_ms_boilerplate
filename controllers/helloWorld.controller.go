package Controllers

import (
	"fmt"
	"log"
	"mailer_ms/config"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
  auth, smtpUrl := config.AuthSmtp()
  from := os.Getenv("SMTP_FROM")
  to := []string{os.Getenv("SMTP_TO")}

  subject := "Hello world"
  headers := config.SetHeaders(&from, &to, &subject)

  body := "This text is send from my Go Api !!"
  message := ""
  for key, value := range headers {
    message += fmt.Sprintf("%s: %s \r\n", key, value)
  }
  message += "\r\n" + body

  err := smtp.SendMail(smtpUrl, auth, from, to, []byte(message))

  if err != nil {
    log.Fatal(err)
		c.JSON(200, gin.H{
			"success": false,
      "message": err,
		})
  } else {
    log.Println("Email Successfully sent")   
    c.JSON(200, gin.H{
      "success" : true,
    })
  }
}
