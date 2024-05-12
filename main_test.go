package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestFormHandler 检查 formHandler 函数是否能正确处理带有表单数据的 POST 请求。
func TestFormHandler(t *testing.T) {
	// 创建一个新的 POST 请求，并包含表单数据
	req, err := http.NewRequest("POST", "/form", strings.NewReader("name=John&address=123Street"))
	if err != nil {
		t.Fatal(err) // 如果创建请求失败，则结束测试
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	rr := httptest.NewRecorder()             // 创建一个 HTTP 记录器用于保存响应
	handler := http.HandlerFunc(formHandler) // 转换 formHandler 为 HTTP 处理函数

	handler.ServeHTTP(rr, req) // 使用处理函数处理请求

	// 检查响应状态码是否为 http.StatusOK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 预期响应正文
	expected := "POST request successfulName  = John Address 123Street"
	// 检查响应正文是否符合预期
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHelloHandler 检查 helloHandler 函数对 GET 请求的处理是否正确。
func TestHelloHandler(t *testing.T) {
	// 创建一个新的 GET 请求
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	// 检查响应状态码是否为 http.StatusOK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// 预期响应正文
	expected := "hello!\n"
	// 检查响应正文是否符合预期
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// TestHelloHandlerReturns404ForNonExistentRoute 测试当访问不存在的路由时 helloHandler 是否返回 404 状态码。
func TestHelloHandlerReturns404ForNonExistentRoute(t *testing.T) {
	// 创建一个指向不存在路由的 GET 请求
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	// 检查响应状态码是否为 http.StatusNotFound
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

// TestHelloHandlerReturns404ForNonGetMethods 测试当使用非 GET 方法访问 /hello 路径时，helloHandler 是否返回 404 状态码。
func TestHelloHandlerReturns404ForNonGetMethods(t *testing.T) {
	// 创建一个 POST 请求指向 /hello 路径
	req, err := http.NewRequest("POST", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	// 检查响应状态码是否为 http.StatusNotFound
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
