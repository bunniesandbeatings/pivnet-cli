package commands

import (
	"github.com/bunniesandbeatings/pivnet-cli/rc"
	"fmt"
	"github.com/bunniesandbeatings/pivnet-cli/repositories"
)

type SearchCommand struct {

}

func (command *SearchCommand) Execute(args []string) error {

	config, err := rc.Read()

	if err != nil {
		return err
	}

	repos, err := repositories.Load(config.Repositories)
	if err != nil {
		return err
	}

	fmt.Println("Listing objects...\n")

	for repositoryName, repository := range *repos {
		fmt.Printf("-------\nIn repository: '%s'\n", repositoryName)
		err = repository.Search()
		if err != nil {
			return err
		}
	}

	return nil
}
