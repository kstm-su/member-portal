# member-portal
kstmメンバーであることを確認し、また事務処理を簡潔化するためのポータルサイト

[Wiki](https://kstm-su.github.io/member-portal/)

## Frontend
### Docker起動方法
1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3000:3000 -d <image_name>` を実行
4. `http://localhost:3000` にアクセス

※ ローカルで起動する場合は、`npm run dev` を実行

## Backend

### Docker起動方法
1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3001:8080 -d <image_name>` を実行
4. `http://localhost:3001` にアクセス


## Common

### APIドキュメント確認方法
1. `cd swagger`で`swagger`ディレクトリに移動します
2. `npx @redocly/cli preview-docs documentation.yml`を実行します
3. `http://localhost:8080/`にアクセス

### Mockサーバーの立て方
1. `swagger/README.md` の "Getting started" の手順を行う
2. "Mock API" の手順を `documentation.yml` と同じディレクトリにて行う
3. `http://localhost:4010`にモックサーバーが立つ



