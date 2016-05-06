package repositories

import (
	"github.com/bunniesandbeatings/pivnet-cli/rc"
	"fmt"
)

type Repository interface {
	Search() error
}

type Repositories map[string]Repository

func Load(repositoryList *[]rc.Repository) (*Repositories, error) {
	repositories := make(Repositories)

	for _, repositoryDefinition := range *repositoryList {
		repoType := repositoryDefinition.Type
		switch repoType {
		case "s3":
			repositories[repositoryDefinition.Name] = NewS3Repository(repositoryDefinition)
		default:
			return nil, fmt.Errorf("%s is an unrecognized repository type", repoType);
		}
	}
	return &repositories, nil
}