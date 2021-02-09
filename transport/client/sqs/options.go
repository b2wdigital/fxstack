package sqs

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/lann/builder"
)

type Options struct {
	Timeout time.Duration
}

type optionsBuilder builder.Builder

func (b optionsBuilder) Timeout(value time.Duration) optionsBuilder {
	return builder.Set(b, "Timeout", value).(optionsBuilder)
}

func (b optionsBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var OptionsBuilder = builder.Register(optionsBuilder{}, Options{}).(optionsBuilder)

func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := giconfig.UnmarshalWithPath("fxstack.transport.client.sqs", o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
