package main

import (
	"fmt"
	"golang-test-task/task_one"
)

func main() {

	// sample_str := "23-ok-48-cab-56-haha"

	// fmt.Println(task_one.TestValidity(sample_str))
	// fmt.Println(task_one.AverageNumber(sample_str))
	// fmt.Println(task_one.WholeStory(sample_str))
	// fmt.Println(task_one.StoryStats(sample_str))
	// fmt.Println()
	fmt.Println(task_one.Generate(true))
	fmt.Println()
	fmt.Println(task_one.Generate(false))
	fmt.Println("---")

	another_str := task_one.Generate(true)
	fmt.Printf("String is: %q ", another_str)
	fmt.Println()
	fmt.Println(task_one.TestValidity(another_str))
	fmt.Println(task_one.AverageNumber(another_str))
	fmt.Println(task_one.WholeStory(another_str))
	fmt.Println(task_one.StoryStats(another_str))

}
