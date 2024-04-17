package evaluator

import (
	"bolang/lexer"
	"bolang/object"
	"bolang/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"Hello, World!"`, "Hello, World!"},
		{`"Hello" + ", " + "World!"`, "Hello, World!"},
		{`"Hello" + ", " + "World" + "!"`, "Hello, World!"},
		{`"BO" * 3`, "BOBOBO"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected)
	}
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
		return false
	}
	return true
}

func TestEvalFloatExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.5", 5.5},
		{"10.5", 10.5},
		{"-5.5", -5.5},
		{"-10.5", -10.5},
		{"5.5 + 5.5 + 5.5 + 5.5 - 10.5", 11.5},
		{"2.5 * 2.5 * 2.5 * 2.5 * 2.5", 97.65625},
		{"-50.5 + 100.5 + -50.5", -0.5},
		{"5.5 * 2.5 + 10.5", 24.25},
		{"5.5 + 2.5 * 10.5", 31.75},
		{"20.5 + 2.5 * -10.5", -5.75},
		{"50.5 / 2.5 * 2.5 + 10.5", 61.0},
		{"2.5 * (5.5 + 10.5)", 40},
		{"3.5 * 3.5 * 3.5 + 10.5", 53.375},
		{"3.5 * (3.5 * 3.5) + 10.5", 53.375},
		{"(5.5 + 10.5 * 2.5 + 15.5 / 3.5) * 2.5 + -10.5", 79.94642857142858},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func testFloatObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
		return false
	}
	return true
}

func TestEvalFloatAndIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5 + 5.5", 10.5},
		{"10.5 + 10", 20.5},
		{"-5 + -5.5", -10.5},
		{"-10.5 + -10", -20.5},
		{"5 * 5.5", 27.5},
		{"10.5 * 10", 105},
		{"-5 * -5.5", 27.5},
		{"-10.5 * -10", 105},
		{"5 / 5.5", 0.9090909090909091},
		{"10.5 / 10", 1.05},
		{"-5 / -5.5", 0.9090909090909091},
		{"-10.5 / -10", 1.05},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testFloatObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"(1 >= 1) == true", true},
		{"(1 >= 1) == false", false},
		{"(1 <= 1) == true", true},
		{"(1 <= 1) == false", false},
		{"(1 <= 2) == true", true},
		{"(1 <= 2) == false", false},
		{"(1 >= 2) == true", false},
		{"(1 >= 2) == false", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func TestEvalBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!!true", true},
		{"!!false", false},
		{"!5", false},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestEvalIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNilObject(t, evaluated)
		}
	}
}

func testNilObject(t *testing.T, obj object.Object) bool {
	if obj != NIL {
		t.Errorf("object is not NIL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}

func TestEvalReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`if (10 > 1) { if (10 > 1) { return 10; } return 1; }`, 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"5 + true;",
			"unknown operator: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"unknown operator: INTEGER + BOOLEAN",
		},
		{
			`"Hello" - "World"`,
			"unknown operator: STRING - STRING",
		},
		{
			`"Hello" / "World"`,
			"unknown operator: STRING / STRING",
		},
		{
			`10 * "World"`,
			"unknown operator: INTEGER * STRING",
		},
		{
			"-false",
			"unknown operator: -BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
				if (10 > 1) {
					if (10 > 1) {
						return true + false;
					}
					return 1;
				}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`{"name": "Bo"}[fn(x) { x }];`,
			"unusable as map key: FUNCTION",
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expected {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expected, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
		{"let name = \"Bo\"; name;", "Bo"},
		{"let name = \"Bo\"; let b = name; b;", "Bo"},
		// {"let name = \"Bo\"; let b = name; let c = name + b + \"lang\"; c;", "BoBolang"},
		// {"let name = \"Bo\"; let b = name; let c = name + b + \"lang\"; let d = c + \"v0.1.0\"; d;", "BoBolangv0.1.0"},
		{"let a = 5; let b = 5.5; a + b;", 10.5},
		{"let a = 5; let b = 5.5; a * b;", 27.5},
		{"let a = 5; let b = 5.5; a / b;", 0.9090909090909091},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			testStringObject(t, evaluated, expected)
		}
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { return x; }"
	evaluated := testEval(input)

	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "return x;"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5.5);", 5.5},
		{"let identity = fn(x) { return x; }; identity(\"Bo\");", "Bo"},
		{"let double = fn(x) { return x * 2; }; double(5);", 10},
		{"let double = fn(x) { return x * 2; }; double(5.5);", 11.0},
		{"let add = fn(x, y) { return x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { return x + y; }; add(5.5, 5.5);", 11.0},
		{"let add = fn(x, y) { return x + y; }; add(5, 5.5);", 10.5},
		{"fn(x){ x; }(5)", 5},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			testStringObject(t, evaluated, expected)
		}
	}
}

func TestClosures(t *testing.T) {
	input := `
		let newAdder = fn(x) {
			fn(y) { x + y };
		};
		let addTwo = newAdder(2);
		addTwo(2);
	`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 4)
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong number of elements. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[1, 2, 3][0]", 1},
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][2]", 3},
		{"let i = 0; [1][i];", 1},
		{"let arr = [1, 2, 3]; arr[0];", 1},
		{"let arr = [1, 2, 3]; arr[1];", 2},
		{"let arr = [1, 2, 3]; arr[2];", 3},
		{"let arr = [1, 2, 3]; let i = arr[0]; arr[i];", 2},
		{"[1, 2, 3][3]", nil},
		{"[1, 2, 3][-1]", nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case nil:
			testNilObject(t, evaluated)
		}
	}
}

func TestMapLiterals(t *testing.T) {
	input := `
		let two = "two";
		{
			"one": 10 - 9,
			two: 1 + 1,
			"thr" + "ee": 6 / 2,
			4: 4,
			3.0: 3.0,
			true: 5,
			false: 6,
		}
	`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Map)
	if !ok {
		t.Fatalf("object is not Map. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.MapKey]int64{
		(&object.String{Value: "one"}).MapKey():   1,
		(&object.String{Value: "two"}).MapKey():   2,
		(&object.String{Value: "three"}).MapKey(): 3,
		(&object.Integer{Value: 4}).MapKey():      4,
		(&object.Float{Value: 3.0}).MapKey():      3.0,
		TRUE.MapKey():                             5,
		FALSE.MapKey():                            6,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("map has wrong number of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		switch v := pair.Value.(type) {
		case *object.Integer:
			testIntegerObject(t, &object.Integer{Value: v.Value}, expectedValue)
		case *object.Float:
			testFloatObject(t, &object.Float{Value: v.Value}, float64(expectedValue))
		}
	}
}

func TestMapIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`let two = "two"; { "one": 10, two: 2 }["one"]`, 10},
		{`let two = "two"; { "one": 10, two: 2 }[two]`, 2},
		{`{ 1: 10, 2: 2 }[1]`, 10},
		{`{ 1: 10, 2: 2 }[2]`, 2},
		{`let two = "two"; { "one": 10, two: 2 }["thr" + "ee"]`, nil},
		{`{ 1: 10, 2: 2 }[3]`, nil},
		{`{ 1: 10, 2: 2 }[-1]`, nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case nil:
			testNilObject(t, evaluated)
		}
	}
}
