package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"runtime"
	"time"
	"unicode"

	//"os/user"
	"example/coretypes"
)

const (
	goodScore = 450
	lowScoreRatio = 10
	goodScoreRatio = 20
)

var(
	ErrCreditScore = errors.New("Invalid credit score")
	ErrIncome = errors.New("income invalid")
	ErrLoanAmount = errors.New("loan amount invalid")
	ErrLoanTerm = errors.New("loan term not a multiple of 12")
)

func main() {
	// var conferenceName string = "Go Conference"
	// const conferenceTickets int = 50
	// var remainingTickets int = 50

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	// var bookings []string
	// fmt.Println("test print")
	// for {
	// 	var firstName string
	// 	var lastName string
	// 	var userTickets int
	// 	var email string

	// 	fmt.Println("Enter your first name:")
	// 	fmt.Scan(&firstName)

	// 	fmt.Println("Enter your second name:")
	// 	fmt.Scan(&lastName)

	// 	fmt.Println("Enter your email address:")
	// 	fmt.Scan(&email)

	// 	fmt.Println("Enter number of tickets:")
	// 	fmt.Scan(&userTickets)

	// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// 	isValidEmail := strings.Contains(email, "@")
	// 	isUserTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// 	if isValidName && isUserTicketNumber && isValidEmail {

	// 		remainingTickets = remainingTickets - userTickets
	// 		// bookings[5] = firstName + " " + lastName
	// 		bookings = append(bookings, firstName+""+lastName)

	// 		fmt.Printf("The whole array: %v\n", bookings)
	// 		fmt.Printf("The first value: %v\n", bookings[0])
	// 		fmt.Printf("Array type: %T\n", bookings)
	// 		fmt.Printf("Length of array: %v\n", len(bookings))

	// 		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	// 		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// 		firstNames := []string{}
	// 		for _, booking := range bookings {
	// 			var names = strings.Fields(booking)
	// 			firstNames = append(firstNames, names[0])
	// 		}
	// 		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	// 		var noTicketsRemaining bool = remainingTickets > 0 || remainingTickets == 0
	// 		if noTicketsRemaining {
	// 			//end program
	// 			fmt.Println("Our conference is booked out. Come back next year.")
	// 			break
	// 		}
	// 	} else {
	// 		if !isValidName {
	// 			fmt.Println("first name or last name you entered is too short")
	// 		}
	// 		if !isValidEmail {
	// 			fmt.Println("Email address you entered doesn't contain @")
	// 		}
	// 		if !isUserTicketNumber {
	// 			fmt.Println("Number of tickets you entered is invalid!")
	// 		}
	// 		fmt.Println("Your input data is invalid, try again!")
	// 	}
	// }
	// TestIfElse()
	// Print1to100()
	// DayBorn()
	// MultipleCaseValues()
	// InitialStatement()
	// PrintBasicLoop()
	// LoopNames()
	// LoopOverMap()
	// BreakContinue()
	// BubbleSort()
	// fmt.Println("Test again")
	// if PasswordChecker("") {
	// 	fmt.Println("Password good")
	// } else {
	// 	fmt.Println("Password bad")
	// }
	// if PasswordChecker("This!I5A") {
	// 	fmt.Println("Password good")
	// } else {
	// 	fmt.Println("Password bad")
	// }
	// MemoryTest()
	// FloatingPoint()
	// OverFlow()
	// BigNumbers()
	// StringText()
	// MultiByteCharacter()
	// SafelyLoopOverString()
	// LengthOfMultiByte()
	// TheNilValue()

	// MemoryTest()
	// MemoryTest()
	// MemoryTest()
	// MemoryTest()
//	taxTotal := .0
	//taxTotal += SalesTaxCalculator(.99, 0.075)
//	taxTotal := .0
//	taxTotal := .0
//	taxTotal := .0

	//taxTotal += SalesTaxCalculator(2.75, 0.015)

	//taxTotal += SalesTaxCalculator(.87, 0.02)
	//fmt.Println("Sales Tax Total:", taxTotal)
	//LoanCalculator()
//	ComplexTypes()
//	var message string= coretypes.TestCoreTypes()
//	fmt.Print(message)
//	fmt.Printf("%#v\n", coretypes.DefineArray())
//	comp1, comp2, comp3 := coretypes.CompArrays()
//	fmt.Println("[5]int == [5]int{0}       :", comp1)
//	fmt.Println("[5]int == [...]int{0, 0, 0, 0, 0} :", comp2)
//	fmt.Println("[5]int == [5]int{0, 0, 0, 0, 9} :", comp3)
//	comp1, comp2, arr3 := coretypes.CompArrays2()
//	fmt.Println("[10]int == [...]int{9:0} :", comp1)
//	fmt.Println("[10]int == [10]int{1,9 :10, 4: 5}:", comp2)
//	fmt.Println("arr3 :", arr3)
//	fmt.Print(coretypes.Message())
//	var arr [10]int
//	arr = coretypes.FillArray(arr)
//	arr = coretypes.OpArray(arr)
//	fmt.Println(arr)
	if longest := coretypes.FindLongest(coretypes.GetPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}

}


func TestIfElse() {
	var num int
	fmt.Println("Enter the num for calculations")
	fmt.Scanf("%d", &num)
	fmt.Printf("Number entered is %d\n", num)
	if err := Validate(num); err != nil {
		fmt.Println(err)
	}
}

func Validate(input int) error {
	if input < 0 {
		return errors.New("Input can't be a negative number")
	} else if input > 100 {
		return errors.New("Input can't be over 100")
	} else if input%7 == 0 {
		return errors.New("Input can't be divided by 7")
	} else {
		return nil
	}
}

func Print1to100() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz.")
		} else if i%5 == 0 {
			fmt.Println("Buzz.")
		} else if i%3 == 0 {
			fmt.Println("Fizz.")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}

func DayBorn() {
	dayBorn := time.Monday

	switch dayBorn {
	case time.Monday:
		fmt.Println("Baby is born on Monday")
	case time.Tuesday:
		fmt.Println("Baby is born on Tuesday")
	case time.Wednesday:
		fmt.Println("Baby is born on Wednesday")
	case time.Thursday:
		fmt.Println("Baby is born on Thursday")
	case time.Friday:
		fmt.Println("Baby is born on Friday")
	case time.Saturday:
		fmt.Println("Baby is born on Saturday")
	case time.Sunday:
		fmt.Println("Baby is born on Sunday")
	default:
		fmt.Println("Baby born day is not valid!")
	}
}

func MultipleCaseValues() {
	dayBorn := time.Monday
	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Baby is born on Weekday!!")
	case time.Saturday, time.Sunday:
		fmt.Println("Baby is born on Weekend!!")
	default:
		//do nothing
	}
}

func InitialStatement() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Baby is born on Weekend!!")
	default:
		fmt.Println("Baby is born on other day!!")
	}
}

func PrintBasicLoop() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func LoopNames() {
	names := []string{"Jim", "Joes", "Jerry", "John", "Jess"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func LoopOverMap() {
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}
	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}

func BreakContinue() {
	rand.Seed(time.Now().UnixNano()) // Set a different seed value each time
	r := rand.Intn(10)               // Generate a random number between 0 and 9
	// fmt.Println(r)
	for {
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		} else {
			fmt.Println(r)
		}
	}
}

func BubbleSort() {
	nums := []int{5, 4, 6, 3, 2, 0, 8, 7, 9, 1}
	fmt.Println("Before:", nums)
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}
	fmt.Println("After:", nums)
}

func PasswordChecker(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	if len(pwR) > 15 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func MemoryTest() {
	var list []int8
	for i := 0; i <= 100000000; i++ {
		list = append(list, 100)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}

func FloatingPoint() {
	var a int = 100
	var b float32 = 100
	var c float64 = 100

	fmt.Println(a / 3)
	fmt.Println(b / 3)
	fmt.Println(c / 3)

	fmt.Println((a / 3) * 3)
	fmt.Println((b / 3) * 3)
	fmt.Println((c / 3) * 3)
}

func OverFlow() {
	var a int8 = 125
	var b uint8 = 253
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
}

func BigNumbers() {
	intA := math.MaxInt64
	intA += 1
	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64 :", math.MaxInt64)
	fmt.Println("Int :", intA)
	fmt.Println("Big Int :", bigA.String())
}

func StringText() {
	comment1 := `This is the BEST
	thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"

	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	comment1 = `In "Windows" the user directory is "C:\Users\"`
	comment2 = "In \"Windows\" the user directory is \"C:\\Users\\\""

	fmt.Println(comment1)

	fmt.Println(comment2)

}

func MultiByteCharacter() {
	username := "Sir_King_ärya"
	runes := []rune(username)

	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}

func SafelyLoopOverString() {
	logLevel := "コ案ヨさ"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}

func LengthOfMultiByte() {
	username := "Sir_King_ärya"
	fmt.Println("Bytes :", len(username))
	fmt.Println("Runes :", len([]rune(username)))

	//Limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}

func TheNilValue() {
	m := "This is text value"
	var message *string
	message = &m
	if message == nil {
		fmt.Println("error, unexpected nil value")
		return
	}
	fmt.Println(*message)
}

func SalesTaxCalculator(cost float64, taxRate float64) float64 {
	return cost * taxRate
}

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	interest :=20.0
	if creditScore >= goodScore {
		interest = 15.0
	}
	if creditScore < 1 {
		return ErrCreditScore
	}
	if loanAmount < 1 {
		return ErrLoanAmount
	}
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}
	rate := interest/100

	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)
	totalInterest := (payment * loanTerm) - loanAmount
	approved :=false
	if income > payment { 
		ratio := (payment / income) * 100 
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = false
		}
	}

	fmt.Println("Credit Score :", creditScore)
	fmt.Println("Income :", income)
	fmt.Println("Loan Amount : ", loanAmount)
	fmt.Println("Loan Term :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate :", interest)
	fmt.Println("Total Cost :", totalInterest)
	fmt.Println("Approved :", approved)
	fmt.Println("")

	return nil

}


func LoanCalculator(){
	//Approved
	fmt.Println("Applicant 1")
	fmt.Println("------------")
	err := checkLoan(500, 1000, 1000, 24)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//Denied
	fmt.Println("Applicant 2")
	fmt.Println("------------")
	err = checkLoan(350, 1000, 10000, 12)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
