package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.AlbumDetail, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumDetail, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
}

type albumService struct {
	albumRepository repository.AlbumRepository
	singerRepository repository.SingerRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository,singerRepository repository.SingerRepository) *albumService {
	return &albumService{albumRepository: albumRepository,singerRepository: singerRepository}
}

func (a *albumService) GetAlbumListService(ctx context.Context) ([]*model.AlbumDetail, error) {
	albums, err := a.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	singers, err := a.singerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	albumDetails := make([]*model.AlbumDetail, 0, len(albums))
	for _, album := range albums {
		albumDetails = append(albumDetails, model.NewAlbumDetail(album.ID,album.Title,singers[album.SingerID - 1]))
	}

	return albumDetails, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumDetail, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	singer, err := s.singerRepository.Get(ctx, album.SingerID)
	if err != nil {
		return nil, err
	}
	albumDetail := model.NewAlbumDetail(album.ID,album.Title,singer)
	return albumDetail, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}