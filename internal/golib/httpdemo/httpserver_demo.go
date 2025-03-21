package httpdemo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type routerManger struct{}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonResp struct {
	Items []Person `json:"items,omitempty"`
	Code  int      `json:"code"`
	Msg   string   `json:"msg"`
}

var Persons PersonResp

func (r *routerManger) HandelePerson(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.URL.Host, req.URL.Path, req.URL.Query())
	if req.Method == "POST" {
		r.PostPersonInfo(w, req)
	} else if req.Method == "GET" {
		r.GetPersonInfo(w, req)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorContent := PersonResp{Code: int(http.StatusNoContent), Msg: "Cannot support this method. "}
		if data, err := json.Marshal(errorContent); err == nil {
			fmt.Println(string(data))
			w.Write(data)
		}
	}
}

// GetPersonInfo 获取创建的用户
func (r *routerManger) GetPersonInfo(w http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.Header)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Println(req.URL.Host, req.URL.Path, req.URL.Query())

	if data, err := json.Marshal(Persons); err == nil {
		w.Write(data)
	}
}

// PostPersonInfo 创建用户
func (r *routerManger) PostPersonInfo(w http.ResponseWriter, req *http.Request) {
	var person Person
	data, _ := io.ReadAll(req.Body)
	w.Header().Set("Content-Type", "application/json")

	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContent := PersonResp{Code: int(http.StatusInternalServerError), Msg: "Body is not empty."}
		if errData, err := json.Marshal(errorContent); err == nil {
			fmt.Println(string(errData))
			w.Write(errData)
			return
		}
	}

	if err := json.Unmarshal(data, &person); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorContent := PersonResp{Code: int(http.StatusBadRequest), Msg: "illegal json body."}
		if errData, err := json.Marshal(errorContent); err == nil {
			w.Write(errData)
		}
		return
	}

	Persons.Items = append(Persons.Items, person)
	successContent := PersonResp{Code: int(http.StatusOK), Msg: "success."}
	if successData, err := json.Marshal(successContent); err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Println(string(successData))
		w.Write(successData)
		return
	}
}

func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, req.URL.Path, req.URL.RawQuery)
}

func HandleHttpRequest() {
	port := 8989
	r := routerManger{}

	mux := http.NewServeMux()
	mux.HandleFunc("/person", r.HandelePerson)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, pattern := mux.Handler(r)
		fmt.Println(pattern)
		if pattern == "" {
			NotFoundHandler(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	})

	fmt.Printf("Server is running on http://0.0.0.0:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), handler))
}
