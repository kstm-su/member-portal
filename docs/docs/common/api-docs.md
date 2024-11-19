# APIドキュメントについて

## 目的
OpenAPIを利用してAPIドキュメントを作成することで、フロントエンドとバックエンドの実装を隠蔽し、APIの定義をインターフェースとして利用することにより、開発効率を向上させる。

## 仕様
- OpenAPI 3.0.0
- ファイル形式: YAML
- ファイル名: documentation.yaml
- ファイルの配置場所: /swagger

## ファイル構成
`/components`: OpenAPIのコンポーネントを定義する <br>
`/paths`: APIのエンドポイントを定義する <br>
`documentation.yaml`: OpenAPIの定義を記述するファイル

## APIドキュメントの確認方法
### Github Pagesで確認する方法
masterブランチのAPIドキュメントについては[こちら](/member-portal/redoc/index.html)で確認できます。

### ローカルで確認する方法
1. `cd swagger`で`swagger`ディレクトリに移動します
2. `npx @redocly/cli preview-docs documentation.yml`を実行します
3. `http://localhost:8080/`にアクセス