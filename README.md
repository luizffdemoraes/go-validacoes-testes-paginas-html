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

<h2 align="center">
Testes
</h2>

### 🧪 Testes utilizando Postman
```json
pm.test("Status code da requisição deve ser 200", function () {
    pm.response.to.have.status(200);
});
pm.test("Verificando o conteúdo da Resposta", function () {
    var jsonData = pm.response.json();
    pm.response.to.have.body('{"API diz:":"E ai Luiz, tudo beleza?"}');
});
```

Instalando lib validação do stretchr testify

```go get github.com/stretchr/testify```


### 🛠 Tecnologias

As seguintes ferramentas foram usadas na construção do projeto:

- [GoLang 1.20](https://go.dev/)
- [Gin 1.9.1](https://github.com/gin-gonic/gin)
- [Gorm](https://gorm.io/index.html)
- [Validator](https://pkg.go.dev/gopkg.in/validator.v2)
- [Testify](https://github.com/stretchr/testify)