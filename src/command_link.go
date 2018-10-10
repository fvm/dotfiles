package dotfiles

func Link(repoInput string, specified []string, dry bool) error {
	repo, err := absolutePathToRepo(repoInput)
	if err != nil {
		return err
	}

	m, err := GetMappings(repo.Join(".dotfiles"))
	if err != nil {
		return err
	}

	if specified == nil || len(specified) == 0 {
		err = m.CreateAllLinks(dry)
		if e, ok := err.(*NothingLinkedError); ok {
			e.RepoPath = repo.String()
		}
		return err
	} else {
		return m.CreateSomeLinks(specified, dry)
	}
}
