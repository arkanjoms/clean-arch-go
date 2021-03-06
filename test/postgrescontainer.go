package test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"log"
	"os"
)

var (
	pgC testcontainers.Container
)

const (
	testDbHost          = "localhost"
	testDbName          = "testdb"
	testDbUser          = "test"
	testDbPassword      = "test"
	testPostgresSvcPort = "5432"
)

func ContainerDBStart(basePath string) (context.Context, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13-alpine",
		ExposedPorts: []string{testPostgresSvcPort},
		Env: map[string]string{
			"POSTGRES_DB":       testDbName,
			"POSTGRES_USER":     testDbUser,
			"POSTGRES_PASSWORD": testDbPassword,
		},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort(testPostgresSvcPort),
			wait.ForSQL(testPostgresSvcPort, "postgres", func(port nat.Port) string {
				return fmt.Sprintf(
					"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
					testDbHost,
					port.Port(),
					testDbUser,
					testDbPassword,
					testDbName,
				)
			}),
		),
	}
	var err error
	pgC, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	testDbPort, _ := pgC.MappedPort(ctx, testPostgresSvcPort)
	log.Printf("Database started as port: %s", testDbPort)

	setEnv(testDbPort, basePath)
	return ctx, nil
}

func setEnv(testDbPort nat.Port, basePath string) {
	_ = os.Setenv("DB_HOST", testDbHost)
	_ = os.Setenv("DB_PORT", testDbPort.Port())
	_ = os.Setenv("DB_NAME", testDbName)
	_ = os.Setenv("DB_USERNAME", testDbUser)
	_ = os.Setenv("DB_PASSWORD", testDbPassword)
	_ = os.Setenv("MIGRATION_SOURCE_URL", getMigrationsPath(basePath))
}

func getMigrationsPath(basePath string) string {
	return fmt.Sprintf("%s/migrations", basePath)
}

func ContainerDBStop(ctx context.Context) {
	err := pgC.Terminate(ctx)
	if err != nil {
		log.Println(err.Error())
	}
}

func DatasetTest(db *sql.DB, basePath string, clearDataFileName string, scripts ...string) error {
	if clearDataFileName != "" {
		err := cleanDatabase(db, basePath, clearDataFileName)
		if err != nil {
			return err
		}
	}

	if scripts != nil {
		for _, s := range scripts {
			script, err := loadScript(basePath, s)
			if err != nil {
				return fmt.Errorf("could not load script: %w", err)
			}
			err = execScript(db, script)
			if err != nil {
				return fmt.Errorf("could not execute script: %w", err)
			}
		}
	}
	return nil
}

func cleanDatabase(db *sql.DB, basePath string, clearDataFileName string) error {
	logrus.Info("cleaning database")
	script, err := clearDataScript(basePath, clearDataFileName)
	if err != nil {
		return err
	}
	err = execScript(db, script)
	if err != nil {
		return err
	}
	logrus.Info("database cleaning successfully")
	return nil
}

func execScript(db *sql.DB, script string) error {
	_, err := db.Exec(script)
	if err != nil {
		return err
	}
	return nil
}

func clearDataScript(basePath string, fileName string) (string, error) {
	return loadScript(basePath, fileName)
}

func loadScript(basePath string, fileName string) (string, error) {
	filePath := fmt.Sprintf("%s/test/sql/%s", basePath, fileName)
	c, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(c), nil
}
