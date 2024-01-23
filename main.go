package main

import (
	"bufio"
	"fmt"
	"os"
	_ "reflect"
	"strconv"
	"strings"
)

func solve(a, b, op string, x, y *int, roman *bool) (result int) {
    var aRoman, bRoman bool

    //проверили на системы
    firstOperand, err := strconv.Atoi(a)
    if err != nil {
        aRoman = true
        firstOperand = 1
    }
    secondOperand, err := strconv.Atoi(b)
    if err != nil {
        bRoman = true
        secondOperand = 1
    }
    if aRoman != bRoman {
        panic("Числам из разных систем тут не рады")
    }
    if firstOperand > 10 || secondOperand > 10 || firstOperand < 1 || secondOperand < 1 {
        panic("Арабские числа больше 10 и меньше 1 вызывают у меня дискомфорт")
    }

    if aRoman == true && bRoman == true {
        if romanToInt(a) > 10 || romanToInt(a) < 1 || romanToInt(b) > 10 || romanToInt(b) < 1 {
            panic("Римские числа больше 10 и меньше 1 меня пугают")
        }
        *x = romanToInt(a) 
        *y = romanToInt(b)
        *roman = true
    } else {
        *x = firstOperand

        *y = secondOperand

        *roman = false
    }

    switch {
    case op == "+":
        result = *x + *y
    case op == "-":
        result = *x - *y
    case op == "*":
        result = *x * *y
    case op == "/":
        result = *x / *y
    }


    return result
}

func decompose(operation string, operators [4]string, a, b, op *string){
    for _, operator := range operators {
        before, after, found := strings.Cut(operation, operator)
        if found == false {
            continue
        }
        
        *a = strings.TrimSpace(strings.ToUpper(before))


        *b = strings.TrimSpace(strings.ToUpper(after))


        *op = strings.TrimSpace(strings.ToUpper(operator))

        break
    }
}

func main() {
    operators := [4]string{"/", "*", "+", "-"}
    var a, b, op string
    var roman bool
    var x, y int
    var result int

    fmt.Println("Welcome to kata-calculator")
    for {
        reader := bufio.NewReader(os.Stdin)
        console, err := reader.ReadString('\n')
        if err != nil {
            panic(err)
        }




        decompose(console, operators, &a, &b, &op) //после этой функции понятен оператор, а также записаны операнды
        result = solve(a, b, op, &x, &y, &roman)
        if roman == true {
            fmt.Print(intToRoman(&result), "\n", "Введите следующее выражение \n")
            
        } else {
            fmt.Print(result, "\n", "Введите следующее выражение \n")
        }
    }
}




func romanToInt(s string) int {
    
    compares := map[string]int{
        "I" : 1,
        "V" : 5,
        "X" : 10,
        "L" : 50,
        "C" : 100,
    }

    integer := 0

    for i := 0; i < len(s); i++ {
        current := string((s)[i])
        next := ""

        if i < (len(s)-1) {
            next = string((s)[i+1])
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

func intToRoman(num *int) string {
    if *num < 1 {
        panic("Римляне не в курсе про отрицательные числа и ноль")
    }

    comparesMap := map[int]string {
    1     :   "I" ,
    5     :   "V" ,
    10    :   "X" ,
    50    :   "L" ,
    100   :   "C" ,
    500   :   "D" ,
    1000  :   "M" ,
    }
    
    var roman string
    dec := 1 
    nums := strconv.Itoa(*num)

    for i := (len(nums)-1); i>=0; i-- {
        current, _ := strconv.Atoi(string(nums[i]))

        if current == 5 {
            roman = comparesMap[5*dec] + roman
        } else if current == 0 {
            dec*=10
            continue
        } else if current == 4 {
            roman = comparesMap[dec] + comparesMap[dec*5] + roman
        } else if current == 9 {
            roman = comparesMap[dec] + comparesMap[dec*10] + roman
        } else {
            if current > 5 {
                number := comparesMap[dec*5]
                for j:=0; j<(current-5); j++ {
                    number = number + comparesMap[dec]
                }
                roman = number + roman
            } else { 
                for j:=0; j<current; j++ {
                    roman = comparesMap[dec] + roman
                }
            }
        }

        dec *= 10
    }


    return roman
}
//
//
//func main() {
//    //считать операнды
//	reader := bufio.NewReader(os.Stdin)
//
//    var firstOperand    interface{}
//    var secondOperand   interface{}
//    var operator        string
//    var operation       []string
//    var oType           interface{}
//
//	
//    stroke, err := reader.ReadString('\n')
//    if err != nil {
//        panic(err)
//    }
//    operation = strings.Split(stroke, " ")
//    firstOperand = operation[0]
//    secondOperand = operation[2]
//    operator = operation[1]
//    
//
//    if reflect.TypeOf(firstOperand) != reflect.TypeOf(secondOperand) {
//        panic(err)
//    }
//    if reflect.TypeOf(firstOperand) == reflect.TypeOf(0) && strings.Atoi(firstOperand) > 10 {
//
//    }
//
//
//    switch {
//    case firstOperand.(type) == int && secondOperand.(type) == int:
//        
//    case firstOperand.(type) == string:
//        oType = operationType
//    }
//
//
//
//    switch firstOperand = operation[0]{
//    case 
//    }
//}



// DONE считать строку через буфио
// DONE разделить строку на операнды и оператор посредством пробела "операнд1" "оператор" "операнд2"
// DONE если арабские операнды больше 10 - выдать панику
// DONE если один операнд римский, а другой арабский - выдать панику
// 

// распознать оператор
// выполнить расчеты
// если операнды заданы римскимии числами - отдать резульат римским числом
// если операнды римские и результат <=0 - выдать панику
// отдать результат в консоль

