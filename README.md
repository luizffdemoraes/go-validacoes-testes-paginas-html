<h1 align="center">
Go: validações, testes e páginas HTML
</h1>

<h2 align="center">
Instalando e criando a primeira rota com Gin
</h2>

* Carregamos o projeto base e criamos a imagem do banco de dados no Docker;
* Criamos nossas validações na struct de Aluno, garantindo que um campo não fique em branco e tenha uma quantidade específica de caracteres;
* Aplicamos essa validação no controller no momento que criamos ou editamos um aluno

### Adicionar lib de validação.:

```go get gopkg.in/validator.v2```

### 🛠 Tecnologias

As seguintes ferramentas foram usadas na construção do projeto:

- [GoLang 1.20](https://go.dev/)
- [Gin 1.9.1](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/index.html)
- [Validator](https://pkg.go.dev/gopkg.in/validator.v2)