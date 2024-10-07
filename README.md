# algorithms-with-go

# challenges

## 01 - Determine if a number is in a list [code]

Source file: `num_in_list.go`
Function def: `NumInList(list []int, num int) bool`

Given a list of numbers, determine if a specific number is in that list.

Ex:

```go
NumInList([]int{1,2,3,4,5}, 5) // true
NumInList([]int{3,3,3,3,3}, 5) // false
NumInList([]int{3,5,3,5,3}, 5) // true
NumInList([]int{4,2,22,-10,8}, -10) // true

// empty lists!
NumInList(nil, 5) // false
NumInList([]int{}, 5) // false
```

## 02 - Sum all the numbers in a list [code]

Source file: `sum.go`
Function def: `Sum(numbers []int) int`

Given a list of numbers, sum them all up and return the sum.

Ex:

```go
Sum([]int{1,2,3,4,5}) // 15
Sum([]int{3,3}) // 6
Sum([]int{3,5,3,5,3}) // 19
Sum([]int{4,2,22,-10,8}) // 26

// let's just return 0 for empty lists
Sum(nil) // 0
Sum([]int{}) // 0
```

## 03 - Reverse a string [code]

Source file: `reverse.go`
Function def: `Reverse(word string) string`

Given a string, return its reverse.

Ex

```go
Reverse("cat") // "tac"
Reverse("ca t") // "t ac"
Reverse("alphabet") // "tebahpla"
```

## 04 - FizzBuzz [code]

Source file: `fizz_buzz.go`
Function def: `FizzBuzz(n int)`

Given a number N, print out all the numbers from 1 to N but when a number is divisible by 3 print out "Fizz" instead of the number, and when a number is divisible by 5 print out "Buzz" instead of the number. For numbers divisible by both 3 and 5 print "Fizz Buzz".

Ex

```go
FizzBuzz(1)
// 1
FizzBuzz(5)
// 1, 2, Fizz, 4, Buzz
FizzBuzz(15)
// 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz
```

## 05 - Decimal to another base [code]

Source file: `dec_to_base.go`
Function def: `DecToBase(dec, base int) string`

You can convert a number from one base to another by:

1. Take the number you have in decimal and get the remainder when it is divided by the new base. The remainder is the next "digit" in the new base. You add this digit to the LEFT of the new number as you create it.

_NOTE: If the digit is >=10, then you need to use a letter like A, B, C, D, ..._

2. Divide your number by the new base, then repeat from step (1) if the new number is > 0.

Example 1:

```
14 to base 2:

Step 1: 14 % 2 = 0, so next digit is "0".
  Our new number is currently "0"

Step 2: 14 / 2 = 7, so we update our number to be 7
  7 is > 0, so we go to step 1...

Step 1: 7 % 2 = 1; next digit is "1"
  Our new number is currently "10" (we add the new digit to the left)

Step 2: 7 / 2 = 3 (round down always), so we update our number to be 3
  3 > 0, so go to step 1

Step 1: 3 % 2 = 1; next digit is "1"
  Our new number is "110"

Step 2: 3 / 2 = 1; Go to step 1

Step 1: 1 % 1 = 1; next digit is "1"
  Our new number is "1110"

Step 2: 1 / 2 = 0; stop here!

Final number is "1110", which is 14 in base 2!
```

Example 2:

```
15 to base 16:

Step 1: 15 % 16 = 15, so next digit is "F"
  Current new number: "F"

Step 2: 15 / 16 = 0. stop!

Final new number: "F"
```

_Note: If it is easier, you can use the `Reverse` function from earlier to help solve this. Breaking problems into smaller problems and reusing solutions is often a great way to make them easier._

Ex

```go
DecToBase(1, 2) // "1"
DecToBase(2, 2) // "10"
DecToBase(7, 3) // "21"
DecToBase(14, 2) // "1110"
DecToBase(14, 16) // "E"
DecToBase(17, 16) // "11"
```
