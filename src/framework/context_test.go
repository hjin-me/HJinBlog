package fw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestFwContextOutput(t *testing.T) {

	testStr := "hello world"
	testContentType := "text/plain"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawCtx, _ := context.WithTimeout(context.Background(), time.Second)

		ctx := WithHttp(rawCtx, w, r, map[string]string{})
		ctx.Output(testStr, testContentType)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	ct, ok := res.Header["Content-Type"]
	if !ok {
		t.Error("Response no Content-Type")
	}
	if ct[0] != testContentType {
		t.Error("Content Type not", testContentType)
	}
	if string(greeting) != testStr {
		t.Error("response error")
	}

	fmt.Printf("%s", greeting)

}

func TestFwContextJSON(t *testing.T) {

	type TD struct {
		Hello string
		Num   []int
	}
	testData := TD{"world", []int{3, 2, 1}}
	testContentType := "application/json"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawCtx, _ := context.WithTimeout(context.Background(), time.Second)

		ctx := WithHttp(rawCtx, w, r, map[string]string{})
		ctx.Json(testData)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	ct, ok := res.Header["Content-Type"]
	if !ok {
		t.Error("Response no Content-Type")
	}
	if ct[0] != testContentType {
		t.Error("Content Type not", testContentType)
	}
	tstr, err := json.Marshal(testData)
	if err != nil {
		t.Fatal(err)
	}
	if string(greeting) != string(tstr) {
		t.Error("response error")
	}

	fmt.Printf("%s", greeting)

}

func TestFwContextTpl5xx(t *testing.T) {

	type TD struct {
		Hello string
		Num   []int
	}
	testData := TD{"world", []int{3, 2, 1}}
	testHtml := "Template file error"
	testContentType := "text/html"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawCtx, _ := context.WithTimeout(context.Background(), time.Second)

		ctx := WithHttp(rawCtx, w, r, map[string]string{})
		ctx.Tpl("abc.tpl", testData)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	ct, ok := res.Header["Content-Type"]
	if !ok {
		t.Error("Response no Content-Type")
	}
	if ct[0] != testContentType {
		t.Error("Content Type not", testContentType)
	}
	////tstr, err := json.Marshal(testData)
	////if err != nil {
	////	t.Fatal(err)
	////}
	if string(greeting) != string(testHtml) {
		t.Error("response error", string(greeting))
	}

	fmt.Printf("%s", greeting)

}

func TestFwContextTpl(t *testing.T) {

	type TD struct {
		Hello string
		Num   []int
	}
	testData := TD{"world", []int{3, 2, 1}}
	testHtml := "Template file error"
	testContentType := "text/html"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawCtx, _ := context.WithTimeout(context.Background(), time.Second)

		ctx := WithHttp(rawCtx, w, r, map[string]string{})
		ctx.Tpl("abc.tpl", testData)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	ct, ok := res.Header["Content-Type"]
	if !ok {
		t.Error("Response no Content-Type")
	}
	if ct[0] != testContentType {
		t.Error("Content Type not", testContentType)
	}
	////tstr, err := json.Marshal(testData)
	////if err != nil {
	////	t.Fatal(err)
	////}
	if string(greeting) != string(testHtml) {
		t.Error("response error", string(greeting))
	}

	fmt.Printf("%s", greeting)

}
