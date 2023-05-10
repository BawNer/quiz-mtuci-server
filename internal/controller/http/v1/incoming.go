package v1

import (
	"fmt"
	"net/http"
	"quiz-mtuci-server/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"

	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/pkg/logger"
)

type serviceRoutes struct {
	t usecase.UseCase
	l logger.Interface
}

func newQuizRoutes(handler *gin.RouterGroup, t usecase.UseCase, l logger.Interface) {
	r := &serviceRoutes{t, l}

	h := handler.Group("/quiz")
	{
		h.GET("/", r.GetAllQuiz)
		h.GET("/:id", r.GetQuizById)
		h.POST("/", r.SaveQuiz)
	}
}

func (s *serviceRoutes) GetAllQuiz(c *gin.Context) {
	quizzes, err := s.t.GetAllQuiz(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.QuizzesResponseUI{
		Success:     true,
		Description: "",
		Quizzes:     quizzes,
	})
}

func (s *serviceRoutes) GetQuizById(c *gin.Context) {
	quizID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "problem when get id params")
		return
	}

	quiz, err := s.t.GetQuizById(c, quizID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.QuizResponseUI{
		Success:     true,
		Description: "",
		Quiz:        quiz,
	})
}

func (s *serviceRoutes) SaveQuiz(c *gin.Context) {
	var request entity.QuizUI

	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parse request json, %s", err))
		return
	}

	quiz, err := s.t.SaveQuiz(c, &request)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, entity.QuizResponseUI{
		Success:     true,
		Description: "Опрос создан!",
		Quiz:        quiz,
	})
}
