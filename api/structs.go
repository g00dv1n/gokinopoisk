package api

// FilmInfo film info struct
type FilmInfo struct {
	IsHasSimilarFilms uint       `json:"isHasSimilarFilms"`
	RatingData        RatingInfo `json:"ratingData"`
	FilmID            string     `json:"filmID"`
	WebURL            string     `json:"webURL"`
	NameRU            string     `json:"nameRU"`
	NameEN            string     `json:"nameEN"`
	PosterURL         string     `json:"posterURL"`
	Country           string     `json:"country"`
	Slogan            string     `json:"slogan"`
	Genre             string     `json:"genre"`
	Year              uint32     `json:"year,string"`
}

//RatingInfo film rating info
type RatingInfo struct {
	Rating     float32 `json:"rating,string"`
	RatingIMDb float32 `json:"ratingIMDb,string"`
}

//GalleryTemp temp sctuct for response
type GalleryInfo struct {
	Gallery Gallery `json:"gallery"`
}

type Gallery struct {
	Kadr   []KadrInfo `json:"kadr"`
	KadrSp []KadrInfo `json:"kadr_sp"`
	Poster []KadrInfo `json:"poster"`
}

type KadrInfo struct {
	Image   string `json:"image"`
	Preview string `json:"preview"`
}

type GalleryInfoForSave struct {
	Kadr   []KadrInfo
	KadrSp []KadrInfo
	Poster []KadrInfo
	FilmID string
}
