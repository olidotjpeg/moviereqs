package main

var newMovies = []Movie{
	{ID: "6", Title: "The Lion King", Genre: "", ReleaseYear: 2019, Director: "Jon Favreau", LiveAction: 1},
	{ID: "7", Title: "The Terminator", Genre: "", ReleaseYear: 1984, Director: "James Cameron", LiveAction: 1},
	{ID: "8", Title: "The Lion, the Witch and the Wardrobe", Genre: "", ReleaseYear: 2005, Director: "Andrew Adamson", LiveAction: 1},
}

func makeRecommendations() {
	movies := getAllMovies()
	ratings := getAllRatingsForUserById(1)

	// Create a map of movie IDs to indices in the factor matrix
	movieIdxMap := make(map[string]int)
	for i, m := range movies {
		movieIdxMap[m.ID] = i
	}

	var xData [][]float64
	var yData []float64

	for _, movie := range movies {
		var row []float64

		row = append(row, float64(movie.LiveAction))
		row = append(row, float64(movie.ReleaseYear))

		if movie.Genre == "Action" {
			row = append(row, 1.0)
		} else {
			row = append(row, 0.0)
		}
		if movie.Genre == "Comedy" {
			row = append(row, 1.0)
		} else {
			row = append(row, 0.0)
		}
		if movie.Genre == "Drama" {
			row = append(row, 1.0)
		} else {
			row = append(row, 0.0)
		}

		xData = append(xData, row)

		rating := getRatingForMovieAndUser(movie.ID, ratings)
		yData = append(yData, float64(rating))
	}
}

// func createRegressionTree(xData [][]float64, yData []float64) *Node {
// 	rand.Seed(0)

// 	root := &Node {
// 		xData: xData,
// 		yData: yData,
// 	}

// 	splitNode(root)

// 	return root
// }

// func splitNode() {

// }

func getRatingForMovieAndUser(movieID string, ratings []Rating) int {
	for _, rating := range ratings {
		if rating.MovieId == movieID {
			return rating.Rating
		}
	}
	return 0
}
