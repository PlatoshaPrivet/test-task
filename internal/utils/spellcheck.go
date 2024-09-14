package utils

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SpellResult struct {
	Errors []Error `xml:"error"`
}

type Error struct {
	Code string `xml:"code,attr"`
}

func ValidateSpelling(text string) bool {
	formattedText := strings.ReplaceAll(text, " ", "+")

	url := fmt.Sprintf("https://speller.yandex.net/services/spellservice/checkText?text=%s", formattedText)

	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var result SpellResult
	if err := xml.Unmarshal(body, &result); err != nil {
		return false
	}

	for _, err := range result.Errors {
		if err.Code != "0" {
			return false
		}
	}

	return true
}
