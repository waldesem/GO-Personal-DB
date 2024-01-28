package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"backend/app/models"
	"backend/pkg/configs"
	"backend/pkg/middlewares"
	"backend/pkg/routes"
	"backend/pkg/utils"
	"backend/platform/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Print("No .env file found")
	}
	cli := &cli.App{
		Name:  "staffsec cli",
		Usage: "Command line interface for StaffSec",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create database and default data",
				Action: func(c *cli.Context) error {
					createDefault()
					return nil
				},
			},
			{
				Name:  "start",
				Usage: "Start server",
				Action: func(c *cli.Context) error {
					start()
					return nil
				},
			},
			{
				Name:  "test",
				Usage: "Test cli command",
				Action: func(c *cli.Context) error {
					log.Println("Test")
					return nil
				},
			},
		},
	}

	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func start() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)
	app.Use(cache.New())
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(favicon.New())
	app.Use(recover.New())
	app.Use(logger.New())

	// Static files.
	app.Static("/", "./static/",
		fiber.Static{
			ByteRange: true,
			Browse:    true,
			Index:     "index.html",
		},
	)

	// Middlewares.
	middlewares.FiberMiddleware(app)

	// Routes.
	routes.LoginRoutes(app)
	routes.MessageRoutes(app)
	routes.AdminRoutes(app)
	routes.PublicRoutes(app)
	routes.FileRoutes(app)
	routes.ConnectRoutes(app)
	routes.NotFoundRoute(app)

	log.Fatal(app.Listen(":3000"))
}

func createDefault() {
	basePath, err := utils.MakeBasePath()
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(basePath)
	if err != nil {
		err = os.Mkdir(basePath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	letters := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
	for _, letter := range letters {
		letterPath := filepath.Join(basePath, string(letter))
		_, err = os.Stat(letterPath)
		if err != nil {
			err = os.Mkdir(letterPath, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	db := database.OpenDb()
	err = db.AutoMigrate(
		&models.Group{}, &models.Role{}, &models.User{}, &models.Message{},
		&models.Region{}, &models.Category{}, &models.Status{},
		&models.Person{}, &models.Document{}, &models.Address{}, &models.Workplace{},
		&models.Contact{}, &models.Staff{}, &models.Affilation{}, &models.Relation{},
		&models.Conclusion{}, &models.Check{}, &models.Poligraf{},
		&models.Investigation{}, &models.Inquiry{}, &models.Connection{},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, region := range utils.Regions {
		db.Create(&models.Region{
			NameRegion: region,
		})
	}
	for _, status := range utils.Statuses {
		db.Create(&models.Status{
			NameStatus: status,
		})
	}
	for _, category := range utils.Categories {
		db.Create(&models.Category{
			NameCategory: category,
		})
	}
	for _, group := range utils.Groups {
		db.Create(&models.Group{
			NameGroup: group,
		})
	}
	for _, role := range utils.Roles {
		db.Create(&models.Role{
			NameRole: role,
		})
	}

	user := models.User{
		UserName: "superadmin",
		Password: utils.GeneratePassword("88888888"),
	}
	roles := []models.Role{}
	db.Where("name_role = ?", "admin").Find(&roles)
	user.Roles = roles

	groups := []models.Group{}
	db.Where("name_group = ?", "admins").Find(&groups)
	user.Groups = groups

	db.Create(&user)

	db.Create(&models.Person{
		CategoryID:       models.Category.GetID(models.Category{}, utils.Categories["candidate"]),
		RegionID:         models.Region.GetID(models.Region{}, utils.Regions["MAIN_OFFICE"]),
		FullName:         "Бендер Остап Сулеман",
		PreviousFullName: "Ильф и Петров",
		BirthDate:        time.Now().Format("2006-01-02"),
		BirthPlace:       "г.Нью-Васюки",
		Citizen:          "Россия",
		ExCitizen:        "Турция",
		Snils:            "12345678901",
		Inn:              "123456789012",
		Education:        "Университет Джордано Бруно",
		MaritalStatus:    "женат",
		AdditionalInfo:   "Холодный философ и свободный художник",
		PathToDocs:       basePath,
		StatusID:         models.Status.GetID(models.Status{}, utils.Statuses["new"]),
	})
	log.Println("done")
}
