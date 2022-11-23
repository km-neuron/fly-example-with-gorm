package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rg-app/app/model"
)

func (api *API) IndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := api.BaseViewPath()

	fmt.Println(r.Host)

	var client = &http.Client{}
	// https://rg-app.fly.dev/api/student/read
	req, err := http.NewRequest("GET", "https://rg-app.fly.dev/api/student/read", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Membaca response body dari server
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var student []model.StudentView
	err = json.Unmarshal(body, &student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println("Data read: ", student)

	data := map[string]interface{}{
		"Region": os.Getenv("FLY_REGION"),
		"Data":   student,
	}

	tmpl.ExecuteTemplate(w, "index.html.tmpl", data)
}
