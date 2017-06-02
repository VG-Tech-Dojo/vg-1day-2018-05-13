package main

import (
	"fmt"
	"io/ioutil"
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
		t.Fatalf("failed to get response: %s", err)
	}
	defer resp.Body.Close()

	if expected := 200; resp.StatusCode != expected {
		t.Fatalf("status code expected %d but not, actual %d", expected, resp.StatusCode)
	}
}

func TestAPIがpingに応答する(t *testing.T) {
	resp, err := http.Get(ts.URL + "/api/ping")
	if err != nil {
		t.Fatalf("failed to get response: %s", err)
	}
	defer resp.Body.Close()

	if expected := 200; resp.StatusCode != expected {
		t.Fatalf("status code expected %d but not, actual %d", expected, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read http response, %s", err)
	}

	if expected := "pong"; string(b) != expected {
		t.Fatalf("response body expected %s but not, actual %s", expected, string(b))
	}
}

func TestAPIがメッセージを全て返す(t *testing.T) {}

func TestAPIが指定したIDのメッセージを返す(t *testing.T) {}

func TestAPIが新しいメッセージを作成する(t *testing.T) {}

func TestAPIが指定したIDのメッセージを更新する(t *testing.T) {}

func TestAPIが指定したIDのメッセージを削除するする(t *testing.T) {}
