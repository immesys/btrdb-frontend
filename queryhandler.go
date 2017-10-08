//+build ignore

package main

import (
	"net/http"
	"regexp"
	"strings"
)

const SELECT = "SELECT"
const FROM = "FROM"

func tokenize(q string) []string {
	rv := []string{}
	parts := strings.Split(q, " ")
	for _, p := range parts {
		rv = append(rv, strings.TrimSpace(p))
	}
	return rv
}
func queryhandler(w http.ResponseWriter, req *http.Request) {
	//Actual query comes from "q" parameter
	q, ok := req.URL.Query()["q"]
	if !ok || len(q) != 1 {
		w.WriteHeader(404)
		w.Write([]byte("no Q parameter"))
	}
	tokens := tokenize(q[0])
	if strings.EqualFold(tokens[0], SELECT) {
		doselect(q[0], w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("I didn't get the query"))
	}
}

func doselect(q string, w http.ResponseWriter) {

	//Syntax:
	// SELECT field,field,field FROM measurement
	// SELECT * FROM measurement
	// SELECT field FROM measurement...
	// SELECT field FROM measurement... WHERE annkey=annval,annkey=annval
	// SELECT MAX(field) etc
	// SELECT MIN(field) etc
	// SELECT MEAN
	die := func(msg string) {
		w.WriteHeader(500)
		w.Write([]byte(msg))
	}
	exp := regexp.MustCompile(`(?i)SELECT\s+(?P<field>[^\s]+)\s+FROM\s+(?P<measurement>[^\s]+)(\s+WITH\s+(?P<with>[^\s]+))?(\s+AFTER\s+(?P<after>[^\s]+))?(\s+BEFORE\s+(?P<before>[^\s]+))?(\s+WINDOW\s+(?P<window>[^\s]+))?`)
	match := exp.FindStringSubmatch(q)
	result := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	allfields := true
	var fields []string
	if tk[1] != "*" {
		fields = strings.Split(tk[1], ",")
		allfields = false
	}
	if !strings.EqualFold(tk[2], FROM) {
		die("bad query")
		return
	}
	collection := tk[3]
	prefix := false
	if strings.HasSuffix(collection, "...") {
		collection = strings.TrimSuffix(collection, "...")
		prefix = true
	}

	_ = fields
	_ = prefix
	_ = collection
	_ = allfields
}
