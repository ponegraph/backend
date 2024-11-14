package helper

import "fmt"

func GetSongByIdQuery(songId int) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		SELECT DISTINCT ?songName ?releaseDate ?bpm ?key ?mode ?spotifyStream ?spotifyPlaylistCount ?applePlaylistCount
		?deezerPlaylistCount ?spotifyChart ?appleChart ?deezerChart ?shazamChart ?danceability ?energy ?valence ?acousticness ?instrumentalness ?liveness ?speechiness
		WHERE {
			?song v:hasSongId %d ;
			rdf:type owl:Song ;
			rdfs:label ?songName ;
			v:hasReleaseDate ?releaseDate ;
			v:hasBpm ?bpm ;
			v:hasTotalSpotifyStream ?spotifyStream ;
			v:hasSpotifyPlaylistCount ?spotifyPlaylistCount ;
			v:hasApplePlaylistCount ?applePlaylistCount ;
			v:hasDeezerPlaylistCount ?deezerPlaylistCount ;
			v:hasSpotifyChart ?spotifyChart ;
			v:hasAppleChart ?appleChart ;
			v:hasDeezerChart ?deezerChart ;
			v:hasAcousticness ?acousticness ;
			v:hasDanceability ?danceability ; 
			v:hasEnergy ?energy ;
			v:hasInstrumentalness ?instrumentalness ;
			v:hasLiveness ?liveness ;
			v:hasSpeechiness ?speechiness ;
			v:hasValence ?valence ;
			v:hasMode ?mode .
			
			OPTIONAL { ?song v:hasKey ?key }
			OPTIONAL { ?song v:hasShazamChart ?shazamChart }
		} limit 1
	`, songId)

	return query
}

func GetTopKSongUnitQuery(k int) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		SELECT DISTINCT ?songName ?songId ?releaseDate 
		WHERE {
			?song v:hasSongId ?songId ;
			rdf:type owl:Song ;
			rdfs:label ?songName ;
			v:hasReleaseDate ?releaseDate ;
			v:hasTotalSpotifyStream ?spotifyStream .
		}
		ORDER BY DESC(?spotifyStream)
		LIMIT %d
	`, k)
	return query
}

func GetSongUnitByIdQuery(songId int) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		SELECT DISTINCT ?songName ?songId ?releaseDate
		WHERE {
			?song v:hasSongId ?songId ;
			rdf:type owl:Song ;
			rdfs:label ?songName ;
			v:hasReleaseDate ?releaseDate ;
			v:hasSongId %d .
		}
		LIMIT 1
	`, songId)

	return query
}

func GetAllSongUnitFromArtistIdQuery(artistId string) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		PREFIX mbid: <http://musicbrainz.org/artist/>
		
		SELECT DISTINCT ?songName ?songId ?releaseDate
		WHERE {
			?song v:hasSongId ?songId ;
			rdf:type owl:Song ;
			rdfs:label ?songName ;
			v:hasReleaseDate ?releaseDate ;
			v:hasArtist ?artist .
		
			?artist v:hasMbid mbid:%s .
		}
	`, artistId)

	return query
}

func GetAllSongFeatureQuery() string {
	query := `
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		SELECT DISTINCT ?songName ?songId ?bpm ?danceability ?energy ?valence ?acousticness ?instrumentalness ?liveness ?speechiness
		WHERE {
			?song v:hasSongId ?songId ;
			rdf:type owl:Song ;
			rdfs:label ?songName ;
			v:hasBpm ?bpm ;
			v:hasAcousticness ?acousticness ;
			v:hasDanceability ?danceability ; 
			v:hasEnergy ?energy ;
			v:hasInstrumentalness ?instrumentalness ;
			v:hasLiveness ?liveness ;
			v:hasSpeechiness ?speechiness ;
			v:hasValence ?valence .
		} 
		ORDER BY ?songId
	`

	return query
}

func GetAllSongIdFromSameArtistQuery(songId int) string {
	// songId is excluded from the result
	query := fmt.Sprintf(`
		PREFIX v: <http://example.com/vocab#>
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		
		SELECT DISTINCT ?songId
		WHERE {
			?targetSong v:hasSongId %d ;
			rdf:type owl:Song ;
			v:hasArtist ?artist .
			
			?song v:hasArtist ?artist ;
			v:hasSongId ?songId .
			
			FILTER(?songId != %d)
		}
	`, songId, songId)

	return query
}
