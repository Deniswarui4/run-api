package main

import (
	"fmt"
	"log"

	"github.com/warui/event-ticketing-api/internal/auth"
	"github.com/warui/event-ticketing-api/internal/config"
	"github.com/warui/event-ticketing-api/internal/database"
	"github.com/warui/event-ticketing-api/internal/models"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("🗑️  Cleaning up existing users...")

	// Delete all users and related data
	// Note: This will cascade delete related records due to foreign key constraints
	
	// Delete organizer balances first
	result := db.Exec("DELETE FROM organizer_balances")
	log.Printf("Deleted %d organizer balance records", result.RowsAffected)

	// Delete tickets
	result = db.Exec("DELETE FROM tickets")
	log.Printf("Deleted %d ticket records", result.RowsAffected)

	// Delete transactions
	result = db.Exec("DELETE FROM transactions")
	log.Printf("Deleted %d transaction records", result.RowsAffected)

	// Delete withdrawals
	result = db.Exec("DELETE FROM withdrawals")
	log.Printf("Deleted %d withdrawal records", result.RowsAffected)

	// Delete ticket types
	result = db.Exec("DELETE FROM ticket_types")
	log.Printf("Deleted %d ticket type records", result.RowsAffected)

	// Delete events
	result = db.Exec("DELETE FROM events")
	log.Printf("Deleted %d event records", result.RowsAffected)

	// Finally delete users
	result = db.Exec("DELETE FROM users")
	log.Printf("Deleted %d user records", result.RowsAffected)

	log.Println("✅ Database cleanup completed!")

	// Create new test admin
	log.Println("👑 Creating new test admin...")

	// Hash password
	hashedPassword, err := auth.HashPassword("Admin@123")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	// Create admin user
	admin := &models.User{
		Email:       "admin@eventtickets.com",
		Password:    hashedPassword,
		FirstName:   "System",
		LastName:    "Administrator",
		Phone:       "+1234567890",
		Role:        models.RoleAdmin,
		IsActive:    true,
		IsVerified:  true, // Admin is pre-verified
	}

	if err := db.Create(admin).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	// Create test moderator
	log.Println("🛡️  Creating test moderator...")

	hashedModPassword, err := auth.HashPassword("Moderator@123")
	if err != nil {
		log.Fatalf("Failed to hash moderator password: %v", err)
	}

	moderator := &models.User{
		Email:       "moderator@eventtickets.com",
		Password:    hashedModPassword,
		FirstName:   "Test",
		LastName:    "Moderator",
		Phone:       "+1234567891",
		Role:        models.RoleModerator,
		IsActive:    true,
		IsVerified:  true, // Pre-verified for testing
	}

	if err := db.Create(moderator).Error; err != nil {
		log.Fatalf("Failed to create moderator user: %v", err)
	}

	// Create test organizer
	log.Println("🎭 Creating test organizer...")

	hashedOrgPassword, err := auth.HashPassword("Organizer@123")
	if err != nil {
		log.Fatalf("Failed to hash organizer password: %v", err)
	}

	organizer := &models.User{
		Email:       "organizer@eventtickets.com",
		Password:    hashedOrgPassword,
		FirstName:   "Test",
		LastName:    "Organizer",
		Phone:       "+1234567892",
		Role:        models.RoleOrganizer,
		IsActive:    true,
		IsVerified:  true, // Pre-verified for testing
	}

	if err := db.Create(organizer).Error; err != nil {
		log.Fatalf("Failed to create organizer user: %v", err)
	}

	// Create organizer balance record
	balance := &models.OrganizerBalance{
		OrganizerID: organizer.ID,
	}
	if err := db.Create(balance).Error; err != nil {
		log.Printf("Warning: Failed to create organizer balance: %v", err)
	}

	fmt.Println("\n🎉 Database reset completed successfully!")
	fmt.Println("==========================================")
	fmt.Println("✅ Test Users Created:")
	fmt.Println("")
	fmt.Println("👑 ADMIN:")
	fmt.Println("   Email: admin@eventtickets.com")
	fmt.Println("   Password: Admin@123")
	fmt.Println("   Status: ✅ Verified & Active")
	fmt.Println("")
	fmt.Println("🛡️  MODERATOR:")
	fmt.Println("   Email: moderator@eventtickets.com")
	fmt.Println("   Password: Moderator@123")
	fmt.Println("   Status: ✅ Verified & Active")
	fmt.Println("")
	fmt.Println("🎭 ORGANIZER:")
	fmt.Println("   Email: organizer@eventtickets.com")
	fmt.Println("   Password: Organizer@123")
	fmt.Println("   Status: ✅ Verified & Active")
	fmt.Println("")
	fmt.Println("==========================================")
	fmt.Println("🚀 Ready to test the platform!")
	fmt.Println("   - All users are pre-verified")
	fmt.Println("   - Can login immediately")
	fmt.Println("   - Database is clean")
}
