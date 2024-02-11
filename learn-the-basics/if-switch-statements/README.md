# Conditional Statements

Conditional statements are used to run code only if a certain condition is true; go supports:

- **if** statements
- **if / else** statements
- **switch** **case** statements

# If

Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

[See this example.](./examples/if/main.go)

# If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

[See this example.](./examples/if-with-a-short-statement/main.go) (Try using v in the last return statement.)

# If and else

Variables declared inside an if short statement are also available inside any of the else blocks.

[See this example.](./examples/if-and-else/main.go) (Both calls to pow return their results before the call to fmt.Println in main begins.)

# Another examples if/else

Branching with if and else in Go is straight-forward.

You can have an if statement without an else.

Logical operators like && and || are often useful in conditions.

A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.

[See this example.](./examples/another-examples-if-else/main.go)

# Switch

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

[See this example.](./examples/switch/main.go)

# Switch evaluation order

Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

```
switch i {
case 0:
case f():
}
```

does not call f if i==0.)

Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.

[See this example.](./examples/switch-evaluation-order/main.go)

# Switch with no condition

Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else chains.

[See this example.](./examples/switch-with-no-codition/main.go)

# Reference(s)

[A Tour of Go](https://go.dev/tour/flowcontrol/5)

[Go by Example: If/Else](https://gobyexample.com/if-else)
