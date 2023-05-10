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
	result, err := s.t.GetAllQuiz(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.QuizesResponse{
		Success:     true,
		Description: "",
		Quizes:      result,
	})
}

func (s *serviceRoutes) GetQuizById(c *gin.Context) {
	quizID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "problem when get id params")
		return
	}

	result, err := s.t.GetQuizById(c, quizID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.QuizResponse{
		Success:     true,
		Description: "",
		Quiz:        result,
	})
}

func (s *serviceRoutes) SaveQuiz(c *gin.Context) {
	var request entity.Quiz

	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parse request json, %s", err))
		return
	}

	quiz, err := s.t.SaveQuiz(c, &request)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, entity.QuizResponse{
		Success:     true,
		Description: "Опрос создан!",
		Quiz:        quiz,
	})
}
