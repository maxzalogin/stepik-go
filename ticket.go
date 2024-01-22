/*
Условные конструкции

Определите является ли билет счастливым.  Счастливым считается билет, в шестизначном номере
которого сумма первых трёх цифр совпадает с суммой трёх последних.

Формат входных данных
На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".
*/
package main

import "fmt"

func main() {

	const (
		Ten			   int = 10
		OneHundred     int = 100
		Thousand       int = 1000
		TenThousand	   int = 10000
	)

	var input int
	fmt.Scan(&input)

	right := ((input % Thousand) / OneHundred) + ((input % Thousand) % OneHundred / Ten) + ((input % Thousand % OneHundred) % Ten)
	left := ((input / TenThousand	) / Ten) + ((input / TenThousand	) % Ten) + ((input / Thousand) % Ten) 

	if left == right{
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}