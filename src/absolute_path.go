package dotfiles

import (
	"fmt"
	"os"

	"github.com/rhysd/abspath"
)

func absolutePathToRepo(repo string) (abspath.AbsPath, error) {
	if repo == "" {
		repo = os.Getenv("DOTFILES_REPO_PATH")
	}

	if repo == "" {
		repo = "."
		fmt.Fprintln(os.Stderr, "No repository was specified nor $DOTFILES_REPO_PATH was not set. Assuming current repository is a dotfiles repository.\n")
	}

	p, err := abspath.ExpandFrom(repo)
	if err != nil {
		return abspath.AbsPath{}, err
	}

	if !p.IsDir() {
		return abspath.AbsPath{}, fmt.Errorf("'%s' is not a directory. Please specify your dotfiles directory.", p.String())
	}

	return p, nil
}
