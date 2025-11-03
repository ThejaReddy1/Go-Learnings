
# Operators in Go

In Go, there are **five different types of operators**.

## Operators and Operands

Example:

```go
a + b
```

Here, `a` and `b` are **operands**, and `+` is the **operator**.

## 1. Comparison Operators

* `==` — equal to
* `!=` — not equal to
* `<` — less than
* `<=` — less than or equal to
* `>` — greater than
* `>=` — greater than or equal to

## 2. Arithmetic Operators

* `+` — addition
* `-` — subtraction
* `*` — multiplication
* `/` — division
* `%` — modulus (remainder)
* `++` — increment
* `--` — decrement

## 3. Assignment Operators

* `=` — assignment
* `+=` — add and assign
* `-=` — subtract and assign
* `*=` — multiply and assign
* `/=` — divide and assign
* `%=` — modulus and assign

## 4. Bitwise Operators

* `&` — bitwise AND
* `|` — bitwise OR
* `^` — bitwise XOR
* `<<` — left shift
* `>>` — right shift

## 5. Logical Operators

* `&&` — logical AND
* `||` — logical OR
* `!` — logical NOT

---

### Notes

* Use `==` to compare values; `=` is the assignment operator.
* In Go, `++` and `--` are **statements**, not expressions — they cannot be used inside expressions.

---

See the Go source files in this repository for concrete examples of using these operators.
