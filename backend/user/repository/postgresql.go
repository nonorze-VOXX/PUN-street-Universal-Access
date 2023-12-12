package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type postgresqlUserRepo struct {
	db *sql.DB
}

func NewPostgressqlUserRepo(db *sql.DB) domain.UserRepo {
	return &postgresqlUserRepo{db}
}

func (p *postgresqlUserRepo) Login(ctx context.Context, email string, password string) (int, error) {
	sqlStatement := `
	SELECT password, authority FROM user_data WHERE email = $1;
	`
	var hashedPassword string
	var authority int
	err := p.db.QueryRow(sqlStatement, email).Scan(&hashedPassword, &authority)
	if err != nil {
		return 0, err
	}
	if hashedPassword != password {
		return 0, errors.New("wrong password")
	}

	return authority, nil
}

func (p *postgresqlUserRepo) GetByID(ctx context.Context, id int64) (*swagger.UserData, error) {
	row := p.db.QueryRow("SELECT * FROM user_data WHERE user_id = $1", id)
	s := &swagger.UserData{}
	if err := row.Scan(&s.UserId, &s.UserName, &s.Password, &s.UserEmail, &s.Address, &s.Phone, &s.Birthday, &s.Authority, &s.CartId, &s.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (p *postgresqlUserRepo) GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error) {
	sqlStatement := `
		SELECT 	email, user_id, name, authority, status
		FROM user_data;
	`

	rows, err := p.db.Query(sqlStatement)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	l := []swagger.UserDataShort{}
	for rows.Next() {
		s := swagger.UserDataShort{}
		err := rows.Scan(&s.UserEmail, &s.UserId, &s.UserName, &s.Authority, &s.Status)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		l = append(l, s)
	}

	return l, nil
}
