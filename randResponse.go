package main

import (
	"fmt"
	"math/rand"
	"time"
)

//used

var ques = []string{
	"People say I look like both my mother and father.",
	"Father was a teacher.",
	"I was my father’s favourite.",
	"I'm looking forward to the weekend.",
	"My grandfather was French!",
}

// https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
func ElizaResponse() string {
	/*if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?" //checks for the word "father" and gives this response
	}

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?") // checks for "I am" and gives this response
	}

	re2 := regexp.MustCompile(`(?i)I'm ([^.?!]*)[.?!]?`)
	if matched := re2.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?") // checks for "I'm" and gives this response
	}*/
	answers := []string{ // randon answers given if there are no matches
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	rand.Seed(time.Now().UTC().UnixNano())

	return answers[rand.Intn(len(answers))]
}

func main() {
	/*var input string // variables used
	fmt.Println("=======================")
	fmt.Println("         ELIZA         ")
	fmt.Println("=======================")
	fmt.Println("Enter the word :") // user prompt
	fmt.Scanf("%s\n", &input)
	// This tests whether a pattern matches a string.
	fmt.Println(input)
	fmt.Println(ElizaResponse(input))
	fmt.Println() // realised i didnt need user input
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println() //matches father

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println() //matches father

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println() //matches father

	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println() //matches father

	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println() // matches I am

	fmt.Println("I'm happy.")
	fmt.Println(ElizaResponse("I'm happy."))
	fmt.Println() // matches I'm

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println() // matches I am

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println() // matches I am

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println() // matches I am
	*/
	for i := 0; i < 5; i++ {

		fmt.Println("Q:" + ques[i])
		fmt.Println("Re:" + ElizaResponse())
	}
} // main
