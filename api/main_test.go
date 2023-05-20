package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// HTTPリクエストを作成
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// HTTPレスポンスのモックを作成
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	// handlerを使ってreq（リクエスト）を処理して、レスポンスをrrに書き込む
	handler.ServeHTTP(rr, req)

	// レスポンスのステータスコードが正しいか
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, rr.Code)
	}

	// レスポンスのボディの内容が正しいかチェック
	expected := "Hello World!!!"
	if rr.Body.String() != expected {
		t.Errorf("expected body %v but got %v", expected, rr.Body.String())
	}
}
