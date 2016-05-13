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
		c.Insert(film)
	}
	wg.Done()
}

func scrapper() {
	for atomic.LoadUint64(&atomicCounter) < end {
		atomic.AddUint64(&atomicCounter, 1)
		film, err := api.GetFilm(atomic.LoadUint64(&atomicCounter) - 1)
		if err == nil {
			c.Insert(film)
		}
	}
	wg.Done()
}

var (
	c             *mgo.Collection
	end           uint64
	wg            sync.WaitGroup
	atomicCounter uint64
)

func main() {
	fmt.Println("Start scrapper..")

	/*var conf config
	err := conf.GetConfig("config.json")
	end = conf.EndID
	atomicCounter = conf.StartID

	if err != nil {
		fmt.Println("Cannot find config.json")
		return
	}

	session, err := mgo.Dial(conf.MongoURL)
	defer session.Close()
	if err != nil {
		fmt.Println("Cannot find config.json")
		return
	}
	c = session.DB("kinopoisk").C("films")

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		scrapper()
	}

	wg.Wait()*/

	gallery, _ := api.GetGallery(714888)

	fmt.Println(gallery.Gallery.Kadr[5])

	fmt.Println("Done!")
}
