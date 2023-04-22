package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/Polidoro-root/go-inventory-management/configs"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupDatabase(t *testing.T) *sql.DB {
	configs := configs.LoadTestEnv(t)

	ctx := context.Background()

	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		t.Fatal("unable to get the current filename")
	}

	dirname := filepath.Dir(filename)

	migrationsDir := filepath.Join(dirname, "..", "..", "sql", "migrations")

	dir, err := os.ReadDir(migrationsDir)

	if err != nil {
		t.Fatal(err)
	}

	var migrations []string

	for _, file := range dir {
		if strings.HasSuffix(file.Name(), "up.sql") {
			migrations = append(migrations, filepath.Join(migrationsDir, file.Name()))
		}
	}

	var initScripts []testcontainers.ContainerFile

	for _, script := range migrations {
		initScripts = append(initScripts, testcontainers.ContainerFile{
			HostFilePath:      script,
			ContainerFilePath: "/docker-entrypoint-initdb.d/" + filepath.Base(script),
			FileMode:          0755,
		})
	}

	req := testcontainers.ContainerRequest{
		Image: "postgres:15-alpine",
		Env: map[string]string{
			"POSTGRES_DB":       configs.DBName,
			"POSTGRES_USER":     configs.DBUser,
			"POSTGRES_PASSWORD": configs.DBPassword,
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort("5432/tcp"),
			wait.ForLog("database system is ready to accept connections"),
		),
		Files: initScripts,
	}

	dbContainer, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := dbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	})

	host, err := dbContainer.Host(ctx)

	if err != nil {
		t.Fatal(err)
	}

	port, err := dbContainer.MappedPort(ctx, "5432")

	if err != nil {
		t.Fatal(err)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s %s",
		configs.DBUser,
		configs.DBPassword,
		host,
		port.Port(),
		configs.DBName,
		configs.DBOptions,
	)

	db, err := sql.Open(
		configs.DBDriver,
		connStr,
	)

	if err != nil {
		t.Fatal(err)
	}

	if db == nil {
		t.Fatal("db must be initiated")
	}

	return db
}
