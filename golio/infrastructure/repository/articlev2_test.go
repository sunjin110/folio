package repository_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/infrastructure/chatgpt"
	"github.com/sunjin110/folio/golio/infrastructure/gcp/custom_search_api"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
)

// go test -v -count=1 -timeout 30s -run ^Test_articleV2_ChangeBodyByAI_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_articleV2_ChangeBodyByAI_Real(t *testing.T) {
	SkipConvey("Test_articleV2_ChangeBodyByAI_Real", t, func() {

		ctx := context.Background()
		client := chatgpt.NewClient("")

		customSearchClient, err := custom_search_api.NewClient(ctx, "")
		So(err, ShouldBeNil)

		googleCustomSearchRepo := repository.NewGoogleCustomSearch(customSearchClient)

		articleV2Repo := repository.NewArticleV2(context.Background(), nil, client, googleCustomSearchRepo)

		_, err = articleV2Repo.ChangeBodyByAI(ctx, &model.Article{
			Body: `
## phrase

- But it is even more difficult when it is done in what is known as "alpine style"
- Himalayas were being climbed. 
- They would then climb their way from one camp to the next
- It took them only three days to reach toe top of montain.
- It took them only three days to ~

## words
- alternative: 別の
- siege: 包囲
- dozens: 数十
- series of camps equipped
- supplies: 用品
- eventually: 最終的に
- hire: 雇う
- **set out: 出発する**
- Gasherbrum: 山の名前
- in this way: このように
- to took me {xxx days} to ~
- the first ever: 史上初めて
- Since it takes less time: 時間があまりかからないので
	- Sinceは~なのでという意味で使われることもある: ニュアンスとしてはすでに明確でわかっていること
- being caught: 捕まる
- avalanche: 雪雪崩
- there is a smaller chance of ~ : ~の可能性は低い
- expeditions: 遠征
- trace: 痕跡
- accomplishments: 成果
- 
`,
		}, "phraseの欄に日本語翻訳を追加してください、元の英文の右に日本語翻訳を付与してください")

		So(err, ShouldBeNil)

	})
}

// go test -v -count=1 -timeout 30s -run ^Test_articleV2_GenerateBodyByAI_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_articleV2_GenerateBodyByAI_Real(t *testing.T) {
	SkipConvey("Test_articleV2_GenerateBodyByAI_Real", t, func() {

		ctx := context.Background()
		client := chatgpt.NewClient("")

		customSearchClient, err := custom_search_api.NewClient(ctx, "")
		So(err, ShouldBeNil)

		googleCustomSearchRepo := repository.NewGoogleCustomSearch(customSearchClient)

		articleV2Repo := repository.NewArticleV2(context.Background(), nil, client, googleCustomSearchRepo)

		result, err := articleV2Repo.GenerateBodyByAI(ctx, "今日の天気を教えてください")
		So(err, ShouldBeNil)
		fmt.Println("result is ", result)

		// 		_, err = articleV2Repo.ChangeBodyByAI(ctx, &model.Article{
		// 			Body: `
		// ## phrase

		// - But it is even more difficult when it is done in what is known as "alpine style"
		// - Himalayas were being climbed.
		// - They would then climb their way from one camp to the next
		// - It took them only three days to reach toe top of montain.
		// - It took them only three days to ~

		// ## words
		// - alternative: 別の
		// - siege: 包囲
		// - dozens: 数十
		// - series of camps equipped
		// - supplies: 用品
		// - eventually: 最終的に
		// - hire: 雇う
		// - **set out: 出発する**
		// - Gasherbrum: 山の名前
		// - in this way: このように
		// - to took me {xxx days} to ~
		// - the first ever: 史上初めて
		// - Since it takes less time: 時間があまりかからないので
		// 	- Sinceは~なのでという意味で使われることもある: ニュアンスとしてはすでに明確でわかっていること
		// - being caught: 捕まる
		// - avalanche: 雪雪崩
		// - there is a smaller chance of ~ : ~の可能性は低い
		// - expeditions: 遠征
		// - trace: 痕跡
		// - accomplishments: 成果
		// -
		// `,
		// 		}, "phraseの欄に日本語翻訳を追加してください、元の英文の右に日本語翻訳を付与してください")

		// 		So(err, ShouldBeNil)

	})
}
