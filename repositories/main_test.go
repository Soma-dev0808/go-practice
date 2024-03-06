package repositories_test

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

var (
	dbUser 		= "docker"
	dbPassword 	= "docker"
	dbDatabase 	= "sampledb"
	dbConn 		= fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// テスト前のテーブル削除
func cleanupDB() error {
	var stderr bytes.Buffer
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-P", "3307", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("MySQL Error:", stderr.String())
		return err
	}
	return nil
}

// テスト前のテーブル作成＆データ挿入
func setupTestData() error {
	var stderr bytes.Buffer
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-P", "3307", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setupDB.sql")
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("MySQL Error:", stderr.String())
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}

	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup error")
		return err
	}

	if err := setupTestData(); err != nil {
		fmt.Println("setup error")
		return err
	}
	return nil
}

// 前テスト共通の後処理
func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}