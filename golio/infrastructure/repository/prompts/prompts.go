package prompts

import (
	_ "embed"
)

var (
	// 命令通りに記事のBodyを変更するプロンプト
	//go:embed chatgpt/change_article_body.txt
	ChangeArticleBodyChatGPT string
)
