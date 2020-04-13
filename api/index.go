package api

import (
	"net/http"
	"os"

	f "github.com/mbraeutig/faunadb-go/tree/release-2.11.0-semverfix/v2"

)

func Handler(w http.ResponseWriter, r *http.Request) {
	secret = os.Getenv("FAUNADB_SECRET")
	adminClient := newAdminClient(secret)
	io.WriteString(w, fmt.Sprintf("Client: %v", adminClient)
}

func newAdminClient(faunaSecret string) *f.FaunaClient {
	return f.NewFaunaClient(faunaSecret)
}
