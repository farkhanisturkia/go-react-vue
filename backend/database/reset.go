package database

import (
	"fmt"
	"log"

	"go-react-vue/backend/models"
	"gorm.io/gorm"
)

func ResetAll(seedAfterReset bool) {
	fmt.Println("🔥 Starting full database reset (users, courses, user_courses)...")

	// Urutan PENTING: hapus tabel ANAK dulu (yang punya foreign key), baru tabel INDUK
	tableNames := []string{"user_courses", "courses", "users"} // nama tabel sesuai GORM default

	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	defer DB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	for _, tableName := range tableNames {
		fmt.Printf("🗑️  Clearing table: %s\n", tableName)
		
		// Coba TRUNCATE dulu (lebih cepat & reset auto-increment)
		err := DB.Exec("TRUNCATE TABLE `"+tableName+"`").Error
		if err != nil {
			fmt.Printf("  TRUNCATE failed, trying DELETE...\n")
			// Fallback: DELETE semua data
			switch tableName {
				case "users":
					err = DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{}).Error
				case "courses":
					err = DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Course{}).Error
				case "user_courses":
					err = DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.UserCourse{}).Error
			}
			
			if err != nil {
				log.Fatalf("❌ Failed to clear table %s: %v", tableName, err)
			}
			
			// Reset auto-increment manual (MySQL)
			DB.Exec("ALTER TABLE `"+tableName+"` AUTO_INCREMENT = 1")
		} else {
			fmt.Printf("  ✅ %s cleared with TRUNCATE\n", tableName)
		}
	}

	// Reset auto-increment untuk PostgreSQL (opsional, uncomment jika pakai PG)
	// ResetPostgreSQLSequences()

	fmt.Println("✅ Full database reset completed!")

	if seedAfterReset {
		fmt.Println("🌱 Starting re-seeding after reset...")
		Seed()
	} else {
		fmt.Println("⏹️  Reset completed. No seeding performed.")
	}
}

// Untuk PostgreSQL (uncomment jika pakai PostgreSQL)
func ResetPostgreSQLSequences() {
	sequences := []string{
		"users_id_seq",
		"courses_id_seq", 
		"user_courses_id_seq",
	}
	for _, seq := range sequences {
		DB.Exec("ALTER SEQUENCE IF EXISTS " + seq + " RESTART WITH 1")
	}
}