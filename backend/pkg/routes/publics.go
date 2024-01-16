package routes

import (
	"github.com/gofiber/fiber/v2"

	"backend/app/controllers"
	"backend/pkg/middlewares"
)

func PublicRoutes(a *fiber.App) {

	a.Get("/classes", controllers.GetClasses)
	a.Post("/information", controllers.PostInformation)
	a.Post(
		"/index/:item/:page",
		controllers.PostIndex, middlewares.AuthRequired([]string{}, []string{}),
	)

	resumeGroup := a.Group(
		"/resume",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	resumeGroup.Get("/:action/:person_id", controllers.GetResume)
	resumeGroup.Post("/:action", controllers.PostResume)
	resumeGroup.Delete("/:action/:person_id", controllers.DeleteResume)

	staffGroup := a.Group(
		"/staff/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	staffGroup.Get("/", controllers.GetStaffs)
	staffGroup.Post("/", controllers.PostStaffs)
	staffGroup.Patch("/", controllers.PatchStaffs)
	staffGroup.Delete("/", controllers.DeleteStaffs)

	docsGroup := a.Group("/document/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	docsGroup.Get("/", controllers.GetDocs)
	docsGroup.Post("/", controllers.PostDocs)
	docsGroup.Delete("/", controllers.DeleteDocs)
	docsGroup.Patch("/", controllers.PatchDocs)

	addressGroup := a.Group(
		"/address/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	addressGroup.Get("/", controllers.GetAddress)
	addressGroup.Post("/", controllers.PostAddress)
	addressGroup.Delete("/", controllers.DeleteAddress)
	addressGroup.Patch("/", controllers.PatchAddress)

	contactGroup := a.Group(
		"/contact/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	contactGroup.Get("/", controllers.GetContact)
	contactGroup.Post("/", controllers.PostContact)
	contactGroup.Delete("/", controllers.DeleteContact)
	contactGroup.Patch("/", controllers.PatchContact)

	workGroup := a.Group(
		"/workplace/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	workGroup.Get("/", controllers.GetWorkplace)
	workGroup.Post("/", controllers.PostWorkplace)
	workGroup.Delete("/", controllers.DeleteWorkplace)
	workGroup.Patch("/", controllers.PatchWorkplace)

	affilationGroup := a.Group(
		"/affilation/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	affilationGroup.Get("/", controllers.GetAffilation)
	affilationGroup.Post("/", controllers.PostAffilation)
	affilationGroup.Delete("/", controllers.DeleteAffilation)
	affilationGroup.Patch("/", controllers.PatchAffilation)

	relationGroup := a.Group(
		"/relation/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	relationGroup.Get("/", controllers.GetRelation)
	relationGroup.Post("/", controllers.PostRelation)
	relationGroup.Delete("/", controllers.DeleteRelation)
	relationGroup.Patch("/", controllers.PatchRelation)

	checkGroup := a.Group(
		"/check/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	checkGroup.Get("/", controllers.GetCheck)
	checkGroup.Patch("/", controllers.PatchCheck)
	checkGroup.Delete("/", controllers.DeleteCheck)

	robotGroup := a.Group(
		"/robot",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	robotGroup.Get("/", controllers.GetRobot)
	robotGroup.Post("/", controllers.PostRobot)
	robotGroup.Delete("/", controllers.DeleteRobot)

	investigationGroup := a.Group(
		"/investigation/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	investigationGroup.Get("/", controllers.GetInvestigation)
	investigationGroup.Post("/", controllers.PostInvestigation)
	investigationGroup.Patch("/", controllers.PatchInvestigation)
	investigationGroup.Delete("/", controllers.DeleteInvestigation)

	poligrafGroup := a.Group(
		"/poligraf/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	poligrafGroup.Get("/", controllers.GetPoligraf)
	poligrafGroup.Post("/", controllers.PostPoligraf)
	poligrafGroup.Patch("/", controllers.PatchPoligraf)
	poligrafGroup.Delete("/", controllers.DeletePoligraf)

	inquiryGroup := a.Group(
		"/inquiry/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	inquiryGroup.Get("/", controllers.GetInquiry)
	inquiryGroup.Post("/", controllers.PostInquiry)
	inquiryGroup.Patch("/", controllers.PatchInquiry)
	inquiryGroup.Delete("/", controllers.DeleteInquiry)

	fileGroup := a.Group(
		"/file/:action/:item_id",
		middlewares.AuthRequired([]string{"user"}, []string{"staffsec"}),
	)
	fileGroup.Get("/", controllers.GetFile)
	fileGroup.Post("/", controllers.PostFile)
}
