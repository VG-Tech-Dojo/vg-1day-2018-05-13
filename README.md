vg-1day-2018-05-13
---

## 概要

vg-1day-2018-05-13 で使う予定のサンプルアプリです。

サーバーサイドはgolang、クライアントサイドはHTML, Vue.jsで実装されています。

## Docker 環境

インターンに必須ではありませんが、 go 1.10.2 が入ったdocker環境を用意しました。

```
$ pwd
/Users/s-sasamoto/src/github.com/VG-Tech-Dojo/vg-1day-2018-05-13
$ make docker/build

# xxx は 設定した nickname
$ make docker/deps/xxx
$ make docker/run/xxx
```

## Vagrant 環境

こちらもインターンに必須ではありませんが、 Vagrantfile も用意しました。

```
$ pwd
/Users/s-sasamoto/src/github.com/VG-Tech-Dojo/vg-1day-2018-05-13
$ vagrant up
$ vagrant ssh

$ cd go/src/github.com/VG-Tech-Dojo/vg-1day-2018-05-13
```

## 参考リンク

[gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.](https://github.com/gin-gonic/gin)


[A progressive, incrementally-adoptable JavaScript framework for building UI on the web.](https://jp.vuejs.org)

## golangの勉強に役立つリンク

[A Tour of Go](https://tour.golang.org/welcome/1)

まずはこれをやりましょう。基礎的な文法がわかります。

[How to Write Go Code - The Go Programming Language](https://golang.org/doc/code.html)

初めてのgoプロジェクトを作る際の参考になります。

[Effective Go - The Go Programming Language](https://golang.org/doc/effective_go.html)

goらしい書き方、Tipsを学べます。

[CodeReviewComments · golang/go Wiki](https://github.com/golang/go/wiki/CodeReviewComments)

goに慣れてきたらこれも読むと良いでしょう。

goコードのレビュー時によく指摘されることがまとまっています。

位置づけ的には Effective Go の補足です。

[GoDoc](https://godoc.org/)

標準ライブラリの使い方やフレームワークの使い方を見る時はgodocを見ましょう。

