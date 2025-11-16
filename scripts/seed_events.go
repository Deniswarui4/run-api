package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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

	log.Println("Seeding sample events and categories...")

	// Get organizer user
	var organizer models.User
	if err := db.Where("email = ?", "organizer@eventtickets.com").First(&organizer).Error; err != nil {
		log.Fatalf("Organizer user not found. Please run seed_admin.go first: %v", err)
	}

	// Create categories
	categories := []models.Category{
		{
			ID:       uuid.New(),
			Name:     "Music",
			Icon:     "🎵",
			Color:    "#FF6B6B",
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Sports",
			Icon:     "⚽",
			Color:    "#4ECDC4",
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Technology",
			Icon:     "💻",
			Color:    "#45B7D1",
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Business",
			Icon:     "💼",
			Color:    "#96CEB4",
			IsActive: true,
		},
		{
			ID:       uuid.New(),
			Name:     "Food & Drink",
			Icon:     "🍕",
			Color:    "#FFA07A",
			IsActive: true,
		},
	}

	for _, category := range categories {
		var existing models.Category
		if err := db.Where("name = ?", category.Name).First(&existing).Error; err != nil {
			// Category doesn't exist, create it
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create category %s: %v", category.Name, err)
			} else {
				log.Printf("✅ Created category: %s", category.Name)
			}
		} else {
			log.Printf("⏭️  Category already exists: %s", category.Name)
		}
	}

	// Create sample events
	now := time.Now()
	events := []models.Event{
		{
			ID:          uuid.New(),
			OrganizerID: organizer.ID,
			Title:       "Summer Music Festival 2025",
			Description: "Join us for the biggest music festival of the year! Featuring top artists from around the world.",
			Category:    "Music",
			Venue:       "Central Park Arena",
			Address:     "123 Park Avenue",
			City:        "New York",
			Country:     "USA",
			StartDate:   now.AddDate(0, 2, 0),  // 2 months from now
			EndDate:     now.AddDate(0, 2, 3),  // 3 days event
			Status:      models.EventStatusPublished,
			IsFeatured:  true,
		},
		{
			ID:          uuid.New(),
			OrganizerID: organizer.ID,
			Title:       "Tech Conference 2025",
			Description: "The premier technology conference bringing together innovators, developers, and industry leaders.",
			Category:    "Technology",
			Venue:       "Convention Center",
			Address:     "456 Tech Boulevard",
			City:        "San Francisco",
			Country:     "USA",
			StartDate:   now.AddDate(0, 1, 15), // 1.5 months from now
			EndDate:     now.AddDate(0, 1, 16),
			Status:      models.EventStatusPublished,
			IsFeatured:  true,
		},
		{
			ID:          uuid.New(),
			OrganizerID: organizer.ID,
			Title:       "Champions League Final",
			Description:  "Watch the biggest football match of the year live at the stadium!",
			Category:    "Sports",
			Venue:       "National Stadium",
			Address:     "789 Sports Complex",
			City:        "London",
			Country:     "UK",
			StartDate:   now.AddDate(0, 3, 10),
			EndDate:     now.AddDate(0, 3, 10),
			Status:      models.EventStatusPublished,
			IsFeatured:  true,
		},
		{
			ID:          uuid.New(),
			OrganizerID: organizer.ID,
			Title:       "Food & Wine Expo",
			Description: "Taste the finest cuisines and wines from around the world at this exclusive expo.",
			Category:    "Food & Drink",
			Venue:       "Grand Exhibition Hall",
			Address:     "321 Culinary Street",
			City:        "Paris",
			Country:     "France",
			StartDate:   now.AddDate(0, 1, 20),
			EndDate:     now.AddDate(0, 1, 22),
			Status:      models.EventStatusPublished,
			IsFeatured:  false,
		},
		{
			ID:          uuid.New(),
			OrganizerID: organizer.ID,
			Title:       "Startup Pitch Competition",
			Description: "Watch innovative startups pitch their ideas to top investors. Network with entrepreneurs and VCs.",
			Category:    "Business",
			Venue:       "Innovation Hub",
			Address:     "555 Startup Lane",
			City:        "Austin",
			Country:     "USA",
			StartDate:   now.AddDate(0, 2, 5),
			EndDate:     now.AddDate(0, 2, 5),
			Status:      models.EventStatusPublished,
			IsFeatured:  false,
		},
	}

	for _, event := range events {
		if err := db.Create(&event).Error; err != nil {
			log.Printf("Failed to create event %s: %v", event.Title, err)
			continue
		}

		// Create ticket types for each event
		ticketTypes := []models.TicketType{
			{
				ID:          uuid.New(),
				EventID:     event.ID,
				Name:        "General Admission",
				Description: "Standard entry ticket",
				Price:       50.00,
				Quantity:    1000,
				Sold:        0,
				IsActive:    true,
			},
			{
				ID:          uuid.New(),
				EventID:     event.ID,
				Name:        "VIP",
				Description: "VIP access with premium seating and backstage pass",
				Price:       150.00,
				Quantity:    100,
				Sold:        0,
				IsActive:    true,
			},
			{
				ID:          uuid.New(),
				EventID:     event.ID,
				Name:        "Early Bird",
				Description: "Limited early bird special pricing",
				Price:       35.00,
				Quantity:    200,
				Sold:        0,
				IsActive:    true,
			},
		}

		for _, ticketType := range ticketTypes {
			if err := db.Create(&ticketType).Error; err != nil {
				log.Printf("Failed to create ticket type for event %s: %v", event.Title, err)
			}
		}

		featuredStatus := ""
		if event.IsFeatured {
			featuredStatus = " (Featured)"
		}
		log.Printf("✅ Created event: %s%s with %d ticket types", event.Title, featuredStatus, len(ticketTypes))
	}

	fmt.Println("\n🎉 Sample events and categories seeded successfully!")
	fmt.Println("==========================================")
	fmt.Println("Created:")
	fmt.Println("- 5 Categories")
	fmt.Println("- 5 Events (3 featured)")
	fmt.Println("- 15 Ticket Types")
	fmt.Println("==========================================")
	fmt.Println("\n✅ You can now browse events at http://localhost:3000")
}
