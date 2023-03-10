package data

func DataGenerator() []map[string]interface{} {
	QuizData := []map[string]interface{}{
		{
			"topic":          "Linux",
			"question":       "What is the command to list all files and directories in the current directory?",
			"answer_options": []string{"ls", "pwd", "cd", "touch"},
			"correct_answer": "ls",
		},
		{
			"topic":          "Linux",
			"question":       "What is the command to change the current working directory?",
			"answer_options": []string{"chdir", "cd", "pwd", "ls"},
			"correct_answer": "cd",
		},
		{
			"topic":          "Linux",
			"question":       "What is the command to create a new directory?",
			"answer_options": []string{"mkdir", "touch", "rm", "cp"},
			"correct_answer": "mkdir",
		},
		{
			"topic":          "Linux",
			"question":       "What is the command to remove a file?",
			"answer_options": []string{"rm", "rmdir", "mv", "cp"},
			"correct_answer": "rm",
		},
		{
			"topic":          "Linux",
			"question":       "What is the command to remove a directory?",
			"answer_options": []string{"rmdir", "rm", "mv", "cp"},
			"correct_answer": "rmdir",
		},
		{
			"topic":          "Linux",
			"question":       "What is the command to copy a file?",
			"answer_options": []string{"cp", "mv", "rm", "rmdir"},
			"correct_answer": "cp",
		},
		{
			"topic":          "Linux",
			"question":       "What is the command to move a file?",
			"answer_options": []string{"mv", "cp", "rm", "rmdir"},
			"correct_answer": "mv",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is NOT a valid way to create a new slice?",
			"answer_options": []string{
				"var s []int",
				"s := make([]int)",
				"s := []int{1, 2, 3}",
				"s := []int",
			},
			"correct_answer": "s := []int",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is the correct way to create a new channel?",
			"answer_options": []string{
				"var c chan int",
				"c := make(chan int)",
				"c := chan int{}",
				"c := chan int",
			},
			"correct_answer": "c := make(chan int)",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is the correct way to create a new map?",
			"answer_options": []string{
				"var m map[string]int",
				"m := make(map[string]int)",
				"m := map[string]int{}",
				"m := map[string]int",
			},
			"correct_answer": "m := make(map[string]int)",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is NOT a valid way to create a new struct?",
			"answer_options": []string{
				"var s struct{}",
				"s := make(struct{})",
				"s := struct{}{}",
				"s := struct{}",
			},
			"correct_answer": "s := make(struct{})",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is NOT a valid way to create a new function?",
			"answer_options": []string{
				"func f() {}",
				"var f func()",
				"f := make(func())",
				"f := func(){}",
			},
			"correct_answer": "f := make(func())",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a for loop in Golang?",
			"answer_options": []string{
				"for i := 0; i < 10; i++ {}",
				"for i = 0; i < 10; i++ {}",
				"for (i = 0; i < 10; i++) {}",
				"for i in 0..10 {}",
			},
			"correct_answer": "for i := 0; i < 10; i++ {}",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for an if statement in Golang?",
			"answer_options": []string{
				"if x > 0 {}",
				"if (x > 0) {}",
				"if x > 0 then {}",
				"if x > 0 {} else {}",
			},
			"correct_answer": "if x > 0 {}",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to add all changes to the staging area?",
			"answer_options": []string{"git add", "git commit", "git push", "git pull"},
			"correct_answer": "git add",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to commit changes?",
			"answer_options": []string{"git commit", "git add", "git push", "git pull"},
			"correct_answer": "git commit",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to push changes to a remote repository?",
			"answer_options": []string{"git push", "git add", "git commit", "git pull"},
			"correct_answer": "git push",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to pull changes from a remote repository?",
			"answer_options": []string{"git pull", "git add", "git commit", "git push"},
			"correct_answer": "git pull",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to view the commit history?",
			"answer_options": []string{"git log", "git status", "git diff", "git show"},
			"correct_answer": "git log",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to view the changes in the working directory?",
			"answer_options": []string{"git diff", "git log", "git status", "git show"},
			"correct_answer": "git diff",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to view the status of the working directory and staging area?",
			"answer_options": []string{"git status", "git diff", "git log", "git show"},
			"correct_answer": "git status",
		},
		{
			"topic":          "Git",
			"question":       "What is the correct command to view the changes in a commit?",
			"answer_options": []string{"git show", "git diff", "git log", "git status"},
			"correct_answer": "git show",
		},
		{
			"topic":    "Git",
			"question": "What is the correct command to switch to a different branch?",
			"answer_options": []string{
				"git switch",
				"git checkout",
				"git branch",
				"git merge",
			},
			"correct_answer": "git checkout",
		},
		{
			"topic":    "Git",
			"question": "What is the correct command to create a new branch?",
			"answer_options": []string{
				"git new",
				"git create",
				"git branch",
				"git merge",
			},
			"correct_answer": "git branch",
		},
		{
			"topic":    "Git",
			"question": "What is the correct command to merge changes from one branch into another?",
			"answer_options": []string{
				"git merge",
				"git new",
				"git create",
				"git branch",
			},
			"correct_answer": "git merge",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is NOT a valid way to create a new interface?",
			"answer_options": []string{
				"var i interface{}",
				"i := make(interface{})",
				"i := interface{}{}",
				"i := interface{}",
			},
			"correct_answer": "i := make(interface{})",
		},
		{
			"topic":    "Golang",
			"question": "Which of the following is NOT a valid way to create a new function?",
			"answer_options": []string{
				"func f() {}",
				"var f func()",
				"f := make(func())",
				"f := func(){}",
			},
			"correct_answer": "f := make(func())",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a switch statement in Golang?",
			"answer_options": []string{
				"switch x { case 0: case 1: }",
				"switch { case x == 0: case x == 1: }",
				"switch x { case 0; case 1; }",
				"switch { case x = 0; case x = 1; }",
			},
			"correct_answer": "switch x { case 0: case 1: }",
		},
		{
			"topic":          "Golang",
			"question":       "What is the correct syntax for a type assertion in Golang?",
			"answer_options": []string{"x.(T)", "(T)x", "x as T", "x is T"},
			"correct_answer": "x.(T)",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a type switch in Golang?",
			"answer_options": []string{
				"switch x.(type) { case int: case string: }",
				"switch type(x) { case int: case string: }",
				"switch x { case int: case string: }",
				"switch (type)x { case int: case string: }",
			},
			"correct_answer": "switch x.(type) { case int: case string: }",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a select statement in Golang?",
			"answer_options": []string{
				"select { case <-c: }",
				"select c { case <-c: }",
				"select { case c <-: }",
				"select <-c { case c: }",
			},
			"correct_answer": "select { case <-c: }",
		},
		{
			"topic":          "Golang",
			"question":       "What is the correct syntax for a go statement in Golang?",
			"answer_options": []string{"go f()", "go f", "f() go", "f go"},
			"correct_answer": "go f()",
		},
		{
			"topic":          "Golang",
			"question":       "What is the correct syntax for a deferred function call in Golang?",
			"answer_options": []string{"defer f()", "defer f", "f() defer", "f defer"},
			"correct_answer": "defer f()",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a return statement in Golang?",
			"answer_options": []string{
				"return x",
				"return (x)",
				"return x;",
				"return (x);",
			},
			"correct_answer": "return x",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a for range loop in Golang?",
			"answer_options": []string{
				"for i, v := range x {}",
				"for i, v in x {}",
				"for (i, v) in x {}",
				"for i, v of x {}",
			},
			"correct_answer": "for i, v := range x {}",
		},
		{
			"topic":    "Golang",
			"question": "What is the correct syntax for a type definition in Golang?",
			"answer_options": []string{
				"type T struct{}",
				"struct T{}",
				"type T {}",
				"struct{} T",
			},
			"correct_answer": "type T struct{}",
		},
	}
	return QuizData
}

// GetUniqueTopic - Get unique topics from quiz data
func GetUniqueTopic(quizData []map[string]interface{}) []string {
	uniqueTopics := make(map[string]int)
	uniqueTopicsSlice := make([]string, 0)

	for _, quiz := range quizData {
		topic := quiz["topic"].(string)
		uniqueTopics[topic]++
	}
	for topic := range uniqueTopics {
		uniqueTopicsSlice = append(uniqueTopicsSlice, topic)
	}
	return uniqueTopicsSlice
}

// FilterQuizData - Filter quiz data based on topic
func FilterQuizData(quizData []map[string]interface{}, topic string) []map[string]interface{} {
	filteredData := make([]map[string]interface{}, 0)

	for _, quiz := range quizData {
		if quiz["topic"] == topic {
			filteredData = append(filteredData, quiz)
		}
	}
	return filteredData
}
