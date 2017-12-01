package main

import (
	"fmt"
	"math/rand"
	"time"
	"regexp"
	"strings"
)

// Adapted from: https://golang.org/pkg/math/rand/


func ElizaResponse(inputStr string) string {
	

	input := inputStr
	rand.Seed(time.Now().UTC().UnixNano())
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?"
	}
	re := regexp.MustCompile("(?i)" + `(?i)i\'?(?:\s?am|m)([^.?!]*)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}
	answers := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
		//"How do you know that you are not sure that I understand the effect my questions are having on you.",
	}
	response := "User input :" + input + " \nOutput :" + answers[rand.Intn(len(answers))]
	return response
}

func Reflect(inputStr string) string {
	// Split the input on word boundaries.
	input := inputStr
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(input, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `me`},
		{`my`, `your`},
		{`your`, `my`},
		
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}

func main() {
	userInput := []string{
				/*"People say I look like both my mother and father.",
				"Father was a teacher.",
				"I was my father’s favourite.",
				"I'm looking forward to the weekend.",
				"My grandfather was French!",
				"I am happy.",
				"I am not happy with your responses.",*/
				"I am not sure that you understand the effect that your questions are having",
				//"I am supposed to just take what you're saying at face value?",

			}
		rand.Seed(time.Now().UTC().UnixNano())
		
		word := ElizaResponse(userInput[rand.Intn(len(userInput))])
		
		fmt.Print(word)

	/*rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponse("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponse("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	fmt.Println()

	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println()

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()*/
}