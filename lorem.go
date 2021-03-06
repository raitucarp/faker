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

func (l *Lorem) Wordʹ() string {
	word := getData("Lorem", "Words")
	l.Word = sample(word)
	return l.Word
}

func (l *Lorem) Wordsʹ(params ...interface{}) []string {
	length := 3

	types := kindOf(params...)

	if len(types) >= 1 && types[0] == reflect.Int {
		length = params[0].(int)
	}

	var words []string

	for i := 0; i < length; i++ {
		words = append(words, l.Wordʹ())
	}

	l.Words = words

	return l.Words
}

func (l *Lorem) Sentenceʹ(params ...interface{}) string {
	wordCount := random(3, 10)

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		wordCount = params[0].(int)
	}

	words := l.Wordsʹ(wordCount)
	sentence := strings.Join(words, " ")
	sentence = strings.ToLower(sentence)
	sentence = strings.ToUpper(string(sentence[0])) + sentence[1:]
	sentence += "."

	l.Sentence = sentence

	return l.Sentence
}

func (l *Lorem) Sentencesʹ(params ...interface{}) string {
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
		sentence := l.Sentenceʹ()
		sentences = append(sentences, sentence)
	}

	l.Sentences = strings.Join(sentences, separator)

	return l.Sentences
}

func (l *Lorem) Paragraphʹ(params ...interface{}) string {
	sentenceCount := 3

	types := kindOf(params...)
	if len(types) >= 1 && types[0] == reflect.Int {
		sentenceCount = params[0].(int)
	}

	sentenceCount += random(1, 3)
	l.Paragraph = l.Sentencesʹ(sentenceCount)

	return l.Paragraph

}

func (l *Lorem) Paragraphsʹ(params ...interface{}) string {
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
		paragraphs = append(paragraphs, l.Paragraphʹ())
	}
	l.Paragraphs = strings.Join(paragraphs, separator)
	return l.Paragraphs
}

func (l *Lorem) Text(params ...interface{}) (val string) {
	//method := []string{"word", "words", "sentence", "sentences", "paragraph", "paragraphs"}
	rnd := rand.Intn(6)
	switch rnd {
	case 0:
		l.Wordʹ()
	case 1:
		l.Wordsʹ()
	case 2:
		l.Sentenceʹ()
	case 3:
		l.Sentencesʹ()
	case 4:
		l.Paragraphʹ()
	case 5:
		l.Paragraphsʹ()
	}

	return
}

// Fake Generate random data for all field
func (l *Lorem) Fake() {
	l.Wordʹ()
	l.Wordsʹ()
	l.Sentenceʹ()
	l.Sentencesʹ()
	l.Paragraphʹ()
	l.Paragraphsʹ()
}

// ToJSON Encode name and its fields to JSON.
func (l *Lorem) ToJSON() (string, error) {
	return ToJSON(l)
}

// ToXML Encode name and its fields to XML.
func (l *Lorem) ToXML() (string, error) {
	return ToXML(l)
}
