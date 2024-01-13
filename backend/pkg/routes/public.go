package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func PublicRoutes(a *fiber.App) {

	a.Get("/classes", controllers.GetClasses)
	a.Post("/index/:item/:page", middlewares.JWTProtected(), controllers.PostIndex)

	resumeGroup := a.Group("/resume")
	resumeGroup.Post("/:action", middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}), controllers.PostResume)
	resumeGroup.Get("/:action/:person_id", controllers.GetResume)
	resumeGroup.Delete("/:action/:person_id", controllers.DeleteResume)

	staffGroup := a.Group("/staff/:action/:item_id")
	staffGroup.Get("/", middlewares.JWTProtected(), controllers.GetStaffs)
	staffGroup.Post("/", middlewares.JWTProtected(), controllers.PostStaffs)
	staffGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchStaffs)
	staffGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteStaffs)

	docsGroup := a.Group("/document/:action/:item_id")
	docsGroup.Post("/", middlewares.JWTProtected(), controllers.PostDocs)
	docsGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteDocs)
	docsGroup.Get("/", middlewares.JWTProtected(), controllers.GetDocs)
	docsGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchDocs)

	addressGroup := a.Group("/address/:action/:item_id")
	addressGroup.Post("/", middlewares.JWTProtected(), controllers.PostAddress)
	addressGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteAddress)
	addressGroup.Get("/", middlewares.JWTProtected(), controllers.GetAddress)
	addressGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchAddress)

	contactGroup := a.Group("/contact/:action/:item_id")
	contactGroup.Post("/", middlewares.JWTProtected(), controllers.PostContact)
	contactGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteContact)
	contactGroup.Get("/", middlewares.JWTProtected(), controllers.GetContact)
	contactGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchContact)

	workGroup := a.Group("/workplace/:action/:item_id")
	workGroup.Post("/", middlewares.JWTProtected(), controllers.PostWorkplace)
	workGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteWorkplace)
	workGroup.Get("/", middlewares.JWTProtected(), controllers.GetWorkplace)
	workGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchWorkplace)

	affilationGroup := a.Group("/affilation/:action/:item_id")
	affilationGroup.Post("/", middlewares.JWTProtected(), controllers.PostAffilation)
	affilationGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteAffilation)
	affilationGroup.Get("/", middlewares.JWTProtected(), controllers.GetAffilation)
	affilationGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchAffilation)

	relationGroup := a.Group("/relation/:action/:item_id")
	relationGroup.Post("/", middlewares.JWTProtected(), controllers.PostRelation)
	relationGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteRelation)
	relationGroup.Get("/", middlewares.JWTProtected(), controllers.GetRelation)
	relationGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchRelation)

}
