# jan_api_proxy

## 概要

- JANコードから商品名を取得するAPI
    - Yahoo Shopping APIを利用
    - Gemini APIを利用して商品名のみを抽出
    - Redisをキャッシュとして利用
    - Clean Architectureで実装

## 環境変数

- `YAHOO_APP_ID`: Yahoo Shopping APIのアプリケーションID
- `GEMINI_API_KEY`: Gemini APIのAPIキー

`.env.example`をコピーして`.env`を作成し，環境変数を設定してください

## Tech Stack

- Go
- Redis
- Docker
- Docker Compose
- Yahoo Shopping API
- Gemini API
- Clean Architecture

## routes

- GET `localhost:8080/api/v1/product/:jan_code`
  - JANコードから商品名を取得する
  - キャッシュがあればキャッシュを返す
  - キャッシュがなければAPIを叩いて商品名を取得し，キャッシュに保存して返す
- GET `localhost:8080/api/v1/products`
    - キャッシュに保存されている商品名を全て取得する
- DELETE `localhost:8080/api/v1/product/:jan_code`
    - キャッシュに保存されている商品名を削除する
- POST `localhost:8080/api/v1/product`
    - input: `{"jan_code": "1234567890123", "product_name": "product_name"}`
    - キャッシュに商品名を保存する

## reference

- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [ローカルRedisをDocker Composeを使って起動する](https://zenn.dev/ring_belle/articles/docker-compose-redis)
- [GoでRedisを使ったときの基本の備忘録](https://qiita.com/tsukasaI/items/8f043af2db69c41f9724)
- [Go言語でYahoo!ショッピングの商品検索を行う。](https://qiita.com/takishita2nd/items/8358b83ef2653ad371d3)
- [（Go）Google Gemini APIを試す](https://zenn.dev/moutend/articles/aed9763635b32d)