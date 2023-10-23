# Buf Tour

This repository is used in the Buf Tour described [here](https://docs.buf.build/tour/introduction).

The tour introduces you to the `buf` CLI and the Buf Schema Registry ([BSR](https://docs.buf.build/bsr/overview)).
Along the way, you will enforce lint standards, detect breaking changes, generate code, create a
[module](https://docs.buf.build/bsr/overview#module), manage a non-trivial dependency graph, and publish the module
to the BSR so that it can be consumed by others. The tour takes approximately 20 minutes to complete.


# メモ
- 経緯としては、protoの出力先を上手く指定できなかったので buf へ変更してみた
- Google の protobuf コンパイラの最新の代替品である [Bufを組み込み中](https://buf.build/docs/tutorials/getting-started-with-buf-cli#generate-code)、connectプラグインを利用することになった
- connectプラグインによる出力指定で期待した結果が得られた（internal/proto/hoge.protoがinternal/grpc/internal/proto/hoge.protoにならないように出来る）
- 「ブラウザーおよび gRPC 互換の HTTP API を構築するためのスリム ライブラリ」とは何ぞや
  - gRPC-GatewayはHTTPサーバーとgRPCサーバー両方を建てて、HTTPリクエストを仲介する感じ？
  - サンプルを見る限り、connectはHTTP2サーバーを建てる？　gRPCやconnectプロトコルのリクエストを受け、内部ではHTTPハンドラーで処理する感じ？
