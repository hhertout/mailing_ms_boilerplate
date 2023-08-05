package Controllers

import (
	"fmt"
	"log"
	"mailer_ms/config"
	GoMailer "mailer_ms/services"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
  auth, smtpUrl := config.SmtpAuthentification()
  from := os.Getenv("SMTP_FROM")
  to := []string{os.Getenv("SMTP_TO")}

  subject := "Hello world"
  headers := config.SetHeaders(&from, &to, &subject)

  body := "This text is send from my Go Api !!"
  message := ""
  for k, v := range headers {
    message += fmt.Sprintf("%s: %s \r\n", k, v)
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

func HelloWorldWithHtml(c *gin.Context) {
  templatePath := "helloworld.html"
  templateData := struct{
    Name string
    URL string
  } {
    Name: "toto",
    URL: "http://go.dev",
  }
  r := GoMailer.NewRequest("Hello from go Controller", []string {"jane.doe@gmail.com"})
  r.ParseHTMLTemplate(templatePath, templateData).SendEmail()
}

