package v0

import "log"

// New creates a new EmailTaskService with the given InitAttribute.
// It will panic if the InitAttribute is invalid.
func New(attr InitAttribute) *EmailTaskService {
	if !checkRepository(attr.Repo) {
		log.Panicf("[EmailTaskService][New] missing repo %+v", attr.Repo)
	}

	return &EmailTaskService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

// checkRepository checks if the EmailTaskRepo repository attribute is set.
//
// Parameters:
//   - repoAttr: The repository attributes containing the EmailTaskRepo repository.
//
// Returns:
//   - bool: true if the EmailTaskRepo repository is non-nil, indicating a valid repository; otherwise, false.
func checkRepository(repoAttr RepoAttribute) bool {
	return repoAttr.EmailTaskRepo != nil
}
