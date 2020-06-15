package kinesis

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

type Options struct {
	RandomPartitionKey bool `config:"randompartitionkey"`
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath("fxstack.provider.kinesis", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
