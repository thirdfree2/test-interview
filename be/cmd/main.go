package main

import (
	"be-interview-app/config"
	"be-interview-app/internal/delivery/http"
	"be-interview-app/internal/delivery/validate"
	"be-interview-app/internal/repository"
	"be-interview-app/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	cfg := config.LoadConfig()
	
	db, err := repository.InitDB(cfg)
	if err != nil {
		log.Fatalf("Critical Error: %v", err)
	}
	

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("password", validate.PasswordValidator)
	}
	
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
    userUsecase := usecase.NewUserUsecase(userRepo)
    userHandler := http.NewUserHandler(userUsecase)

	http.MapRoutes(r, userHandler)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "UP",
			"database": "Connected",
		})
	})

	log.Printf("Server running at %s...", cfg.AppPort)

	// 3. รัน Server (ค่าเริ่มต้นคือ port 8080)
	r.Run() 
}