package postgresql

import (
	"database/sql"

	"tadeobennett.net/quotation/pkg/models"
)

type QuoteModel struct {
	DB *sql.DB
}

func (m *QuoteModel) Insert(author, category, body string) (int, error) {
	var id int

	s := `
	INSERT INTO quotations(author_name, category, quote)
	VALUES ($1, $2, $3)
	RETURNING quotation_id		
	`

	err := m.DB.QueryRow(s, author, category, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// return a slice of quotes but each of the quotes, return a pointer to that quote instead of copying them all
func (m *QuoteModel) Read() ([]*models.Quote, error) {

	s := `
	SELECT author_name, category, quote
	FROM quotations
	LIMIT 10
	`

	rows, err := m.DB.Query(s) //returns the rows of results
	if err != nil {
		//because the slice was empty
		return nil, err
	}
	//clean up before leave Read()
	defer rows.Close()

	quotes := []*models.Quote{}

	//Iterate over rows (a result set)
	for rows.Next() {

		//has to be initialized to empty
		q := &models.Quote{}

		err = rows.Scan(&q.Author_name, &q.Category, &q.Body)
		if err != nil {
			//the slice is empty
			return nil, err
		}

		//Append to quotes
		quotes = append(quotes, q)
	}

	//Always check the rows.Err()
	//instead of checking for errors and seeing if it is nil, we do it in one line
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}
