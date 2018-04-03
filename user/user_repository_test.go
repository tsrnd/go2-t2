package user

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tsrnd/trainning/shared/repository"
	"github.com/tsrnd/trainning/shared/test"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	uuid        = "12345678-1234-1234-1234-112233445566"
	uuidInvalid = "12345678-1234-1234-1234-112233445566-9999"
	qrSelect    = `SELECT * FROM "user_app" WHERE "user_app".deleted_at IS NULL AND (("uuid" = $1)) ORDER BY "user_app"."id_user_app" ASC LIMIT 1`
	qrInsert    = `INSERT INTO "user_app" ("uuid","user_name","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5) RETURNING "user_app"."id_user_app"`
)

func TestFindOrCreateSuccessWhenUserIsExist(t *testing.T) {
	loggerConfig := test.NewLogger()
	bh := &repository.BaseRepository{Logger: loggerConfig}
	mock, db := test.NewSQLMock()
	r := NewRepository(bh, db, db, nil)

	rows := sqlmock.NewRows([]string{"id_user_app", "uuid", "user_name", "created_at", "updated_at"}).
		AddRow(1, "12345678-1234-1234-1234-012345678901", "", time.Now(), time.Now())
	mock.ExpectQuery(test.FixedFullRe(qrSelect)).WillReturnRows(rows)

	user, err := r.FindOrCreate(uuid)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindOrCreateSuccessWhenUserIsNotExist(t *testing.T) {
	loggerConfig := test.NewLogger()
	bh := &repository.BaseRepository{Logger: loggerConfig}
	mock, db := test.NewSQLMock()
	r := NewRepository(bh, db, db, nil)

	rSelectResult := sqlmock.NewRows([]string{"id_user_app"})
	rInsertResult := sqlmock.NewRows([]string{"id_user_app"}).AddRow(1)
	mock.ExpectQuery(test.FixedFullRe(qrSelect)).WillReturnRows(rSelectResult)
	mock.ExpectQuery(test.FixedFullRe(qrInsert)).WillReturnRows(rInsertResult)

	user, err := r.FindOrCreate(uuid)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindOrCreateFailWhenUUIDCharaterMoreThan36(t *testing.T) {
	loggerConfig := test.NewLogger()
	bh := &repository.BaseRepository{Logger: loggerConfig}
	mock, db := test.NewSQLMock()
	r := NewRepository(bh, db, db, nil)

	rSelectResult := sqlmock.NewRows([]string{"id_user_app"})
	mock.ExpectQuery(test.FixedFullRe(qrSelect)).WillReturnRows(rSelectResult)
	mock.ExpectQuery(test.FixedFullRe(qrInsert)).WillReturnError(fmt.Errorf("Some error when insert data with uuid charater more than 36"))

	_, err := r.FindOrCreate(uuidInvalid)
	assert.Error(t, err)
}
