package helper

import "fmt"

func GetAllArtistUnitBySongIdQuery(songId int) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
	
		SELECT DISTINCT ?artistName ?artistId ?mbUrl
		WHERE {
			?song v:hasSongId %d ;
			rdf:type owl:Song ;
			v:hasArtist ?artist .
	
			?artist rdfs:label ?artistName .
			OPTIONAL {
			?artist v:hasMbid ?mbUrl .
			?mbUrl rdfs:label ?artistId .
			}
		}
	`, songId)

	return query
}

func GetArtistByIdQuery(artistId string) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		PREFIX mbid: <http://musicbrainz.org/artist/>
		
		SELECT DISTINCT ?artistName ?artistId ?mbUrl ?countryName ?totalLastfmListeners ?totalLastfmScrobbles (GROUP_CONCAT(?tagName; separator=", ") AS ?tags)
		WHERE {
			?artist v:hasMbid mbid:%s ;
				rdf:type owl:Artist ;
				rdfs:label ?artistName ;
				v:hasMbid ?mbUrl .
			
			?mbUrl rdfs:label ?artistId .
		
			OPTIONAL { 
				?artist v:hasCountry ?country .
				?country rdfs:label ?countryName
			}
			OPTIONAL { ?artist v:hasTotalLastfmListeners ?totalLastfmListeners }
			OPTIONAL { ?artist v:hasTotalLastfmScrobbles ?totalLastfmScrobbles }
			OPTIONAL { 
				?artist v:hasTag ?tag .
				?tag rdfs:label ?tagName .
			}
		}
		GROUP BY ?artistName ?artistId ?mbUrl ?countryName ?totalLastfmListeners ?totalLastfmScrobbles
		LIMIT 1
	`, artistId)

	return query
}

func GetAllArtistUnitByTagQuery(tag string) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		SELECT DISTINCT ?artistName ?artistId ?mbUrl 
		WHERE {
			?artist rdf:type owl:Artist ;
			rdfs:label ?artistName ;
			v:hasMbid ?mbUrl ;
			v:hasTag ?tag .
			
			?tag rdfs:label ?tagLabel
			
			FILTER(REGEX(?tagLabel, "%s", "i"))
			
			?mbUrl rdfs:label ?artistId .
		}
	`, tag)

	return query
}

func GetAllArtistUnitByNameQuery(artistName string) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		SELECT DISTINCT ?artistName ?artistId ?mbUrl
		WHERE {
			?artist rdf:type owl:Artist ;
			rdfs:label ?artistName ;
			v:hasMbid ?mbUrl ;
			v:hasNormalizedName ?artistSearchName .
		
			FILTER(REGEX(?artistSearchName, "%s", "i"))
		
			?mbUrl rdfs:label ?artistId .
		}
	`, artistName)

	return query
}

func GetArtistInfoFromDbpediaQuery(mbUrl string) string {
	query := fmt.Sprintf(`
		PREFIX dbo: <http://dbpedia.org/ontology/>
		PREFIX foaf: <http://xmlns.com/foaf/0.1/>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>

		SELECT DISTINCT ?description ?externalReference ?imageUrl
		WHERE {
			?artist owl:sameAs <%s> ;
					dbo:abstract ?description ;
					dbo:thumbnail ?imageUrl ;
					foaf:isPrimaryTopicOf ?externalReference .
			FILTER (LANG(?description) = "en")
		}
		LIMIT 1
	`, mbUrl)
	return query
}

func GetTopKArtistUnitQuery(k int) string {
	query := fmt.Sprintf(`
		PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
		PREFIX rdfs: <http://www.w3.org/2000/01/rdf-schema#>
		PREFIX owl: <http://www.w3.org/2002/07/owl#>
		PREFIX v: <http://example.com/vocab#>
		
		PREFIX xsd: <http://www.w3.org/2001/XMLSchema#>
		SELECT DISTINCT ?artistName ?artistId ?mbUrl
		WHERE {
			?artist rdf:type owl:Artist ;
			rdfs:label ?artistName ;
			v:hasMbid ?mbUrl ;
			v:hasTotalLastfmListeners ?totalLastfmListeners .
					
			?mbUrl rdfs:label ?artistId .
		}
		ORDER BY DESC(xsd:integer(?totalLastfmListeners))
		LIMIT %d
	`, k)
	return query
}
