// Go offers built-in support for [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// Here are some examples of  common regexp-related tasks
// in Go.

package main

import (
	"fmt" // imports
	"math/rand"
	"regexp" //used
)

func ElizaResponse(input string) string {
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?"
	}

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}

	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	return answers[rand.Intn(len(answers))]
}

func main() {
	var input string // variables used
	fmt.Println("=======================")
	fmt.Println("         ELIZA         ")
	fmt.Println("=======================")
	fmt.Println("Enter the word :") // user prompt
	fmt.Scanf("%s\n", &input)
	// This tests whether a pattern matches a string.
	fmt.Println(ElizaResponse(input))
	fmt.Println()

} // main
