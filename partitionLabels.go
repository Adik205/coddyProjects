package coddy

import "strings"

// Функция, которая делит строку на подстроки так, чтобы каждая буква встречалась только в одной части
func PartitionLabels(s string) []int {
	var result []int
	lastIndexes := make(map[byte]int)

	// Для каждой английской буквы находим его последнее вхождение в строке s
	for i := 'a'; i <= 'z'; i++ {
		lastIndexes[byte(i)] = strings.LastIndexByte(s, s[i])
	}

	for i := 0; i < len(s); i++ {
		lastIndex := lastIndexes[s[i]] // Находим последнее вхождение буквы с индексом i

		if i == lastIndex { // Если последнее вхождение всех букв до lastIndex не превышает lastIndex
			result = append(result, lastIndex+1-i) // Добавляем длину подстроки, которая получилось в итоге
		}
		i = lastIndex + 1 // Меняем значение i на индекс начала следующей подчасти
	}
	return result
}

/*
"abacbd"

i = 4
lastIndex = 4


*/
