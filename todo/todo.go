package todo

import (
	"log"

	// f "github.com/mbraeutig/faunadb-go/v2/faunadb"
	f "github.com/mbraeutig/faunadb-v2/faunadb"
)

type Todo struct {
	Text     string `fauna:"text"`
	Complete bool   `fauna:"complete"`
}

func Delete(client *f.FaunaClient, id string) error {
	_, err := client.Query(f.Delete(f.RefCollection(f.Collection("todos"), id)))
	return err
}

func Get(client *f.FaunaClient, id string) (*Todo, error) {
	result, err := client.Query(f.Get(f.RefCollection(f.Collection("todos"), id)))
	if err != nil {
		return nil, err
	}

	var t Todo
	result.At(f.ObjKey("data")).Get(&t)

	return &t, nil
}

func GetAll(client *f.FaunaClient) ([]*Todo, error) {
	todos := make([]*Todo, 0)
	result, err := client.Query(f.Get(f.Collection("todos")))
	if err != nil {
		return nil, err
	}
	log.Printf("%+v: ", result)
	return todos, nil
}

func Update(client *f.FaunaClient, id string, todo Todo) error {
	_, err := client.Query(
		f.Update(
			f.RefCollection(f.Collection("todos"), id),
			f.Obj{"data": f.Obj{"tags": todo}},
		),
	)
	return err
}

func Replace(client *f.FaunaClient, id string, todo Todo) error {
	_, err := client.Query(
		f.Replace(
			f.RefCollection(f.Collection("todos"), id),
			f.Obj{"data": todo},
		),
	)

	return err
}

func Create(client *f.FaunaClient, todo Todo) (string, error) {
	v, err := client.Query(
		f.Create(
			f.Collection("todos"),
			f.Obj{"data": todo},
		),
	)
	if err != nil {
		return "", err
	}

	var ref f.RefV
	v.At(f.ObjKey("ref")).Get(&ref)

	return ref.ID, nil
}
