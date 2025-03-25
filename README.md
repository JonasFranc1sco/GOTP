# Projeto: Autenticação 2FA com Go

Este é um sistema de autenticação de dois fatores (2FA) desenvolvido em Go. Ele gera um código TOTP e o envia por e-mail para o usuário como uma camada extra de segurança.

## Tecnologias Utilizadas

- Go
- Biblioteca `pquerna/otp/totp` para geração do código 2FA.
- Biblioteca `joho/godotenv` para carregar variáveis de ambiente.
- `net/smtp` para envio de e-mails.
- `golang.org/x/term` para entrada segura de senha sem que pessoas ao redor consigam ver o que está sendo digitado.

## Funcionalidades

- Cadastro de usuário (nome, e-mail e senha).
- Encriptação segura da senha na entrada.
- Geração e envio de um código de autenticação TOTP por e-mail.
- Validação do código inserido pelo usuário.

## Como Usar

### 1. Configurar Variáveis de Ambiente
Crie um arquivo `.env` e adicione as credenciais do e-mail que será utilizado para o envio do TOTP:

```
BOT_EMAIL=seuemail@gmail.com
BOT_PASSWORD=suasenha
```

### 2. Executar o Projeto

Compile e execute o programa com:
```sh
go run main.go
```

O sistema solicitará nome, e-mail e senha do usuário. Em seguida, enviará um código 2FA para o e-mail informado. O usuário deverá inseri-lo para concluir a autenticação.

## Melhorias Futuras
Em breve pretendo adaptar este projeto para transformá-lo em uma API, para aprimorar meus estudos e deixar o projeto útil. 

## Autor

Desenvolvido por Jonas Francisco Gouveia.

