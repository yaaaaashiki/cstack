# c-stack

Escapism var3
現実逃避の時間です ∩( ´∀｀)∩ﾄﾞｳｿﾞ (っ´∀｀)っ))ﾖﾛｼ!ｸ

Go + React + Redux

## dbconfig.yml の設置

ルートディレクトリに設置します.
git 上で ignore しているので各々よろしくお願いします.


## 依存パッケージ管理 

[glide](https://github.com/Masterminds/glide) で管理します. 
新たなパッケージを追加時は、
```
glide install
```
上記コマンドをターミナルで打ちます.


すると glide.yml が更新され、vendor 以下にパッケージが書き込まれます.



## 起動

```
make run
```



## テスト
```
make test
``
