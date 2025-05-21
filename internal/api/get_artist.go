package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://www.theaudiodb.com/api/v1/json/2/"

// Funciones para obtener estructuras usando la api de theaudiodb.com
// Casi todas siguen un formato similar maybe se puede usar un middleware
// para DRYear el codigo

func GetArtist(artistId *string) (Artist, error) {
	fullUrl := baseUrl + "artist.php?i="
	if artistId == nil {
		return Artist{}, fmt.Errorf("Nothing to retrieve")
	}
	fullUrl += *artistId
	res, err := http.Get(fullUrl)
	if err != nil {
		return Artist{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Artist{}, fmt.Errorf("%v", res.Status)
	}

	artistDetail := Artist{}

	if err := json.Unmarshal(body, &artistDetail); err != nil {
		return Artist{}, err
	}
	return artistDetail, nil
}

func GetAlbums(artistId *string) (ArtistAlbums, error) {
	fullUrl := baseUrl + "album.php?i="
	if artistId == nil {
		return ArtistAlbums{}, fmt.Errorf("Nothing to retrieve")
	}
	fullUrl += *artistId

	res, err := http.Get(fullUrl)
	if err != nil {
		return ArtistAlbums{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return ArtistAlbums{}, fmt.Errorf("%v", res.Status)
	}

	albumsDetail := ArtistAlbums{}

	if err := json.Unmarshal(body, &albumsDetail); err != nil {
		return ArtistAlbums{}, err
	}
	for key := range albumsDetail.Album {
		if len(albumsDetail.Album[key].StrAlbumThumb) == 0 {
			albumsDetail.Album[key].StrAlbumThumb = "/app/assets/images/not_available.png"
		}
	}
	return albumsDetail, nil
}

func GetAlbumSongs(albumId *string) (AlbumTracks, error) {
	fullUrl := baseUrl + "track.php?m="
	if albumId == nil {
		return AlbumTracks{}, fmt.Errorf("Nothing to retrieve")
	}
	fullUrl += *albumId
	res, err := http.Get(fullUrl)
	if err != nil {
		return AlbumTracks{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return AlbumTracks{}, fmt.Errorf("%v", res.Status)
	}

	albumTracksDetail := AlbumTracks{}

	if err := json.Unmarshal(body, &albumTracksDetail); err != nil {
		return AlbumTracks{}, err
	}
	return albumTracksDetail, nil
}
