package main

func reformatDate(date string) string {
	result := ""

	monthMap := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	if len(date) == 12 {
		date = "0" + date
	}

	year := date[8:]
	month := monthMap[date[5:8]]
	day := date[0:2]
	result = year + "-" + month + "-" + day

	return result
}

//func main() {
//	result := reformatDate("20th Oct 2052")
//	print(result)
//}
