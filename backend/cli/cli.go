package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"

	"backend/app/models"
	"backend/pkg/utils"
	"backend/platform/database"
)

func main() {
	app := &cli.App{
		Name:  "staffsec",
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
				Name:  "test",
				Usage: "Test",
				Action: func(c *cli.Context) error {
					log.Println("Test")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createDefault() {
	basePath := MakeBasePath()
	_, err := os.Stat(basePath)
	if err != nil {
		os.Mkdir(basePath, 0755)
	}
	letters := "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

	for _, letter := range letters {
		letterPath := filepath.Join(basePath, string(letter))
		_, err = os.Stat(letterPath)
		if err != nil {
			os.Mkdir(letterPath, 0755)
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
		Password: generateBcryptHash("88888888"),
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

func generateBcryptHash(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return hashedPassword
}

func MakeBasePath() string {
	cur, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(cur, "..", "..", "..", "persons")
}
