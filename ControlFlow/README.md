# Control Flow in Go

This directory contains examples demonstrating various control flow mechanisms in Go programming language.

## Files Overview

### 1. Conditional Statements

#### [ifConditional.go](ifConditional.go)
Basic if-else conditional statements
- Simple if-else structure
- Condition evaluation with string comparison
- Proper else block syntax (must be on same line as closing brace)

#### [if-else-ifConfitional.go](if-else-ifConfitional.go)
Extended conditional statements with multiple conditions
- if-else-if chains
- Multiple condition testing
- Short variable declaration with `:=`

### 2. Switch Statements

#### [switchStatement.go](switchStatement.go)
Traditional switch statements with expression matching
- Basic switch syntax with expression
- Multiple case values in single case
- `fallthrough` keyword usage
- Default case handling

#### [switchWithConditional.go](switchWithConditional.go)
Switch statements with conditional expressions
- Switch without expression (condition-based)
- Boolean expressions in case statements
- Implicit break behavior in Go
- Range-based conditions

### 3. Loops

#### [loopingWithFor.go](loopingWithFor.go)
For loop variations and usage
- Standard for loop with initialization, condition, and post statement
- Simplified for loop (while-like behavior)
- Infinite loop structure
- Loop counter operations

#### [breakAndContinue.go](breakAndContinue.go)
Loop control statements
- `break` statement for early loop termination
- `continue` statement for skipping iterations
- Practical examples with conditional logic

## Key Go Control Flow Concepts

### If Statements
```go
if condition {
    // code block
} else if anotherCondition {
    // code block
} else {
    // code block
}
```

### Switch Statements
```go
// Expression-based switch
switch variable {
case value1:
    // code
case value2, value3:
    // multiple values
default:
    // default case
}

// Condition-based switch
switch {
case condition1:
    // code
case condition2:
    // code
}
```

### For Loops
```go
// Standard for loop
for i := 0; i < 10; i++ {
    // code
}

// While-like loop
for condition {
    // code
}

// Infinite loop
for {
    // code
}
```

### Loop Control
- `break`: Exit loop immediately
- `continue`: Skip current iteration
- `fallthrough`: Continue to next case in switch (explicit)

## Important Notes

1. **Brace Placement**: In Go, the opening brace `{` must be on the same line as the control statement
2. **Implicit Break**: Switch statements in Go have implicit break - no need to add `break` at the end of each case
3. **Short Declaration**: Use `:=` for short variable declaration within control structures
4. **Boolean Expressions**: Use switch without expression for condition-based logic
5. **Multiple Values**: Switch cases can handle multiple values separated by commas

## Running the Examples

To run any example:
```bash
go run filename.go
```

Each file contains detailed comments explaining the concepts and syntax used.