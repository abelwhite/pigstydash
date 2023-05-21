// Filename: internal/models/pigs.go
package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Room struct {
	RoomID        int64
	Name          string
	NumOfPigsties int64
	Temperature   string
	Humidity      string
	CreatedAt     time.Time
}

type RoomModel struct {
	DB *sql.DB
}

func (m *RoomModel) Insert(name string, num_of_pigsty int64, temperature string, humidity string) (int64, error) {
	var id int64
	// build the query
	statement := `
	             INSERT INTO rooms(name, num_of_pigsty, temperature, humidity )
							 VALUES($1, $2, $3, $4)
							 RETURNING room_id
	             `
	// build a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// write options to the database
	err := m.DB.QueryRowContext(ctx, statement, name, num_of_pigsty, temperature, humidity).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (m *RoomModel) Read() ([]*Room, error) {
	//create SQL statement
	statement := `
		SELECT *
		FROM rooms
		
	`
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	//cleanup before we leave our read method
	defer rows.Close()

	rooms := []*Room{} //this will contain the pointer to all quotes

	for rows.Next() {
		q := &Room{}
		err = rows.Scan(&q.RoomID, &q.Name, &q.NumOfPigsties, &q.Temperature, &q.Humidity, &q.CreatedAt)

		if err != nil {
			return nil, err
		}
		rooms = append(rooms, q) //contain first row
	}
	//check to see if there were error generated

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil

}

func (m *RoomModel) Delete(roomID int) error {
	// create SQL statement to delete a quote with a given ID
	statement := `
		DELETE FROM rooms
		WHERE room_id = $1
	`
	// execute the delete statement and check for errors
	_, err := m.DB.Exec(statement, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (m *RoomModel) Update(q *Room) (int64, error) {
	var id int64
	// create SQL statement
	statement := `
        UPDATE rooms
        SET name=$1,num_of_pigsty=$2,temperature=$3, humidity=$4
        WHERE room_id=$5
    `
	//sets the timeout for the DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, statement, q.Name, q.NumOfPigsties, q.Temperature, q.Humidity, q.RoomID)
	fmt.Println(q.RoomID)

	if err != nil {
		fmt.Println(err)
		return 0, err

	}

	return id, err
}

func (m *RoomModel) Readd(room_id int) ([]*Room, error) {
	//create SQL statement
	statement := `
		SELECT *
		FROM rooms
		WHERE room_id = $1
	`
	rows, err := m.DB.Query(statement, room_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//cleanup before we leave our read method
	defer rows.Close()
	// pointer to every piece of data in the form
	rooms := []*Room{}

	rows.Next()

	q := &Room{}
	err = rows.Scan(&q.RoomID, &q.Name, &q.NumOfPigsties, &q.Temperature, &q.Humidity, &q.CreatedAt)

	rooms = append(rooms, q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//check to see if there were error generated
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rooms, nil

}
