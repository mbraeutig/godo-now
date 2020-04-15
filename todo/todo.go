package todo

import (
	f "github.com/mbraeutig/faunadb-v2/faunadb"
)

type Todo struct {
	Text     string `fauna:"text"`
	Complete bool   `fauna:"complete"`
	ID       string `fauna:"id"`
}

func Delete(client *f.FaunaClient, id string) error {
	_, err :=
		client.Query(f.Delete(f.RefCollection(f.Collection("todos"), id)))
	return err
}

func Read(client *f.FaunaClient, id string) (*Todo, error) {
	result, err :=
		client.Query(f.Get(f.RefCollection(f.Collection("todos"), id)))
	if err != nil {
		return nil, err
	}

	var t Todo
	result.At(f.ObjKey("data")).Get(&t)

	return &t, nil
}

func ReadAll(client *f.FaunaClient) ([]*Todo, error) {
	result, err :=
		client.Query(
			f.Map(f.Paginate(f.Match(f.Index("all_todos"))), f.Lambda("X", f.Get(f.Var("X")))))
	if err != nil {
		return nil, err
	}

	var elements f.ArrayV
	result.At(f.ObjKey("data")).Get(&elements)

	todos := make([]*Todo, 0)
	for _, element := range elements {

		var object f.ObjectV
		element.At(f.ObjKey("data")).Get(&object)

		var ref f.RefV
		element.At(f.ObjKey("ref")).Get(&ref)

		var t Todo
		object.Get(&t)

		t.ID = ref.ID

		todos = append(todos, &t)
	}

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
	type createTodo struct {
		Text     string
		Complete bool
	}
	ct := createTodo{Text: todo.Text, Complete: todo.Complete}

	v, err := client.Query(
		f.Create(
			f.Collection("todos"),
			f.Obj{"data": ct},
		),
	)
	if err != nil {
		return "", err
	}

	var ref f.RefV
	v.At(f.ObjKey("ref")).Get(&ref)

	return ref.ID, nil
}
