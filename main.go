package main

import (
	"fmt"
	"golang-test-task/task"
)

func main() {

	// sample_str := "23-ok-48-cab-56-haha"

	// fmt.Println(task.TestValidity(sample_str))
	// fmt.Println(task.AverageNumber(sample_str))
	// fmt.Println(task.WholeStory(sample_str))
	// fmt.Println(task.StoryStats(sample_str))
	// fmt.Println()
	fmt.Println(task.Generate(true))
	fmt.Println(task.Generate(false))
	fmt.Println("---")

	another_str := task.Generate(true)
	fmt.Printf("String is: %q ", another_str)
	fmt.Println()
	fmt.Println(task.TestValidity(another_str))
	fmt.Println(task.AverageNumber(another_str))
	fmt.Println(task.WholeStory(another_str))
	fmt.Println(task.StoryStats(another_str))

}
