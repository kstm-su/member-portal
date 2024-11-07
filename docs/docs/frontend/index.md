# フロントエンドについて

## フロントエンドの概要

Next.jsを使用し、backendと連携して動作するフロントエンドです。
メンバーの管理および、メンバーの情報を表示するためのページを提供します。

## 起動方法

### Dockerでの起動方法
#### Localでビルドする場合

1. Dockerfileのあるディレクトリに移動
2. `docker build -t <image_name> .` を実行（image_nameは自由に設定）
3. `docker run -p 3000:3000 -d <image_name>` を実行
4. `http://localhost:3000` にアクセス

#### GitHub container registryからpullする場合
1. `docker pull ghcr.io/kstm-su/member-portal/frontend:latest` を実行
2. `docker run -p 3000:3000 -d ghcr.io/kstm-su/member-portal/frontend:latest` を実行
3. `http://localhost:3000` にアクセス

### Nodeでローカルで起動する場合
1. `cd member-portal-frontend` でディレクトリに移動
2. `npm run dev` を実行
3. `http://localhost:3000` にアクセス

