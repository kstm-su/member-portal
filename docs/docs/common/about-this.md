# 当ドキュメントについて

このドキュメントは、[MkDocs-material](https://squidfunk.github.io/mkdocs-material/)を使用して作成されました。

## 目的

今後の開発の参入障壁を低くするために、どのような技術スタックを利用し、何を目的としているかを示すため。

## 編集するには
`docs/docs`にドキュメントの本体が配置されているため、そちらを編集してください。<br>
また、新たにファイルを追加する場合は、`mkdocs.yml`にある`nav`に追加してください。

## 表示方法

### ローカルでの表示

以下のコマンドを実行することで、ローカルでの表示が可能です。

#### 事前準備

```bash
pip install mkdocs-material
```

#### Linux

```bash
cd docs
./start.sh
```

#### Windows

`docs/docs`に`swagger`をコピーしたのち、以下のコマンドを実行してください。

```bash
cd docs
mkdocs serve
```

### GitHub Pagesでの表示

GitHub Pagesでの表示は、[こちら](https://kstm-su.github.io/member-portal/)からご覧いただけます。