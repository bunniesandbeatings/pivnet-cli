package rc

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/bunniesandbeatings/pivnet-cli/util"
	"gopkg.in/yaml.v2"
)

type S3 struct {
	Bucket string `yaml:"bucket,omitempty"`
	AccessKeyId string `yaml:"access_key_id,omitempty"`
	SecretAccessKey string `yaml:"secret_access_key,omitempty"`
	Region string `yaml:"region,omitempty"`
}

type rcYAML struct {
	S3 S3 `yaml:"s3,omitempty"`
}

func Read() (*rcYAML, error) {
	var rc *rcYAML

	pivnetRC := filepath.Join(util.UserHomeDir(), ".pivnetrc")
	if _, err := os.Stat(pivnetRC); err == nil {
		pivnetRCBytes, err := ioutil.ReadFile(pivnetRC)
		if err != nil {
			return nil, fmt.Errorf("could not read %s", pivnetRC)
		}

		err = yaml.Unmarshal(pivnetRCBytes, &rc)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal %s", pivnetRC)
		}
	}

	if rc == nil {
		return nil, fmt.Errorf("%s was empty?", pivnetRC)
	}

	return rc, nil
}
