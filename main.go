package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/pquerna/otp/totp"
	"golang.org/x/term"
)

// Estrutura do usuário (Nome, email e senha)
type users struct {
	Username string
	Email    string
	Password string
}

// Func de configuração do envio de email
func sendEmail(to, subject, body string) error {

	botEmail := os.Getenv("BOT_EMAIL")
	botPassword := os.Getenv("BOT_PASSWORD")

	from := botEmail
	Password := botPassword
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Converter a mensagem de Email em formato MIME
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)

	//Conexão com o servidor SMTP usado para enviar Emails
	auth := smtp.PlainAuth("", from, Password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil

}

func generateTOTP(user users) {

	// Gerar uma chave TOTP
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Jonas Francisco",
		AccountName: user.Email,
		SecretSize:  5,
		Period:      60,
	})
	if err != nil {
		log.Fatalf("Erro ao gerar chave TOTP: %v", err)
	}

	//Corpo do Email
	EmailBody := fmt.Sprintf(`
	Olá, %s sua chave de acesso para o sistema!
	Chave secreta: %s`, user.Username, key.Secret())

	//Envio do email para o usuário
	err = sendEmail(user.Email, "Sua chave para meu sistema", EmailBody)
	if err != nil {
		log.Fatalf("Erro ao enviar Email %v", err)
	}

	//Validação do código inserido pelo usuário
	validateTOTP := key.Secret()
	fmt.Println("Retorne a chave enviada no Email")
	var verifyTOTP string
	fmt.Scan(&verifyTOTP)
	if validateTOTP == verifyTOTP {
		fmt.Println("Validado com sucesso.")
	} else {
		fmt.Println("Inválido")
	}
}

func main() {
	godotenv.Load()

	//Cadastro do usuário
	user := users{}
	fmt.Println("Olá! Você está no sistema de 2FA.")
	fmt.Println("Insira seu nome:")
	fmt.Scan(&user.Username)

	fmt.Println("Digite seu Email:")
	fmt.Scan(&user.Email)

	fmt.Println("Digite sua senha:")
	//Faz com que a senha do usuário não seja mostrada na tela, para isso é necessário transformar a senha em bytes (int) e depois retornar para string
	psswrdByte, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Deu erro! %v", err)
		return
	}
	user.Password = string(psswrdByte)

	fmt.Println("Confirme sua senha:")
	confirmByte, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Erro na confirmação: %v", err)
	}
	confirmPassword := string(confirmByte)
	if user.Password != confirmPassword {
		fmt.Println("Senha incorreta.")
		return
	} else {
		fmt.Println("Senha correta. Em alguns instantes iremos enviar um Email de confirmação.")
	}

	generateTOTP(user)

}
