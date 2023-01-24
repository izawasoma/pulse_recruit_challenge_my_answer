# マップ 

他言語で連想配列やハッシュに当たり、keyとvalueから成るデータ群のことを指す

<br>

## 宣言

<br>

これは、最も基本的なmapの宣言です。
```go
var items = map[string]int{
    "やくそう" : 80,
    "どくけしそう" : 50,
    "おなべのふた" : 160,
}

fmt.Println(items) 
// map[やくそう:80 どくけしそう:50 おなべのふた:160]
```

<br>

一方で、make関数でもmapを作成することは可能です。<br>
追加や取得、削除に関しては以下の通りです。<br>
取得できなかった際にエラーを受け取ったり、削除はdelete関数を用いることなどはポイントです。
```go
//空のmapを作成
items := make(map[string]int)

//情報を追加
items["やくそう"] = 80
items["どくけしそう"] = 50

//登録していないデータを取り出す

item, ok := items["おなべのふた"] //okにエラー情報が入る

if !ok {
    fmt.Println("登録されていないアイテムです。")
}
else{
    fmt.Println(item)
}

//BEFORE
fmt.Println(items) // map[やくそう:80 どくけしそう:50]
//要素を削除する
delete(items,"どくけしそう")
//AFTER
fmt.Println(items) // map[やくそう:80]

```

<br>

〈補足〉mapも範囲式forを用いることが可能です。

```go
var items = map[string]int{
    "やくそう" : 80,
    "どくけしそう" : 50,
    "おなべのふた" : 160,
}

for itemName, price range := items {
    fmt.Println(itemName, price)
}
```