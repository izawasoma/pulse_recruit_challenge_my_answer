# サーバーエンジニア向け 2024新卒採用事前課題

## 課題3
アルバムを管理するAPIを新規作成しましょう。

### 3-1
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]
```

### 3-2
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer_id":1}
```

### 3-3
アルバムを追加するAPI
```
curl -X POST -d '{"id":10,"title":"Chris 1st","singer_id":3}' http://localhost:8888/albums

# このようなレスポンスを期待しています
{"id":10,"title":"Chris 1st","singer_id":3}

# そして、アルバムを取得するAPIでは、追加したものが存在するように
curl http://localhost:8888/albums/10
```

### 3-4
アルバムを削除するAPI
```
curl -X DELETE http://localhost:8888/albums/1
```

## 課題4
アルバムを取得するAPIでは、歌手の情報も付加するように改修しましょう。

### 4-1
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 4-2
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

## 課題5
歌手とそのアルバムを管理するという点で、現状のAPIの改善点を検討し思いつく限り書き出してください。

- singer情報を削除した後、albumがヒットしないのでnot foundになってしまう
- curl http://localhost:8888/albums/aaa のエラーメッセージが404 page not foundになっている

## 学習記録

- [スライスについて](https://github.com/izawasoma/pulse_recruit_challenge_my_answer/blob/main/note/slice.md "スライスについて")
- [マップについて](https://github.com/izawasoma/pulse_recruit_challenge_my_answer/blob/main/note/map.md "マップについて")
- [繰り返し制御について](https://github.com/izawasoma/pulse_recruit_challenge_my_answer/blob/main/note/roop.md "繰り返し制御について")
- [構造体について](https://github.com/izawasoma/pulse_recruit_challenge_my_answer/blob/main/note/struct.md "構造体について")
- [その他について](https://github.com/izawasoma/pulse_recruit_challenge_my_answer/blob/main/note/other.md "その他について")
- [デフォルトパッケージについて](https://github.com/izawasoma/pulse_recruit_challenge_my_answer/blob/main/note/default_package.md "デフォルトパッケージについて")