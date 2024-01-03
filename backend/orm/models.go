package orm

import (
	"time"

	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

type Group struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey json:id serialize:json"`
	NameGroup string `gorm:"size(256) json:group serialize:json"`
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
	ID       uint   `gorm:"primaryKey json:id serialize:json"`
	NameRole string `gorm:"size(256) json:role serialize:json"`
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
	ID        uint      `gorm:"primaryKey; autoIncrement; not null; unique; json:id serialize:json"`
	FullName  string    `gorm:"size(256) json:fullname serialize:json"`
	UserName  string    `gorm:"size(256) json:username serialize:json"`
	Password  byte      `gorm:"json:password serialize:json"`
	Email     string    `gorm:"size(256) json:email serialize:json"`
	CreatedAt time.Time `json:"created serialize:json"`
	UpdatedAt time.Time `json:"updated serialize:json"`
	LastLogin time.Time `json:"last_login serialize:json"`
	Blocked   bool      `gorm:"default:false json:blocked serialize:json"`
	Attempt   int       `gorm:"default:0 json:attempt serialize:json"`
	Groups    []Group   `gorm:"many2many:user_groups; json:groups serialize:json"`
	Roles     []Role    `gorm:"many2many:user_roles; json:roles serialize:json"`
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
	ID             uint      `gorm:"primaryKey json:id serialize:json"`
	Title          string    `gorm:"size(256) json:title serialize:json"`
	MessageContent string    `gorm:"size(256) json:message serialize:json"`
	StatusRead     string    `gorm:"size(256) json:status serialize:json"`
	CreatedAt      time.Time `json:"created serialize:json"`
	UserID         uint
}

type Category struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey json:id serialize:json"`
	NameCategory string `gorm:"size(256) json:category serialize:json"`
	Persons      []Person
}

func (category Category) GetID() uint {
	return category.ID
}

type Status struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey json:id serialize:json"`
	NameStatus string `gorm:"size(256) json:status serialize:json"`
	Persons    []Person
}

func (status Status) GetID() uint {
	return status.ID
}

type Region struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey json:id serialize:json"`
	NameRegion string `gorm:"size(256) json:region serialize:json"`
	Persons    []Person
}

func (region Region) GetID() uint {
	return region.ID
}

type Person struct {
	gorm.Model
	ID               uint       `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	CategoryID       uint       `gorm:"json:category_id serialize:json"`
	RegionID         uint       `gorm:"json:region_id serialize:json"`
	FullName         string     `gorm:"not null; index json:fullname serialize:json"`
	PreviousFullName string     `gorm:"json:previous serialize:json"`
	BirthDate        *time.Time `gorm:"not null json:birthday serialize:json"`
	BirthPlace       string     `gorm:"json:birth_place serialize:json"`
	Citizen          string     `gorm:"size(256) json:country serialize:json"`
	ExCitizen        string     `gorm:"size(256) json:ex_citizen serialize:json"`
	Snils            string     `gorm:"size(11) json:snils serialize:json"`
	Inn              string     `gorm:"size(12) json:inn serialize:json"`
	Education        string     `gorm:"json:education serialize:json"`
	MaritalStatus    string     `gorm:"son:marital serialize:json"`
	AdditionalInfo   string     `gorm:"json:addition serialize:json"`
	PathToDocs       string     `gorm:"json:path serialize:json"`
	StatusID         uint       `gorm:"json:status_id serialize:json"`
	CreatedAt        time.Time  `json:"created serialize:json"`
	UpdatedAt        time.Time  `json:"updated serialize:json"`
	Documents        []Document
	Addresses        []Address
	Workplaces       []Workplace
	Contacts         []Contact
	Staffs           []Staff
	Affiliations     []Affilation
	Relations        []Relations
	Checks           []Check
	Inquiries        []Inquiry
	Investigations   []Investigation
	Robots           []Robot
	Poligrafs        []Poligraf
}

type Staff struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Position   string `gorm:"json:position serialize:json"`
	Department string `gorm:"json:department serialize:json"`
	PersonID   uint
}

type Document struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	View     string    `gorm:"size(256) json:view serialize:json"`
	Series   string    `gorm:"size(256) json:series serialize:json"`
	Number   string    `gorm:"size(256) json:number serialize:json"`
	Agency   string    `gorm:"json:agency serialize:json"`
	Issue    time.Time `gorm:"json:issue serialize:json"`
	PersonID uint
}

type Address struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	View     string `gorm:"size(256) json:view serialize:json"`
	Region   string `gorm:"size(256) json:region serialize:json"`
	Address  string `gorm:"json:address serialize:json"`
	PersonID uint
}

type Contact struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	View     string `gorm:"size(256) json:view serialize:json"`
	Contact  string `gorm:"size(256) json:contact serialize:json"`
	PersonID uint
}

type Workplace struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	View      string `gorm:"json:view serialize:json"`
	Workplace string `gorm:"json:workplace serialize:json"`
	Address   string `gorm:"json:address serialize:json"`
	Position  string `gorm:"json:position serialize:json"`
	Reason    string `gorm:"json:reason serialize:json"`
	PersonID  uint
}

type Affilation struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	View     string    `gorm:"size(256) json:view serialize:json"`
	Name     string    `gorm:"size(256) json:name serialize:json"`
	Inn      string    `gorm:"size(12) json:inn serialize:json"`
	Position string    `gorm:"json:position serialize:json"`
	Deadline time.Time `gorm:"autoCreateTime; autoUpdateTime json:deadline serialize:json"`
	PersonID uint
}

type Relations struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	View     string `gorm:"size(256) json:relation serialize:json"`
	Relation uint   `gorm:"foreignKey:PersonID json:relation_id serialize:json"`
	PersonID uint
}

type Check struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	CheckWorkplace string    `gorm:"json:workplace serialize:json"`
	CheckEmployee  string    `gorm:"json:employee serialize:json"`
	CheckDocument  string    `gorm:"json:document serialize:json"`
	CheckInn       string    `gorm:"json:inn serialize:json"`
	Debt           string    `gorm:"json:debt serialize:json"`
	Bankruptcy     string    `gorm:"json:bankruptcy serialize:json"`
	BKI            string    `gorm:"json:bki serialize:json"`
	Courts         string    `gorm:"json:courts serialize:json"`
	Affiliation    string    `gorm:"json:affiliation serialize:json"`
	Terrorist      string    `gorm:"json:terrorist serialize:json"`
	MVD            string    `gorm:"json:mvd serialize:json"`
	Internet       string    `gorm:"json:internet serialize:json"`
	Cronos         string    `gorm:"json:cronos serialize:json"`
	CROS           string    `gorm:"json:cros serialize:json"`
	Comments       string    `gorm:"json:comments serialize:json"`
	Addition       string    `gorm:"json:addition serialize:json"`
	ConclusionID   uint      `gorm:"foreignKey:ConclusionID serialize:json json:conclusion_id"`
	Officer        string    `gorm:"json:officer serialize:json"`
	Deadline       time.Time `gorm:"autoCreateTime; autoUpdateTime json:deadline serialize:json"`
	PersonID       uint
}

type Robot struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey"`
	Employee   string    `gorm:"json:employee serialize:json"`
	Inn        string    `gorm:"json:inn serialize:json"`
	Debt       string    `gorm:"json:debt serialize:json"`
	Bankruptcy string    `gorm:"json:bankruptcy serialize:json"`
	BKI        string    `gorm:"json:bki serialize:json"`
	Courts     string    `gorm:"json:courts serialize:json"`
	Terrorist  string    `gorm:"json:terrorist serialize:json"`
	MVD        string    `gorm:"json:mvd serialize:json"`
	Deadline   time.Time `gorm:"autoCreateTime; autoUpdateTime"`
	PersonID   uint
}

type Conclusion struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Conclusion string `gorm:"size(256) json:conclusion serialize:json"`
	Checks     []Check
}

type Poligraf struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Theme    string    `gorm:"size(256) json:theme serialize:json"`
	Results  string    `gorm:"json:results serialize:json"`
	Officer  string    `gorm:"size(256) json:officer serialize:json"`
	Deadline time.Time `gorm:"autoCreateTime; autoUpdateTime json:deadline serialize:json"`
	PersonID uint
}

type Investigation struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Theme    string    `gorm:"size(256) json:theme serialize:json"`
	Info     string    `gorm:"json:info serialize:json"`
	Officer  string    `gorm:"size(256) json:officer serialize:json"`
	Deadline time.Time `gorm:"autoCreateTime; autoUpdateTime json:deadline serialize:json"`
	PersonID uint
}

type Inquiry struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Info      string    `gorm:"json:info serialize:json"`
	Initiator string    `gorm:"size(256) json:initiator serialize:json"`
	Source    string    `gorm:"size(256) json:source serialize:json"`
	Officer   string    `gorm:"size(256) json:officer serialize:json"`
	Deadline  time.Time `gorm:"autoCreateTime; autoUpdateTime json:deadline serialize:json"`
	PersonID  uint
}

type Connection struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Company  string    `gorm:"size(256) json:company serialize:json"`
	City     string    `gorm:"size(256) json:city serialize:json"`
	Fullname string    `gorm:"size(256) json:fullname serialize:json"`
	Phone    string    `gorm:"size(256) json:phone serialize:json"`
	Adding   string    `gorm:"size(256) json:adding serialize:json"`
	Mobile   string    `gorm:"size(256) json:mobile serialize:json"`
	Mail     string    `gorm:"size(256) json:mail serialize:json"`
	Comment  string    `gorm:"json:comment serialize:json"`
	Data     time.Time `gorm:"autoCreateTime; autoUpdateTime json:data serialize:json"`
}

type Anketa struct {
	ID         uint   `gorm:"primaryKey autoIncrement unique not null json:id serialize:json"`
	Fullname   string `gorm:"size(256) json:fullname serialize:json"`
	Birthday   string `gorm:"size(256) json:birthday serialize:json"`
	Birthplace string `gorm:"size(256) json:birthplace serialize:json"`
	Snils      string `gorm:"size(11) json:snils serialize:json"`
	Inn        string `gorm:"size(12) json:inn serialize:json"`
	Series     string `gorm:"size(256) json:series serialize:json"`
	Number     string `gorm:"size(256) json:number serialize:json"`
	Agency     string `gorm:"size(256) json:agency serialize:json"`
	Issue      string `gorm:"size(256) json:issue serialize:json"`
	Address    string `gorm:"size(256) json:address serialize:json"`
}
