package cmd

import "github.com/b2wdigital/fxstack/n/channel"

type Cmd struct {
	channs []channel.Channel
}

func New() *Cmd {
	return &Cmd{}
}

func (c *Cmd) WithChan(ch channel.Channel) *Cmd {
	c.channs = append(c.channs, ch)
	return c
}