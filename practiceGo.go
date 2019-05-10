package main

//note that this practice is inspired by https://blog.csdn.net/caileigood/article/details/84107879, most of the functions are trying to solve the problems in the beginnning part of the pracitce

//this is just a small practice for golang, hopefully will have some better projects in go in the future
import(
	"fmt"
	"math"
	"strings"
)

//1. get ride of space and change line in a string
func modifyString(){
	test_string := "what\n is in your shose?\n i don't know!"
	test_string = strings.Replace(test_string, "\n", "", -1)
	test_string = strings.Replace(test_string, " ", "", -1)
	fmt.Println(test_string)
}

//2.finding if the num is even or odd
func judgeEven(num int){
	if num%2 == 0{
		fmt.Println("it is even")
	}else{
		fmt.Println(" it is odd")
	}
}

//3. finding the prime number withing 100, note that prime numbers can only be divided by 1 or itself
//probably have to try finding the number one by one, just need to check nums less than itself, for loop
func printPrime(){
	isPrime := true
	for i := 1; i < 100; i++{
		for j := 1; j < i; j++{
			if i != j && j != 1 && i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime{
			fmt.Println(" %d is a prime", i)
		}
		isPrime = true   //reset the boolean
	}
}

//4. multiplication talbe withing 9*9
func multiplication(){
	for i := 1; i <= 9; i++{
		for j := 1; j <= i; j++{
			fmt.Printf(" %d*%d=%d\t", i, j, i*j)//if wanting to print multiple numbers, must do the c style printing line		
		}
		fmt.Println()			//for formating
	}
}


//5.hui yang's triangle, will print out the number of lines with the triangle depending on the num of the rows
func triangleNum(rows int){
	numbers := make([]int, rows)  
	//making an array length of rows, the difference between make and new is that make can give the uninitialized first element the correct type, not nill
	for i := 0; i < rows; i++{
		for j:= 0; j < (i + 1); j++{
			//if it is the first row, first column, print 1
			//if it is the first column or the last column, print out 1
			//if not, it equals to the previous line's above 2 value's sum
			// the number row here is used to recode numbers
			var num int
			var numLen = len(numbers)
			if j ==0 || j == i{
				num = 1
			}else{
				num = numbers[numLen - i] + numbers[ numLen - i - 1]
			}
			numbers = append(numbers, num)	//the way of appending requires the original array and the new elements	
			fmt.Print(num, " ")
		}
		fmt.Println("")
	}//if using new to initialize, it will tell you it is nill and can't add more, also, new can't be initialized with a size 
}

//6. to separate the students' scores by a switch statement, depending on the score value
func separate_scores(score int) string{
	var letter_grade string
	switch{
		case score >= 90: letter_grade = "A"
		case score < 90 && score >= 60: letter_grade = "B"
		case score < 60: letter_grade = "C"
		default: letter_grade = "D"
	}
	return letter_grade
}

//7. define a structure, including x and y axis, finding the 2d distance between those two points
type loc struct{
	x float64
	y float64
}

func distance_calculator(p1, p2 loc)float64{
	return math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)

}


//8.using iota to form a week, using 1-7 to represent the week
//iota can only be used in the const expression, can't be used separately
//also note that const need to use small bracket
const(
	start = iota    //not sure in the original version it uses "_ int = iota"
	monday 
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)


//9. define 2 slices, compare those two slice, can compare the lenth and the elements inside of the two slices
//the compared type will be set as strings
//not sure how to use generic in go, seems like the documentation is missing or i didn't find them
func sliceComparator(ar1, ar2 []string) bool{
	if len(ar1) != len(ar2){
		return false
	}else{
		for i,_ := range ar1{	//only need the index here, didn't use the value, well, can also use value if u really want
			if ar1[i] != ar2[i]{
				return false
			}
		}
		return true
	}
}

//10. define a map, use student name as the key and the score as the value, find the value by keys and delete the index in which the corresponding key is not found
//need to use pointer here in order to change the actual data from outside the function
func stuSearch(stuN *[]string, name_score map[string]int){
	for i, name := range *stuN{
		score, exist := name_score[name]
		if !exist && score <= 60{
			remove(stuN, i)
		}
	}
}

func remove(stuN *[]string, i int){		//basically shift everything behind the index i to the previous one position
	var temp []string = *stuN    //copy the same content
	for j := i; j < len(temp)-1; j++{
		temp[j] = temp[j+1]
	}
	*stuN = temp[:len(temp)-1]		//copy everything down to the original array
}


func main(){
	//modifyString()
	//judgeEven(2)
	//printPrime()
	//multiplication()
	//triangleNum(10)
	/*var score string= separate_scores(88)
	fmt.Println(score)*/
	/*p1 := loc{-2, 3}			//need to put loc in front to specify the data type
	p2 := loc{1, 2}
	result := distance_calculator(p1, p2)
	fmt.Println(result)*/
	//fmt.Println(monday, tuesday, sunday)
	/*ar1 := []string{"a", "b", "c"}
	ar2 := make([]string, 3)
	copy(ar2, ar1)		//or according to the websiteonline, u can use ar2 := ar1[:]
	result := sliceComparator(ar1, ar2)
	fmt.Println(result)*/
	stu_arr := []string{"pierogie", "dumpling", "bum", "general tzuo's chicken", "taco"}
	score_map := map[string]int{"pierogie":70, "dumpling":90, "rosted duck":65}
	stuSearch(&stu_arr, score_map)		//sending the address of the stu_arr to the fuction
	fmt.Println(stu_arr)
	
}