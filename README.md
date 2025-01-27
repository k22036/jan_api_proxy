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

## reference

- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [ローカルRedisをDocker Composeを使って起動する](https://zenn.dev/ring_belle/articles/docker-compose-redis)
- [GoでRedisを使ったときの基本の備忘録](https://qiita.com/tsukasaI/items/8f043af2db69c41f9724)
- [Go言語でYahoo!ショッピングの商品検索を行う。](https://qiita.com/takishita2nd/items/8358b83ef2653ad371d3)
- [（Go）Google Gemini APIを試す](https://zenn.dev/moutend/articles/aed9763635b32d)