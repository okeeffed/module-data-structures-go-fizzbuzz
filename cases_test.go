package fizzbuzz

type testCase struct {
	description string
	input       int
	expected    string
}

var testCases = []testCase{
	{
		description: "should return fizz",
		input:       3,
		expected:    "fizz",
	},
	{
		description: "should return buzz",
		input:       10,
		expected:    "buzz",
	},
	{
		description: "should return fizzbuzz",
		input:       15,
		expected:    "fizzbuzz",
	},
	{
		description: "should return number",
		input:       2,
		expected:    "2",
	},
}
