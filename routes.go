package main

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", data, nil)
}

func addDiagnosis(r *http.Request) *web.Response {
	switch r.Method {
	case http.MethodGet:
		return web.HTML(http.StatusOK, html, "add-diagnosis.html", nil, nil)
	case http.MethodPost:
		r.ParseForm()
		var newDiagnosedDate = r.Form.Get("date")
		var allergic = r.Form.Get("allergic") == "on"
		updateAllergies(newDiagnosedDate, allergic)
		return web.HTML(http.StatusOK, html, "main-screen.html", data, nil)
	}
	return web.Empty(http.StatusNotImplemented)
}
