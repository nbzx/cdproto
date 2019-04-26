// Package webaudio provides the Chrome DevTools Protocol
// commands, types, and events for the WebAudio domain.
//
// This domain allows inspection of Web Audio API.
// https://webaudio.github.io/web-audio-api/.
//
// Generated by the cdproto-gen command.
package webaudio

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/nbzx/cdproto/cdp"
)

// EnableParams enables the WebAudio domain and starts sending context
// lifetime events.
type EnableParams struct{}

// Enable enables the WebAudio domain and starts sending context lifetime
// events.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes WebAudio.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandEnable, nil, nil)
}

// DisableParams disables the WebAudio domain.
type DisableParams struct{}

// Disable disables the WebAudio domain.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes WebAudio.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDisable, nil, nil)
}

// GetRealtimeDataParams fetch the realtime data from the registered
// contexts.
type GetRealtimeDataParams struct {
	ContextID ContextID `json:"contextId"`
}

// GetRealtimeData fetch the realtime data from the registered contexts.
//
// parameters:
//   contextID
func GetRealtimeData(contextID ContextID) *GetRealtimeDataParams {
	return &GetRealtimeDataParams{
		ContextID: contextID,
	}
}

// GetRealtimeDataReturns return values.
type GetRealtimeDataReturns struct {
	RealtimeData *ContextRealtimeData `json:"realtimeData,omitempty"`
}

// Do executes WebAudio.getRealtimeData against the provided context.
//
// returns:
//   realtimeData
func (p *GetRealtimeDataParams) Do(ctxt context.Context) (realtimeData *ContextRealtimeData, err error) {
	// execute
	var res GetRealtimeDataReturns
	err = cdp.Execute(ctxt, CommandGetRealtimeData, p, &res)
	if err != nil {
		return nil, err
	}

	return res.RealtimeData, nil
}

// Command names.
const (
	CommandEnable          = "WebAudio.enable"
	CommandDisable         = "WebAudio.disable"
	CommandGetRealtimeData = "WebAudio.getRealtimeData"
)
