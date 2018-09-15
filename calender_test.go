package chicmi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	const f = "./json/calender.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/calendar_in_city/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		b, err := getTestFile(f)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", f)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient)
	resp, _, err := client.Search(&SearchParams{
		City:         "new-york",
		Types:        "sample-sales",
		FeaturedOnly: "yes",
		DateFrom:     "2018/08/26",
		DateTo:       "2017/09/03",
		MaxResults:   1,
		Source:       "embed",
	})
	assert.Len(t, resp.Values.Events, 1)
	assert.Equal(t, "", err.Error())
	// Events
	assert.Equal(t, "6 Prince Street", resp.Values.Events[0].AddressBusinessName)
	assert.Equal(t, "6 Prince Street", resp.Values.Events[0].AddressStreet1)
	assert.Equal(t, "", resp.Values.Events[0].AddressStreet2)
	assert.Equal(t, "New York", resp.Values.Events[0].AddressCity)
	assert.Equal(t, "10012", resp.Values.Events[0].AddressZip)
	assert.Equal(t, "2085", resp.Values.Events[0].CountryID)
	assert.Equal(t, "15259", resp.Values.Events[0].EventID)
	assert.Equal(t, "United Nude Pop Up Sample Sale", resp.Values.Events[0].EventNameEn)
	// Meta
	assert.Equal(t, "1.0", resp.Meta.Version)
	assert.Equal(t, 4, resp.Meta.Step)
	assert.Equal(t, "en", resp.Meta.Language)
	assert.Equal(t, true, resp.Meta.Remote)
}
