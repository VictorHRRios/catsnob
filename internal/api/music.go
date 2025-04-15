package api

type Artist struct {
	Artists []struct {
		IDArtist       string  `json:"idArtist"`
		StrArtist      string  `json:"strArtist"`
		StrLabel       string  `json:"strLabel"`
		IDLabel        string  `json:"idLabel"`
		IntFormedYear  string  `json:"intFormedYear"`
		StrGenre       string  `json:"strGenre"`
		StrBiographyEN *string `json:"strBiographyEN"`
		StrBiographyES *string `json:"strBiographyES"`
		StrCountry     string  `json:"strCountry"`
		StrCountryCode string  `json:"strCountryCode"`
		StrArtistThumb string  `json:"strArtistThumb"`
		StrArtistLogo  string  `json:"strArtistLogo"`
	} `json:"artists"`
}

type ArtistAlbums struct {
	Album []struct {
		IDAlbum          string `json:"idAlbum"`
		IDArtist         string `json:"idArtist"`
		IDLabel          string `json:"idLabel"`
		StrAlbum         string `json:"strAlbum"`
		IntYearReleased  string `json:"intYearReleased"`
		StrGenre         string `json:"strGenre"`
		StrLabel         string `json:"strLabel"`
		StrAlbumThumb    string `json:"strAlbumThumb"`
		StrDescriptionEN string `json:"strDescriptionEN,omitempty"`
		StrDescriptionES any    `json:"strDescriptionES"`
	} `json:"album"`
}

type AlbumTracks struct {
	Track []struct {
		IDTrack        string `json:"idTrack"`
		IDAlbum        string `json:"idAlbum"`
		IDArtist       string `json:"idArtist"`
		StrTrack       string `json:"strTrack"`
		IntDuration    string `json:"intDuration"`
		StrGenre       string `json:"strGenre"`
		IntTrackNumber string `json:"intTrackNumber"`
	} `json:"track"`
}
