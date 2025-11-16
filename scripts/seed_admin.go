package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/warui/event-ticketing-api/internal/auth"
	"github.com/warui/event-ticketing-api/internal/config"
	"github.com/warui/event-ticketing-api/internal/database"
	"github.com/warui/event-ticketing-api/internal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Check if admin already exists
	var existingAdmin models.User
	if err := db.Where("role = ?", models.RoleAdmin).First(&existingAdmin).Error; err == nil {
		fmt.Println("Admin user already exists:")
		fmt.Printf("Email: %s\n", existingAdmin.Email)
		fmt.Printf("Name: %s %s\n", existingAdmin.FirstName, existingAdmin.LastName)
		os.Exit(0)
	}

	// Create admin user
	password := "Admin@123"
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	admin := &models.User{
		Email:      "admin@eventtickets.com",
		Password:   hashedPassword,
		FirstName:  "System",
		LastName:   "Administrator",
		Phone:      "+1234567890",
		Role:       models.RoleAdmin,
		IsActive:   true,
		IsVerified: true,
	}

	if err := db.Create(admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	fmt.Println("✅ Admin user created successfully!")
	fmt.Println("==========================================")
	fmt.Printf("Email: %s\n", admin.Email)
	fmt.Printf("Password: %s\n", password)
	fmt.Println("==========================================")
	fmt.Println("⚠️  Please change the password after first login!")

	// Create a moderator user
	moderatorPassword := "Moderator@123"
	hashedModPassword, _ := auth.HashPassword(moderatorPassword)

	moderator := &models.User{
		Email:      "moderator@eventtickets.com",
		Password:   hashedModPassword,
		FirstName:  "Test",
		LastName:   "Moderator",
		Phone:      "+1234567891",
		Role:       models.RoleModerator,
		IsActive:   true,
		IsVerified: true,
	}

	if err := db.Create(moderator).Error; err != nil {
		log.Printf("Warning: Failed to create moderator user: %v", err)
	} else {
		fmt.Println("\n✅ Moderator user created successfully!")
		fmt.Println("==========================================")
		fmt.Printf("Email: %s\n", moderator.Email)
		fmt.Printf("Password: %s\n", moderatorPassword)
		fmt.Println("==========================================")
	}

	// Create an organizer user
	organizerPassword := "Organizer@123"
	hashedOrgPassword, _ := auth.HashPassword(organizerPassword)

	organizer := &models.User{
		Email:      "organizer@eventtickets.com",
		Password:   hashedOrgPassword,
		FirstName:  "Test",
		LastName:   "Organizer",
		Phone:      "+1234567892",
		Role:       models.RoleOrganizer,
		IsActive:   true,
		IsVerified: true,
	}

	if err := db.Create(organizer).Error; err != nil {
		log.Printf("Warning: Failed to create organizer user: %v", err)
	} else {
		// Create organizer balance
		balance := &models.OrganizerBalance{
			OrganizerID: organizer.ID,
		}
		db.Create(balance)

		fmt.Println("\n✅ Organizer user created successfully!")
		fmt.Println("==========================================")
		fmt.Printf("Email: %s\n", organizer.Email)
		fmt.Printf("Password: %s\n", organizerPassword)
		fmt.Println("==========================================")
	}

	// Create an attendee user
	attendeePassword := "Attendee@123"
	hashedAttPassword, _ := auth.HashPassword(attendeePassword)

	attendee := &models.User{
		Email:      "attendee@eventtickets.com",
		Password:   hashedAttPassword,
		FirstName:  "Test",
		LastName:   "Attendee",
		Phone:      "+1234567893",
		Role:       models.RoleAttendee,
		IsActive:   true,
		IsVerified: true,
	}

	if err := db.Create(attendee).Error; err != nil {
		log.Printf("Warning: Failed to create attendee user: %v", err)
	} else {
		fmt.Println("\n✅ Attendee user created successfully!")
		fmt.Println("==========================================")
		fmt.Printf("Email: %s\n", attendee.Email)
		fmt.Printf("Password: %s\n", attendeePassword)
		fmt.Println("==========================================")
	}

	fmt.Println("\n🎉 Database seeding completed!")
	fmt.Println("You can now start the API server with: go run cmd/api/main.go")
}
