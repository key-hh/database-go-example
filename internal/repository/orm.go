package repository

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-database/ent/entgen"
	"go-database/ent/entgen/migrate"
	"go-database/ent/entgen/user"
)

type ORMRepository struct {
	db     *sql.DB
	client *entgen.Client
	dbx    *sqlx.DB
}

func NewORMRepository() *ORMRepository {
	return &ORMRepository{}
}

func (ss *ORMRepository) Init(ctx context.Context) error {
	db, err := sql.Open("mysql", "test:123!@#@tcp(localhost:3306)/test")
	if err != nil {
		return err
	}
	ss.db = db
	ss.client = entgen.NewClient(entgen.Driver(entsql.OpenDB("mysql", db)))
	ss.dbx = sqlx.NewDb(db, "mysql")

	if err := ss.client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		return err
	}
	return nil
}

func (ss *ORMRepository) Close() error {
	if ss.client == nil {
		return nil
	}
	return ss.client.Close()
}

func (ss *ORMRepository) Create(ctx context.Context) error {
	ss.client.Pet.Delete().Exec(ctx)
	ss.client.User.Delete().Exec(ctx)

	p1, err := ss.client.Pet.
		Create().
		SetName("pet1").
		SetDesc("petdesc1").
		Save(ctx)
	if err != nil {
		return err
	}
	p2, err := ss.client.Pet.
		Create().
		SetName("pet2").
		SetDesc("petdesc2").
		Save(ctx)
	if err != nil {
		return err
	}
	p3, err := ss.client.Pet.
		Create().
		SetName("pet3").
		SetDesc("petdesc3").
		Save(ctx)
	if err != nil {
		return err
	}
	p4, err := ss.client.Pet.
		Create().
		SetName("pet4").
		SetDesc("petdesc4").
		Save(ctx)
	if err != nil {
		return err
	}

	_, err = ss.client.User.
		Create().
		SetName("test1").
		SetAge(10).
		AddPets(p1, p2).
		Save(ctx)
	if err != nil {
		return err
	}
	users := []*entgen.UserCreate{
		ss.client.User.
			Create().
			SetName("testb1").
			SetAge(11).AddPets(p3, p4),
		ss.client.User.
			Create().
			SetName("testb2").
			SetAge(21),
	}
	_, err = ss.client.User.CreateBulk(users...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (ss *ORMRepository) Get(ctx context.Context, name string) (interface{}, error) {
	//i, _ := strconv.Atoi(id)
	//return ss.client.User.Get(ctx, i)
	return ss.client.User.
		Query().
		Where(user.Name(name)).
		WithPets().
		First(ctx)
}

func (ss *ORMRepository) List(ctx context.Context) (interface{}, error) {
	return ss.client.User.
		Query().
		Where(user.Name("testb1")).
		WithPets().
		All(ctx)
}

func (ss *ORMRepository) ListX(ctx context.Context, minAge int) ([]*entgen.User, error) {
	var users []*entgen.User
	err := ss.dbx.Select(&users, "select * from users where age > ?", minAge)
	return users, err
}
