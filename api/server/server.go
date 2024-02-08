package server

import (
	"359/ports"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	handler ports.Handler
}

func NewService(handler ports.Handler) *Server {
	return &Server{handler: handler}
}

// Initialize initializes the server by setting up the routes and handlers for various endpoints.
func (server *Server) Initialize() {
	app := fiber.New()
	app.Use(cors.New())

	// Endpoints work as follows:
	// app.Group("/path") creates a group of endpoints with the same prefix
	// for example, app.Group("/users") creates a group of endpoints that all start with /users

	// Login endpoints
	login := app.Group("/login")
	login.Post("/", server.handler.Login)           // Handles user login
	login.Post("/admin", server.handler.LoginAdmin) // Handles admin login

	// User endpoints
	users := app.Group("/users")

	// Admin endpoints
	admin := users.Group("/admin")
	admin.Get("/", server.handler.GetAdmin)                         // Retrieves admin information
	admin.Get("/petsNumber", server.handler.GetNumberOfCatsAndDogs) // Retrieves the number of cats and dogs
	admin.Get("/money", server.handler.GetMoney)                    // Retrieves (as json) the total amount of money for keepers and the app {keeper: money[0], app: money[1]}
	admin.Get("/usersNumber", server.handler.GetNumberOfUsers)      // Retrieves (as json) the number of owners and keepers {owners: owners, keepers: keepers}

	// Owner endpoints
	owners := users.Group("/owners")
	owners.Get("/", server.handler.GetOwners)                       // Retrieves all owners
	owners.Post("/", server.handler.RegisterOwner)                  // Registers a new owner
	owners.Put("/:id", server.handler.UpdateOwner)                  // Updates owner information
	owners.Get("/:id", server.handler.GetOwner)                     // Retrieves owner information by ID
	owners.Delete("/:id", server.handler.DeleteOwner)               // Deletes an owner
	owners.Get("/:id/findKeepers", server.handler.AvailableKeepers) // Finds available keepers for an owner

	// base on orderBy param in the query, it will order the keepers by distance or duration to get there by car
	// using the google maps api
	owners.Get("/:id/orderKeepers", server.handler.OrderClosestKeepers) // Orders closest keepers for an owner
	owners.Get("/:id/bookings", server.handler.GetBookingsByOwner)      // Retrieves bookings for an owner

	// Keeper endpoints
	keepers := users.Group("/keepers")
	keepers.Get("/", server.handler.GetKeepers)                                    // Retrieves all keepers
	keepers.Post("/", server.handler.RegisterKeeper)                               // Registers a new keeper
	keepers.Put("/:id", server.handler.UpdateKeeper)                               // Updates keeper information
	keepers.Get("/:id", server.handler.GetKeeper)                                  // Retrieves keeper information by ID
	keepers.Delete("/:id", server.handler.DeleteKeeper)                            // Deletes a keeper
	keepers.Get("/:id/bookings", server.handler.GetBookingsByKeeperId)             // Retrieves bookings for a keeper
	keepers.Get("/:id/reviews", server.handler.GetReviewsByKeeper)                 // Retrieves reviews for a keeper
	keepers.Get("/:id/bookingsNumber", server.handler.GetBookingsNumberByKeeperId) // Retrieves the number of bookings for a keeper
	keepers.Get("/:id/petKeepersDays", server.handler.GetPetKeepersDays)           // Retrieves the number of days a pet has been with keepers

	// Pet endpoints
	pets := app.Group("/pets")
	pets.Get("/", server.handler.GetPets)                           // Retrieves all pets
	pets.Post("/", server.handler.RegisterPet)                      // Registers a new pet
	pets.Put("/:id", server.handler.UpdatePet)                      // Updates pet information
	pets.Get("/:id", server.handler.GetPet)                         // Retrieves pet information by ID
	pets.Get("/:type/:breed", server.handler.GetPetsByTypeAndBreed) // Retrieves pets by type and breed
	pets.Get("/:type/", server.handler.GetNumberOfCats)             // Retrieves the number of cats
	pets.Get("/:type/", server.handler.GetNumberOfDogs)             // Retrieves the number of dogs

	// Booking endpoints
	bookings := app.Group("/bookings")
	bookings.Get("/", server.handler.GetBookings)         // Retrieves all bookings
	bookings.Post("/", server.handler.CreateBooking)      // Creates a new booking
	bookings.Put("/:id", server.handler.UpdateBooking)    // Updates booking information
	bookings.Get("/:id", server.handler.GetBooking)       // Retrieves booking information by ID
	bookings.Delete("/:id", server.handler.DeleteBooking) // Deletes a booking

	// Review endpoints
	reviews := app.Group("/reviews")
	reviews.Post("/", server.handler.CreateReview) // Creates a new review

	// Message endpoints
	messages := app.Group("/messages")
	messages.Post("/", server.handler.CreateMessage)        // Creates a new message
	messages.Get("/", server.handler.GetMessagesByUsername) // Retrieves messages by username

	// Additional endpoints
	app.Delete("petDeletion/:pet_id", server.handler.DeletePet)          // Deletes a pet by ID
	app.Put("petWeight/:pet_id/:weight", server.handler.UpdatePetWeight) // Updates the weight of a pet

	log.Fatal(app.Listen(":3000"))
}
