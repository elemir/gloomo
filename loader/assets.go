package loader

import (
	"iter"

	"github.com/elemir/gloomo/container"
)

type asset[T any] struct {
	value    T
	isLoaded bool
	taskNum  int
}

// Assets is a container of loading assets tasks and it's result.
type Assets[T any] struct {
	taskQueue []string
	data      map[string]asset[T]
}

// Load adds task for loading asset and return it if it loaded successful.
func (a *Assets[T]) Load(path string) (T, bool) {
	if a.data == nil {
		a.data = make(map[string]asset[T])
	}

	loadedAsset, ok := a.data[path]
	if !ok {
		taskNum := len(a.taskQueue)

		a.taskQueue = append(a.taskQueue, path)
		a.data[path] = asset[T]{
			taskNum: taskNum,
		}

		return container.Zero[T](), false
	}

	return loadedAsset.value, loadedAsset.isLoaded
}

// Put result into Asset container.
func (a *Assets[T]) Put(path string, value T) {
	if a.data == nil {
		a.data = make(map[string]asset[T])
	}

	if !a.data[path].isLoaded {
		a.markAsDone(path)
	}

	a.data[path] = asset[T]{
		value:    value,
		taskNum:  0,
		isLoaded: true,
	}
}

// markAsDone can be used only with not loaded paths that has a task in queue.
func (a *Assets[T]) markAsDone(path string) {
	asst := a.data[path]

	taskNum := asst.taskNum

	lastTask, taskQueue := popSlice(a.taskQueue)
	a.taskQueue = taskQueue

	if taskNum != len(taskQueue) {
		a.taskQueue[taskNum] = lastTask

		lastAsset := a.data[path]
		lastAsset.taskNum = taskNum

		a.data[lastTask] = lastAsset
	}
}

func popSlice[T any](slice []T) (T, []T) {
	idx := len(slice)

	return slice[idx-1], slice[:idx-1]
}

// NotLoadedPaths is an iterator of all not loaded paths.
func (a *Assets[T]) NotLoadedPaths() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, task := range a.taskQueue {
			if !yield(task) {
				return
			}
		}
	}
}
