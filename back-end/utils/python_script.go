package utils

// 以下转换函数需要根据你的具体需求来实现
func GenderToString(gender bool) string {
	if gender {
		return "Male"
	}
	return "Female"
}

func MarriedToString(married bool) string {
	if married {
		return "Yes"
	}
	return "No"
}

func EducationToString(education bool) string {
	if education {
		return "Graduate"
	}
	return "Not Graduate"
}

func SelfEmployedToString(selfEmployed bool) string {
	if selfEmployed {
		return "Yes"
	}
	return "No"
}

func CityToString(city bool) string {
	if city {
		return "Urban"
	}
	return "Rural"
}
