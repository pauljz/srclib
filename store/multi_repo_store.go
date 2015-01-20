package store

import (
	"sourcegraph.com/sourcegraph/srclib/graph"
	"sourcegraph.com/sourcegraph/srclib/unit"
)

// MultiRepoStore provides access to RepoStores for multiple
// repositories.
//
// Using this interface instead of directly accessing a single
// RepoStore allows aliasing repository URIs and supporting both ID
// and URI lookups.
type MultiRepoStore interface {
	// Repo gets a single repository from the store.
	Repo(string) (string, error)

	// Repos returns all repositories that match the RepoFilter.
	Repos(RepoFilter) ([]string, error)

	// RepoStore's methods call the corresponding methods on the
	// RepoStore of each repository contained within this multi-repo
	// store. The combined results are returned (in undefined order).
	RepoStore
}

// A MultiRepoImporter imports srclib build data for a repository's
// source unit at a specific version into a RepoStore.
type MultiRepoImporter interface {
	// Import imports srclib build data for a source unit at a
	// specific version into the store.
	Import(repo, commitID string, unit *unit.SourceUnit, data graph.Output) error
}

// A MultiRepoStoreImporter implements both MultiRepoStore and
// MultiRepoImporter.
type MultiRepoStoreImporter interface {
	MultiRepoStore
	MultiRepoImporter
}

// TODO(sqs): What should the Repo type be? Right now it is just string.

// checkDefKeyValidForRepoStore returns an *InvalidKeyError if the def
// key is underspecified for use in (RepoStore).Def.
func checkDefKeyValidForMultiRepoStore(key graph.DefKey) error {
	if err := checkDefKeyValidForRepoStore(key); err != nil {
		return err
	}
	if key.Repo == "" {
		return &InvalidKeyError{"empty DefKey.Repo"}
	}
	return nil
}
