package repository

import giconfig "github.com/b2wdigital/goignite/config"

type Options struct {
	Provider string
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath("fxstack.repository.event", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
