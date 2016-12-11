package models

import (
	"database/sql"
)

const (
	QInitTeamTable = `create table if not exists teams (
		id integer primary key,
		name text unique,
		members text,
		score integer,
		token text unique,
		last_valid_submission timestamp
);`

	QInitSubmitted = `create table if not exists submitted (
		id integer primary key,
		team_id integer,
		flag_id integer,
		foreign key(team_id) references teams(id)
	);`

	QInitSessionsTable = `create table if not exists sessions (
		token varchar(16) primary key,
		created_at timestamp,
		expires_at timestamp
	);`

	QGetTeams = `
select id, name, members, score, token, last_valid_submission
from teams;`

	QCreateTeam = `
insert into teams (
	name, members, score, token, last_valid_submission
) values (
	?, ?, 0, ?, datetime(0, 'unixepoch', 'localtime')
);`

	QFindTeamBySubmissionToken = `
select id, name, members, score, last_valid_submission
from teams
where token = ?;`

	QUpdateTeam = `
update teams
set score = ?, token = ?, last_valid_submission = ?
where id = ?;`

	QFindSubmission = `
select id
from submitted
where team_id = ? and flag_id = ?;`

	QFindSessionToken = `
select created_at, expires_at
from sessions
where token = ?;`

	QSaveSubmission = `
insert into submitted (
	team_id, flag_id
) values (
	?, ?
);`
)

// InitTables initializes the database tables.
func InitTables(db *sql.DB) error {
	_, err := db.Exec(QInitTeamTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(QInitSessionsTable)
	if err != nil {
		return err
	}
	_, err = db.Exec(QInitSubmitted)
	return err
}