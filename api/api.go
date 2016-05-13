package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetInstanse - get instanse by id
func getInstanse(id uint64, instanseQuery string, instanse interface{}) error {
	query := fmt.Sprintf(instanseQuery, id)

	res, err := http.Get(query)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if string(body) == "null" {
		err := errors.New("Null json")
		return err
	}

	err = json.Unmarshal(body, instanse)

	return err
}

// GetFilm - get film by id
func GetFilm(id uint64) (FilmInfo, error) {
	var film FilmInfo
	err := getInstanse(id, getFilmQuery, &film)
	return film, err
}

func GetGallery(id uint64) (GalleryInfo, error) {
	var gallery GalleryInfo
	err := getInstanse(id, getGalleryQuery, &gallery)
	return gallery, err
}
