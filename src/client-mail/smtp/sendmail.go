package smtp

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
	mail "github.com/xhit/go-simple-mail"
)

func SendMail(w http.ResponseWriter, r *http.Request) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./src/client-mail/smtp")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No hay fichero de configuración \n", err)
		os.Exit(1)
	}

	host := viper.GetString("smtp.host")
	port := viper.GetInt("smtp.port")
	username := viper.GetString("smtp.username")
	password := viper.GetString("smtp.password")
	from := viper.GetString("smtp.from")
	to := viper.GetString("smtp.to")
	subject := viper.GetString("smtp.subject")
	body := viper.GetString("smtp.body")

	server := mail.NewSMTPClient()

	// Servidor SMTP
	server.Host = host
	server.Port = port

	// Credenciales
	server.Username = username
	server.Password = password

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	email := mail.NewMSG()

	// Quien manda el correo
	email.SetFrom(from)
	// Quien recibe el correo
	email.AddTo(to)

	// Asunto
	email.SetSubject(subject)

	// Mensaje
	email.SetBody(mail.TextPlain, body)

	err = email.Send(smtpClient)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Correo enviado")
	fmt.Fprintf(w, "Correo enviado")
}
