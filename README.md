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

## 06 - Another base to decimal [code]

Source file: `base_to_dec.go`
Function def: `BaseToDec(value string, base int) int`

Each digit gets multiplied by the base raised to a power. Eg for 1110 in base 2:

```
1     2     3
10^2  10^1  10^0

1    1    1    0
2^3  2^2  2^1  2^0

so this is: 1*2^3 + 1*2^2 + 1*2^1 + 0*2^0 = 8 + 4 + 2 + 0 = 14
```

```go
BaseToDec("1", 2) // 1
BaseToDec("10", 2) // 2
BaseToDec("21", 3) // 7
BaseToDec("1110", 2) // 14
BaseToDec("E", 16) // 14
BaseToDec("11", 16) // 17
```

## 07 - Any base to any base [code]

Source file: `base_to_base.go`
Function def: `BaseToBase(value string, base, newBase int) string`

Break it into smaller problems we have already solved!

Ex:

```go
BaseToBase("E", 16, 2) // "1110"
BaseToBase("8831A383B", 12, 16) // "DEADBEEF"
```

## 08 - Find two that sum [code]

Source file: `find_two_that_sum.go`
Function def: `FindTwoThatSum(numbers []int, sum int) (int, int)`

Find two numbers in a list that sum to a given amount.

FindTwoThatSum will look for two numbers in the provided numbers list that sum up to the sum argument. It will then return the indices of each of these numbers.

If there are multiple correct answers, any of them may be used. If there are no correct answers, (-1, -1) should be returned.

Ex:

```go
FindTwoThatSum([]int{1, 2, 3, 4}, 7) // (2, 3)
FindTwoThatSum([]int{0, -1, 1}, 0) // (1, 2)
FindTwoThatSum([]int{0, 1, 1}, 0) // (-1, -1)
```

Or if we have duplicate answers any of them are okay. All of the following are correct.

```go
FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4) // (5, 6)
FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4) // (1, 3)
FindTwoThatSum([]int{10, 1, 12, 3, 7, 2, 2, 1}, 4) // (1, 7)
```

Each number will only be used once, and the original numbers list should not be rearranged or altered in any way.

If you want to challenge yourself further, try adding different criteria for which numbers will be used when there are multiple correct answers.
Eg:

1. Always return the lowest index possible for the first value.
2. Always return the index of lowest value possible for the first return value.
3. Always return the index of the two values who have a minimal difference. Eg prefer the values 2, 2 over 1, 3 over 0, 4 for the sum of 4.

## 09 - Prime factorization (primes provided) [code]

Source file: `factor.go`
Function def: `Factor(primes []int, number int) []int`

Factor takes in a list of primes and a number and factors that number with the provided primes.

The returned numbers can be in any order; tests will sort them in increasing order to make checking easier.

Bonus: Any remainder (aside from 1) that can't be factored will be treated as a prime in the results.

Examples:

```go
Factor([]int{2,3,5,7}, 30) // []int{2,3,5}
Factor([]int{2,3,5,7}, 28) // []int{2,2,7}
Factor([]int{2,3,5,7}, 720) // []int{2,2,2,3,3,5}
```

Examples with remainders:

```go
Factor([2,5], 30) // []int{2,5,3}
Factor([3,5], 720) // []int{3,3,5,16}
Factor([], 4) // []int{4}
```
