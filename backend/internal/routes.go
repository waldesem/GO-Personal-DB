package internal

import (
	orm "backend/orm"
	"encoding/json"
	"strconv"
)

func HandleIndex(item string, page string, resp string) ([]byte, bool, bool) {
	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 1
	}

	db := orm.OpenDb()
	var persons []orm.Person
	var pagination = 16
	var hasPrev, hasNext bool

	switch item {
	case "new":
		statusNew := orm.Status{}.GetID(orm.Statuses["new"])
		statusUpd := orm.Status{}.GetID(orm.Statuses["update"])
		statusRep := orm.Status{}.GetID(orm.Statuses["repeat"])
		db.
			Where("status_id = ? OR status_id = ? OR status_id = ?", statusNew, statusUpd, statusRep).
			Where("full_name LIKE ?", "%"+resp+"%").
			Find(&persons).
			Limit(pagination).
			Offset(pagination * (intPage - 1))
	case "officer":
		var checks []orm.Check
		statusFin := orm.Status{}.GetID(orm.Statuses["finish"])
		statusCan := orm.Status{}.GetID(orm.Statuses["cancel"])
		db.
			Find(&persons).
			Where("status_id != ? OR status_id != ?", statusFin, statusCan).
			Limit(pagination).
			Offset(pagination * (intPage - 1)).
			Joins("JOIN checks ON checks.person_id = people.id").
			Where(orm.Check{Officer: "current"}).
			Find(&checks)
	case "search":
		db.Where("fullname LIKE ?", "%"+resp+"%").Find(&persons).Limit(10).Offset(10 * (intPage - 1))
	default:
		return nil, false, false
	}

	jsonData, err := json.Marshal(persons)
	if err != nil {
		return nil, false, false
	}

	if intPage > 1 {
		hasPrev = true
	}
	if len(persons) == pagination {
		hasNext = true
	}

	return jsonData, hasPrev, hasNext
}
