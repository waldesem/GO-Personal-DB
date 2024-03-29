package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type NameChange struct {
	Reason                   string `json:"reason"`
	FirstNameBeforeChange    string `json:"firstNameBeforeChange"`
	LastNameBeforeChange     string `json:"lastNameBeforeChange"`
	HasNoMidNameBeforeChange bool   `json:"hasNoMidNameBeforeChange"`
	YearOfChange             int    `json:"yearOfChange"`
	NameChangeDocument       string `json:"nameChangeDocument"`
}

type Education struct {
	EducationType   string `json:"educationType"`
	InstitutionName string `json:"institutionName"`
	BeginYear       int    `json:"beginYear"`
	EndYear         int    `json:"endYear"`
	Specialty       string `json:"specialty"`
}

type Experience struct {
	BeginDate                         string `json:"beginDate"`
	EndDate                           string `json:"endDate,omitempty"`
	CurrentJob                        bool   `json:"currentJob,omitempty"`
	Name                              string `json:"name"`
	Address                           string `json:"address"`
	Phone                             string `json:"phone"`
	ActivityType                      string `json:"activityType"`
	Position                          string `json:"position"`
	IsPositionMatchEmploymentContract bool   `json:"isPositionMatchEmploymentContract,omitempty"`
	EmploymentContractPosition        string `json:"employmentContractPosition,omitempty"`
	FireReason                        string `json:"fireReason,omitempty"`
}

type Organization struct {
	View     string `json:"view"`
	Inn      string `json:"inn"`
	OrgType  string `json:"orgType"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

type Person struct {
	PositionName                      string         `json:"positionName"`
	Department                        string         `json:"department"`
	StatusDate                        string         `json:"statusDate"`
	LastName                          string         `json:"lastName"`
	FirstName                         string         `json:"firstName"`
	MidName                           string         `json:"midName"`
	HasNameChanged                    bool           `json:"hasNameChanged"`
	NameWasChanged                    []NameChange   `json:"nameWasChanged"`
	Birthday                          string         `json:"birthday"`
	Birthplace                        string         `json:"birthplace"`
	Citizen                           string         `json:"citizen"`
	HasAdditionalCitizenship          bool           `json:"hasAdditionalCitizenship"`
	AdditionalCitizenship             string         `json:"additionalCitizenship"`
	MaritalStatus                     string         `json:"maritalStatus"`
	RegAddress                        string         `json:"regAddress"`
	ValidAddress                      string         `json:"validAddress"`
	ContactPhone                      string         `json:"contactPhone"`
	HasNoRussianContactPhone          bool           `json:"hasNoRussianContactPhone"`
	Email                             string         `json:"email"`
	HasInn                            bool           `json:"hasInn"`
	Inn                               string         `json:"inn"`
	HasSnils                          bool           `json:"hasSnils"`
	Snils                             string         `json:"snils"`
	PassportSerial                    string         `json:"passportSerial"`
	PassportNumber                    string         `json:"passportNumber"`
	PassportIssueDate                 string         `json:"passportIssueDate"`
	PassportIssuedBy                  string         `json:"passportIssuedBy"`
	Education                         []Education    `json:"education"`
	HasJob                            bool           `json:"hasJob"`
	Experience                        []Experience   `json:"experience"`
	HasPublicOfficeOrganizations      bool           `json:"hasPublicOfficeOrganizations"`
	PublicOfficeOrganizations         []Organization `json:"publicOfficeOrganizations"`
	HasStateOrganizations             bool           `json:"hasStateOrganizations"`
	StateOrganizations                []Organization `json:"stateOrganizations"`
	HasRelatedPersonsOrganizations    bool           `json:"hasRelatedPersonsOrganizations"`
	RelatedPersonsOrganizations       []Organization `json:"relatedPersonsOrganizations"`
	HasMtsRelatedPersonsOrganizations bool           `json:"hasMtsRelatedPersonsOrganizations"`
	MtsRelatedPersonsOrganizations    []Organization `json:"mtsRelatedPersonsOrganizations"`
	HasOrganizations                  bool           `json:"hasOrganizations"`
	Organizations                     []Organization `json:"organizations"`
}

type Anketa struct {
	Resume      map[string]string
	Staff       map[string]string
	Document    map[string]string
	Addresses   []map[string]string
	Workplaces  []map[string]string
	Contacts    []map[string]string
	Affilations []map[string]string
}

func JsonParse(jsonPath string) Anketa {
	f, err := os.Open(jsonPath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	var person Person

	jsonData, err := io.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(jsonData, &person)
	if err != nil {
		log.Println(err)
	}

	resume := map[string]string{
		"fullname":   person.parseFullname(),
		"previous":   person.parsePrevious(),
		"birthday":   person.Birthday,
		"birthplace": person.Birthplace,
		"citizen":    person.Citizen,
		"exCitizen":  person.AdditionalCitizenship,
		"marital":    person.MaritalStatus,
		"education":  person.parseEducation(),
		"inn":        person.Inn,
		"snils":      person.Snils,
	}

	staff := map[string]string{
		"position":   person.PositionName,
		"department": person.Department,
	}

	document := map[string]string{
		"view":   "Паспорт РФ",
		"series": person.PassportSerial,
		"number": person.PassportNumber,
		"issue":  person.PassportIssueDate,
		"agency": person.PassportIssuedBy,
	}

	addresses := []map[string]string{
		{
			"view":         "Адрес проживания",
			"live_address": person.RegAddress,
		},
		{
			"view":        "Адрес регистрации",
			"reg_address": person.ValidAddress,
		},
	}

	contacts := []map[string]string{
		{
			"view":  "Телефон",
			"phone": person.ContactPhone,
		},
		{
			"view":  "Электронная почта",
			"email": person.Email,
		},
	}

	works := person.parseWorkplace()
	workplaces := []map[string]string{}
	for _, item := range works {
		workplaces = append(workplaces, map[string]string{
			"beginDate": item.BeginDate,
			"endDate":   item.EndDate,
			"name":      item.Name,
			"address":   item.Address,
			"position":  item.Position,
			"fire":      item.FireReason,
		})
	}

	affs := person.parseAffilation()
	affilations := []map[string]string{}
	for _, item := range affs {
		affilations = append(affilations, map[string]string{
			"view":     item.View,
			"name":     item.Name,
			"position": item.Position,
			"inn":      item.Inn,
		})
	}

	return Anketa{
		Resume:      resume,
		Staff:       staff,
		Document:    document,
		Addresses:   addresses,
		Workplaces:  workplaces,
		Contacts:    contacts,
		Affilations: affilations,
	}
}

func (person Person) parseFullname() string {
	name := fmt.Sprintf(
		"%s %s %s",
		strings.TrimSpace(person.LastName),
		strings.TrimSpace(person.FirstName),
		strings.TrimSpace(person.MidName),
	)
	return strings.ToTitle(strings.TrimSpace(name))
}

func (person Person) parsePrevious() string {
	var previous []string
	if person.HasNameChanged {
		for _, item := range person.NameWasChanged {
			previous = append(previous, fmt.Sprintf("%s - %s %d %s, %s",
				item.FirstNameBeforeChange, item.LastNameBeforeChange,
				item.YearOfChange, item.NameChangeDocument, item.Reason,
			))
		}
	}
	return strings.Join(previous, "")
}

func (person Person) parseEducation() string {
	var education []string
	if len(person.Education) > 0 {
		for _, item := range person.Education {
			education = append(education, fmt.Sprintf("%s, %s, %d, %d",
				item.EducationType, item.InstitutionName, item.BeginYear, item.EndYear))
		}
	}
	return strings.Join(education, "")
}

func (person Person) parseWorkplace() []Experience {
	var expirience []Experience
	if len(person.Experience) > 0 {
		for _, item := range person.Experience {
			expirience = append(expirience, Experience{
				BeginDate:  item.BeginDate,
				EndDate:    item.EndDate,
				Name:       item.Name,
				Address:    item.Address,
				Position:   item.Position,
				FireReason: item.FireReason,
			})
		}
	}
	return expirience
}

func (person Person) parseAffilation() []Organization {
	var affilation []Organization
	if person.HasPublicOfficeOrganizations {
		for _, item := range person.PublicOfficeOrganizations {
			affilation = append(affilation, Organization{
				View:     "Являлся государственным или муниципальным служащим",
				Name:     item.Name,
				Position: item.Position,
			})
		}
	}
	if person.HasStateOrganizations {
		for _, item := range person.StateOrganizations {
			affilation = append(affilation, Organization{
				View:     "Являлся государственным должностным лицом",
				Name:     item.Name,
				Position: item.Position,
			})
		}
	}
	if person.HasRelatedPersonsOrganizations {
		for _, item := range person.RelatedPersonsOrganizations {
			affilation = append(affilation, Organization{
				View:     "Связанные лица работают в госудраственных организациях",
				Name:     item.Name,
				Position: item.Position,
				Inn:      item.Inn,
			})
		}
	}
	if person.HasOrganizations {
		for _, item := range person.Organizations {
			affilation = append(affilation, Organization{
				View:     "Участвует в деятельности коммерческих организаций",
				Name:     item.Name,
				Position: item.Position,
				Inn:      item.Inn,
			})
		}
	}
	return affilation
}
