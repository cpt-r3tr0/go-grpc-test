package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Error type
type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func EvaluateExp(expr string) float64 {

	// If only simple expression, evaluate directly
	if isSimpleExp(expr) {
		return evaluateSimpleExp(expr)
	}

	// Otherwise, use full expression evaluation
	return evaluateComplexExp(expr)

}

// Check if only simple expression
func isSimpleExp(expr string) bool {
	tokens := strings.Fields(expr)
	return len(tokens) == 3 && isOperand(tokens[0]) && isOperand(tokens[2])
}

// Evaluate simple expression
func evaluateSimpleExp(expr string) float64 {
	tokens := strings.Fields(expr)

	op1, _ := strconv.ParseFloat(tokens[0], 64)
	op2, _ := strconv.ParseFloat(tokens[2], 64)

	switch tokens[1] {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	case "/":
		return op1 / op2
	}

	return 0
}

func evaluateComplexExp(expr string) float64 {

	// // Split expression into tokens
	regex := regexp.MustCompile(`[\d\.]+|[\+\-\*\/]`)
	tokens := regex.FindAllString(expr, -1)

	// // Convert tokens to postfix notation
	postfix := shuntingYard(tokens)

	log.Println("tokens, ", tokens)

	// // Evaluate postfix expression
	return evaluatePostfix(postfix)
}

// Shunting yard algorithm to convert to postfix notation
func shuntingYard(tokens []string) []string {

	var stack []string
	var queue []string

	for _, token := range tokens {

		switch token {

		// If left parenthesis, push to stack
		case "(":
			stack = append(stack, token)

		// If right parenthesis, pop stack to queue until left parenthesis
		case ")":
			for len(stack) > 0 {
				op := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if op == "(" {
					break
				}
				queue = append(queue, op)
			}

		// If operator, pop operators from stack to queue until lower precedence
		case "+", "-", "*", "/":
			for len(stack) > 0 {
				op := stack[len(stack)-1]
				if isHigherPrecedence(op, token) {
					stack = stack[:len(stack)-1]
					queue = append(queue, op)
					continue
				}
				break
			}
			stack = append(stack, token)

		// If operand, add to queue
		default:
			queue = append(queue, token)
		}
	}

	// Pop remaining operators from stack to queue
	for len(stack) > 0 {
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		queue = append(queue, op)
	}

	return queue
}

// Evaluate postfix expression
func evaluatePostfix(postfix []string) float64 {

	var stack []float64

	for _, token := range postfix {

		// If operand, convert to float and push to stack
		if isOperand(token) {
			val, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, val)
		} else {
			// If operator, pop top two operands
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			var result float64
			switch token {
			case "+":
				result = op1 + op2
			case "-":
				result = op1 - op2
			case "*":
				result = op1 * op2
			case "/":
				if op2 == 0 {
					return math.Inf(1)
				}
				result = op1 / op2
			}

			// Push result back to stack
			stack = append(stack, result)
		}
	}

	return stack[0]
}

func isOperand(val string) bool {
	_, err := strconv.ParseFloat(val, 64)
	return err == nil
}

func isHigherPrecedence(op1, op2 string) bool {
	precedence := map[string]int{
		"(": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	return precedence[op1] >= precedence[op2]
}
