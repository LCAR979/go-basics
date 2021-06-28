package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int {
	return 42
}

func init() {
	WhatIsThe = 0
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}

/*
AnswerToLife() is guaranteed to run before init() is called,
and init() is guaranteed to run before main() is called.

Keep in mind that init() is always called, regardless if there's main or not,
so if you import a package that has an init function, it will be executed.

Additionally, you can have multiple init() functions per package;
they will be executed in the order they show up in the file
(after all variables are initialized of course).
If they span multiple files, they will be executed in lexical file name order
*/
