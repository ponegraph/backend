package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	ErrFailedDatabaseQuery  = "failed to query the database"
	ErrSongNotFound         = "song not found"
	ErrArtistNotFound       = "artist not found"
	ErrUnableToReadDatabase = "unable to read data from the database"
)
