package main

import (
	"clean_architecture/internal/config"
	"clean_architecture/internal/entity"
	"clean_architecture/pkg/database"
	"fmt"
	"log"
	"reflect"
	"time"

	"gorm.io/gorm"
)

func main() {
	db, err := database.NewPostgresDB(config.Load())
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Get underlying sql.DB to handle connection properly
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
		return
	}
	defer sqlDB.Close()

	log.Println("Starting migration...")

	err = AutoMigrate(db, &entity.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
		return
	}

	// err = DropTable(db, &entity.User{})
	// if err != nil {
	// 	log.Fatal("Drop table failed:", err)
	// 	return
	// }

	log.Println("Migration completed successfully")

}

func AutoMigrate(db *gorm.DB, entity ...interface{}) error {
	if err := db.AutoMigrate(entity...); err != nil {
		return err
	}

	for _, v := range entity {
		tableName := reflect.TypeOf(v).Elem().Name()
		fmt.Printf("✅ Migrate Table %s at %s\n",
			tableName,
			time.Now().Format("2006-01-02 15:04:05"))
	}

	return nil
}

func DropTable(db *gorm.DB, entity ...interface{}) error {
	for _, v := range entity {
		if err := db.Migrator().DropTable(v); err != nil {
			return err
		}
		tableName := reflect.TypeOf(v).Elem().Name()
		fmt.Printf("🗑️ Drop Table %s at %s\n",
			tableName,
			time.Now().Format("2006-01-02 15:04:05"))
	}
	return nil
}
