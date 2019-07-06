package utils

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
	"log"
	"path"
)

type GitProject string

func (p GitProject) GetRemoteBranchNames() []string{
	repo, err := git.PlainOpen(string(p))

	if err != nil {
		log.Fatal(err)
	}

	refs, err := repo.Storer.IterReferences()

	if err != nil {
		log.Fatal(err)
	}

	branches := storer.NewReferenceFilteredIter(func(ref *plumbing.Reference) bool {
		return ref.Name().IsRemote()
	}, refs)

	return formatBranchName(branches)
}

func formatBranchName(branches storer.ReferenceIter) []string {
	var branchNames []string

	branches.ForEach(func (b *plumbing.Reference) error {
		branchName := path.Base(string(b.Name()))

		if branchName != "HEAD" {
			branchNames = append(branchNames, branchName)
		}

		return nil
	})

	return branchNames
}

func GetProjectName() string {
	return "/Users/haodawang/Documents/repositories/edu-saas-workbench"
}