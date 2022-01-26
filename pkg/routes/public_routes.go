package routes

import (
	"hera/app/controllers/article"
	"hera/app/controllers/auth"

	fiber "github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/author")

	// get active articles
	route.Get("/article", article.GetArticles) // get list of all books

	// Routes for POST method:
	route.Post("/sign/up", auth.AuthorSignUp) // register a new user
	route.Post("/sign/in", auth.UserSignIn)   // auth, return Access & Refresh tokens
}
