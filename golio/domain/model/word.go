package model

type WordDetail struct {
	Word             string
	Definitions      []*WordDefinition
	Frequency        float64
	PronunciationMap PronunciationMap
}

type WordDefinition struct {
	Definition   string
	PartOfSpeech string // noun, verb...
	Synonyms     []string
	Antonyms     []string
	Examples     []string
}

type PronunciationMap map[string]string // key, noun, verb or all
