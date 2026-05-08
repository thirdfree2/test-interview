package repository

import (
	"be-interview-app/config"
	"be-interview-app/internal/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB ทำหน้าที่เชื่อมต่อ Database และเตรียมความพร้อมของ Table
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถเชื่อมต่อฐานข้อมูลได้: %w", err)
	}

	log.Println("เชื่อมต่อฐานข้อมูล PostgreSQL สำเร็จ")

	// สั่ง AutoMigrate เพื่อสร้างหรือ Update Table อัตโนมัติ
	// GORM จะสร้าง Table ชื่อ "users" ให้จาก struct User
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return nil, fmt.Errorf("การทำ Auto Migration ล้มเหลว: %w", err)
	}

	log.Println("สร้าง/อัปเดต Table สำเร็จ")
	
	return db, nil
}