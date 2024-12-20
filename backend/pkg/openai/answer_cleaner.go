package openai

import "strings"

// prefixesに含まれる言葉で始まる場合、それを除去する
func RemovePrefixes_answer(text string) string {
	prefixes := []string{"なぜならば、", "なぜならば", "なぜなら、", "なぜなら"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(text, prefix) {
			return strings.TrimPrefix(text, prefix)
		}
	}
	return text
}

// charsToRemove_answerに含まれる文字を除去
func RemoveChars_answer(text string) string {
	charsToRemove_answer := []string{"「", "」", "！", "。", "\"", "“", "”"}
	for _, char := range charsToRemove_answer {
		text = strings.ReplaceAll(text, char, "")
	}
	return text
}

// :がある場合、:までの文字を除去
func TrimBeforeColon_answer(text string) string {
	if idx := strings.LastIndex(text, ":"); idx != -1 {
		return text[idx+1:]
	}
	return text
}

//答えの最初と最後にある「」を取り除く
func RemoveKakko_answer(text string) string {
	text = strings.TrimPrefix(text, "「")
	text = strings.TrimSuffix(text, "」")
	return text
}
