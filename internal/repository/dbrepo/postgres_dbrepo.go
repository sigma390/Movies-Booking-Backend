package dbrepo

import (
	"backend/internal/models" // Importing the Movie model structure from the internal package
	"context"
	"database/sql" // Provides database interaction capabilities
	"time"
)

// PostgresDBRepo is a struct that holds the database connection pool
type PostgresDBRepo struct {
	DB *sql.DB
}

// dbTimeout specifies the maximum duration for database queries
const dbTimeout = time.Second * 3

// AllMovies retrieves all movies from the "movies" table
func (m *PostgresDBRepo) AllMovies() ([]*models.Movie, error) {
	// 1. Create a context with a timeout to limit the query execution time
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel() // Ensures the context is canceled after the function returns

	// 2. Define the SQL query to fetch movie details
	query := `
		select
			id, title, release_date, runtime,
			mpaa_rating, description, coalesce(image, ''), -- Handle null images with coalesce
			created_at, updated_at
		from
			movies
		order by
			title -- Order movies alphabetically by title
	`

	// 3. Execute the SQL query using the context for timeout control
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err // Return error if query execution fails
	}
	defer rows.Close() // Ensure that the database rows are closed after reading

	// 4. Initialize a slice to store the list of movies
	var movies []*models.Movie

	// 5. Iterate through the rows returned by the query
	for rows.Next() {
		var movie models.Movie

		// 6. Map the query result fields to the Movie struct fields
		err := rows.Scan(
			&movie.ID,           // Movie ID
			&movie.Title,        // Movie Title
			&movie.ReleaseDate,  // Release date of the movie
			&movie.RunTime,      // Runtime of the movie in minutes
			&movie.MPAARating,   // MPAA rating (e.g., PG, R)
			&movie.Description,  // Description or synopsis of the movie
			&movie.Image,        // Movie poster image (use empty string if null)
			&movie.CreatedAt,    // Timestamp of when the movie was created
			&movie.UpdatedAt,    // Timestamp of when the movie was last updated
		)
		if err != nil {
			return nil, err // Return error if mapping fails
		}

		// 7. Append the current movie to the slice
		movies = append(movies, &movie)
	}

	// 8. Return the list of movies and nil error
	return movies, nil
}


func (m *PostgresDBRepo) Connection() *sql.DB{
	return m.DB; // Return the database connection pool
}