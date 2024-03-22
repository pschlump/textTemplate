package textTemplate

import (
	"fmt"
	"strings"
	"testing"

	parse "github.com/pschlump/textTemplate/parse"
)

// PJS test code - Wed Mar 20 07:04:47 MDT 2019
func TestAvailableTemplates(t *testing.T) {
	tmpl := Must(New("tmpl").Parse(`{{template "tmpl" .}}`))
	av := tmpl.AvailableTemplates()
	// fmt.Printf("PJS1 - %s\n", av)
	want := []string{"tmpl"}
	if len(av) != len(want) {
		t.Errorf("got %d; want %d -- unexpected number of templates returned", len(av), len(want))
	} else {
		for i := 0; i < len(av); i++ {
			if av[i] != want[i] {
				t.Errorf("got %q; want %q", av[i], want[i])
			}
		}
	}
}

// PJS test code - Fri Mar 22 07:15:44 MDT 2024
// Test missing value, reurnd as 'bob'
func TestSetEmptyFunc(t *testing.T) {
	called := false
	cb := func(def string, loc parse.Pos, lf string) {
		called = true
		fmt.Printf("%s %d %s\n", def, loc, lf)
	}
	tmpl := Must(New("tmpl").SetEmpty("bob", true).SetEmptyFunc(cb).Parse(`{{define "sub"}} -->>{{.missing}}<<-- {{end}}`))
	fmt.Printf("tmpl.emptyDataValue ->%s<-\n", tmpl.emptyDataValue)
	// emptyDataValue string // PJS value to use when missing value is found
	// errOnEmpty     bool   // PJS if true(default) reports errors on missing values
	// errOnEmptyFunc func(def string, pos parse.Pos, lf string) // PJS if not null, then on error call this function		-- Call funciton with name of template, lf is line/file in source called from
	data := make(map[string]string)
	var buf strings.Builder
	for _, tt := range tmpl.AvailableTemplates() {
		tmpl.Lookup(tt).SetEmpty("bob", true).SetEmptyFunc(cb)
	}
	err := tmpl.ExecuteTemplate(&buf, "sub", data)
	if !called {
		t.Errorf("Falied to call callback func\n")
	}
	if buf.String() != " -->>bob<<-- " {
		t.Errorf("Incorrect template substitution")
	}
	fmt.Printf("->%s<- err=%s\n", buf.String(), err)
}
