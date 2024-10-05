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
