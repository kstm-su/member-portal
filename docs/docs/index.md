title: Home

# Member portalとは
kstmメンバーであることを確認し、また事務処理を簡潔化するためのポータルサイトです。

## APIエンドポイントについて

### APIドキュメント確認方法



### Mockサーバーの立て方
#### Mockサーバーとは

APIのエンドポイントを実際に立てずに、APIのレスポンスをモックで返すサーバーのことです。

#### 立て方
1. `swagger/README.md` の "Getting started" の手順を行う
2. "Mock API" の手順を `documentation.yml` と同じディレクトリにて行う
3. `http://localhost:4010`にモックサーバーが立つ

## 依存ソフトウェア

API Docmentation: [redocly cli](https://redocly.com/docs/cli) <br>
採用理由 : openapi形式のドキュメントを生成および表示するため。 <br>
利用場面: apiのドキュメント生成および表示

Container runtime: [Docker](https://www.docker.com/) <br>
採用理由: 事実上のデファクトスタンダードであるため。

