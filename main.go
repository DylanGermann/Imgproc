package main

import (
	"Imgproc/filter"
	"Imgproc/task"
)

func main() {
	var f filter.Filter = &filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs", "output", f)

	t.Process()
}
