package faker

import (
	"math/rand"
	"reflect"
	"strings"
)

type Lorem struct {
	Word       string
	Words      []string
	Sentence   string
	Sentences  string
	Paragraph  string
	Paragraphs string
}

func (l *Lorem) Word_() string {
	word := getData("Lorem", "Words")
	l.Word = sample(word)
	return l.Word
}

func (l *Lorem) Words_(params ...interface{}) []string {
	length := 3

	types := kindOf(params...)

	if len(types) >= 1 && types[0] == reflect.Int {
		length = params[0].(int)
	}

	var words []string

	for i := 0; i < length; i++ {
		words = append(words, l.Word_())
	}

	l.Words = words

	return l.Words
}

func (l *Lorem) Sentence_(params ...interface{}) string {
	wordCount := random(3, 10)

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		wordCount = params[0].(int)
	}

	words := l.Words_(wordCount)
	sentence := strings.Join(words, " ")
	sentence = strings.ToLower(sentence)
	sentence = strings.ToUpper(string(sentence[0])) + sentence[1:]
	sentence += "."

	l.Sentence = sentence

	return l.Sentence
}

func (l *Lorem) Sentences_(params ...interface{}) string {
	sentenceCount := random(2, 6)
	separator := " "

	sentences := []string{}

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		sentenceCount = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.String {
		separator = params[1].(string)
	}

	for i := 0; i < sentenceCount; i++ {
		sentence := l.Sentence_()
		sentences = append(sentences, sentence)
	}

	l.Sentences = strings.Join(sentences, separator)

	return l.Sentences
}

func (l *Lorem) Paragraph_(params ...interface{}) string {
	sentenceCount := 3

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		sentenceCount = params[0].(int)
	}

	sentenceCount += random(1, 3)
	l.Paragraph = l.Sentences_(sentenceCount)

	return l.Paragraph

}

func (l *Lorem) Paragraphs_(params ...interface{}) string {
	separator := "\n\r"
	paragraphCount := 3

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.String {
		separator = params[0].(string)
	}

	if len(types) >= 2 && types[1] == reflect.Int {
		paragraphCount = params[1].(int)
	}
	paragraphs := []string{}
	for i := 0; i < paragraphCount; i++ {
		paragraphs = append(paragraphs, l.Paragraph_())
	}
	l.Paragraphs = strings.Join(paragraphs, separator)
	return l.Paragraphs
}

func (l *Lorem) Text(params ...interface{}) string {
	//method := []string{"word", "words", "sentence", "sentences", "paragraph", "paragraphs"}
	immutable := reflect.ValueOf(l)
	numMethod := immutable.NumMethod()
	rnd := rand.Intn(numMethod)
	method := immutable.Method(rnd)
	ret := method.Call([]reflect.Value{})
	kind := ret[0].Kind()
	result := ret[0].Interface()

	var val string
	if kind == reflect.Slice {
		val = strings.Join(result.([]string), ", ")
	} else {
		val = result.(string)
	}

	return val
}
