package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"runtime"
	"time"
	"unicode"
	//"os/user"
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
	MultiByteCharacter()
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
