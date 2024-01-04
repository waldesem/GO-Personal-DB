package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"backend/orm"

	"github.com/urfave/cli/v3"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cmd := &cli.Command{
		Name:  "create",
		Usage: "create new project",
		Flags: []cli.Flag{},
		Action: func(_ context.Context, cmd *cli.Command) error {
			createDefault()
			return nil
		},
	}

	cmd.Run(context.Background(), os.Args)
}

func createDefault() {
	basePath := GetBasePath()
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
	fmt.Println(orm.Category.GetID(orm.Category{}, orm.Categories["candidate"]))

	pswd, err := generateBcryptHash("88888888")
	if err != nil {
		log.Fatal(err)
	}

	user := orm.User{
		UserName: "superadmin",
		Password: pswd,
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
		BirthDate:        time.Now(),
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

func generateBcryptHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func GetBasePath() string {
	cur, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(cur, "..", "..", "..", "persons")
}
