package db

import "database/sql"

type Database struct {
	db *sql.DB
}

func (d *Database) CreateEventsTable() error {
	query := `CREATE TABLE IF NOT EXISTS Events(
        eventKey TEXT Primary Key,
        year INTEGER NOT NULL,
        type INTEGER,
        startDate TEXT,
        endDate TEXT
    );`

	_, err := d.db.Exec(query)
	return err
}

func (d *Database) CreateMatchesTable() error {
	query := `CREATE TABLE IF NOT EXISTS Matches(
        matchId INTEGER,
        eventKey TEXT,
        teams TEXT NOT NULL,
        blueTeams TEXT NOT NULL,
        redTeams TEXT NOT NULL,
        blueScore INTEGER,
        redScore INTEGER,
        blueAutoScore INTEGER,
        redAutoScore INTEGER,
        blueTeleScore INTEGER,
        redTeleScore INTEGER,
        blueRP INTEGER,
        redRP INTEGER,
        PRIMARY KEY (matchId, eventKey),
        FOREIGN KEY (eventKey) REFERENCES Events(eventKey)
    );`

	_, err := d.db.Exec(query)
	return err
}

func (d *Database) CreateTeamsTable() error {
	query := `CREATE TABLE IF NOT EXISTS Teams(
        teamKey TEXT PRIMARY KEY
    );`

	_, err := d.db.Exec(query)
	return err
}

// TODO: figure out how teams should be handled (2-3 strings, separate table?)
// - sqlite json helpers? (for insert)

// TODO: other tables
