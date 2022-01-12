package main

import (
	"fmt"
	"golang-test-task/task_one"
)

func main() {

	sample_str := "23-okaaa-48-caba-56-haha"

	fmt.Println(task_one.TestValidity(sample_str))
	fmt.Println(task_one.AverageNumber(sample_str))
	fmt.Println(task_one.WholeStory(sample_str))
	fmt.Println(task_one.StoryStats(sample_str))

}
