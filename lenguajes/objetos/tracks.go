package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Track struct {
	Title   string
	Artist  string
	Album   string
	Year    int
	Length  time.Duration
	WavPath string
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

var tracks = []*Track{
	{
		Title:   "Bohemian Rhapsody",
		Artist:  "Queen",
		Album:   "A Night at the Opera",
		Year:    1975,
		Length:  time.Minute*5 + time.Second*55,
		WavPath: "./audio/bohemian_rhapsody.wav",
	},
	{
		Title:   "Imagine",
		Artist:  "John Lennon",
		Album:   "Imagine",
		Year:    1971,
		Length:  time.Minute*3 + time.Second*3,
		WavPath: "./audio/imagine.wav",
	},
	{
		Title:   "Hotel California",
		Artist:  "Eagles",
		Album:   "Hotel California",
		Year:    1977,
		Length:  time.Minute*6 + time.Second*30,
		WavPath: "./audio/hotel_california.wav",
	},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length", "File")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "----", "------", "----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length, t.WavPath)
	}
	tw.Flush()
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}
