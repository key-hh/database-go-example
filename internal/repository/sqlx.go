package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Service struct {
	Email   string
	Telcode int
}

type SQLXRepository struct {
	db *sqlx.DB
}

func NewSQLXRepository() *SQLXRepository {
	return &SQLXRepository{}
}

func (ss *SQLXRepository) Init() error {
	db, err := sqlx.Connect("mysql", "test:123!@#@tcp(localhost:3306)/test")
	if err != nil {
		return err
	}
	ss.db = db

	db.MustExec("DROP TABLE IF EXISTS person, service")
	db.MustExec(`CREATE TABLE person (
    first_name varchar(20),
    last_name varchar(20),
    email varchar(20)
)`)
	db.MustExec(`CREATE TABLE service (
    email varchar(20),
    telcode int
)`)
	return nil
}

func (ss *SQLXRepository) Close() error {
	return ss.db.Close()
}

func (ss *SQLXRepository) Create(ctx context.Context) error {
	tx, err := ss.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	personStmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email) ")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer personStmt.Close()

	personStmt.ExecContext(ctx, Person{FirstName: "f1", LastName: "l1", Email: "e1"})
	personStmt.ExecContext(ctx, Person{FirstName: "f2", LastName: "l2", Email: "e2"})
	personStmt.ExecContext(ctx, Person{FirstName: "f3", LastName: "l3", Email: "e3"})

	if _, err := tx.NamedExecContext(ctx, "INSERT INTO service (email, telcode) VALUES (:email, :telcode)", []Service{
		{Email: "e1", Telcode: 111},
		{Email: "e2", Telcode: 222},
	}); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ss *SQLXRepository) Get(ctx context.Context, id string) (interface{}, error) {
	var p Person
	err := ss.db.GetContext(ctx, &p, "SELECT * FROM person where email=?", id)
	return p, err
}

func (ss *SQLXRepository) List(ctx context.Context) (interface{}, error) {
	var p []Person
	err := ss.db.SelectContext(ctx, &p, "SELECT * FROM person")
	return p, err
}
