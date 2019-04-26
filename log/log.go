// Package log provides the Chrome DevTools Protocol
// commands, types, and events for the Log domain.
//
// Provides access to log entries.
//
// Generated by the cdproto-gen command.
package log

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/nbzx/cdproto/cdp"
)

// ClearParams clears the log.
type ClearParams struct{}

// Clear clears the log.
func Clear() *ClearParams {
	return &ClearParams{}
}

// Do executes Log.clear against the provided context.
func (p *ClearParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandClear, nil, nil)
}

// DisableParams disables log domain, prevents further log entries from being
// reported to the client.
type DisableParams struct{}

// Disable disables log domain, prevents further log entries from being
// reported to the client.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Log.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables log domain, sends the entries collected so far to the
// client by means of the entryAdded notification.
type EnableParams struct{}

// Enable enables log domain, sends the entries collected so far to the
// client by means of the entryAdded notification.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Log.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandEnable, nil, nil)
}

// StartViolationsReportParams start violation reporting.
type StartViolationsReportParams struct {
	Config []*ViolationSetting `json:"config"` // Configuration for violations.
}

// StartViolationsReport start violation reporting.
//
// parameters:
//   config - Configuration for violations.
func StartViolationsReport(config []*ViolationSetting) *StartViolationsReportParams {
	return &StartViolationsReportParams{
		Config: config,
	}
}

// Do executes Log.startViolationsReport against the provided context.
func (p *StartViolationsReportParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandStartViolationsReport, p, nil)
}

// StopViolationsReportParams stop violation reporting.
type StopViolationsReportParams struct{}

// StopViolationsReport stop violation reporting.
func StopViolationsReport() *StopViolationsReportParams {
	return &StopViolationsReportParams{}
}

// Do executes Log.stopViolationsReport against the provided context.
func (p *StopViolationsReportParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandStopViolationsReport, nil, nil)
}

// Command names.
const (
	CommandClear                 = "Log.clear"
	CommandDisable               = "Log.disable"
	CommandEnable                = "Log.enable"
	CommandStartViolationsReport = "Log.startViolationsReport"
	CommandStopViolationsReport  = "Log.stopViolationsReport"
)
