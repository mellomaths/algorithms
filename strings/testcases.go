package strings

type stringCheckTest struct {
	input    string
	expected bool
	name     string
}

type stringComparisonTest struct {
	s1       string
	s2       string
	expected bool
	name     string
}

type rotateStringNTimesTest struct {
	s        string
	times    int
	expected string
	name     string
}

var palindromeTests = []stringCheckTest{
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

var nonTrivialRotationTests = []stringComparisonTest{
	{"a", "a", false, "Equal string"},
	{"a", "b", false, "Different characters"},
	{"abc", "bca", true, "Trivial rotation"},
	{"abc", "cab", true, "Trivial rotation"},
	{"abc", "abc", false, "Equal string"},
	{"abc", "bca", true, "Trivial rotation"},
	{"abc", "cab", true, "Trivial rotation"},
	{"abc", "abc", false, "Equal string"},
	{"abc", "bca", true, "Trivial rotation"},
	{"abc", "cab", true, "Trivial rotation"},
	{"abcde", "cdeab", true, "Trivial rotation"},
	{"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "uvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrst", true, "Long string rotation"},
}

var rotateStringNTimesTests = []rotateStringNTimesTest{
	{"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", 20, "uvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrst", "Long string rotation"},
}
