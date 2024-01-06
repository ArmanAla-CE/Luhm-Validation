package LuhmValidator

func IsValid(userCardNumber string) bool {
	var oddCardNumberLenght bool
	userCardNumberInInteger := make([]int, len(userCardNumber))
	for index, character := range userCardNumber {
		userCardNumberInInteger[index] = int(character) - 48
	}
	if len(userCardNumberInInteger)%2 == 0 {
		oddCardNumberLenght = false
	} else {
		oddCardNumberLenght = true
	}

	sumOfAllLuhnNumbers := 0
	for index, value := range userCardNumberInInteger {
		tempValue := value
		if !oddCardNumberLenght && index%2 == 0 {
			tempValue *= 2
			if tempValue >= 10 {
				reaminder := tempValue % 10
				tempValue = reaminder + 1
			}
		} else if oddCardNumberLenght && index%2 != 0 {
			tempValue *= 2
			if tempValue >= 10 {
				reaminder := tempValue % 10
				tempValue = reaminder + 1
			}
		}
		sumOfAllLuhnNumbers += tempValue
	}

	return sumOfAllLuhnNumbers%10 == 0
}
