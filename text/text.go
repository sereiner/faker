package text

import (
	"strings"
	"sync"

	"github.com/sereiner/faker/base"
)

//Markov .
type Markov struct {
	cache    map[string][]string
	words    []string
	wordSize int
	index    int
	wordChan chan []string
	sync.Mutex
}

var markov *Markov

func init() {
	markov = &Markov{
		cache:    map[string][]string{},
		words:    strings.Split(strings.Replace(baseText, "\n", " ", -1), " "),
		wordSize: len(strings.Split(strings.Replace(baseText, "\n", " ", -1), " ")),
		index:    0,
		wordChan: make(chan []string, 1e3),
	}

	go markov.database()
	markov.triples()

}

// RealText 通过马尔可夫链算法生成文本字符串
func RealText(size ...int) string {
	textSize := 10
	if len(size) > 0 && size[0] != 0 {
		textSize = size[0]
	}
	seed := base.RandInt(0, markov.wordSize-3)
	seedWord, nextWord := markov.words[seed], markov.words[seed+1]
	w1, w2 := seedWord, nextWord
	genWords := []string{}
	for i := 0; i < textSize; i++ {
		markov.Lock()
		genWords = append(genWords, w1)
		w1, w2 = w2, base.RandomElement(markov.cache[w1+w2])
		genWords = append(genWords, w2)
		markov.Unlock()
	}

	return strings.Join(genWords, " ")
}

func (m *Markov) triples() {
	if len(m.words) < 3 {
		return
	}
	for i := 0; i < len(m.words)-2; i++ {
		m.wordChan <- []string{m.words[i], m.words[i+1], m.words[i+2]}
	}
}

func (m *Markov) database() {
	for {
		select {
		case v := <-m.wordChan:
			m.Lock()
			key := v[0] + v[1]
			if _, ok := m.cache[key]; ok {
				m.cache[key] = append(m.cache[key], v[2])
			} else {
				m.cache[key] = []string{v[2]}
			}
			m.Unlock()
		}
	}
}
