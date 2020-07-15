package lambda

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

type Options struct {
	Skip bool
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath("fxstack.serverless.lambda", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
