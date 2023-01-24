package model

type AlbumDetail struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	Singer 	 *Singer `json:"singer"`
}

func NewAlbumDetail(id AlbumID,title string,singer *Singer) *AlbumDetail{
	return &AlbumDetail{ID:id,Title:title,Singer:singer}
}