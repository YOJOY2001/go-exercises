// 1. Write a function that takes a slice of integers and returns a new slice with only the even numbers.
// 	Input: [1, 2, 3, 4, 5] Output: [2, 4]

// 2. Write a function that takes a slice of strings and returns a new slice with all strings sorted in alphabetical order.
// 	Input: ["zebra", "monkey", "aardvark"] Output: ["aardvark", "monkey", "zebra"]

// 3. Write a function that takes a string and returns a map with the frequency count of each character in the string.
// 	Input: "hello" Output: {"h": 1, "e": 1, "l": 2, "o": 1}

// 4. Write a function that takes a slice of integers and returns the sum of all the elements in the slice.
// 	Input: [1, 2, 3, 4, 5] Output: 15

// 5. Write a function that takes a slice of integers and returns the largest number in the slice.
// 	Input: [1, 2, 3, 4, 5] Output: 5

// 6. Write a function that takes a slice of strings and returns a new slice with all strings that have more than 5 characters.
// 	Input: ["cat", "dog", "elephant", "lion"] Output: ["elephant"]

// 7. Write a function that takes a string and returns a new string with all vowels removed.
// 	Input: "hello" Output: "hll"

// 8. Write a function that takes a slice of integers and returns a new slice with all duplicates removed.
// 	Input: [1, 2, 3, 2, 4, 5, 5] Output: [1, 3, 4]

// 9. Write a function that takes a slice of integers and returns the average of all the elements in the slice.
// 	Input: [1, 2, 3, 4, 5] Output: 3

// 10. Write a function that takes a slice of strings and a string, and returns true if the string is contained in the slice.
// 	Input: (["cat", "dog", "elephant"], "dog") Output: true

// 11. Write a function that takes a slice of integers and returns the second-largest number in the slice.
// 	Input: [1, 2, 3, 4, 5] Output: 4

// 12. Write a function that takes a slice of integers and returns the index of the first occurrence of a given number.
// 	Input: ([1, 2, 3, 2, 4, 5], 2) Output: 1

// 13. Write a function that takes a slice of integers and a number and returns the number of times the number appears in the slice.
// 	Input: ([1, 2, 3, 2, 4, 5], 2) Output: 2

// 14. Write a function that takes a slice of integers and returns a new slice with the elements in reverse order.
// 	Input: [1, 2, 3, 4, 5] Output: [5, 4, 3, 2, 1]

// 15. Write a function that takes a slice of integers and returns true if the slice is sorted in ascending order.
// 	Input: [1, 2, 3, 4, 5] Output: true

// 16. Write a function that takes a slice of strings and a string, and returns the index of the first occurrence of the string in the slice.
// 	Input: (["cat", "dog", "elephant"], "dog") Output: 1

// 17. Write a function that takes a slice of strings and returns a new slice with all strings that start with a given letter.
// 	Input: (["cat", "dog", "elephant"], "e") Output: ["elephant"]

// 18. Write a function that takes a slice of strings and returns a new slice with all strings that are palindromes (i.e., read the same backwards as forwards).
// 	Input: ["racecar", "level", "hello"] Output: ["racecar", "level"]

// 19. Write a function that takes a string and returns a new string with all the words in reverse order.
// 	Input: "This is a test" Output: "test a is This"

// 20. Write a function that takes a slice of integers and returns a new slice with all the prime numbers.
// 	Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] Output: [2, 3, 5, 7]

// 21. Write a function that takes a slice of integers and returns a new slice with the elements in alternating order.
// 	Input: [1, 2, 3, 4, 5] Output: [1, 3, 5, 2, 4]

// 22. Write a function that takes a slice of integers and returns a new slice with all elements that are divisible by a given number.
// 	Input: ([1, 2, 3, 4, 5], 2) Output: [2, 4]

// 23. Write a function that takes a slice of integers and returns true if the slice contains a given number.
// 	Input: ([1, 2, 3, 4, 5], 3) Output: true

// 24. Write a function that takes a slice of strings and returns a new slice with all strings that are at least a given length.
// 	Input: (["cat", "dog", "elephant", "lion"], 5) Output: ["elephant"]

// 25. Write a function that takes a slice of strings and returns a new slice with all strings that contain a given letter.
// 	Input: (["cat", "dog", "elephant", "lion"], "e") Output: ["elephant"]

// 26. Write a function that takes a slice of strings and returns a new slice with all strings that are palindromes (i.e., read the same backwards as forwards).
// 	Input: ["racecar", "level", "hello"] Output: ["racecar", "level"]

// 27. Write a function that takes a string and returns a new string with all the words in reverse order, with the letters in each word reversed as well.
// 	Input: "This is a test" Output: "tset a si sihT"

// 28. Write a function that takes a slice of integers and returns the smallest number that is divisible by all elements in the slice.
// 	Input: [2, 3, 5] Output: 30

// 29. Write a function that takes a slice of integers and returns a new slice with all elements that are palindromes (i.e., read the same forwards as backwards).
// 	Input: [1, 11, 121, 12321] Output: [11, 121, 12321]

// 30. Write a function that takes a slice of integers and returns true if the slice is sorted in descending order.
// 	Input: [5, 4, 3, 2, 1] Output: true

package main

import (
	"fmt"

	"github.com/exercises/exce"
)

func main() {

	fmt.Println(exce.Filtereven([]int{1,2,3,4,5}))		//1 RETURN EVEN NUMBERS.

	fmt.Println(exce.Sorted([]string{"zebra","monkey","aardvark"}))		//2 RETURN IN ALPHABETICAL ORDER.

	fmt.Println(exce.WordCount("hello"))		//3 RETURN FREQUENCY COUNT OF EACH CHARACTER.

	fmt.Println(exce.Sum([]int{1, 2, 3, 4, 5}))		//4 RETURN SUM OF ALL ELEMENTS.

	fmt.Println(exce.Largest([]int{1,2,3,4,5}))		//5 RETURNS LARGEST NUMBER.

	fmt.Println(exce.Fivecr([]string{"cat", "dog", "elephant", "lion"}))		//6 RETURN ELEMENT CONTAINING MORE THAN 5 CHARACTER.

	fmt.Println(exce.Remvow("hello"))		//7 REMOVE THE VOWELS.

	fmt.Println(exce.Dupnum([]int{1, 2, 3, 2, 4, 5, 5})) 	//8 REMOVE DUPLICATE NUMBERS.

	fmt.Println(exce.Average([]int{1,2,3,4,5}))	//9 RETURNS THE AVERAGE.

	fmt.Println(exce.Check([]string{"cat", "dog", "elephant"},"dog"))	//10 RETURNS TRUE IF STRING IS PRESENT IN SLICE.

	fmt.Println(exce.SecLargest([]int{1,2,3,4,5}))	//11 RETURNS THE SECOND LAGEST NUMBER.
	
	fmt.Println(exce.Indexing([]int{1,2,3,4,5},2))	//12 INDEXING TNE NUMBER THROUGH SLICE.

	fmt.Println(exce.Repnum([]int{1, 2, 3, 2, 4, 5},2))  //13 RETURNS THE NO OF TIMES THE NUMBER REPEATED.

	fmt.Println(exce.Reverse([]int{1,2,3,4,5}))	//14 RETURNS THE ELEMENTS IN REVERSE ORDER.

	fmt.Println(exce.AscendingOrder([]int{1, 2, 3, 4, 5}))	//15 RETURNS IN ASCENDING ORDER.

	fmt.Println(exce.CheckOccur([]string{"cat", "dog", "elephant"},"dog"))	//16 RETURNS THE INDEX OF THE STRING.

	fmt.Println(exce.MatchLetter([]string{"cat", "dog", "elephant"},"e"))	//17 RETURNS STRING STARTING WITH E.

	fmt.Println(exce.PalinDrome([]string{"racecar", "level", "hello"}))	//18 RETURNS THE STRINGS THAT ARE PALLINDROME.

	fmt.Println(exce.PrintBackwards("This is a test"))	//19 RETURNS THE WORDS BACKWARDS.

	fmt.Println(exce.Prime([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))	//20 RETURNS THE PRIME NUMBERS IN SLICE.

	fmt.Println(exce.AlternateNum([]int{1, 2, 3, 4, 5}))	//21 PRINTS THE NUMBER ALTERNATELY PRESENT IN THE STRING.

	fmt.Println(exce.DivisibleByN([]int{1, 2, 3, 4, 5},2))	//22 PRINTS THE NUMBER DIVISIBLE BY N.

	fmt.Println(exce.CheckNum([]int{1, 2, 3, 4, 5},3))	//23 RETURNS TRUE IF SLICE CONTAINS THE GIVEN NUMBER.

	fmt.Println(exce.MinLength([]string{"cat", "dog", "elephant", "lion"},5))	//24 RETURNS THE STRING HAVING MINIMUM 5 CHARACTERS.

	fmt.Println(exce.MatchLetter([]string{"cat", "dog", "elephant"},"e"))	//25 RETURNS THE STRING CONTAINING THE GIVEN CHARACTER.

	fmt.Println(exce.PalinDrome([]string{"racecar", "level", "hello"}))	//26 RETURN PALLINDROME STRINGS.

	fmt.Println(exce.ReverseString("This is a test"))	//27 RETURNS THE STRING IN REVERSE.

	fmt.Println(exce.LCM([]int{2,3,5})) 	//28 RETURNS LCM OF THE NUMBERS.

	fmt.Println(exce.PalNum([]int{11, 12, 121, 12321}))	//29 RETURN PALLINDROME NUMBERS.
	
	fmt.Println(exce.DescOrder([]int{5, 4, 3, 2, 1}))	//30 RETURN TRUE IF PALLINDROME.

}
