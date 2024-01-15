package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func PublicRoutes(a *fiber.App) {

	a.Get("/classes", controllers.GetClasses)
	a.Post("/index/:item/:page", middlewares.AuthRequired([]string{}, []string{}), controllers.PostIndex)
	a.Post("information", controllers.PostInformation)

	resumeGroup := a.Group("/resume")
	resumeGroup.Get("/:action/:person_id", middlewares.AuthRequired([]string{}, []string{}), controllers.GetResume)
	resumeGroup.Post("/:action", middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}), controllers.PostResume)
	resumeGroup.Delete("/:action/:person_id", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteResume)

	staffGroup := a.Group("/staff/:action/:item_id")
	staffGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetStaffs)
	staffGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostStaffs)
	staffGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchStaffs)
	staffGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteStaffs)

	docsGroup := a.Group("/document/:action/:item_id")
	docsGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetDocs)
	docsGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostDocs)
	docsGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteDocs)
	docsGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchDocs)

	addressGroup := a.Group("/address/:action/:item_id")
	addressGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetAddress)
	addressGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostAddress)
	addressGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteAddress)
	addressGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchAddress)

	contactGroup := a.Group("/contact/:action/:item_id")
	contactGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetContact)
	contactGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostContact)
	contactGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteContact)
	contactGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchContact)

	workGroup := a.Group("/workplace/:action/:item_id")
	workGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetWorkplace)
	workGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostWorkplace)
	workGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteWorkplace)
	workGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchWorkplace)

	affilationGroup := a.Group("/affilation/:action/:item_id")
	affilationGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetAffilation)
	affilationGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostAffilation)
	affilationGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteAffilation)
	affilationGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchAffilation)

	relationGroup := a.Group("/relation/:action/:item_id")
	relationGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetRelation)
	relationGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostRelation)
	relationGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteRelation)
	relationGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchRelation)

	checkGroup := a.Group("/check/:action/:item_id")
	checkGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetCheck)
	checkGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchCheck)
	checkGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteCheck)

	robotGroup := a.Group("/robot")
	robotGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetRobot)
	robotGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostRobot)
	robotGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteRobot)

	investigationGroup := a.Group("/investigation/:action/:item_id")
	investigationGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetInvestigation)
	investigationGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostInvestigation)
	investigationGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchInvestigation)
	investigationGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteInvestigation)

	poligrafGroup := a.Group("/poligraf/:action/:item_id")
	poligrafGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetPoligraf)
	poligrafGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostPoligraf)
	poligrafGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchPoligraf)
	poligrafGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeletePoligraf)

	inquiryGroup := a.Group("/inquiry/:action/:item_id")
	inquiryGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetInquiry)
	inquiryGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostInquiry)
	inquiryGroup.Patch("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PatchInquiry)
	inquiryGroup.Delete("/", middlewares.AuthRequired([]string{}, []string{}), controllers.DeleteInquiry)

	fileGroup := a.Group("/file/:action/:item_id")
	fileGroup.Get("/", middlewares.AuthRequired([]string{}, []string{}), controllers.GetFile)
	fileGroup.Post("/", middlewares.AuthRequired([]string{}, []string{}), controllers.PostFile)
}
