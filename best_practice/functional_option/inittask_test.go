package functionaloption

import (
	"fmt"
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask("test", WithDesc("this is a test task"))
	t.Logf("%+v", task)
}

func Testnewtask() {

}

func ExampleNewTask() {
	task := NewTask("test", WithDesc("this is a test task"))
	fmt.Printf("%+v", task.Title)
}
