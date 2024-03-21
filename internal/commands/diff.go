package commands

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/utils/merkletrie"
)

func Diff() (string, error) {
	repo, err := git.PlainOpen("./")
	if err != nil {
		return "", fmt.Errorf("git plain open: %w", err)
	}

	ref, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("git plain open: %w", err)
	}

	refMain, err := repo.Reference("refs/heads/main", false)
	if err != nil {
		return "", fmt.Errorf("git plain open: %w", err)
	}

	commCur, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return "", fmt.Errorf("repo.CommitObject: %w", err)
	}

	treeCur, err := commCur.Tree()
	if err != nil {
		return "", fmt.Errorf("commCur.Tree: %w", err)
	}

	commMain, err := repo.CommitObject(refMain.Hash())
	if err != nil {
		return "", fmt.Errorf("repo.CommitObject(main): %w", err)
	}

	treeMain, err := commMain.Tree()
	if err != nil {
		return "", fmt.Errorf("commCur.Tree (main): %w", err)
	}

	changes, err := object.DiffTree(treeCur, treeMain)
	if err != nil {
		return "", fmt.Errorf("diffTree: %w", err)
	}

	var fileNames []string
	for _, change := range changes {
		action, err := change.Action()
		if err != nil {
			return "", fmt.Errorf("git action parse: %w", err)
		}

		if action == merkletrie.Delete {
			continue
		}

		from, to, err := change.Files()
		if err != nil {
			return "", fmt.Errorf("file changes parse: %w", err)
		}

		fileNames = append(fileNames, from.Name)
		fileNames = append(fileNames, to.Name)
	}

	return strings.Join(fileNames, "|"), nil
}
