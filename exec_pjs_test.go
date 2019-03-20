package textTemplate

import (
	"testing"
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
