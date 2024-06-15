package metadatastore

import (
	"context"
	"database/sql"
	"inx/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Store struct {
	dbUrl string
	db    *sql.DB
}

func NewStore(dbUrl string) *Store {
	return &Store{
		dbUrl: dbUrl,
	}
}

func (s *Store) CreateDB() error {
	db, err := sql.Open("postgres", s.dbUrl)
	if err != nil {
		return err
	}
	s.db = db
	return nil
}
func (s *Store) UpdateOffset(id, o int) error {
	ctx := context.Background()

	var co models.Consumeroffset
	co.ConsumerID = id
	co.ConsumerOffset = null.IntFrom(o)
	err := co.Upsert(ctx, s.db, true /* updateOnConflict */, []string{"consumer_id"}, boil.Whitelist("consumer_offset"), boil.Whitelist("consumer_id", "consumer_offset"))
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetCurrentOffset(id int) (int, error) {
	ctx := context.Background()
	co, err := models.Consumeroffsets(qm.Where("consumer_id=?", id)).One(ctx, s.db)
	if err != nil {
		return 0, err
	}
	return co.ConsumerID, nil
}
