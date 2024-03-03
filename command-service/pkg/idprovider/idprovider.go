package idprovider

import "github.com/bwmarrin/snowflake"

type IdProvider struct {
	sfn *snowflake.Node
}

func New() (*IdProvider, error) {
	sfn, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	return &IdProvider{
		sfn: sfn,
	}, nil
}

func (idp *IdProvider) Generate() string {
	return idp.sfn.Generate().String()
}
