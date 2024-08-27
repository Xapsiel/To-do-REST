package speller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	host string
	url  string
	lang string
}

func NewSpeller() *YaSpeller {
	return &YaSpeller{host: "https://speller.yandex.net", url: "/services/spellservice.json/checkText", lang: "en"}
}

//https://speller.yandex.net/services/spellservice.json/checkText?text=синхрафазатрон+в+дубне

func (ys *YaSpeller) CheckText(text string, lang string) (string, error) {
	newtext := ""
	text = strings.ReplaceAll(text, " ", "+")
	client := http.Client{}
	url := fmt.Sprintf("%s%s", ys.host, ys.url)
	resp, err := client.Post(url+"?text="+text, "text/plain", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var spell []speller

	json.Unmarshal(body, &spell)
	// Выводим ответ
	for _, word := range spell {
		newtext += " " + word.S[0]
	}

	newtext = newtext[1:]

	return newtext, nil

}
