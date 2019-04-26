// Package tethering provides the Chrome DevTools Protocol
// commands, types, and events for the Tethering domain.
//
// The Tethering domain defines methods and events for browser port binding.
//
// Generated by the cdproto-gen command.
package tethering

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/nbzx/cdproto/cdp"
)

// BindParams request browser port binding.
type BindParams struct {
	Port int64 `json:"port"` // Port number to bind.
}

// Bind request browser port binding.
//
// parameters:
//   port - Port number to bind.
func Bind(port int64) *BindParams {
	return &BindParams{
		Port: port,
	}
}

// Do executes Tethering.bind against the provided context.
func (p *BindParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandBind, p, nil)
}

// UnbindParams request browser port unbinding.
type UnbindParams struct {
	Port int64 `json:"port"` // Port number to unbind.
}

// Unbind request browser port unbinding.
//
// parameters:
//   port - Port number to unbind.
func Unbind(port int64) *UnbindParams {
	return &UnbindParams{
		Port: port,
	}
}

// Do executes Tethering.unbind against the provided context.
func (p *UnbindParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandUnbind, p, nil)
}

// Command names.
const (
	CommandBind   = "Tethering.bind"
	CommandUnbind = "Tethering.unbind"
)
