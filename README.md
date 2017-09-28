## c-stack

Escapism var3
現実逃避の時間です ∩( ´∀｀)∩ﾄﾞｳｿﾞ (っ´∀｀)っ))ﾖﾛｼ!ｸ

Go + React + Redux


## install

```
go get -u github.com/yaaaaashiki/cstack
```


## Golang version

* 1.9 


## mysql version

* 5.7 


## 依存パッケージ管理 

[glide](https://github.com/Masterminds/glide) で管理します. 
新たなパッケージを追加時は、
```
glide install
```
上記コマンドをターミナルで打ちます.


すると glide.yml が更新され、vendor 以下にパッケージが書き込まれます.


## GOPATH の設定

```
if [ -x "`which go`" ]; then
    export GOPATH=$HOME/go
    export PATH=$PATH::$GOPATH/bin
fi
```
上記を `bashrc` or `zshrc` に記述・反映させる


## 作業ディレクトリ

```
GOPATH/src/github.com/yaaaaashiki/cstack
```
上記のディレクトリで作業します。
`go get` でインストールすればこちらに cstack が clone されているはず


## client 開発初回時

```
npm run build
```

上記コマンドで、`assets/js` 配下に `bundle.js` が出力される
本アプリでは、`bundle.js` のみ読み込んでいて、すべての client 周りの js は `bundle.js` として出力される


## client 開発時

```
npm run build:watch
```

クライアントサイドのコードを watch するコマンド

## seed 入れ直し (env: develop)

```
make developenv 
```


## 起動

```
make run
```


## seed 入れ直し (env: test)

テスト実行前には必ず実行してください.
```
make testenv
```


## テスト
```
make test
```
