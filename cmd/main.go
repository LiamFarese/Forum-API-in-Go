package main

import (
	"log"
	"net/http"

	"Github.com/Synoptic2023/internal/comment"
	"Github.com/Synoptic2023/internal/listing"
	"Github.com/Synoptic2023/internal/post"
	"Github.com/Synoptic2023/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {

	//start database connection
	db, err := sqlx.Connect("pgx", "postgresql://root:secret@localhost:5432/synoptic?sslmode=disable")

	if err != nil {
		log.Fatalf("could not connect to the database %s", err)
	}

	log.Printf("connected to database")

	//Initialise router
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value for the Max-Age header
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Use(middleware.Heartbeat("/"))

	//user dependancies
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	//register and login routes
	r.Post("/register", userHandler.CreateUser)
	r.Post("/login", userHandler.Login)

	//user related routes
	r.Get("/users", userHandler.ListUsers)
	r.Get("/profile/{userId}", userHandler.UserProfile)

	//listing dependancies
	listingRepo := listing.NewListingRepository(db)
	listingService := listing.NewListingService(listingRepo)
	listingHandler := listing.NewListingHandler(listingService)

	//listing routes
	r.Post("/listing/create", listingHandler.CreateListing)
	r.Get("/listing", listingHandler.ListListings)
	r.Get("/listing/{listingId}", listingHandler.GetListing)
	r.Patch("/listing/{listingId}", listingHandler.CloseListing)

	//post dependancies
	postRepo := post.NewPosteRepository(db)
	postService := post.NewPostService(postRepo)
	postHandler := post.NewPostHandler(postService)

	//post routes
	r.Post("/post", postHandler.CreatePost)
	r.Get("/posts", postHandler.ListPosts)
	r.Get("/post/{postId}", postHandler.GetPost)

	//comment dependanies
	commentRepo := comment.NewCommentRepository(db)
	commentServie := comment.NewCommentService(commentRepo)
	commentHandler := comment.NewCommentHandler(commentServie)

	//comment routes
	r.Post("/comment", commentHandler.CreateComment)
	r.Post("/reply", commentHandler.Reply)
	r.Get("/comments/{postId}", commentHandler.GetCommentsFromPost)

	//start server
	err = http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalf("server has stopped %s", err)
	}
}
