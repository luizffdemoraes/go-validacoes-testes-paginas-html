package main

import (
	"api-go-gin/controllers"
	"api-go-gin/database"
	"api-go-gin/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz":"E ai gui, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListaTodosOsAlunosHanlder(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
