package testutils

import (
	"context"
	"database/sql"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
	"github.com/xmtp/xmtpd/pkg/db/queries"
	"github.com/xmtp/xmtpd/pkg/migrations"
)

const (
	LocalTestDBDSNPrefix = "postgres://postgres:xmtp@localhost:8765"
	LocalTestDBDSNSuffix = "?sslmode=disable"
)

func GetCallerName(depth int) string {
	pc, _, _, ok := runtime.Caller(depth)
	if !ok {
		return "unknown"
	}
	nameFull := runtime.FuncForPC(pc).Name()
	nameEnd := filepath.Ext(nameFull)
	name := strings.TrimPrefix(nameEnd, ".")
	return strings.ToLower(name)
}

func openDB(t testing.TB, dsn string) (*sql.DB, string, func()) {
	config, err := pgx.ParseConfig(dsn)
	require.NoError(t, err)
	db := stdlib.OpenDB(*config)
	return db, dsn, func() {
		err := db.Close()
		require.NoError(t, err)
	}
}

func newCtlDB(t testing.TB) (*sql.DB, string, func()) {
	return openDB(t, LocalTestDBDSNPrefix+LocalTestDBDSNSuffix)
}

func newInstanceDB(t testing.TB, ctx context.Context, ctlDB *sql.DB) (*sql.DB, string, func()) {
	dbName := "test_" + GetCallerName(3) + "_" + RandomStringLower(12)
	t.Logf("creating database %s ...", dbName)
	_, err := ctlDB.Exec("CREATE DATABASE " + dbName)
	require.NoError(t, err)

	db, dsn, cleanup := openDB(t, LocalTestDBDSNPrefix+"/"+dbName+LocalTestDBDSNSuffix)
	require.NoError(t, migrations.Migrate(ctx, db))

	return db, dsn, func() {
		cleanup()
		_, err = ctlDB.Exec("DROP DATABASE " + dbName)
		require.NoError(t, err)
	}
}

func NewDB(t *testing.T, ctx context.Context) (*sql.DB, string, func()) {
	ctlDB, _, ctlCleanup := newCtlDB(t)
	db, dsn, cleanup := newInstanceDB(t, ctx, ctlDB)

	return db, dsn, func() {
		cleanup()
		ctlCleanup()
	}
}

func NewDBs(t *testing.T, ctx context.Context, count int) ([]*sql.DB, func()) {
	ctlDB, _, ctlCleanup := newCtlDB(t)
	dbs := []*sql.DB{}
	cleanups := []func(){}

	for i := 0; i < count; i++ {
		db, _, cleanup := newInstanceDB(t, ctx, ctlDB)
		dbs = append(dbs, db)
		cleanups = append(cleanups, cleanup)
	}

	return dbs, func() {
		for i := 0; i < count; i++ {
			cleanups[i]()
		}
		ctlCleanup()
	}
}

func InsertGatewayEnvelopes(
	t *testing.T,
	db *sql.DB,
	rows []queries.InsertGatewayEnvelopeParams,
	notifyChan ...chan bool,
) {
	q := queries.New(db)
	for _, row := range rows {
		inserted, err := q.InsertGatewayEnvelope(context.Background(), row)
		require.Equal(t, int64(1), inserted)
		require.NoError(t, err)

		if len(notifyChan) > 0 {
			select {
			case notifyChan[0] <- true:
			default:
			}
		}
	}
}

func CreatePayer(t *testing.T, db *sql.DB, address ...string) int32 {
	q := queries.New(db)
	var payerAddress string
	if len(address) > 0 {
		payerAddress = address[0]
	} else {
		payerAddress = RandomString(42)
	}

	id, err := q.FindOrCreatePayer(context.Background(), payerAddress)
	require.NoError(t, err)

	return id
}
