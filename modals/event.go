package modals

import (
	"time"
	"example.com/event-booker/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	User_ID     int64
}

var events = []Event{}

func (e *Event) Save() error {

	quary := `INSERT INTO events(name, description, location, dateTime, user_id) VALUES(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(quary)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.User_ID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}
	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	quary := `SELECT * FROM events`

	rows, err := db.DB.Query(quary)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.User_ID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	quary := `SELECT * FROM events WHERE id = ?`

	row := db.DB.QueryRow(quary, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.User_ID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	quary := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?`

	stmt, err := db.DB.Prepare(quary)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	if err != nil {
		return err
	}

	return nil

}

func (event Event) Delete() error {
	quary := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(quary)

	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(event.ID)

	return nil
}

func (e Event) Register(user_id int64) error {
	quary := `INSERT INTO registration(user_id, event_id) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(quary)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user_id, e.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e Event) CancelRegistration(user_id int64) error {
	quary := `DELETE FROM registration WHERE user_id = ? AND event_id = ?`

	stmt, err := db.DB.Prepare(quary)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, e.ID)

	if err != nil {
		return err
	}

	return nil
}
