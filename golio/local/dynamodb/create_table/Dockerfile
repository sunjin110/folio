FROM amazon/aws-cli:2.15.19

COPY create_table.sh create_table.sh

# 特権ユーザーとして実行しないようにする
RUN yum -y install shadow-utils
RUN groupadd -r nonroot && useradd -r -g nonroot nonroot
USER nonroot

ENTRYPOINT sh create_table.sh
