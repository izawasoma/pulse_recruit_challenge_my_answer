package memorydb

import (
	"context"
	"errors"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type singerRepository struct {
	//データの排他処理に用いる（ロック）
	sync.RWMutex
	// キーが SingerID、値が model.Singer のマップ
	singerMap map[model.SingerID]*model.Singer
}

//作成する構造体がその型を満たしているかをチェックする
//https://sourjp.github.io/posts/go-interface/
var _ repository.SingerRepository = (*singerRepository)(nil)

//コンストラクタ
//singerRepository.singerMapに初期値を設定している
func NewSingerRepository() *singerRepository {
	var initMap = map[model.SingerID]*model.Singer{
		1: {ID: 1, Name: "Alice"},
		2: {ID: 2, Name: "Bella"},
		3: {ID: 3, Name: "Chris"},
		4: {ID: 4, Name: "Daisy"},
		5: {ID: 5, Name: "Ellen"},
	}

	return &singerRepository{
		singerMap: initMap,
	}
}

//func レシーバー メソッド名(型 引数) 返り値の型 { ... }
func (r *singerRepository) GetAll(ctx context.Context) ([]*model.Singer, error) {
	//読み込みロック
	r.RLock()
	//読み込みロック解除
	//defer…この関数が実行終了した際に実行する
	defer r.RUnlock()

	//スライスを作成
	singers := make([]*model.Singer, 0, len(r.singerMap))
	for _, s := range r.singerMap {
		singers = append(singers, s)
	}
	return singers, nil
}

func (r *singerRepository) Get(ctx context.Context, id model.SingerID) (*model.Singer, error) {
	r.RLock()
	defer r.RUnlock()

	//該当のidが存在する場合、singerには該当データが入り、okにはnilが入る
	//一方で該当データが存在しない場合、okにはfalseが入る。
	singer, ok := r.singerMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return singer, nil
}

func (r *singerRepository) Add(ctx context.Context, singer *model.Singer) error {
	//書き込みロック
	r.Lock()
	r.singerMap[singer.ID] = singer
	//書き込みロックを解除
	r.Unlock()
	return nil
}

func (r *singerRepository) Delete(ctx context.Context, id model.SingerID) error {
	r.Lock()
	//mapからkeyがidに相当するものを削除する
	delete(r.singerMap, id)
	r.Unlock()
	return nil
}
