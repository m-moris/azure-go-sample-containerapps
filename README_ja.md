# Go のWebサーバーを爆速で Container Apps にデプロイする方法

ここでは、Go の Web サーバーを爆速で Container Apps にデプロイする方法を紹介します。

## Azure の構成

作成するリソースは以下の通り。

- ユーザー割り当てマネージドID
- Azure コンテナレジストリ
- Azure Container App Env
- Azure Container Apps

ユーザー割り当てマネージドID は、Azure Container Apps から Azure Container Registry にアクセスするために必要ですで、ACR Pull の ロールを割り当てます。

Azure Container App Env は、Container Appsの動作環境で、その配下に Container Apps を作成します。

## Go の Web サーバー

`/` エンドポイントで、`Hello, World!` を返す 単純な Web サーバーです


## デプロイ手順

はじめに `Makefile` の `SUFFIX` を適当に修正して、リソース名がグローバルに重複しないようにします。
その後は以下の手順で進めてください。

```sh
make rg    # ソースグループを作成
make env   # Container Apps 環境とACRの作成

export ACR_NAME=<acr-name>  # ACRの名前を設定

make acr-login              # acr にログイン
make build-image            # イメージをビルド＆プッシュ
make deploy                 # Container Apps にデプロイ
```
