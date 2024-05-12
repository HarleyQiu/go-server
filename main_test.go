package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandleForm 测试 handleForm 函数是否能正确处理 POST 请求并返回预期的响应。
func TestHandleForm(t *testing.T) {
	// 创建一个模拟的 POST 请求，包含表单数据
	form := strings.NewReader("name=John&address=123BakerStreet")
	req, err := http.NewRequest("POST", pathForm, form)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 创建响应记录器用于捕获响应
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleForm)

	// 调用 handleForm 处理函数
	handler.ServeHTTP(rr, req)

	// 检查返回的状态码是否为 http.StatusOK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handleForm returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 检查返回的正文是否符合预期
	expected := fmt.Sprintf("%s\nName = John\nAddress = 123BakerStreet\n", successPostMsg)
	if rr.Body.String() != expected {
		t.Errorf("handleForm returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHandleHello 测试 handleHello 函数是否能正确处理 GET 请求并返回预期的响应。
func TestHandleHello(t *testing.T) {
	req, err := http.NewRequest("GET", pathHello, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHello)

	handler.ServeHTTP(rr, req)

	// 检查返回的状态码是否为 http.StatusOK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handleHello returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 检查返回的正文是否为 "hello!\n"
	expected := "hello!\n"
	if rr.Body.String() != expected {
		t.Errorf("handleHello returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHandleHelloNotFound 测试当访问不存在的路径时 handleHello 是否返回 404。
func TestHandleHelloNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleHello)

	handler.ServeHTTP(rr, req)

	// 检查返回的状态码是否为 http.StatusNotFound
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handleHello with invalid path returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
