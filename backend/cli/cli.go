package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"backend/orm"

	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
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

	db := orm.OpenDb()
	err = db.AutoMigrate(
		&orm.Group{}, &orm.Role{}, &orm.User{}, &orm.Message{},
		&orm.Region{}, &orm.Category{}, &orm.Status{},
		&orm.Person{}, &orm.Document{}, &orm.Address{}, &orm.Workplace{},
		&orm.Contact{}, &orm.Staff{}, &orm.Affilation{}, &orm.Relation{},
		&orm.Conclusion{}, &orm.Check{}, &orm.Poligraf{},
		&orm.Investigation{}, &orm.Inquiry{}, &orm.Connection{},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, region := range orm.Regions {
		db.Create(&orm.Region{
			NameRegion: region,
		})
	}
	for _, status := range orm.Statuses {
		db.Create(&orm.Status{
			NameStatus: status,
		})
	}
	for _, category := range orm.Categories {
		db.Create(&orm.Category{
			NameCategory: category,
		})
	}
	for _, group := range orm.Groups {
		db.Create(&orm.Group{
			NameGroup: group,
		})
	}
	for _, role := range orm.Roles {
		db.Create(&orm.Role{
			NameRole: role,
		})
	}

	user := orm.User{
		UserName: "superadmin",
		Password: generateBcryptHash("88888888"),
	}
	roles := []orm.Role{}
	db.Where("name_role = ?", "admin").Find(&roles)
	user.Roles = roles

	groups := []orm.Group{}
	db.Where("name_group = ?", "admins").Find(&groups)
	user.Groups = groups

	db.Create(&user)

	db.Create(&orm.Person{
		CategoryID:       orm.Category.GetID(orm.Category{}, orm.Categories["candidate"]),
		RegionID:         orm.Region.GetID(orm.Region{}, orm.Regions["MAIN_OFFICE"]),
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
		StatusID:         orm.Status.GetID(orm.Status{}, orm.Statuses["new"]),
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
