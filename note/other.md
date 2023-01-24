# その他メモ

## JSON形式の文字列をデコードし、構造体に合わせたデータに変換する

APIなどを作成する際に、ストリーミングで送られてきたJSONデータなどをデコードして、構造体に置き換えるにはjsonパッケージを用います。<br>
手順としては以下の２ステップです。<br>

1. `json.NewDecorder(json文字列)`でEncoderを作成する
2. 手順1で作成したEncoderを用いて、構造体に変換する

<br>
簡単な例は以下の通りです。<br>
<br>

```go
package main

import (
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
    //〜略〜

    //jsonStringという変数は文字列型で、
    //json文字列が以下のように格納されている。
    //jsonString : {"name" : "hoge", "age" : 11}

	// JSONデコーダーを呼び出す
	dec := json.NewDecoder(jsonString)

	var person Person
    //JSONから構造体に変換する
	err = dec.Decode(&person)
	if err != nil {
		// errorハンドリング
	}

    fmt.Println(person) // { hoge , 11 }
}
```

<br>

`err = dec.Decode(&person)`とありますが、どんな時にエラーが返ってくるかに関して一例を出すと、例えば、以下のようなjsonデータを送った場合、エラーが起きます。

<br>

`{"name":123,"age":11}`

<br>

nameは構造体ではstringであるべきなのに対して、今回の例ではintが送られてきています。このような場合、エラーを返却します。

<br>

一方、以下のようなデータの場合は、エラーにはなりません。

<br>

`{"name":"unknownAgeMan"}`

<br>

このようなデータが送られてきた場合、不足しているデータは各型の初期値が保存されます。

<br>

`fmt.Println(person) //{unknownAgeMan,0}`

<br>

さらに、以上の処理は以下のようにより簡潔に書くことも可能です。

```go
//以下の処理を...
dec := json.NewDecoder(r.Body)
var singer Singer
err = dec.Decode(&singer)
if err != nil {
	...
}

//一行で書く...
if err := json.NewDecoder(r.Body).Decode(&singer); err != nil { ... }
```