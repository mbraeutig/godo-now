package fauna

import (
	"log"
	"os"

	// f "github.com/mbraeutig/faunadb-go/v2/faunadb"
	f "github.com/mbraeutig/faunadb-v2/faunadb"
)

var faunaSecret = os.Getenv("FAUNA_ROOT_KEY")
var adminClient *f.FaunaClient

func newAdminClient(faunaSecret string) *f.FaunaClient {
	return f.NewFaunaClient(faunaSecret)
}

func createKey(client *f.FaunaClient, name string) f.Value {
	key, err := client.Query(
		f.CreateKey(
			f.Obj{"database": f.Database(name), "role": "server"},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

func getSecret(key f.Value) (secret string) {
	err := key.At(f.ObjKey("secret")).Get(&secret)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
	}
	return
}

func newFaunaClient(client *f.FaunaClient, secret string) *f.FaunaClient {
	return client.NewSessionClient(secret)
}

func init() {
	if faunaSecret == "" {
		panic("FAUNA_ROOT_KEY environment variable must be specified")
	}

	// Instantiate the admin client
	adminClient = newAdminClient(faunaSecret)
}

func Client() *f.FaunaClient {
	// Accessing the database
	key := createKey(adminClient, "GoDo")
	// Instantiate the Fauna client
	return newFaunaClient(adminClient, getSecret(key))
}
