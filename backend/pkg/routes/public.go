package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func PublicRoutes(a *fiber.App) {

	a.Get("/classes", controllers.GetClasses)
	a.Post("/index/:item/:page", middlewares.JWTProtected(), controllers.PostIndex)
	a.Post("information", controllers.PostInformation)

	resumeGroup := a.Group("/resume")
	resumeGroup.Get("/:action/:person_id", controllers.GetResume)
	resumeGroup.Post("/:action", middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}), controllers.PostResume)
	resumeGroup.Delete("/:action/:person_id", controllers.DeleteResume)

	staffGroup := a.Group("/staff/:action/:item_id")
	staffGroup.Get("/", middlewares.JWTProtected(), controllers.GetStaffs)
	staffGroup.Post("/", middlewares.JWTProtected(), controllers.PostStaffs)
	staffGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchStaffs)
	staffGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteStaffs)

	docsGroup := a.Group("/document/:action/:item_id")
	docsGroup.Get("/", middlewares.JWTProtected(), controllers.GetDocs)
	docsGroup.Post("/", middlewares.JWTProtected(), controllers.PostDocs)
	docsGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteDocs)
	docsGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchDocs)

	addressGroup := a.Group("/address/:action/:item_id")
	addressGroup.Get("/", middlewares.JWTProtected(), controllers.GetAddress)
	addressGroup.Post("/", middlewares.JWTProtected(), controllers.PostAddress)
	addressGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteAddress)
	addressGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchAddress)

	contactGroup := a.Group("/contact/:action/:item_id")
	contactGroup.Get("/", middlewares.JWTProtected(), controllers.GetContact)
	contactGroup.Post("/", middlewares.JWTProtected(), controllers.PostContact)
	contactGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteContact)
	contactGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchContact)

	workGroup := a.Group("/workplace/:action/:item_id")
	workGroup.Get("/", middlewares.JWTProtected(), controllers.GetWorkplace)
	workGroup.Post("/", middlewares.JWTProtected(), controllers.PostWorkplace)
	workGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteWorkplace)
	workGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchWorkplace)

	affilationGroup := a.Group("/affilation/:action/:item_id")
	affilationGroup.Get("/", middlewares.JWTProtected(), controllers.GetAffilation)
	affilationGroup.Post("/", middlewares.JWTProtected(), controllers.PostAffilation)
	affilationGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteAffilation)
	affilationGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchAffilation)

	relationGroup := a.Group("/relation/:action/:item_id")
	relationGroup.Get("/", middlewares.JWTProtected(), controllers.GetRelation)
	relationGroup.Post("/", middlewares.JWTProtected(), controllers.PostRelation)
	relationGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteRelation)
	relationGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchRelation)

	checkGroup := a.Group("/check/:action/:item_id")
	checkGroup.Get("/", middlewares.JWTProtected(), controllers.GetCheck)
	checkGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchCheck)
	checkGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteCheck)

	robotGroup := a.Group("/robot")
	robotGroup.Get("/", middlewares.JWTProtected(), controllers.GetRobot)
	robotGroup.Post("/", middlewares.JWTProtected(), controllers.PostRobot)
	robotGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteRobot)

	investigationGroup := a.Group("/investigation/:action/:item_id")
	investigationGroup.Get("/", middlewares.JWTProtected(), controllers.GetInvestigation)
	investigationGroup.Post("/", middlewares.JWTProtected(), controllers.PostInvestigation)
	investigationGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchInvestigation)
	investigationGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteInvestigation)

	poligrafGroup := a.Group("/poligraf/:action/:item_id")
	poligrafGroup.Get("/", middlewares.JWTProtected(), controllers.GetPoligraf)
	poligrafGroup.Post("/", middlewares.JWTProtected(), controllers.PostPoligraf)
	poligrafGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchPoligraf)
	poligrafGroup.Delete("/", middlewares.JWTProtected(), controllers.DeletePoligraf)

	inquiryGroup := a.Group("/inquiry/:action/:item_id")
	inquiryGroup.Get("/", middlewares.JWTProtected(), controllers.GetInquiry)
	inquiryGroup.Post("/", middlewares.JWTProtected(), controllers.PostInquiry)
	inquiryGroup.Patch("/", middlewares.JWTProtected(), controllers.PatchInquiry)
	inquiryGroup.Delete("/", middlewares.JWTProtected(), controllers.DeleteInquiry)

	fileGroup := a.Group("/file/:action/:item_id")
	fileGroup.Get("/", middlewares.JWTProtected(), controllers.GetFile)
	fileGroup.Post("/", middlewares.JWTProtected(), controllers.PostFile)
}
