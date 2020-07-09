package cloudevents

import "strings"

type Options struct {
	IDsToDiscard []string
}

func DefaultOptions() (*Options, error) {
	o := new(Options)

	rawIDs := HandleDiscardEventsIDValue()
	ids := strings.Split(rawIDs, ",")
	if nil != ids && len(ids) >= 1 && ids[0] == "" {
		return o, nil
	}

	o.IDsToDiscard = ids

	return o, nil
}
