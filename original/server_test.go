package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const (
	dbconf = "dbconfig.yml"
	env    = "test"
)

// テスト毎に立ち上げるのは時間がかかるのでテスト全体で共有して使いまわします
var ts *httptest.Server

func TestMain(m *testing.M) {
	os.Exit(realMain(m))
}

func realMain(m *testing.M) int {
	ts = defaultTestServer()
	defer func() {
		ts.Close()
	}()
	return m.Run()
}

func defaultTestServer() *httptest.Server {
	s := NewServer()
	if err := s.Init(dbconf, env); err != nil {
		panic(fmt.Sprintf("failed to init server: %v", err))
	}

	return httptest.NewServer(s.Engine)
}

func TestTopページが200を返す(t *testing.T) {
	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatalf("get index page failed: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("fail to get index page: status code is %d", resp.StatusCode)
	}
}
