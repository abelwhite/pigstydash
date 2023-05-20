// Filename: internal/models/pigs.go
package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Pigsty struct {
	PigstyID  int64
	Room      string
	Name      string
	NumOfPigs int64
	CreatedAt time.Time
}

type PigstyModel struct {
	DB *sql.DB
}

func (m *PigstyModel) Insert(room string, name string, num_of_pigs int64) (int64, error) {
	var id int64
	// build the query
	statement := `
	             INSERT INTO pigsty(room, name, num_of_pigs )
							 VALUES($1, $2, $3 )
							 RETURNING pigsty_id
	             `
	// build a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// write options to the database
	err := m.DB.QueryRowContext(ctx, statement, room, name, num_of_pigs).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *PigstyModel) Read() ([]*Pigsty, error) {
	//create SQL statement
	statement := `
		SELECT *
		FROM pigsty
		
	`
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	//cleanup before we leave our read method
	defer rows.Close()

	pigsty := []*Pigsty{} //this will contain the pointer to all quotes

	for rows.Next() {
		q := &Pigsty{}
		err = rows.Scan(&q.PigstyID, &q.Room, &q.Name, &q.NumOfPigs, &q.CreatedAt)

		if err != nil {
			return nil, err
		}
		pigsty = append(pigsty, q) //contain first row
	}
	//check to see if there were error generated

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return pigsty, nil

}

func (m *PigstyModel) Delete(pigID int) error {
	// create SQL statement to delete a quote with a given ID
	statement := `
		DELETE FROM pigsty
		WHERE pigsty_id = $1
	`
	// execute the delete statement and check for errors
	_, err := m.DB.Exec(statement, pigID)
	if err != nil {
		return err
	}
	return nil
}

func (m *PigstyModel) Update(q *Pigsty) (int64, error) {
	var id int64
	// create SQL statement
	statement := `
        UPDATE pigsty
        SET room=$1,name=$2,num_of_pigs=$3
        WHERE pigsty_id=$4
    `
	//sets the timeout for the DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, statement, q.Room, q.Name, q.NumOfPigs, q.PigstyID)
	fmt.Println(q.PigstyID)

	if err != nil {
		fmt.Println(err)
		return 0, err

	}

	return id, err
}

func (m *PigstyModel) Readd(pigsty_id int) ([]*Pigsty, error) {
	//create SQL statement
	statement := `
		SELECT *
		FROM pigsty
		WHERE pigsty_id = $1
	`
	rows, err := m.DB.Query(statement, pigsty_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//cleanup before we leave our read method
	defer rows.Close()
	// pointer to every piece of data in the form
	pigsty := []*Pigsty{}

	rows.Next()

	q := &Pigsty{}
	err = rows.Scan(&q.PigstyID, &q.Room, &q.Name, &q.NumOfPigs, &q.CreatedAt)

	pigsty = append(pigsty, q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//check to see if there were error generated
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return pigsty, nil

}
