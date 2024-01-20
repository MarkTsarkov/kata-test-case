package main

import "fmt"
import "strings"
import "bufio"




func romanToInt(s string) int {
    compares := map[string]int{
        "I" : 1,
        "V" : 5,
        "X" : 10,
        "L" : 50,
        "C" : 100
    }

    integer := 0

    for i := 0; i < len(s); i++ {
        current := string(s[i])
        next := ""

        if i < (len(s)-1) {
            next = string(s[i+1])
        }

        if compares[current] < compares[next] {
            integer = integer + compares[next] - compares[current]
            i++
        } else {
            integer += compares[current]
        }
    }
    return integer
}


func main() {
    //считать операнды
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')

	}
}



// считать строку через буфио
// разделить строку на операнды и оператор посредством пробела "операнд1" "оператор" "операнд2"
// DONE если операнды римские - перевести в арабский инт
// если арабские операнды больше 10 - выдать панику
// если операнды римские и результат <=0 - выдать панику
// если один операнд римский, а другой арабский - выдать панику
// 
// распознать оператор
// выполнить расчеты
// если операнды заданы римскимии числами - отдать резульат римским числом
// отдать результат в консоль

