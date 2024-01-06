package utils

var Regions map[string]string = map[string]string{
	"MAIN_OFFICE":  "Главный офис",
	"SOUTH_REGION": "Региональный центр Юг",
	"WEST_REGION":  "Региональный центр Запад",
	"URAL_REGION":  "Региональный центр Урал",
	"EAST_REGION":  "Региональный центр Восток",
}

var Groups map[string]string = map[string]string{
	"admins":   "admins",
	"staffsec": "staffsec",
	"api":      "api",
}

var Roles map[string]string = map[string]string{
	"admin": "admin",
	"user":  "user",
	"api":   "api",
}

var Categories map[string]string = map[string]string{
	"candidate": "Кандидат",
	"staff":     "Сотрудник",
	"vip":       "ВИП",
}

var Statuses map[string]string = map[string]string{
	"new":      "Новый",
	"repeat":   "Повторный",
	"update":   "Обновлен",
	"manual":   "Проверка",
	"save":     "Сохранен",
	"robot":    "Робот",
	"reply":    "Обработан",
	"poligraf": "ПФО",
	"finish":   "Окончено",
	"cancel":   "Отменено",
	"error":    "Ошибка",
}

var Conclusions map[string]string = map[string]string{
	"agreed":       "Согласовано",
	"with_comment": "Согласовано с комментарием",
	"denied":       "Отказано в согласовании",
	"saved":        "Сохранен",
	"canceled":     "Отменено",
}
