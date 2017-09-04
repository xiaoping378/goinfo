package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_PostValues(t *testing.T) {

	ariti := &Ariticle{Title: "22", Desc: "22", Content: "22"}

	buf, _ := json.Marshal(ariti)
	body := bytes.NewBuffer(buf)

	resp, err := http.Post("http://localhost:8001/ariticle", "application/json", body)

	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("response Status: %s", resp.Status)
		defer resp.Body.Close()
	}
}
