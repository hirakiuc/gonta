package plugin

import (
	"../logger"
	"../slack"
)

type Plugin interface {
	// Get Plugin Instance
	GetInstance() Plugin

	// Check whether event is acceptable or not.
	IsAccept(event slack.Event) bool

	// Notify event
	Notify(session *slack.Session, event slack.Event)
}

type PluginRegistry struct {
	plugins []Plugin
}

var registry *PluginRegistry
var log *logger.Logger

func init() {
	log = logger.GetLogger()
}

func GetRegistry() *PluginRegistry {
	if registry == nil {
		registry = &PluginRegistry{[]Plugin{}}
	}

	return registry
}

func (reg *PluginRegistry) AddPlugin(plugin Plugin) {
	reg.plugins = append(reg.plugins, plugin)
}

func (reg *PluginRegistry) Notify(session *slack.Session, event slack.Event) {
	for _, plugin := range reg.plugins {
		if plugin.IsAccept(event) {
			instance := plugin.GetInstance()

			go func(plg Plugin, ses *slack.Session, evt slack.Event) {
				plg.Notify(ses, evt)
			}(instance, session, event)
		}
	}
}
