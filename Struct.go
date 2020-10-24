package main

import "fmt"

//MARK: Passby Values and Reference 
type Artist struct {
	name, genre string
	songs int
}

func main() {
	
	newSong := &Artist{name: "Kalidoss",genre: "Love",songs: 100}

	fmt.Println(newSong.name)
	fmt.Println(newSong.genre)
	fmt.Println(newRelease(newSong))
	fmt.Println(newSong.name)
	fmt.Println(newSong.songs)

}

func newRelease(a *Artist) int {
a.songs++
a.name = "Test Song"
	return a.songs
}
