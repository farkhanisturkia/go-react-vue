package database

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"go-react-vue/backend/helpers"
	"go-react-vue/backend/models"
)

func Seed() {
	fmt.Println("Starting database seeding...")

	SeedUsers()
	adminID := uint(1) // asumsi admin pertama ID = 1

	SeedCourses(adminID)
	SeedEnrollments(adminID)

	fmt.Println("Seeding completed!")
}

// ── Users ───────────────────────────────────────────────
func SeedUsers() {
	now := time.Now()

	// 1 admin
	admin := models.User{
		Name:     "Admin Utama",
		Username: "admin",
		Email:    "admin@example.com",
		Role:     "admin",
		Password: helpers.HashPassword("admin123"),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := DB.Create(&admin).Error; err != nil {
		log.Printf("Failed to create admin: %v", err)
	} else {
		fmt.Println("Admin created (ID:", admin.Id, ")")
	}

	// 80 user
	var users []models.User
	for i := 1; i <= 80; i++ {
		username := fmt.Sprintf("user%03d", i)
		email := fmt.Sprintf("user%03d@example.com", i)

		user := models.User{
			Name:      fmt.Sprintf("User %03d", i),
			Username:  username,
			Email:     email,
			Role:      "user",
			Password:  helpers.HashPassword("password123"),
			CreatedAt: now,
			UpdatedAt: now,
		}
		users = append(users, user)
	}

	result := DB.Create(&users)
	if result.Error != nil {
		log.Printf("Failed to seed regular users: %v", result.Error)
	} else {
		fmt.Printf("Seeded %d regular users\n", len(users))
	}
}

// ── Courses ─────────────────────────────────────────────
func SeedCourses(creatorID uint) {
	now := time.Now()
	courseTitles := []string{
		"Belajar Golang dari Nol sampai Mahir",
		"REST API dengan Gin & GORM",
		"Microservices dengan Go dan gRPC",
		"Clean Architecture di Golang",
		"Testing di Go: Unit, Integration, E2E",
		"Go Concurrency: Goroutines & Channels",
		"Docker & Kubernetes untuk Developer Go",
		"Web Scraping dengan Colly & Go",
		"Build CLI Tools Profesional dengan Cobra",
		"GraphQL Server dengan gqlgen",
		"Real-time App dengan WebSocket & Go",
		"Machine Learning Dasar dengan Go",
		"Blockchain & Smart Contract di Go",
		"Go untuk Backend High-Performance",
		"Advanced Error Handling di Go",
		"Go dan PostgreSQL: Query Optimization",
		"Build SaaS dengan Go & Next.js",
		"Go Security Best Practices 2026",
		"Event-Driven Architecture dengan Go",
		"Go di Edge Computing & IoT",
	}

	var courses []models.Course
	for i, title := range courseTitles {
		course := models.Course{
			Title:       title,
			Description: generateRandomDescription(i),
			Price:       150000,
			CreatorID:   creatorID,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		courses = append(courses, course)
	}

	result := DB.Create(&courses)
	if result.Error != nil {
		log.Fatalf("Failed to seed courses: %v", result.Error)
	}
	fmt.Printf("Seeded %d courses (creator: admin ID %d)\n", len(courses), creatorID)
}

// ── Enrollments (UserCourse) ────────────────────────────
func SeedEnrollments(creatorID uint) {
	var courseIDs []uint
	DB.Model(&models.Course{}).Pluck("id", &courseIDs)

	var userIDs []uint
	DB.Model(&models.User{}).Where("role = ?", "user").Pluck("id", &userIDs)

	if len(courseIDs) == 0 || len(userIDs) == 0 {
		fmt.Println("Tidak ada course atau user untuk enrollment")
		return
	}

	var enrollments []models.UserCourse
	now := time.Now()

	for _, courseID := range courseIDs {
		// Setiap course diikuti 50–68 orang secara random
		numParticipants := rand.Intn(19) + 50 // 50 sampai 68
		shuffledUsers := shuffleUserIDs(userIDs)

		for j := 0; j < numParticipants && j < len(shuffledUsers); j++ {
			enrollment := models.UserCourse{
				CourseID:      courseID,
				CreatorID:     creatorID,
				ParticipantID: shuffledUsers[j],
				CreatedAt:     now,
				UpdatedAt:     now,
			}
			enrollments = append(enrollments, enrollment)
		}
	}

	result := DB.Create(&enrollments)
	if result.Error != nil {
		log.Printf("Failed to seed enrollments: %v", result.Error)
	} else {
		fmt.Printf("Seeded %d enrollments\n", len(enrollments))
	}
}

// ── Helper kecil ────────────────────────────────────────
func shuffleUserIDs(ids []uint) []uint {
	rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
	return ids
}

func generateRandomDescription(idx int) string {
	templates := []string{
		"Kursus lengkap untuk menguasai %s dari dasar hingga level profesional. Cocok untuk pemula hingga intermediate.",
		"Belajar %s secara intensif dengan proyek nyata dan best practice industri tahun 2026.",
		"Pahami %s secara mendalam + optimasi performa + deployment modern.",
		"Transformasi skill %s Anda dalam waktu singkat dengan materi up-to-date.",
	}
	return fmt.Sprintf(templates[idx%len(templates)], "Golang & backend engineering")
}