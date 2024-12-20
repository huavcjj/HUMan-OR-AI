package openai

import "strings"

//charsToRemoveに含まれる文字をお題(topic)から取り除く（例 「宇宙で一番「困ること」は何？」　→　宇宙で一番困ることは何？）
func RemoveChars_topic(text string) string {
	charsToRemove := []string{"「", "」"}
	for _, char := range charsToRemove {
		text = strings.ReplaceAll(text, char, "")
	}
	return text
}

//お題の最初と最後にある「」を取り除く（例 「宇宙で一番「困ること」は何？」　→　宇宙で一番「困ること」は何？）
func RemoveKakko_topic(text string) string {

	text = strings.TrimPrefix(text, "「")
	text = strings.TrimSuffix(text, "」")
	return text
}
