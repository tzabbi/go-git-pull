package main

import (
	"fmt"

	"github.com/tzabbi/go-git-pull/internal/git"
)

func main() {
	fmt.Println(git.GetGitDir("/home/tom/Documents/github/tzabbi"))
}
