package config

import (
	"fmt"
	"os"

	"golang.org/x/xerrors"
)

// AzureConf is azure config
type AzureConf struct {
	// Azure storage endpoint
	Endpoint string `json:"endpoint"`

	// Azure account name to use. AZURE_STORAGE_ACCOUNT environment variable is used if not specified
	AccountName string `json:"accountName"`

	// Azure account key to use. AZURE_STORAGE_ACCESS_KEY environment variable is used if not specified
	AccountKey string `json:"-"`

	// Azure storage container name
	ContainerName string `json:"containerName"`

	Enabled bool `toml:"-" json:"-"`
}

const (
	azureAccount = "AZURE_STORAGE_ACCOUNT"
	azureKey     = "AZURE_STORAGE_ACCESS_KEY"
)

// Validate configuration
func (c *AzureConf) Validate() (errs []error) {
	if !c.Enabled {
		return
	}

	// overwrite if env var is not empty
	if os.Getenv(azureAccount) != "" {
		c.AccountName = os.Getenv(azureAccount)
	}
	if c.AccountName == "" {
		errs = append(errs, xerrors.Errorf("Azure account name is required"))
	}
	if os.Getenv(azureKey) != "" {
		c.AccountKey = os.Getenv(azureKey)
	}
	if c.AccountKey == "" {
		errs = append(errs, xerrors.Errorf("Azure account key is required"))
	}

	if c.Endpoint == "" {
		c.Endpoint = fmt.Sprintf("https://%s.blob.core.windows.net/", c.AccountName)
	}

	if c.ContainerName == "" {
		errs = append(errs, xerrors.Errorf("Azure storage container name is required"))
	}
	return
}
