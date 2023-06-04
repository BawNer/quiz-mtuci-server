package v1

import (
	"fmt"
	"net/http"
	"quiz-mtuci-server/internal/entity"

	"github.com/gin-gonic/gin"

	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/pkg/logger"
)

type serviceRoutes struct {
	t usecase.UseCase
	l logger.Interface
	m *usecase.MiddlewareStruct
}

func newQuizRoutes(handler *gin.RouterGroup, t usecase.UseCase, l logger.Interface, m *usecase.MiddlewareStruct) {
	r := &serviceRoutes{t, l, m}

	h := handler.Group("/quiz")
	h.Use(m.AuthGuard())
	{
		h.GET("/", r.GetAllQuiz)
		h.GET("/:hash", r.GetQuizByHash)
		h.POST("/respond", r.SaveReviewerRespond)
	}
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.UseCase, l logger.Interface, m *usecase.MiddlewareStruct) {
	r := &serviceRoutes{t, l, m}

	h := handler.Group("/users")
	{
		h.POST("/login", r.GetUserByLoginWithPassword)
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

func (s *serviceRoutes) GetQuizByHash(c *gin.Context) {
	quizHash := c.Param("hash")

	quiz, err := s.t.GetQuizByHash(c, quizHash)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userGroupID := int(c.Value("token").(map[string]interface{})["GroupID"].(float64))
	IsAccess := false

	for _, v := range quiz.AccessFor {
		if v.ID == userGroupID {
			IsAccess = true
			break
		}
	}

	if !IsAccess || !quiz.Active {
		errorResponse(c, http.StatusForbidden, "No access")
		return
	}

	c.JSON(http.StatusOK, entity.QuizResponseUI{
		Success:     true,
		Description: "",
		Quiz:        quiz,
	})
}

func (s *serviceRoutes) GetUserByLoginWithPassword(c *gin.Context) {
	var userLogin entity.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parse body, %v", err))
	}

	user, err := s.t.GetUserByLoginWithPassword(c, userLogin)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", user.Token))
	c.JSON(http.StatusOK, entity.ResponseUserLogin{
		Success:     true,
		Description: "Login success",
		User:        user,
	})
}

func (s *serviceRoutes) SaveReviewerRespond(c *gin.Context) {
	var respond *entity.Reviewers
	if err := c.ShouldBindJSON(&respond); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parse body, %v", err))
		return
	}

	err := s.t.SaveReviewers(c, respond)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error save reviewer %+v", err.Error()))
		return
	}

	c.JSON(http.StatusOK, entity.ReviewersResponse{
		Success:     true,
		Description: "Ответ сохранен!",
	})
}
