package openai

import "strings"

// 回答整形関数
func AnswerCleaner(text string) string {

	//回答の最初と最後にある「」を取り除く
	text = strings.TrimPrefix(text, "「")
	text = strings.TrimSuffix(text, "」")

	// prefixesに含まれる言葉で始まる場合、それを除去する
	prefixes := []string{"なぜならば、", "なぜならば", "なぜなら、", "なぜなら"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(text, prefix) {
			return strings.TrimPrefix(text, prefix)
		}
	}

	// charsToRemove_answerに含まれる文字を除去
	charsToRemove_answer := []string{"「", "」", "！", "。", "\"", "“", "”"}
	for _, char := range charsToRemove_answer {
		text = strings.ReplaceAll(text, char, "")
	}

	// :がある場合、:までの文字を除去
	if idx := strings.LastIndex(text, ":"); idx != -1 {
		return text[idx+1:]
	}

	return text
}
