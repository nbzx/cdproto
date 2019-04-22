// Package overlay provides the Chrome DevTools Protocol
// commands, types, and events for the Overlay domain.
//
// This domain provides various functionality related to drawing atop the
// inspected page.
//
// Generated by the cdproto-gen command.
package overlay

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/nbzx/cdproto/cdp"
	"github.com/nbzx/cdproto/dom"
	"github.com/nbzx/cdproto/runtime"
	"github.com/mailru/easyjson"
)

// DisableParams disables domain notifications.
type DisableParams struct{}

// Disable disables domain notifications.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Overlay.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables domain notifications.
type EnableParams struct{}

// Enable enables domain notifications.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Overlay.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnable, nil, nil)
}

// GetHighlightObjectForTestParams for testing.
type GetHighlightObjectForTestParams struct {
	NodeID cdp.NodeID `json:"nodeId"` // Id of the node to get highlight object for.
}

// GetHighlightObjectForTest for testing.
//
// parameters:
//   nodeID - Id of the node to get highlight object for.
func GetHighlightObjectForTest(nodeID cdp.NodeID) *GetHighlightObjectForTestParams {
	return &GetHighlightObjectForTestParams{
		NodeID: nodeID,
	}
}

// GetHighlightObjectForTestReturns return values.
type GetHighlightObjectForTestReturns struct {
	Highlight easyjson.RawMessage `json:"highlight,omitempty"`
}

// Do executes Overlay.getHighlightObjectForTest against the provided context.
//
// returns:
//   highlight - Highlight data for the node.
func (p *GetHighlightObjectForTestParams) Do(ctxt context.Context, h cdp.Executor) (highlight easyjson.RawMessage, err error) {
	// execute
	var res GetHighlightObjectForTestReturns
	err = h.Execute(ctxt, CommandGetHighlightObjectForTest, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Highlight, nil
}

// HideHighlightParams hides any highlight.
type HideHighlightParams struct{}

// HideHighlight hides any highlight.
func HideHighlight() *HideHighlightParams {
	return &HideHighlightParams{}
}

// Do executes Overlay.hideHighlight against the provided context.
func (p *HideHighlightParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandHideHighlight, nil, nil)
}

// HighlightFrameParams highlights owner element of the frame with given id.
type HighlightFrameParams struct {
	FrameID             cdp.FrameID `json:"frameId"`                       // Identifier of the frame to highlight.
	ContentColor        *cdp.RGBA   `json:"contentColor,omitempty"`        // The content box highlight fill color (default: transparent).
	ContentOutlineColor *cdp.RGBA   `json:"contentOutlineColor,omitempty"` // The content box highlight outline color (default: transparent).
}

// HighlightFrame highlights owner element of the frame with given id.
//
// parameters:
//   frameID - Identifier of the frame to highlight.
func HighlightFrame(frameID cdp.FrameID) *HighlightFrameParams {
	return &HighlightFrameParams{
		FrameID: frameID,
	}
}

// WithContentColor the content box highlight fill color (default:
// transparent).
func (p HighlightFrameParams) WithContentColor(contentColor *cdp.RGBA) *HighlightFrameParams {
	p.ContentColor = contentColor
	return &p
}

// WithContentOutlineColor the content box highlight outline color (default:
// transparent).
func (p HighlightFrameParams) WithContentOutlineColor(contentOutlineColor *cdp.RGBA) *HighlightFrameParams {
	p.ContentOutlineColor = contentOutlineColor
	return &p
}

// Do executes Overlay.highlightFrame against the provided context.
func (p *HighlightFrameParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandHighlightFrame, p, nil)
}

// HighlightNodeParams highlights DOM node with given id or with the given
// JavaScript object wrapper. Either nodeId or objectId must be specified.
type HighlightNodeParams struct {
	HighlightConfig *HighlightConfig       `json:"highlightConfig"`         // A descriptor for the highlight appearance.
	NodeID          cdp.NodeID             `json:"nodeId,omitempty"`        // Identifier of the node to highlight.
	BackendNodeID   cdp.BackendNodeID      `json:"backendNodeId,omitempty"` // Identifier of the backend node to highlight.
	ObjectID        runtime.RemoteObjectID `json:"objectId,omitempty"`      // JavaScript object id of the node to be highlighted.
	Selector        string                 `json:"selector,omitempty"`      // Selectors to highlight relevant nodes.
}

// HighlightNode highlights DOM node with given id or with the given
// JavaScript object wrapper. Either nodeId or objectId must be specified.
//
// parameters:
//   highlightConfig - A descriptor for the highlight appearance.
func HighlightNode(highlightConfig *HighlightConfig) *HighlightNodeParams {
	return &HighlightNodeParams{
		HighlightConfig: highlightConfig,
	}
}

// WithNodeID identifier of the node to highlight.
func (p HighlightNodeParams) WithNodeID(nodeID cdp.NodeID) *HighlightNodeParams {
	p.NodeID = nodeID
	return &p
}

// WithBackendNodeID identifier of the backend node to highlight.
func (p HighlightNodeParams) WithBackendNodeID(backendNodeID cdp.BackendNodeID) *HighlightNodeParams {
	p.BackendNodeID = backendNodeID
	return &p
}

// WithObjectID JavaScript object id of the node to be highlighted.
func (p HighlightNodeParams) WithObjectID(objectID runtime.RemoteObjectID) *HighlightNodeParams {
	p.ObjectID = objectID
	return &p
}

// WithSelector selectors to highlight relevant nodes.
func (p HighlightNodeParams) WithSelector(selector string) *HighlightNodeParams {
	p.Selector = selector
	return &p
}

// Do executes Overlay.highlightNode against the provided context.
func (p *HighlightNodeParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandHighlightNode, p, nil)
}

// HighlightQuadParams highlights given quad. Coordinates are absolute with
// respect to the main frame viewport.
type HighlightQuadParams struct {
	Quad         dom.Quad  `json:"quad"`                   // Quad to highlight
	Color        *cdp.RGBA `json:"color,omitempty"`        // The highlight fill color (default: transparent).
	OutlineColor *cdp.RGBA `json:"outlineColor,omitempty"` // The highlight outline color (default: transparent).
}

// HighlightQuad highlights given quad. Coordinates are absolute with respect
// to the main frame viewport.
//
// parameters:
//   quad - Quad to highlight
func HighlightQuad(quad dom.Quad) *HighlightQuadParams {
	return &HighlightQuadParams{
		Quad: quad,
	}
}

// WithColor the highlight fill color (default: transparent).
func (p HighlightQuadParams) WithColor(color *cdp.RGBA) *HighlightQuadParams {
	p.Color = color
	return &p
}

// WithOutlineColor the highlight outline color (default: transparent).
func (p HighlightQuadParams) WithOutlineColor(outlineColor *cdp.RGBA) *HighlightQuadParams {
	p.OutlineColor = outlineColor
	return &p
}

// Do executes Overlay.highlightQuad against the provided context.
func (p *HighlightQuadParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandHighlightQuad, p, nil)
}

// HighlightRectParams highlights given rectangle. Coordinates are absolute
// with respect to the main frame viewport.
type HighlightRectParams struct {
	X            int64     `json:"x"`                      // X coordinate
	Y            int64     `json:"y"`                      // Y coordinate
	Width        int64     `json:"width"`                  // Rectangle width
	Height       int64     `json:"height"`                 // Rectangle height
	Color        *cdp.RGBA `json:"color,omitempty"`        // The highlight fill color (default: transparent).
	OutlineColor *cdp.RGBA `json:"outlineColor,omitempty"` // The highlight outline color (default: transparent).
}

// HighlightRect highlights given rectangle. Coordinates are absolute with
// respect to the main frame viewport.
//
// parameters:
//   x - X coordinate
//   y - Y coordinate
//   width - Rectangle width
//   height - Rectangle height
func HighlightRect(x int64, y int64, width int64, height int64) *HighlightRectParams {
	return &HighlightRectParams{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

// WithColor the highlight fill color (default: transparent).
func (p HighlightRectParams) WithColor(color *cdp.RGBA) *HighlightRectParams {
	p.Color = color
	return &p
}

// WithOutlineColor the highlight outline color (default: transparent).
func (p HighlightRectParams) WithOutlineColor(outlineColor *cdp.RGBA) *HighlightRectParams {
	p.OutlineColor = outlineColor
	return &p
}

// Do executes Overlay.highlightRect against the provided context.
func (p *HighlightRectParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandHighlightRect, p, nil)
}

// SetInspectModeParams enters the 'inspect' mode. In this mode, elements
// that user is hovering over are highlighted. Backend then generates
// 'inspectNodeRequested' event upon element selection.
type SetInspectModeParams struct {
	Mode            InspectMode      `json:"mode"`                      // Set an inspection mode.
	HighlightConfig *HighlightConfig `json:"highlightConfig,omitempty"` // A descriptor for the highlight appearance of hovered-over nodes. May be omitted if enabled == false.
}

// SetInspectMode enters the 'inspect' mode. In this mode, elements that user
// is hovering over are highlighted. Backend then generates
// 'inspectNodeRequested' event upon element selection.
//
// parameters:
//   mode - Set an inspection mode.
func SetInspectMode(mode InspectMode) *SetInspectModeParams {
	return &SetInspectModeParams{
		Mode: mode,
	}
}

// WithHighlightConfig a descriptor for the highlight appearance of
// hovered-over nodes. May be omitted if enabled == false.
func (p SetInspectModeParams) WithHighlightConfig(highlightConfig *HighlightConfig) *SetInspectModeParams {
	p.HighlightConfig = highlightConfig
	return &p
}

// Do executes Overlay.setInspectMode against the provided context.
func (p *SetInspectModeParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetInspectMode, p, nil)
}

// SetShowAdHighlightsParams highlights owner element of all frames detected
// to be ads.
type SetShowAdHighlightsParams struct {
	Show bool `json:"show"` // True for showing ad highlights
}

// SetShowAdHighlights highlights owner element of all frames detected to be
// ads.
//
// parameters:
//   show - True for showing ad highlights
func SetShowAdHighlights(show bool) *SetShowAdHighlightsParams {
	return &SetShowAdHighlightsParams{
		Show: show,
	}
}

// Do executes Overlay.setShowAdHighlights against the provided context.
func (p *SetShowAdHighlightsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowAdHighlights, p, nil)
}

// SetPausedInDebuggerMessageParams [no description].
type SetPausedInDebuggerMessageParams struct {
	Message string `json:"message,omitempty"` // The message to display, also triggers resume and step over controls.
}

// SetPausedInDebuggerMessage [no description].
//
// parameters:
func SetPausedInDebuggerMessage() *SetPausedInDebuggerMessageParams {
	return &SetPausedInDebuggerMessageParams{}
}

// WithMessage the message to display, also triggers resume and step over
// controls.
func (p SetPausedInDebuggerMessageParams) WithMessage(message string) *SetPausedInDebuggerMessageParams {
	p.Message = message
	return &p
}

// Do executes Overlay.setPausedInDebuggerMessage against the provided context.
func (p *SetPausedInDebuggerMessageParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetPausedInDebuggerMessage, p, nil)
}

// SetShowDebugBordersParams requests that backend shows debug borders on
// layers.
type SetShowDebugBordersParams struct {
	Show bool `json:"show"` // True for showing debug borders
}

// SetShowDebugBorders requests that backend shows debug borders on layers.
//
// parameters:
//   show - True for showing debug borders
func SetShowDebugBorders(show bool) *SetShowDebugBordersParams {
	return &SetShowDebugBordersParams{
		Show: show,
	}
}

// Do executes Overlay.setShowDebugBorders against the provided context.
func (p *SetShowDebugBordersParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowDebugBorders, p, nil)
}

// SetShowFPSCounterParams requests that backend shows the FPS counter.
type SetShowFPSCounterParams struct {
	Show bool `json:"show"` // True for showing the FPS counter
}

// SetShowFPSCounter requests that backend shows the FPS counter.
//
// parameters:
//   show - True for showing the FPS counter
func SetShowFPSCounter(show bool) *SetShowFPSCounterParams {
	return &SetShowFPSCounterParams{
		Show: show,
	}
}

// Do executes Overlay.setShowFPSCounter against the provided context.
func (p *SetShowFPSCounterParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowFPSCounter, p, nil)
}

// SetShowPaintRectsParams requests that backend shows paint rectangles.
type SetShowPaintRectsParams struct {
	Result bool `json:"result"` // True for showing paint rectangles
}

// SetShowPaintRects requests that backend shows paint rectangles.
//
// parameters:
//   result - True for showing paint rectangles
func SetShowPaintRects(result bool) *SetShowPaintRectsParams {
	return &SetShowPaintRectsParams{
		Result: result,
	}
}

// Do executes Overlay.setShowPaintRects against the provided context.
func (p *SetShowPaintRectsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowPaintRects, p, nil)
}

// SetShowScrollBottleneckRectsParams requests that backend shows scroll
// bottleneck rects.
type SetShowScrollBottleneckRectsParams struct {
	Show bool `json:"show"` // True for showing scroll bottleneck rects
}

// SetShowScrollBottleneckRects requests that backend shows scroll bottleneck
// rects.
//
// parameters:
//   show - True for showing scroll bottleneck rects
func SetShowScrollBottleneckRects(show bool) *SetShowScrollBottleneckRectsParams {
	return &SetShowScrollBottleneckRectsParams{
		Show: show,
	}
}

// Do executes Overlay.setShowScrollBottleneckRects against the provided context.
func (p *SetShowScrollBottleneckRectsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowScrollBottleneckRects, p, nil)
}

// SetShowHitTestBordersParams requests that backend shows hit-test borders
// on layers.
type SetShowHitTestBordersParams struct {
	Show bool `json:"show"` // True for showing hit-test borders
}

// SetShowHitTestBorders requests that backend shows hit-test borders on
// layers.
//
// parameters:
//   show - True for showing hit-test borders
func SetShowHitTestBorders(show bool) *SetShowHitTestBordersParams {
	return &SetShowHitTestBordersParams{
		Show: show,
	}
}

// Do executes Overlay.setShowHitTestBorders against the provided context.
func (p *SetShowHitTestBordersParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowHitTestBorders, p, nil)
}

// SetShowViewportSizeOnResizeParams paints viewport size upon main frame
// resize.
type SetShowViewportSizeOnResizeParams struct {
	Show bool `json:"show"` // Whether to paint size or not.
}

// SetShowViewportSizeOnResize paints viewport size upon main frame resize.
//
// parameters:
//   show - Whether to paint size or not.
func SetShowViewportSizeOnResize(show bool) *SetShowViewportSizeOnResizeParams {
	return &SetShowViewportSizeOnResizeParams{
		Show: show,
	}
}

// Do executes Overlay.setShowViewportSizeOnResize against the provided context.
func (p *SetShowViewportSizeOnResizeParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetShowViewportSizeOnResize, p, nil)
}

// Command names.
const (
	CommandDisable                      = "Overlay.disable"
	CommandEnable                       = "Overlay.enable"
	CommandGetHighlightObjectForTest    = "Overlay.getHighlightObjectForTest"
	CommandHideHighlight                = "Overlay.hideHighlight"
	CommandHighlightFrame               = "Overlay.highlightFrame"
	CommandHighlightNode                = "Overlay.highlightNode"
	CommandHighlightQuad                = "Overlay.highlightQuad"
	CommandHighlightRect                = "Overlay.highlightRect"
	CommandSetInspectMode               = "Overlay.setInspectMode"
	CommandSetShowAdHighlights          = "Overlay.setShowAdHighlights"
	CommandSetPausedInDebuggerMessage   = "Overlay.setPausedInDebuggerMessage"
	CommandSetShowDebugBorders          = "Overlay.setShowDebugBorders"
	CommandSetShowFPSCounter            = "Overlay.setShowFPSCounter"
	CommandSetShowPaintRects            = "Overlay.setShowPaintRects"
	CommandSetShowScrollBottleneckRects = "Overlay.setShowScrollBottleneckRects"
	CommandSetShowHitTestBorders        = "Overlay.setShowHitTestBorders"
	CommandSetShowViewportSizeOnResize  = "Overlay.setShowViewportSizeOnResize"
)
