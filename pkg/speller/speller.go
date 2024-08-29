package speller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"test_case/pkg/errors"
	"time"
)

var symbols = "АаБбВвГгДдЕеЁёЖжЗзИиЙйКкЛлМмНнОоПпРрСсТтУуФфХхЦцЧчШшЩщЪъЫыЬьЭэЮюЯяABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type speller struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}
type YaSpeller struct {
	host    string
	url     string
	lang    string
	timeout time.Duration
}

func NewSpeller(timeout time.Duration) *YaSpeller {

	return &YaSpeller{host: "https://speller.yandex.net", url: "/services/spellservice.json/checkText", lang: "en", timeout: timeout}
}

// https://speller.yandex.net/services/spellservice.json/checkText?text=синхрафазатрон+в+дубне
func (ys *YaSpeller) CheckText(text string, lang string) (string, error) {
	result := ""
	text = strings.ReplaceAll(text, " ", "+")
	splitText := strings.Split(text, "+")
	for _, word := range splitText {
		newtext, err := ys.CheckWord(word, lang)
		if err != nil {
			return text, err
		}
		result += " " + newtext
	}
	if len(result) > 1 {
		result = result[1:]
	}
	return result, nil

}

func (ys *YaSpeller) CheckWord(text string, lang string) (string, error) {
	client := http.Client{Timeout: ys.timeout}
	url := fmt.Sprintf("%s%s", ys.host, ys.url)
	resp, err := client.Post(url+"?text="+text, "text/plain", nil)
	if err != nil {
		return text, errors.New("Post requests in YS", err.Error(), http.StatusServiceUnavailable)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return text, errors.New("Reading request body", err.Error(), http.StatusServiceUnavailable)
	}
	var spell []speller

	err = json.Unmarshal(body, &spell)
	if err != nil {
		return text, errors.New("Unmarsling from json", err.Error(), http.StatusServiceUnavailable)
	}
	// Выводим ответ
	if len(spell) == 0 {
		return text, nil
	}

	return spell[0].S[0], nil

}
