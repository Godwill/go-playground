package main

import "fmt"

type Artist struct {
	name, studio string
	awards int
	albums []album
}

type album struct {
	name string
	year, songs int
}

func printName (a *Artist) {
	fmt.Printf("The artist name is %s\n", a.name)
}

func printAlbumCount (a *Artist) {
	fmt.Printf("%s has %d albums.\n", a.name, len(a.albums))
}

func printAlbums (a *Artist) {
	fmt.Printf("%s has the following %+v albums.\n", a.name, a.albums)
}

func releaseAlbum (a *Artist, alb ...album) []album{
	for _, v := range alb {
		a.albums = append(a.albums, v)
	}
	return a.albums
}

func changeName (a *Artist, name string) {
	fmt.Printf("The artist's formerly known as %s is now %s.\n", a.name, name)
	a.name = name
}

func main() {
	singer := Artist{name: "James Moswagdine", awards: 2, studio: "Jim Records",
			albums : []album{
				{name: "Happy Days Vol I", year: 2018, songs: 16},
			}}

			alb := album{songs: 14, name: "Happy Days Vol II", year: 2020}
			alb2 := album{songs: 14, name: "Happy Days Vol III", year: 2021}
			alb3 := album{songs: 14, name: "Happy Days Vol IV", year: 2021}

	printName(&singer)
	changeName(&singer, "Jim")
	releaseAlbum(&singer, alb)
	releaseAlbum(&singer, alb2, alb3)
	printName(&singer)
	printAlbums(&singer)
	printAlbumCount(&singer)
}
