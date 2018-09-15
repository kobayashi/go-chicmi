package chicmi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvent(t *testing.T) {
	const f = "./json/event.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/events_get/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		b, err := getTestFile(f)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", f)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient)
	resp, _, err := client.GetEvent(&EventSearchParams{
		EventID: 16597,
	})
	assert.Equal(t, "", err.Error())
	// Event
	assert.Equal(t, "Global Fashions", resp.Values.AddressBusinessName)
	assert.Equal(t, "49 West 38th Street", resp.Values.AddressStreet1)
	assert.Equal(t, "Main Floor and 2nd Floor", resp.Values.AddressStreet2)
	assert.Equal(t, "New York", resp.Values.AddressCity)
	assert.Equal(t, "10018", resp.Values.AddressZip)
	assert.Equal(t, "2085", resp.Values.CountryID)
	assert.Equal(t, "16597", resp.Values.ID)
	assert.Equal(t, "Global Fashions Group Sample Sale", resp.Values.EventNameEn)
	// Meta
	assert.Equal(t, "1.0", resp.Meta.Version)
	assert.Equal(t, 1, resp.Meta.Step)
	assert.Equal(t, "en", resp.Meta.Language)
	assert.Equal(t, true, resp.Meta.Remote)
}
