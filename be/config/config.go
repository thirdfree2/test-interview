package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config เก็บค่ากำหนดค่าต่างๆ ของระบบ
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

// LoadConfig อ่านค่าจาก .env หรือ Environment Variables
func LoadConfig() *Config {
	// ลองโหลด .env (ถ้าไม่มีจะไม่ Error เพราะบน Production อาจจะตั้งผ่าน OS)
	if err := godotenv.Load(); err != nil {
		log.Println("แจ้งเตือน: ไม่พบไฟล์ .env จะใช้ค่าจาก Environment ของระบบแทน")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "postgres"),
		AppPort:    getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}