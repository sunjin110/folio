#!/bin/bash

# 対象のファイル名
FILE="generate/schema/http/go/openapi/api_golio.go"

# import文に追加するパッケージ
NEW_IMPORTS=(
    "\"errors\";"
    "\"io\";"
)



# 既存のimport文のブロックを特定し、新しいパッケージを追加する
for import in "${NEW_IMPORTS[@]}"; do
    if ! grep -q $import $FILE; then
        # macOSの場合、改行を明示的に挿入する
        sed -i '' "/import (/a\\
    $import" $FILE
    fi
done

# 不要なreflectパッケージの参照を削除
sed -i '' '/"reflect"/d' "generate/schema/http/go/openapi/api_golio.go"
sed -i '' '/"reflect"/d' "generate/schema/http/go/openapi/api.go"
