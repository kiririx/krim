package mapping

func SexGet(sexCode uint) string {
	switch sexCode {
	case 0:
		return "男"
	case 1:
		return "女"
	}
	return "未知"
}

func SexTo(sex string) uint {
	switch sex {
	case "男":
		return 0
	case "女":
		return 1
	}
	return 0
}
