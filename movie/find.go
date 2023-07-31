package movie

func FindName(imdbID string) string {
	switch imdbID {
	case "marv01":
		return "Avengers: Endgame"
	case "marv02":
		return "Black Panther"
	}

	return "not found."
}
