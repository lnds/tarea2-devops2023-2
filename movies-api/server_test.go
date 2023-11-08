package main

import (
	"database/sql"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/lnds/lab-ic-ms/api/database"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}

func TestIndexRoute(t *testing.T) {

	sqlDb, gormDb, mock := DbMock(t)
	defer sqlDb.Close()

	instanceDb := database.NewDbInstance(gormDb)
	database.DB = instanceDb

	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		method string
		route  string

		mockQuery func(mock sqlmock.Sqlmock)
		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description: "success: GET /directors",
			route:       "/directors",
			method:      "GET",
			mockQuery: func(mock sqlmock.Sqlmock) {
				directors := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "John Huston").
					AddRow(2, "Alfred Hitchcok")
				expectedSQL := "SELECT (.+) FROM \"directors\""
				mock.ExpectQuery(expectedSQL).WillReturnRows(directors)
			},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "[{\"id\":1,\"name\":\"John Huston\"},{\"id\":2,\"name\":\"Alfred Hitchcok\"}]",
		},
		{
			description: "success: GET directors/1",
			route:       "/directors/1",
			method:      "GET",
			mockQuery: func(mock sqlmock.Sqlmock) {
				directors := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "John Huston")
				expectedSQL := "SELECT (.+) FROM \"directors\" WHERE \"directors\".\"id\" =(.+)"
				mock.ExpectQuery(expectedSQL).WillReturnRows(directors)
			},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"id\":1,\"name\":\"John Huston\"}",
		},
		{
			description: "failure: GET directors/3",
			route:       "/directors/3",
			method:      "GET",
			mockQuery: func(mock sqlmock.Sqlmock) {
				directors := sqlmock.NewRows([]string{"id", "name"})
				expectedSQL := "SELECT (.+) FROM \"directors\" WHERE \"directors\".\"id\" =(.+)"
				mock.ExpectQuery(expectedSQL).WillReturnRows(directors)
			},
			expectedError: false,
			expectedCode:  500,
			expectedBody:  "{\"message\":\"record not found\"}",
		},
		{
			description: "success: GET /movies",
			route:       "/movies",
			method:      "GET",
			mockQuery: func(mock sqlmock.Sqlmock) {
				directors := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(2, "Alfred Hitchcok").
					AddRow(1, "John Huston")
				movies := sqlmock.NewRows([]string{"id", "title", "description", "year", "director_id"}).
					AddRow(1, "Psicosis", "Drama", 1960, 2).
					AddRow(2, "The Misfits", "Drama", 1961, 1)

				expectedSQL := "SELECT (.+) FROM \"movies\""
				mock.ExpectQuery(expectedSQL).WillReturnRows(movies)
				expectedSQLD := "SELECT (.+) FROM \"directors\" WHERE \"directors\".\"id\" IN(.+)"
				mock.ExpectQuery(expectedSQLD).WillReturnRows(directors)

			},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "[{\"id\":\"1\",\"title\":\"Psicosis\",\"year\":\"1960\",\"description\":\"Drama\",\"director_id\":\"2\",\"director\":{\"id\":2,\"name\":\"Alfred Hitchcok\"}},{\"id\":\"2\",\"title\":\"The Misfits\",\"year\":\"1961\",\"description\":\"Drama\",\"director_id\":\"1\",\"director\":{\"id\":1,\"name\":\"John Huston\"}}]",
		},
		{
			description: "success: GET movies/1",
			route:       "/movies/1",
			method:      "GET",
			mockQuery: func(mock sqlmock.Sqlmock) {
				directors := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(2, "Alfred Hitchcok")
				movies := sqlmock.NewRows([]string{"id", "title", "description", "year", "director_id"}).
					AddRow(1, "Psicosis", "Drama", 1960, 2)
				expectedSQL := "SELECT (.+) FROM \"movies\" WHERE \"movies\".\"id\" =(.+)"
				mock.ExpectQuery(expectedSQL).WillReturnRows(movies)
				expectedSQLD := "SELECT (.+) FROM \"directors\" WHERE \"directors\".\"id\" =(.+)"
				mock.ExpectQuery(expectedSQLD).WillReturnRows(directors)

			},
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"id\":\"1\",\"title\":\"Psicosis\",\"year\":\"1960\",\"description\":\"Drama\",\"director_id\":\"2\",\"director\":{\"id\":2,\"name\":\"Alfred Hitchcok\"}}",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	// Setup the app as it is done in the main function
	app := fiber.New()
	setupRoutes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			test.method,
			test.route,
			nil,
		)

		if test.mockQuery != nil {
			test.mockQuery(mock)
		}

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := io.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
