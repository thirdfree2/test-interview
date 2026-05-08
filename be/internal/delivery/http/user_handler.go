package http

import (
	"be-interview-app/internal/usecase"
	"errors"
	"net/http"
	"strconv"

	customError "be-interview-app/internal/delivery/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// RegisterInput คือโครงสร้างของ JSON ที่รับมาจาก Client
type RegisterInput struct {
	Name string `json:"name" binding:"required,max=20,alphanum"`
	Password string `json:"password" binding:"required,max=32,password"`
}

type LoginInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput

	// validate request
	if err := c.ShouldBindJSON(&input); err != nil {

		var ve validator.ValidationErrors

		if errors.As(err, &ve) {

			errorMessages := map[string]string{}

			for _, fe := range ve {

				switch fe.Field() {

				case "Name":

					switch fe.Tag() {
					case "required":
						errorMessages["name"] = "กรุณากรอก username"

					case "max":
						errorMessages["name"] = "username ต้องไม่เกิน 20 ตัว"

					case "alphanum":
						errorMessages["name"] = "username ใช้ได้เฉพาะ a-z A-Z 0-9 และห้ามเว้นวรรค"
					}

				case "Password":

					switch fe.Tag() {
					case "required":
						errorMessages["password"] = "กรุณากรอก password"

					case "max":
						errorMessages["password"] = "password ต้องไม่เกิน 32 ตัว"

					case "password":
						errorMessages["password"] = "password ต้องเป็นภาษาอังกฤษเท่านั้น"
					}
				}
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"errors": errorMessages,
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	// business logic
	user, err := h.usecase.Register(input.Name, input.Password)
	if err != nil {

		if errors.Is(err, customError.ErrDuplicateUser) {
			c.JSON(http.StatusConflict, gin.H{
				"status": 409,
				"message": "ชื่อผู้ใช้นี้ถูกใช้แล้ว",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "ไม่สามารถลงทะเบียนได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": 201,
		"message": "สมัครสมาชิกสำเร็จ",
		"data": user,
	})
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "invalid user id",
		})
		return
	}

	user, err := h.usecase.UserProfile(id)
	if err != nil {

		// user not found
		if errors.Is(err, customError.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"message": "ไม่พบผู้ใช้งาน",
			})
			return
		}

		// internal error
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "cannot get profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": user,
	})
}	

func (h *UserHandler) GetMe(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	userID := int(userIDValue.(float64))

	user, err := h.usecase.UserProfile(userID)
	if err != nil {

		if errors.Is(err, customError.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *UserHandler) Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	token, err := h.usecase.Login(input.Name, input.Password)
	if err != nil {

		if errors.Is(err, customError.ErrInvalidCredential) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"message": "username หรือ password ไม่ถูกต้อง",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message": "cannot login",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"token": token,
	})
}