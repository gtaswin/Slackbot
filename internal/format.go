package bot

import (
	"log"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/go-ini/ini"
)

//Format for parsing the given text
func Format(text string, cfg *ini.File) string {
	var out string
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Print(err)
	}

	text = reg.ReplaceAllString(text, "")
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	array := strings.Fields(text)
	if array[0] == cfg.Section("main").Key("command").String() {
		out = Command(text, cfg)
	} else {
		out = cfg.Section("chat").Key(text).String()
	}
	if out == "" {
		out = cfg.Section("main").Key("noentry").String()
	} else if utf8.RuneCountInString(out) >= 3000 {
		out = "Out of range"
	}
	return out
}
