#!/bin/bash

# PostgreSQL接続情報
DB_HOST="ep-sweet-poetry-a13hzd7i.ap-southeast-1.aws.neon.tech"
DB_NAME="golio"
DB_USER="golio"
DB_PASSWORD="XX password"

# ローカルディレクトリのパス
LOCAL_DIR="exportdata"

# 環境変数にパスワードを設定
export PGPASSWORD=$DB_PASSWORD

# ディレクトリ内のすべてのCSVファイルをインポート
    psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\copy article_tags FROM 'exportdata/golio/public.article_tags/1/part-00000-12802492-58cf-47f3-948f-10714109a7cf-c000.gz.csv' CSV HEADER"
    psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\copy article_summaries FROM 'exportdata/golio/public.article_summaries/1/part-00000-72d0113c-bc1e-4c2e-8c0b-ab86f60aa9aa-c000.gz.csv' CSV HEADER"
    psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\copy media FROM 'exportdata/golio/public.media/1/part-00000-88057521-77c9-4809-84a2-12b4869b1d5d-c000.gz.csv' CSV HEADER"
    psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\copy schema_migrations FROM 'exportdata/golio/public.schema_migrations/1/part-00000-414aaa1a-2e9d-41e0-beff-5f71bc963db9-c000.gz.csv' CSV HEADER"
    psql -h $DB_HOST -U $DB_USER -d $DB_NAME -c "\copy article_bodies FROM 'exportdata/golio/public.article_bodies/1/part-00000-cb23a169-fc4f-4865-a81a-9546eb1e36a3-c000.gz.csv' CSV HEADER"

# 環境変数のパスワードを削除
unset PGPASSWORD
