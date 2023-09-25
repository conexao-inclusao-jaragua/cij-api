# Conexão Inclusão Jaraguá - API

Olá, bem-vindo(a) à API do CIJ! Aqui você encontrará informações importantes sobre a estrutura da aplicação, os pré-requisitos necessários para executá-la, as instruções de instalação, o banco de dados utilizado e os autores responsáveis pelo desenvolvimento.

## 🧱 Estrutura

- Fiber: 2.49.1
- Gorm: 1.25.4
- Golang-JWT: 3.2.2
- Mysql-driver: 1.5.1

## ✅ Pré-requisitos

- Golang 1.21.0
- MySQL 8+

Antes de prosseguir, certifique-se de ter os seguintes componentes instalados:

## 🛠 Instalação

1. **Clonar o repositório:** Clone o repositório [API](https://github.com/conexao-inclusao-jaragua/cij-api.git) do Github para sua máquina local
2. **Instalar as dependências:** Navegue até o diretório do projeto clonado e execute o seguinte comando para instalar todas as dependências
```
go install 
```
3. **Configurar variáveis de ambiente:** Crie um arquivo `app.env` na raiz do projeto e configure-o com as variáveis disponíveis no arquivo `app.env.example`
4. **Iniciar a aplicação:** Se a instalação das dependências for bem sucedida e as variáveis de ambiente estiverem configuradas, a aplicação está pronta para ser iniciada. Para isso, execute este outro comando
```
go run main.go
```

## ✍ Autores

- [Camilly de Souza Pessotti](https://github.com/pessotticamilly)
- [Camilly Vitória da Rocha Goltz](https://github.com/VitoriaCamilly)
- [Cauã Kath](https://github.com/CauaKath)
- [Kenzo Hideaky Ferreira Sato](https://github.com/Kenzohfs)
