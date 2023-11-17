package main

import (
	"math"
	"romannumeral"
	"strings"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Функция, чтобы проверить, какое число символов в строке x содержится в проверочной строке w
func CountAny (x string, w string) int {
	y := 0
	l := len(w)
	for i := 0; i < l ; i++ {
		y += strings.Count(x, string(w[i]))
	}
	return(y)
}

//функция вычисления

func Calculation (a1 int, a2 int, act string) int {
	y := 0
	switch act {
	case "+":
		y = a1+a2
	case "-":
		y = a1-a2
	case "/":
		y = a1/a2
	case "*":
		y = a1*a2
	}
	return(y)
}

// функция для вывода
func Result_string (x string) string {
	l := len(x)
// проверка наличия искомых операторов в строке
	NVAction := !(strings.ContainsAny(x, "+-*/"))
// проверка количества искомых операторов в строке
	NVActionsNum := !(CountAny(x, "+-*/") == 1)
// место нахождения оператора в строке. -1 заменена на 0
	Ind := int(math.Max( float64 (strings.IndexAny(x, "+-*/")) , 0))
// проверка корректности места нахождения оператора
	NVActionPlace := Ind == l-1 || Ind == 0
// проверка на некорректные символы
	NVSymbol := !(CountAny(x, "0123456789IVXLCDM+-*/")==l)
// проверка на наличие символов, которых не должно быть в выражении с арабскими числами
	arabian := CountAny(x, "0123456789+-*/")==l
// проверка на наличие символов, которых не должно быть в выражении с римскими числами
	roman := CountAny(x, "IVXLCDM+-*/")==l
// флаг наличия одновременно арабских и римских чисел
	RomArabMix := !( arabian || roman )
// поиск математической операции
	act:=string(x[Ind])
// разделям строку на операнды
	a := strings.Split(x, act)
// преобразование арабских операндов в число
	arab1, _ := strconv.Atoi(a[0])
	arab2, _ := strconv.Atoi(a[1])
// преобразование римских операндов в число
	roman1, err1 := romannumeral.StringToInt(a[0])
	roman2, err2 := romannumeral.StringToInt(a[1])
// функция romannumeral.StringToInt содержит ошибку:
// ряд некорректных римских чисел (напр. IIIIII) всё равно преобразуется в число
// поэтому необходима дополнительная проверка
	roman1_check, _ := romannumeral.IntToString(roman1)
	roman2_check, _ := romannumeral.IntToString(roman2)
// флаг некорректности римского числа
	NVRoman := (err1 != nil) || (err2 != nil) || !(roman1_check == a[0]) || !(roman2_check == a[1])
	y := ""
	res := 0
// перевод флагов в коды ошибок
	switch {
	case NVSymbol:
		y = "Вывод ошибки, тaк как строка содержит неожиданный символ."
	case NVAction:
		y = "Вывод ошибки, так как строка не является математической операцией."
	case NVActionsNum||NVActionPlace:
		y = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	case RomArabMix:
		y = "Вывод ошибки, так как используются одновременно разные системы счисления."
// проверка, что арабские числа не меньше 1
	case arabian && (arab1<1 || arab2<1):
		y = "Вывод ошибки, так как одно из введённых чисел меньше 1."
// проверка, что арабские числа не больше 10
	case arabian && (arab1>10 || arab2>10):
		y = "Вывод ошибки, так как одно из введённых чисел больше 10."
// вычисление результатов
	case arabian:
		res = Calculation(arab1, arab2, act)
		y = strconv.Itoa (res)
	case NVRoman:
		y = "Вывод ошибки, так как одно из римских чисел некорректно."
// проверка, что римское число не больше 10
	case roman1>10 || roman2>10:
		y = "Вывод ошибки, так как одно из введённых римских чисел больше 10."
// вычисления
	case Calculation(roman1, roman2, act)<0:
		y = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	case Calculation(roman1, roman2, act)==0:
		y = "Вывод ошибки, так как в римской системе нет нуля."
	default:
		y, _ = romannumeral.IntToString(Calculation(roman1, roman2, act))
	}
	return(y)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println(Result_string(text))
}