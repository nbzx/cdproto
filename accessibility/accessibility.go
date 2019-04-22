// Package accessibility provides the Chrome DevTools Protocol
// commands, types, and events for the Accessibility domain.
//
// Generated by the cdproto-gen command.
package accessibility

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/nbzx/cdproto/cdp"
	"github.com/nbzx/cdproto/runtime"
)

// DisableParams disables the accessibility domain.
type DisableParams struct{}

// Disable disables the accessibility domain.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Accessibility.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables the accessibility domain which causes AXNodeIds to
// remain consistent between method calls. This turns on accessibility for the
// page, which can impact performance until accessibility is disabled.
type EnableParams struct{}

// Enable enables the accessibility domain which causes AXNodeIds to remain
// consistent between method calls. This turns on accessibility for the page,
// which can impact performance until accessibility is disabled.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Accessibility.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnable, nil, nil)
}

// GetPartialAXTreeParams fetches the accessibility node and partial
// accessibility tree for this DOM node, if it exists.
type GetPartialAXTreeParams struct {
	NodeID         cdp.NodeID             `json:"nodeId,omitempty"`         // Identifier of the node to get the partial accessibility tree for.
	BackendNodeID  cdp.BackendNodeID      `json:"backendNodeId,omitempty"`  // Identifier of the backend node to get the partial accessibility tree for.
	ObjectID       runtime.RemoteObjectID `json:"objectId,omitempty"`       // JavaScript object id of the node wrapper to get the partial accessibility tree for.
	FetchRelatives bool                   `json:"fetchRelatives,omitempty"` // Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
}

// GetPartialAXTree fetches the accessibility node and partial accessibility
// tree for this DOM node, if it exists.
//
// parameters:
func GetPartialAXTree() *GetPartialAXTreeParams {
	return &GetPartialAXTreeParams{}
}

// WithNodeID identifier of the node to get the partial accessibility tree
// for.
func (p GetPartialAXTreeParams) WithNodeID(nodeID cdp.NodeID) *GetPartialAXTreeParams {
	p.NodeID = nodeID
	return &p
}

// WithBackendNodeID identifier of the backend node to get the partial
// accessibility tree for.
func (p GetPartialAXTreeParams) WithBackendNodeID(backendNodeID cdp.BackendNodeID) *GetPartialAXTreeParams {
	p.BackendNodeID = backendNodeID
	return &p
}

// WithObjectID JavaScript object id of the node wrapper to get the partial
// accessibility tree for.
func (p GetPartialAXTreeParams) WithObjectID(objectID runtime.RemoteObjectID) *GetPartialAXTreeParams {
	p.ObjectID = objectID
	return &p
}

// WithFetchRelatives whether to fetch this nodes ancestors, siblings and
// children. Defaults to true.
func (p GetPartialAXTreeParams) WithFetchRelatives(fetchRelatives bool) *GetPartialAXTreeParams {
	p.FetchRelatives = fetchRelatives
	return &p
}

// GetPartialAXTreeReturns return values.
type GetPartialAXTreeReturns struct {
	Nodes []*Node `json:"nodes,omitempty"` // The Accessibility.AXNode for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
}

// Do executes Accessibility.getPartialAXTree against the provided context.
//
// returns:
//   nodes - The Accessibility.AXNode for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
func (p *GetPartialAXTreeParams) Do(ctxt context.Context, h cdp.Executor) (nodes []*Node, err error) {
	// execute
	var res GetPartialAXTreeReturns
	err = h.Execute(ctxt, CommandGetPartialAXTree, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// GetFullAXTreeParams fetches the entire accessibility tree.
type GetFullAXTreeParams struct{}

// GetFullAXTree fetches the entire accessibility tree.
func GetFullAXTree() *GetFullAXTreeParams {
	return &GetFullAXTreeParams{}
}

// GetFullAXTreeReturns return values.
type GetFullAXTreeReturns struct {
	Nodes []*Node `json:"nodes,omitempty"`
}

// Do executes Accessibility.getFullAXTree against the provided context.
//
// returns:
//   nodes
func (p *GetFullAXTreeParams) Do(ctxt context.Context, h cdp.Executor) (nodes []*Node, err error) {
	// execute
	var res GetFullAXTreeReturns
	err = h.Execute(ctxt, CommandGetFullAXTree, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Nodes, nil
}

// Command names.
const (
	CommandDisable          = "Accessibility.disable"
	CommandEnable           = "Accessibility.enable"
	CommandGetPartialAXTree = "Accessibility.getPartialAXTree"
	CommandGetFullAXTree    = "Accessibility.getFullAXTree"
)
