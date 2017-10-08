package main

// Stats is base structure for either anime/manga stats; see AnimeStats or MangaStats
type stats struct {
	// TODO: move common fields for AnimeStats/MangaStats to here
}

// AnimeStats is structure for holding anime statistics
type AnimeStats struct {
	inProgress int
	completed  int
	onhold     int
	dropped    int
	planned    int
	rewatched  int
	days       float64
	meanScore  float64
	episodes   int
}

// MangaStats is structure for holding manga statistics
type MangaStats struct {
	inProgress int
	completed  int
	onhold     int
	dropped    int
	planned    int
	rewatched  int
	days       float64
	meanScore  float64
	chapters   int
	volumes    int
}
