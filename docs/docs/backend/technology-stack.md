# 技術スタック

## 採用ライブラリ

Web framework: [echo](https://echo.labstack.com/) <br>
採用理由: openapi-generatorを利用することを想定し、親和性の高いechoを採用。

ORM: [gorm](https://gorm.io/) <br>
採用理由: GitHubのスター数が多く、開発者が多いため。

Command line tool: [cobra](https://cobra.dev/) <br>
採用理由: サブコマンドの実装が容易であるため また、k8sでも採用されているため、

Config: [viper](https://github.com/spf13/viper) <br>
採用理由: cli toolとして採用した、cobraとの親和性が高いため、

## 依存ソフトウェア

Language: [Go](https://golang.org/) <br>
採用理由: 本プロジェクトの言語として採用。

Linter & Formatter: [golangci-lint](https://golangci-lint.run/) <br>
採用理由: GitHubのスター数が多く、開発者が多いため。また、複数のlinterを統合しているため。

Database: [PostgreSQL](https://www.postgresql.org/) [Optional]<br>
採用理由: 慣習的に利用されるため。 また、ORMのgormが対応しているため。 (ConfigでSQLite3を利用することも可能)

Secret Manager(検討中): [Hashicorp Vault](https://www.vaultproject.io/)　<br>
採用理由: シークレット管理のため<br>
利用場面: パスワードハッシュ化時のpepperの管理および、OAuth2クライアントのシークレット管理の際に利用のため。<br>
その他の候補: [Cyber Ark conjur](https://www.conjur.org/), [Infisical](https://infisical.com/)

Object storage: [MinIO](https://min.io/) [Optional]<br>
採用理由: S3互換のため。







