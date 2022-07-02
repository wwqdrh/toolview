package main

// package p
// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"
// )

// func TestConf(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r, _ := http.NewRequest("GET", "/api/etcd/conf/status", nil)
// 	Engine.ServeHTTP(w, r)
// 	fmt.Println(w.Body.String())

// 	w = httptest.NewRecorder()
// 	body, _ := json.Marshal(map[string]interface{}{
// 		"endpoints": []string{"  ", "127.0.0.1:2379"},
// 		"username":  "root",
// 		"password":  "123456",
// 	})
// 	r, _ = http.NewRequest("POST", "/api/etcd/conf/update", bytes.NewReader(body))
// 	Engine.ServeHTTP(w, r)
// 	fmt.Println(w.Body.String())

// 	w = httptest.NewRecorder()
// 	r, _ = http.NewRequest("GET", "/api/etcd/conf/status", nil)
// 	Engine.ServeHTTP(w, r)
// 	fmt.Println(w.Body.String())

// 	r, _ = http.NewRequest("POST", "/api/etcd/conf/update", bytes.NewReader(body))
// 	Engine.ServeHTTP(w, r)
// 	fmt.Println(w.Body.String())

// 	w = httptest.NewRecorder()
// 	r, _ = http.NewRequest("GET", "/api/etcd/conf/verify", nil)
// 	Engine.ServeHTTP(w, r)
// 	fmt.Println(w.Body.String())

// 	if os.Getenv("LOCAL") != "" {
// 		// 如果环境中有etcd且连接成功 那么做后面的测试
// 		w = httptest.NewRecorder()
// 		r, _ = http.NewRequest("GET", "/api/etcd/key/list?prefix=%2F", nil)
// 		Engine.ServeHTTP(w, r)
// 		fmt.Println(w.Body.String())

// 		w = httptest.NewRecorder()
// 		r, _ = http.NewRequest("GET", "/api/etcd/key/list", nil)
// 		Engine.ServeHTTP(w, r)
// 		fmt.Println(w.Body.String())
// 	}
// }
//
