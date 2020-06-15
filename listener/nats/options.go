package nats

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

type Options struct {
	Subjects []string
	Queue    string
}

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath("fxstack.listener.nats", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
