package routes

import (
	"hera/app/controllers/article"
	"hera/pkg/middleware"

	fiber "github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")

	// Routes for POST method:
	route.Post("/article", middleware.JWTProtected(), article.CreateArticle)
	route.Get("/article/:id", article.GetArticleByID)

}
