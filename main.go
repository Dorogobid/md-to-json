package main

import (
	"bufio"
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	fileName := os.Args[1]
	if fileName == "" {
		panic("Enter file name")
	}

	lines, err := readLines(fileName)
	if err != nil {
		panic("cannot read file")
	}

	questions := parseLines(lines)

	b, err := json.Marshal(questions)
    if err != nil {
        log.Println(err)
        return
    }
	err = os.WriteFile(fileName[0:len(fileName)-3] + ".json", b, 0644)
	if err != nil {
        log.Println(err)
        return
    }
	log.Println("File converted!")
}

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func parseLines(lines []string) []Question {
	var questions []Question
	var answers []Answer

	var questionText string
	var codeType string
	var codeBlockTmp string
	var codeBlock string

	isCodeBlock := false
	isAnswerBlock := false

	for pos, line := range lines {
		if line != "" {
			if len(line) > 3 && !isCodeBlock {
				if line[0:4] == "####" {
					for pos, char := range line {
						if char == '.' {
							questionText = line[pos+2:]
						}
					}
				} 
			}
			if len(line) >= 3 {
				if len(line) > 3 && line[0:3] == "```" {
					codeType = line[3:]
					isCodeBlock = true
				} else if line == "```" {
					isCodeBlock = false
					codeBlock = codeBlockTmp
					codeBlockTmp = ""
				}
			}
			if isCodeBlock {
				if line[0] != '`' {
					codeBlockTmp = codeBlockTmp + line + "\n"
				}
				
			}
			if len(line) > 3 {
				if line[0:3] == "- [" {
					isTrue := false
					isAnswerBlock = true
					if line[3:4] == "x" {
						isTrue = true
					}
					answers = append(answers, Answer{AnswerText: line[6:], IsTrue: isTrue})
				}
			}
		}
		if (isAnswerBlock && line == "") || pos == len(lines)-1 {
			isAnswerBlock = false
			questions = append(questions, Question{QuestionText: questionText, CodeBlock: codeBlock, CodeType: codeType, Answers: answers, CategoryID: 0})
			answers = []Answer{}
			questionText = ""
			codeBlock = ""
			codeType = ""
		}
	}
	return questions
}