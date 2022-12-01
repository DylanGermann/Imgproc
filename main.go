package main

import (
	"Imgproc/filter"
	"Imgproc/task"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var srcDir = flag.String("src", "", "Input directory")
	var dstDir = flag.String("dst", "", "Output directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskType = flag.String("task", "waitgrp", "waitgrp/channel")
	var poolSize = flag.int("poolsize", 4, "Workers pool size for the channel task")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = &filter.Grayscale{}
	case "blur":
		f = &filter.Blur{}
	default:
		os.Exit(1)
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	t := task.NewChanTask(*srcDir, *dstDir, f, 16)

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)
}
