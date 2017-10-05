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


## MySQL version

* 5.7 


## node version

* 8.5.0

## npm version

* 5.3.0

## 依存パッケージ管理 

[glide](https://github.com/Masterminds/glide) で管理します
新たなパッケージを追加時は

```
glide get github.com/yaaaaashiki/cstack
glide install
```

上記コマンドをターミナルで打ちます


すると glide.yml が更新され、vendor 以下にパッケージが書き込まれます

パッケージを最新バージョンに更新するときは

```
glide up
```

で更新します

## GOPATH の設定

```sh
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
上記のディレクトリで作業します

`go get` でインストールすると上記ディレクトリに clone されているはず


## dbconfig.yml

`dbconfig.yml` の設定は各々の MySQL の設定に依存する

```yml
datasource: root:@/cstack_dev?parseTime=true&collation=utf8_general_ci&interpolateParams=true

datasource: root:@/cstack_test?parseTime=true&collation=utf8_general_ci&interpolateParams=true
```

上記は各々修正してください


(ex: password を設定している場合)
```yml
datasource: root:password@tcp(localhost:3306)/cstack_dev?parseTime=true&collation=utf8_general_ci&interpolateParams=true
```


## client 開発初回時

```
cd client

npm install

npm run build
```

上記コマンドで、`node_modules` ディレクトリの中に必要なパッケージがインストールされ `assets/js` 配下に `bundle.js` が出力される

本アプリでは、`bundle.js` のみ読み込んでいて、すべての client 周りの js は `bundle.js` として出力される


## client 開発時

```
cd client

npm run build:watch
```

クライアントサイドのコードを watch するコマンド


## 初回 DB 構築時

```
make migrate/init

make migrate/init/test
```


## seed 入れ直し (develop)

```
make developenv 
```


## 起動

```
make run
```


## seed 入れ直し (test)

テスト実行前には必ず実行してください
```
make testenv
```


## テスト
```
make test
```
