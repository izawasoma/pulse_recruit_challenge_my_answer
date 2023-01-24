# スライス

## スライスの宣言と参照

```go
//宣言
var sl []int
//初期値込みで宣言
var sl2 []int = []int{100,200,300}
//暗黙的な宣言
sl3 = []int{100,200,300}

//取り出し
fmt.Println(sl3[2]) // 300
```

## make関数を用いたスライスの生成

```go
//make関数を使ってスライスを宣言
sl4 := make([]int,5)
fmt.Println(sl4) //[0,0,0,0,0]
```

## より高度な参照

```go
var slice []string = []string{"0A","1B","2C","3D","4E"}
//index2番目より前を取り出し
fmt.Println(slice[:2]) //["0A","1B"]
//index2番目以降を取り出し
fmt.Println(slice[2:]) //["2C","3D","4E"]
//index2番目以降、index4番目より前を取得
fmt.Println(slice[2:4]) //["2C","3D"]

//要素数を求める
fmt.Println(len(slice)) //5

//最後の要素を取り出す
fmt.Println(slice[len(slice) - 1]) //"4E"
```

## 要素の追加

```go
//宣言
var fruits []string = []string{"apple","banana"}

fmt.Println(fluits)
//["apple","banana"]

//一件追加
fruits = append(fluits,"lemon")
fmt.Println(fluits) 
//["apple","banana","lemon"]

//複数件追加
fruits = append(fluits,"strawberry","mango")
fmt.Println(fluits) 
//["apple","banana","lemon","strawberry","mango"]
```

## capacity（容量）
基本的にスライスには最大要素数は設定されていませんが、make関数を用いることで最大容量を設定することでメモリの消費を抑えることができます。なお、最大容量を超えた際、エラーが起こることはなく、最大容量が設定値の倍数に拡張されます。

```go
sl := make([]int,5,10)

fmt.Println(sl) //[0,0,0,0,0]
fmt.Println(len(sl)) //5
fmt.Println(cap(sl)) //10

sl = append(sl,1,2,3,4,5,6)
fmt.Println(sl) //[0,0,0,0,0,1,2,3,4,5,6]
fmt.Println(len(sl)) //11
fmt.Println(cap(sl)) //20 :メモリの範囲を倍に増やした
```