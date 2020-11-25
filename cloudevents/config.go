package cloudevents

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	cloudeventsConfig = "fxstack.cloudevents"

	MiddlewareNewRelicEnabled       = "fxstack.cloudevents.middleware.newrelic.enabled"
	MiddlewareNewRelicTxName        = "fxstack.cloudevents.middleware.newrelic.txname"
	MiddlewareLogEnabled            = "fxstack.cloudevents.middleware.log.enabled"
	MiddlewareEventPublisherEnabled = "fxstack.cloudevents.middleware.eventpublisher.enabled"

	HandleDiscardEventsID = cloudeventsConfig + ".handle.discard.ids"
)

func init() {

	log.Println("getting configurations for fxstack cloudevents")

	giconfig.Add(MiddlewareNewRelicEnabled, false, "cloudevents newrelic middleware enable/disable")
	giconfig.Add(MiddlewareNewRelicTxName, "changeme", "cloudevents newrelic middleware tx name")
	giconfig.Add(MiddlewareLogEnabled, true, "cloudevents log middleware enable/disable")
	giconfig.Add(MiddlewareEventPublisherEnabled, true, "cloudevents event publisher middleware enable/disable")

	giconfig.Add(HandleDiscardEventsID, "", "cloudevents events id that will not be processed, comma separated")
}

func MiddlewareNewRelicEnabledValue() bool {
	return giconfig.Bool(MiddlewareNewRelicEnabled)
}

func MiddlewareNewRelicTxNameValue() string {
	return giconfig.String(MiddlewareNewRelicTxName)
}

func MiddlewareLogEnabledValue() bool {
	return giconfig.Bool(MiddlewareLogEnabled)
}

func MiddlewareEventPublisherEnabledValue() bool {
	return giconfig.Bool(MiddlewareEventPublisherEnabled)
}

func HandleDiscardEventsIDValue() string {
	return giconfig.String(HandleDiscardEventsID)
}
