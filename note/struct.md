# 構造体
他言語で言うところの「クラス」の役割を持つもの。



## 構造体の基本
以下は最もシンプルな構造体の１つです。

```go
//構造体の宣言
type User struct {
    name string `ユーザーの名前`
    age int `ユーザーの年齢`
}

func main() {
    //structで宣言した型を用いて変数を宣言します
    var user1 User
    //構造体が持つ変数や関数は宣言と同時に利用可能になります。
    user1.name = "一郎"
    user2.age = 22
    fmt.Println(user1) // { 一郎 , 22 }
}
```

<br>

## 構造体のメソッド
<br>

次は、メソッドに関してです。基本的には<br>
``func (別名 レシーバー) メソッド名 (型 引数) 返り値の型 {}``<br>
で定義できます。<br>
<br>
ここで注意したいのはレシーバーにおける値の渡し方です。<br>
構造体の値をメソッドを用いて変更するためにはポインタ型で渡す必要があります。<br>


```go
//構造体の宣言
type User struct {
    name string `ユーザーの名前`
    age int `ユーザーの年齢`
}

//セッター
//更新を行うのでポインタ型でレシーバーに参照渡しを行う
func (u *User) SetName(string name) {　
    u.name = name
}

//自己紹介を表示するメソッド
func (u User) SayMyProfile () {
    fmt.Sprintf("私は%vです。%v歳です。", u.name , u.age)
}

func main() {
    //暗示的に構造体から変数を宣言するには以下のように行います。
    user1 := User{name:"七海", age:23}

    user1.SayMyProfile() //私は七海です。23歳です。

    user1.SetName("若菜")
    user1.SayMyProfile() //私は若菜です。23歳です。
}
```

<br>

## 構造体のコンストラクタ
<br>

```go
//構造体の宣言
type User struct {
    name string `ユーザーの名前`
    age int `ユーザーの年齢`
}

//コンストラクタ
func NewUser(name string, age int) *User {
    return &User{name: name, age: age}
}


func main() {
    //暗示的に構造体から変数を宣言するには以下のように行います。
    user1 := NewUser("篤志",18)
    fmt.Printf(user1) // { 篤志 , 18 }
}
```
