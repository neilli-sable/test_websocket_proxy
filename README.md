# test_websocket_proxy

## これは何？

- Golang の httputil.ReverseProxy で WebSocket がプロキシーできるか試したもの
- 結果: できた

## 使い方

### goのエコーサーバーでテストする
- cd 01-echoserver; go run main.go
- cd 02-reverseproxy; go run main.go
- wscat -c ws://localhost:8080/echo -o localhost:8080

### Socket.ioのチャットサーバーでテストする
- cd 03-socket.io; node index.js
- cd 02-reverseproxy; go run main.go
- ブラウザからlocalhost:12345を開いてテスト

## 参考

- net/http/httputil: add WebSocket support to ReverseProxy #26937

  - https://github.com/golang/go/issues/26937

- golang.org/x/net/websocket の使い方と kuiperbelt
  - https://qiita.com/mackee_w/items/b6036d132888caaf5f98
- websockets/wscat
  - https://github.com/websockets/wscat

- Socket.IO
  - https://socket.io/