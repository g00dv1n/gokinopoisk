package main

import (
	"fmt"
	"gokinopoisk/api"
	"gopkg.in/mgo.v2"
	"sync"
	"sync/atomic"
)

func runner(id uint64) {
	film, err := api.GetFilm(id)
	if err == nil {
		filmsCollection.Insert(film)
	}
	wg.Done()
}

func scrapper() {
	for atomic.LoadUint64(&atomicCounter) < end {
		var (
			counter     uint64
			film        api.FilmInfo
			gallery     api.GalleryInfo
			saveGallery api.GalleryInfoForSave
			err         error
		)

		atomic.AddUint64(&atomicCounter, 1)
		counter = atomic.LoadUint64(&atomicCounter) - 1
		film, err = api.GetFilm(counter)
		gallery, err = api.GetGallery(counter)

		saveGallery = api.GalleryInfoForSave{
			Kadr:   gallery.Gallery.Kadr,
			KadrSp: gallery.Gallery.KadrSp,
			Poster: gallery.Gallery.Poster,
			FilmID: film.FilmID,
		}

		if err == nil && film.RatingData.RatingIMDb > minImdb {
			filmsCollection.Insert(film)
			imagesCollection.Insert(saveGallery)
		}
	}
	wg.Done()
}

var (
	filmsCollection  *mgo.Collection
	imagesCollection *mgo.Collection
	end              uint64
	wg               sync.WaitGroup
	atomicCounter    uint64
	minImdb          float32
)

func main() {
	fmt.Println("Start scrapper..")

	var conf config
	err := conf.GetConfig("config.json")
	end = conf.EndID
	atomicCounter = conf.StartID
	minImdb = conf.MinImdb

	if err != nil {
		fmt.Println("Cannot find config.json")
		return
	}

	/*mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{conf.MongoURL},
		Database: conf.Database,
		Username: conf.User,
		Password: conf.Password,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)*/

	session, err := mgo.Dial("localhost")
	defer session.Close()
	if err != nil {
		fmt.Println("Cannot connect to mongo")
		return
	}
	filmsCollection = session.DB(conf.Database).C("films")
	imagesCollection = session.DB(conf.Database).C("images")

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		scrapper()
	}

	wg.Wait()

	fmt.Println("Done!")
}
