package main

import "fmt"

// Write SearchByArtist here (return all albums by the given artist)
func (lib Library) SearchByArtist(artist string) []Album {
	var list []Album

	for _, album := range lib {
		if album.Artist == artist {
			list = append(list, album)
		}
	}

	return list
}

// Write GenreCount here (count songs per genre across all albums)
func (lib Library) GenreCount() map[string]int {
	count := make(map[string]int)

	for _, album := range lib {
		for i := range album.Songs {
			count[album.Songs[i].Genre]++
		}
	}

	return count
}

// Write TotalSongs here (total number of songs across all albums)
func (lib Library) TotalSongs() int {
	var list []Song

	for _, album := range lib {
		for i := range album.Songs {
			list = append(list, album.Songs[i])
		}
	}

	return len(list)
}

// Define Library as a named map type (map albums by title)
type Library map[string]Album

// Write NewLibrary here
func NewLibrary() Library {
	return make(Library)
}

// Write AddAlbum method here
func (lib Library) AddAlbum(a Album) {
	lib[a.Title] = Album{
		Title:  a.Title,
		Artist: a.Artist,
		Year:   a.Year,
		Songs:  a.Songs,
	}
}

// Write FindAlbum method here (return Album and bool)
func (lib Library) FindAlbum(title string) (Album, bool) {
	album, found := lib[title]
	return album, found
}

// Write RemoveAlbum method here (return bool)
func (lib Library) RemoveAlbum(title string) bool {
	_, found := lib[title]
	if found {
		delete(lib, title)
	}

	return found
}

// Write AllSongs method here (return every song from every album)
func (lib Library) AllSongs() []Song {
	var list []Song

	for _, value := range lib {
		list = append(list, value.Songs...)
	}

	return list
}

/*

 */
// Define your Song struct here
type Song struct {
	Title    string
	Artist   string
	Duration int
	Genre    string
}

// Define your Album struct here
type Album struct {
	Title  string
	Artist string
	Year   int
	Songs  []Song
}

// Write NewSong here
func NewSong(title, artist string, duration int, genre string) Song {
	return Song{
		Title:    title,
		Artist:   artist,
		Duration: duration,
		Genre:    genre,
	}
}

// Write NewAlbum here (start with empty Songs)
func NewAlbum(title, artist string, year int) Album {
	return Album{
		Title:  title,
		Artist: artist,
		Year:   year,
		Songs:  nil,
	}
}

// Write AddSong method here (pointer receiver)
func (a *Album) AddSong(s Song) {
	a.Songs = append(a.Songs, s)
}

// Write TotalDuration method here (value receiver)
func (a Album) TotalDuration() int {
	sum := 0

	for i := range a.Songs {
		sum += a.Songs[i].Duration
	}

	return sum
}

func main() {
	lib := NewLibrary()

	a := NewAlbum("Abbey Road", "The Beatles", 1969)
	a.AddSong(NewSong("Come Together", "The Beatles", 259, "rock"))
	a.AddSong(NewSong("Something", "The Beatles", 182, "rock"))
	fmt.Printf("%s (%d) - %d songs, %ds total\n", a.Title, a.Year, len(a.Songs), a.TotalDuration())

	album, found := lib.FindAlbum("Abbey Road")
	fmt.Printf("Found: %v, %s by %s\n", found, album.Title, album.Artist)
	fmt.Printf("All songs: %d\n", len(lib.AllSongs()))

	abbey := NewAlbum("Abbey Road", "The Beatles", 1969)
	abbey.AddSong(NewSong("Come Together", "The Beatles", 259, "rock"))
	abbey.AddSong(NewSong("Something", "The Beatles", 182, "rock"))
	lib.AddAlbum(abbey)

	thriller := NewAlbum("Thriller", "Michael Jackson", 1982)
	thriller.AddSong(NewSong("Thriller", "Michael Jackson", 357, "pop"))
	thriller.AddSong(NewSong("Beat It", "Michael Jackson", 258, "rock"))
	lib.AddAlbum(thriller)

	fmt.Printf("Beatles albums: %d\n", len(lib.SearchByArtist("The Beatles")))
	fmt.Printf("Genre counts: %v\n", lib.GenreCount())
	fmt.Printf("Total songs: %d\n", lib.TotalSongs())

}
