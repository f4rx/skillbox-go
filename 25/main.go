package main

import (
	"flag"
	"fmt"
	"unicode/utf8"

	logger "github.com/f4rx/logger-zap-wrapper"
	"go.uber.org/zap"
)

var slog *zap.SugaredLogger //nolint:gochecknoglobals

func init() {
	slog = logger.NewSugaredLogger()
	slog.Sync() //nolint:errcheck
}

func main() {
	sourceString := flag.String("str", "", "Source string")
	subString := flag.String("substr", "", "Substring for searc")

	flag.Parse()
	slog.Debugw("Parsed vars",
		"sourceString", sourceString,
		"subString", subString)

	if *subString == "" || *sourceString == "" {
		slog.Warn("Please set --str and --substr")
		return
	}
	slog.Info("started")

	c := searchSubstring(*sourceString, *subString)
	slog.Info("Contains: ", c)
}

func searchSubstring(str, subString string) (contains bool) {
	firstRune, size := utf8.DecodeRuneInString(subString)
	slog.Debug(fmt.Sprintf("First rune: %c %v", firstRune, size))
	for l, r := range str {
		slog.Debug(fmt.Sprintf("%c %v", r, l))
		if r == firstRune && l+len(subString) <= len(str) {
			tmpSubString := str[l : l+len(subString)]
			slog.Debugw("tmpSubString var",
				"tmpSubString", tmpSubString,
				"SubString", subString,
			)

			if contains = compareSubstringWithRunes(tmpSubString, subString); contains {
				return
			}
		}
	}
	return
}

func compareSubstringWithRunes(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	for i, w := 0, 0; i < len(str1); i += w {
		runeValueFromStr1, width1 := utf8.DecodeRuneInString(str1[i:])
		runeValueFromStr2, _ := utf8.DecodeRuneInString(str1[i:])
		if runeValueFromStr1 != runeValueFromStr2 {
			return false
		}
		w = width1
	}
	return true
}
