package api

const endPoint = "http://api.kinopoisk.cf/"

// get film info
const getFilmQuery = endPoint + "getFilm?filmID=%v"

// get film screenshots
const getGalleryQuery = endPoint + "getGallery?filmID=%v"

//get film staff ( actors, operators etc)
const getStaffQuery = endPoint + "getStaff?filmID=%v"
