# jan_api_proxy

## 概要

- JANコードから商品名を取得するAPI
    - Yahoo Shopping APIを利用
    - Gemini APIを利用して商品名のみを抽出
    - Redisをキャッシュとして利用
    - Clean Architectureで実装

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