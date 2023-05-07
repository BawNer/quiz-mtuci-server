package v1

import (
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
	}
}

func (s *serviceRoutes) GetAllQuiz(c *gin.Context) {
	result, err := s.t.GetAllQuiz(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
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
	}

	c.JSON(http.StatusOK, entity.QuizResponse{
		Success:     true,
		Description: "",
		Quiz:        result,
	})
}