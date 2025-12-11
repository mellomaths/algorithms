package strings

type palindromeTest struct {
	input    string
	expected bool
	name     string
}

var palindromeTests = []palindromeTest{
	{"", true, "Empty string"},
	{"a", true, "Single character"},
	{"aa", true, "Two same characters"},
	{"ab", false, "Two different characters"},
	{"aba", true, "Odd length palindrome"},
	{"abba", true, "Even length palindrome"},
	{"abc121cba", true, "Palindrome with numbers"},
	{"abc", false, "Odd length non-palindrome"},
	{"abcd", false, "Even length non-palindrome"},
	{"A man a plan a canal Panama", true, "Palindrome with spaces and mixed case"},
	{"No 'x' in Nixon", true, "Palindrome with punctuation and mixed case"},
	{"Hello, World!", false, "Non-palindrome with punctuation"},
}
