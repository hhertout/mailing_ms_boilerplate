package repository

type MailQuery struct {
	ID      int    `db:"_id"`
	Date    string `db:"date"`
	To      string `db:"to"`
	Subject string `db:"subject"`
	Sent    bool   `db:"sent"`
	Error   string `db:"error"`
	Viewed  string `db:"viewed"`
}

type MailSaveWithoutError struct {
	To      string `db:"to"`
	Subject string `db:"subject"`
	Sent    bool   `db:"sent"`
}

type MailSaveWithError struct {
	To      string `db:"to"`
	Subject string `db:"subject"`
	Sent    bool   `db:"sent"`
	Error   string `db:"error"`
}

func (r Repository) SaveWithoutError(to []string, subject string) error {
	toConverted := r.convertToToString(to)
	_, err := r.dbPool.Exec(`
		INSERT INTO mail ("to", subject, sent) 
		VALUES ($1, $2, $3)
		`, toConverted, subject, true,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) SaveWithError(to []string, subject string, error error) error {
	toConverted := r.convertToToString(to)
	_, err := r.dbPool.Exec(`
		INSERT INTO mail ("to", subject, sent, error) 
		VALUES ($1, $2, $3, $4)
		`, toConverted, subject, false, error.Error(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) ListMail() ([]MailQuery, error) {
	var res []MailQuery

	rows, err := r.dbPool.Query(`SELECT * FROM "mail" LIMIT 1000`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mail MailQuery
		if err := rows.Scan(&mail.ID, &mail.Date, &mail.To, &mail.Subject, &mail.Sent, &mail.Error, &mail.Viewed); err != nil {
			return nil, err
		}
		res = append(res, mail)
	}

	return res, nil
}

func (r Repository) convertToToString(args []string) string {
	res := ""
	for _, dest := range args {
		res += dest
	}
	return res
}
