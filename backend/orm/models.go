package orm

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDb() *gorm.DB {
	var dsn = "host=localhost user=flask password=flask dbname=personal port=5432 sslmode=disable"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type Group struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	NameGroup string `gorm:"size(256)" json:"group" serialize:"json"`
	Users     []User `gorm:"many2many:user_groups;"`
}

func (group Group) GetID(name string) uint {
	groupId := uint(0)
	if group.NameGroup == name {
		groupId = group.ID
	}
	return groupId
}

type Role struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	NameRole string `gorm:"size(256)" json:"role" serialize:"json"`
	Users    []User `gorm:"many2many:user_roles;"`
}

func (role Role) GetID(name string) uint {
	roleId := uint(0)
	if role.NameRole == name {
		roleId = role.ID
	}
	return roleId
}

type User struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	FullName  string    `gorm:"size(256)" json:"fullname" serialize:"json"`
	UserName  string    `gorm:"size(256)" json:"username" serialize:"json"`
	Password  []byte    `json:"password" serialize:"json"`
	Email     string    `gorm:"size(256)" json:"email" serialize:"json"`
	CreatedAt time.Time `json:"created" serialize:"json"`
	UpdatedAt time.Time `json:"updated" serialize:"json"`
	LastLogin time.Time `json:"last_login" serialize:"json"`
	Blocked   bool      `gorm:"default:false" json:"blocked" serialize:"json"`
	Attempt   int       `gorm:"default:0" json:"attempt" serialize:"json"`
	Groups    []Group   `gorm:"many2many:user_groups" json:"groups" serialize:"json"`
	Roles     []Role    `gorm:"many2many:user_roles" json:"roles" serialize:"json"`
	Messages  []Message
}

func (user User) HasGroup(groups []string) bool {
	for _, g := range user.Groups {
		for _, group := range groups {
			if g.NameGroup == group {
				return true
			}
		}
	}
	return false
}

func (user User) HasRole(roles []string) bool {
	for _, r := range user.Roles {
		for _, role := range roles {
			if r.NameRole == role {
				return true
			}
		}
	}
	return false
}

type Message struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Title          string    `gorm:"size(256)" json:"title" serialize:"json"`
	MessageContent string    `gorm:"size(256)" json:"message" serialize:"json"`
	StatusRead     string    `gorm:"size(256)" json:"status" serialize:"json"`
	CreatedAt      time.Time `json:"created" serialize:"json"`
	UserID         uint
}

type Category struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	NameCategory string `gorm:"size(256)" json:"category" serialize:"json"`
	Persons      []Person
}

func (category Category) GetID(name string) uint {
	categoryId := uint(0)
	if category.NameCategory == name {
		categoryId = category.ID
	}
	return categoryId
}

type Status struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	NameStatus string `gorm:"size(256)" json:"status" serialize:"json"`
	Persons    []Person
}

func (status Status) GetID(name string) uint {
	statusId := uint(0)
	if status.NameStatus == name {
		statusId = status.ID
	}
	return statusId
}

type Region struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	NameRegion string `gorm:"size(256)" json:"region" serialize:"json"`
	Persons    []Person
}

func (region Region) GetID(name string) uint {
	regionId := uint(0)
	if region.NameRegion == name {
		regionId = region.ID
	}
	return regionId
}

type Person struct {
	gorm.Model
	ID               uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	CategoryID       uint      `json:"category_id" serialize:"json"`
	RegionID         uint      `json:"region_id" serialize:"json"`
	FullName         string    `gorm:"not null; index" json:"fullname" serialize:"json"`
	PreviousFullName string    `json:"previous" serialize:"json"`
	BirthDate        time.Time `gorm:"not null" json:"birthday" serialize:"json"`
	BirthPlace       string    `json:"birth_place" serialize:"json"`
	Citizen          string    `gorm:"size(256)" json:"country" serialize:"json"`
	ExCitizen        string    `gorm:"size(256)" json:"ex_citizen" serialize:"json"`
	Snils            string    `gorm:"size(11)" json:"snils" serialize:"json"`
	Inn              string    `gorm:"size(12)" json:"inn" serialize:"json"`
	Education        string    `json:"education" serialize:"json"`
	MaritalStatus    string    `gorm:"son:marital" serialize:"json"`
	AdditionalInfo   string    `json:"addition" serialize:"json"`
	PathToDocs       string    `json:"path" serialize:"json"`
	StatusID         uint      `json:"status_id" serialize:"json"`
	CreatedAt        time.Time `json:"created" serialize:"json"`
	UpdatedAt        time.Time `json:"updated" serialize:"json"`
	Documents        []Document
	Addresses        []Address
	Workplaces       []Workplace
	Contacts         []Contact
	Staffs           []Staff
	Affiliations     []Affilation
	Relations        []Relation
	Checks           []Check
	Inquiries        []Inquiry
	Investigations   []Investigation
	Robots           []Robot
	Poligrafs        []Poligraf
}

type Staff struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Position   string `json:"position" serialize:"json"`
	Department string `json:"department" serialize:"json"`
	PersonID   uint
}

type Document struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	View     string    `gorm:"size(256)" json:"view" serialize:"json"`
	Series   string    `gorm:"size(256)" json:"series" serialize:"json"`
	Number   string    `gorm:"size(256)" json:"number" serialize:"json"`
	Agency   string    `json:"agency" serialize:"json"`
	Issue    time.Time `json:"issue" serialize:"json"`
	PersonID uint
}

type Address struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	View     string `gorm:"size(256)" json:"view" serialize:"json"`
	Region   string `gorm:"size(256)" json:"region" serialize:"json"`
	Address  string `json:"address" serialize:"json"`
	PersonID uint
}

type Contact struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	View     string `gorm:"size(256)" json:"view" serialize:"json"`
	Contact  string `gorm:"size(256)" json:"contact" serialize:"json"`
	PersonID uint
}

type Workplace struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	View      string `json:"view" serialize:"json"`
	Workplace string `json:"workplace" serialize:"json"`
	Address   string `json:"address" serialize:"json"`
	Position  string `json:"position" serialize:"json"`
	Reason    string `json:"reason" serialize:"json"`
	PersonID  uint
}

type Affilation struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	View     string    `gorm:"size(256)" json:"view" serialize:"json"`
	Name     string    `gorm:"size(256)" json:"name" serialize:"json"`
	Inn      string    `gorm:"size(12)" json:"inn" serialize:"json"`
	Position string    `json:"position" serialize:"json"`
	Deadline time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"deadline" serialize:"json"`
	PersonID uint
}

type Relation struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	View     string `gorm:"size(256)" json:"relation" serialize:"json"`
	Relation uint   `gorm:"foreignKey:PersonID" json:"relation_id" serialize:"json"`
	PersonID uint
}

type Conclusion struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Conclusion string `gorm:"size(256)" json:"conclusion" serialize:"json"`
	Checks     []Check
}

type Check struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	CheckWorkplace string    `json:"workplace" serialize:"json"`
	CheckEmployee  string    `json:"employee" serialize:"json"`
	CheckDocument  string    `json:"document" serialize:"json"`
	CheckInn       string    `json:"inn" serialize:"json"`
	Debt           string    `json:"debt" serialize:"json"`
	Bankruptcy     string    `json:"bankruptcy" serialize:"json"`
	BKI            string    `json:"bki" serialize:"json"`
	Courts         string    `json:"courts" serialize:"json"`
	Affiliation    string    `json:"affiliation" serialize:"json"`
	Terrorist      string    `json:"terrorist" serialize:"json"`
	MVD            string    `json:"mvd" serialize:"json"`
	Internet       string    `json:"internet" serialize:"json"`
	Cronos         string    `json:"cronos" serialize:"json"`
	CROS           string    `json:"cros" serialize:"json"`
	Comments       string    `json:"comments" serialize:"json"`
	Addition       string    `json:"addition" serialize:"json"`
	ConclusionID   uint      `gorm:"foreignKey:ConclusionID" serialize:"json" json:"conclusion_id"`
	Officer        string    `json:"officer" serialize:"json"`
	Deadline       time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"deadline" serialize:"json"`
	PersonID       uint
}

type Robot struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	Employee   string    `json:"employee" serialize:"json"`
	Inn        string    `json:"inn" serialize:"json"`
	Debt       string    `json:"debt" serialize:"json"`
	Bankruptcy string    `json:"bankruptcy" serialize:"json"`
	BKI        string    `json:"bki" serialize:"json"`
	Courts     string    `json:"courts" serialize:"json"`
	Terrorist  string    `json:"terrorist" serialize:"json"`
	MVD        string    `json:"mvd" serialize:"json"`
	Deadline   time.Time `gorm:"autoCreateTime; autoUpdateTime"`
	PersonID   uint
}

type Poligraf struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Theme    string    `gorm:"size(256)" json:"theme" serialize:"json"`
	Results  string    `json:"results" serialize:"json"`
	Officer  string    `gorm:"size(256)" json:"officer" serialize:"json"`
	Deadline time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"deadline" serialize:"json"`
	PersonID uint
}

type Investigation struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Theme    string    `gorm:"size(256)" json:"theme" serialize:"json"`
	Info     string    `json:"info" serialize:"json"`
	Officer  string    `gorm:"size(256)" json:"officer" serialize:"json"`
	Deadline time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"deadline" serialize:"json"`
	PersonID uint
}

type Inquiry struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Info      string    `json:"info" serialize:"json"`
	Initiator string    `gorm:"size(256)" json:"initiator" serialize:"json"`
	Source    string    `gorm:"size(256)" json:"source" serialize:"json"`
	Officer   string    `gorm:"size(256)" json:"officer" serialize:"json"`
	Deadline  time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"deadline" serialize:"json"`
	PersonID  uint
}

type Connection struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Company  string    `gorm:"size(256)" json:"company" serialize:"json"`
	City     string    `gorm:"size(256)" json:"city" serialize:"json"`
	Fullname string    `gorm:"size(256)" json:"fullname" serialize:"json"`
	Phone    string    `gorm:"size(256)" json:"phone" serialize:"json"`
	Adding   string    `gorm:"size(256)" json:"adding" serialize:"json"`
	Mobile   string    `gorm:"size(256)" json:"mobile" serialize:"json"`
	Mail     string    `gorm:"size(256)" json:"mail" serialize:"json"`
	Comment  string    `json:"comment" serialize:"json"`
	Data     time.Time `gorm:"autoCreateTime; autoUpdateTime" json:"data" serialize:"json"`
}

type Anketa struct {
	ID         uint   `gorm:"primaryKey; autoIncrement; not null; unique" json:"id" serialize:"json"`
	Fullname   string `gorm:"size(256)" json:"fullname" serialize:"json"`
	Birthday   string `gorm:"size(256)" json:"birthday" serialize:"json"`
	Birthplace string `gorm:"size(256)" json:"birthplace" serialize:"json"`
	Snils      string `gorm:"size(11)" json:"snils" serialize:"json"`
	Inn        string `gorm:"size(12)" json:"inn" serialize:"json"`
	Series     string `gorm:"size(256)" json:"series" serialize:"json"`
	Number     string `gorm:"size(256)" json:"number" serialize:"json"`
	Agency     string `gorm:"size(256)" json:"agency" serialize:"json"`
	Issue      string `gorm:"size(256)" json:"issue" serialize:"json"`
	Address    string `gorm:"size(256)" json:"address" serialize:"json"`
}
