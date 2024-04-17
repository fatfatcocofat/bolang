package evaluator

import (
	"bo/ast"
	"bo/evaluator/builtins"
	"bo/object"
	"fmt"
)

var (
	NIL   = &object.Nil{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.LetStatement:
		val := Eval(node.Value, env) // Evaluate the expression in the LetStatement node
		if isError(val) {
			return val
		}

		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *ast.PostfixExpression:
		return evalPostfixExpression(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.Return{Value: val}
	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Body: body, Env: env}
	case *ast.CallExpression:
		fn := Eval(node.Function, env)
		if isError(fn) {
			return fn
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(fn, args)
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}

		return &object.Array{Elements: elements}
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}

		return evalIndexExpression(left, index)
	case *ast.MapLiteral:
		return evalMapLiteral(node, env)
	case *ast.PrintStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		return evalPrintStatement(val)
	}

	return NIL
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.Return:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalPrintStatement(val object.Object) object.Object {
	switch val := val.(type) {
	case *object.Integer:
		fmt.Println(val.Inspect())
	case *object.Float:
		fmt.Println(val.Inspect())
	case *object.String:
		fmt.Println(val.Inspect())
	case *object.Boolean:
		fmt.Println(val.Inspect())
	case *object.Array:
		fmt.Println(val.Inspect())
	case *object.Map:
		fmt.Println(val.Inspect())
	case *object.Nil:
		fmt.Println(val.Inspect())
	default:
		fmt.Println(val.Type())
	}

	return NIL
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if ok {
		return val
	}

	if fn, ok := builtins.BuiltinsFn[node.Value]; ok {
		return fn
	}

	return newError("identifier not found: " + node.Value)
}

func evalMapLiteral(node *ast.MapLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.MapKey]object.MapPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		mapKey, ok := key.(object.Mapable)
		if !ok {
			return newError("unusable as map key: %s", key.Type())
		}

		pairs[mapKey.MapKey()] = object.MapPair{Key: key, Value: value}
	}

	return &object.Map{Pairs: pairs}
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.MAP_OBJ:
		return evalMapIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalMapIndexExpression(mapObj, index object.Object) object.Object {
	mapObject := mapObj.(*object.Map)
	key, ok := index.(object.Mapable)
	if !ok {
		return newError("unusable as map key: %s", index.Type())
	}

	pair, ok := mapObject.Pairs[key.MapKey()]
	if !ok {
		return NIL
	}

	return pair.Value
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NIL
	}

	return arrayObject.Elements[idx]
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right)
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalMixedInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalMixedInfixExpression(operator, left, right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String)

	switch operator {
	case "+":
		rightVal, ok := right.(*object.String)
		if !ok {
			return newError("non-string argument for string concatenation: %s", right.Type())
			// rightVal = &object.String{Value: right.Inspect()}
		}

		return &object.String{Value: leftVal.Value + rightVal.Value}
	case "*":
		rightVal, ok := right.(*object.Integer)
		if !ok {
			return newError("non-integer argument for string multiplication: %s", right.Type())
		}

		result := ""
		for i := 0; i < int(rightVal.Value); i++ {
			result += leftVal.Value
		}

		return &object.String{Value: result}
	case "==":
		rightVal, ok := right.(*object.String)
		if !ok {
			return FALSE
		}

		return nativeBoolToBooleanObject(leftVal.Value == rightVal.Value)
	case "!=":
		rightVal, ok := right.(*object.String)
		if !ok {
			return TRUE
		}

		return nativeBoolToBooleanObject(leftVal.Value != rightVal.Value)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalMixedInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := float64(0)
	rightVal := float64(0)

	switch left := left.(type) {
	case *object.Integer:
		leftVal = float64(left.Value)
	case *object.Float:
		leftVal = left.Value
	}

	switch right := right.(type) {
	case *object.Integer:
		rightVal = float64(right.Value)
	case *object.Float:
		rightVal = right.Value
	}

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value

	switch operator {
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalPostfixExpression(node *ast.PostfixExpression, env *object.Environment) object.Object {
	operator := node.Operator
	switch operator {
	case "++":
		return evalPostfixIncrementExpression(node, env)
	case "--":
		return evalPostfixDecrementExpression(node, env)
	default:
		return newError("unknown operator: %s", operator)
	}
}

func evalPostfixIncrementExpression(node *ast.PostfixExpression, env *object.Environment) object.Object {
	val, ok := env.Get(node.Token.Literal)
	if !ok {
		return newError("identifier not found: " + node.Token.Literal)
	}

	switch val := val.(type) {
	case *object.Integer:
		oldValue := val.Value
		val.Value++
		env.Set(node.Token.Literal, val)
		return &object.Integer{Value: oldValue}
	case *object.Float:
		oldValue := val.Value
		val.Value++
		env.Set(node.Token.Literal, val)
		return &object.Float{Value: oldValue}
	default:
		return newError("invalid left-hand side expression in postfix operation")
	}
}

func evalPostfixDecrementExpression(node *ast.PostfixExpression, env *object.Environment) object.Object {
	val, ok := env.Get(node.Token.Literal)
	if !ok {
		return newError("identifier not found: " + node.Token.Literal)
	}

	switch val := val.(type) {
	case *object.Integer:
		oldValue := val.Value
		val.Value--
		env.Set(node.Token.Literal, val)
		return &object.Integer{Value: oldValue}
	case *object.Float:
		oldValue := val.Value
		val.Value--
		env.Set(node.Token.Literal, val)
		return &object.Float{Value: oldValue}
	default:
		return newError("invalid left-hand side expression in postfix operation")
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NIL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -right.Value}
	case *object.Float:
		return &object.Float{Value: -right.Value}
	}

	return newError("unknown operator: -%s", right.Type())
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	}

	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}

	return NIL
}
