## ADMIN

1. na diagrafonte kai ta bookings otan diagrafete user
2. logout ston admin


## OWNER

1. na xwrisoume keepers home / owners home  (done)
2. (isws) na ftiaksoume to UI sto keepers home (done)
3. petid
4. onomata sta tables

## KEEPER
1. bookings accept/reject (done)
2. see bookings (done)
4. messages
5. reviews
6. messages

front owner booking 2 keeper

login := app.Group("/login")
	login.Post("/", server.handler.Login)

	login.Post("/admin", server.handler.LoginAdmin)

	users := app.Group("/users")

	admin := users.Group("/admin")
	admin.Get("/", server.handler.GetAdmin)
	endpoint used for getting admin

	admin.Get("/petsNumber", server.handler.GetNumberOfCatsAndDogs)
	endpoint used for getting the number of cats and dogs

	admin.Get("/money", server.handler.GetMoney)
	endpoint used for getting the money

	admin.Get("/usersNumber", server.handler.GetNumberOfUsers)
	endpoint used for getting the number of users

	owners := users.Group("/owners")
	owners.Get("/", server.handler.GetOwners)
	endpoint used for getting owners

	owners.Post("/", server.handler.RegisterOwner)
	endpoint used for registering owners

	owners.Put("/:id", server.handler.UpdateOwner)
	endpoint used for updating owners

	owners.Get("/:id", server.handler.GetOwner)
	endpoint used for getting owners

	owners.Delete("/:id", server.handler.DeleteOwner)
	endpoint used for deleting owners

	owners.Get("/:id/findKeepers", server.handler.AvailableKeepers)
	endpoint used for finding keepers

	owners.Get("/:id/orderKeepers", server.handler.OrderClosestKeepers)
	endpoint used for sorting keepers by distance

	owners.Get("/:id/bookings", server.handler.GetBookingsByOwner)
	endpoint used for getting bookings

	keepers := users.Group("/keepers")
	keepers.Get("/", server.handler.GetKeepers)
	endpoint used for getting keepers

	keepers.Post("/", server.handler.RegisterKeeper)
	endpoint used for registering keepers

	keepers.Put("/:id", server.handler.UpdateKeeper)
	endpoint used for updating keepers

	keepers.Get("/:id", server.handler.GetKeeper)
	endpoint used for getting keeper

	keepers.Delete("/:id", server.handler.DeleteKeeper)
	endpoint used for deleting keeper

	keepers.Get("/:id/bookings", server.handler.GetBookingsByKeeperId)
	endpoint used for getting bookings

	pets := app.Group("/pets")
	pets.Get("/", server.handler.GetPets)
	endpoint used for getting pets

	pets.Post("/", server.handler.RegisterPet)
	endpoint used for registering pets

	pets.Put("/:id", server.handler.UpdatePet)
	endpoint used for updating pets

	pets.Get("/:id", server.handler.GetPet)
	endpoint used for getting pets

	pets.Get("/:type/:breed", server.handler.GetPetsByTypeAndBreed)
	endpoint used for getting pets by type and breed

	pets.Get("/:type/", server.handler.GetNumberOfCats)
	endpoint used for getting the number of cats

	pets.Get("/:type/", server.handler.GetNumberOfDogs)
	endpoint used for getting the number of dogs

	bookings := app.Group("/bookings")
	bookings.Get("/", server.handler.GetBookings)
	endpoint used for getting bookings

	bookings.Post("/", server.handler.CreateBooking)
    endpoint used for creating bookings

	bookings.Put("/:id", server.handler.UpdateBooking)
    endpoint used for updating bookings

	bookings.Get("/:id", server.handler.GetBooking)
    endpoint used for getting bookings

	bookings.Delete("/:id", server.handler.DeleteBooking)
    endpoint used for deleting bookings

	reviews := app.Group("/reviews")


	reviews.Post("/", server.handler.CreateReview)
    endpoint used for creating reviews

	reviews.Get("/:id", server.handler.GetReviewsByKeeper)
    endpoint used for getting reviews
