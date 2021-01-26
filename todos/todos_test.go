package todos

import "testing"

func TestAddTodo(t *testing.T) {
	var todo Todos
	todo.Title = "Testing"
	todo.Completed = true
	if AddTodo(todo) != true {
		t.Error("Test failed to add simple todo")
	}
}
