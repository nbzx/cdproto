// Package headlessexperimental provides the Chrome DevTools Protocol
// commands, types, and events for the HeadlessExperimental domain.
//
// This domain provides experimental commands only supported in headless
// mode.
//
// Generated by the cdproto-gen command.
package headlessexperimental

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"
	"encoding/base64"

	"github.com/nbzx/cdproto/cdp"
)

// BeginFrameParams sends a BeginFrame to the target and returns when the
// frame was completed. Optionally captures a screenshot from the resulting
// frame. Requires that the target was created with enabled BeginFrameControl.
// Designed for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gl/3zHXhB for more background.
type BeginFrameParams struct {
	FrameTimeTicks   float64           `json:"frameTimeTicks,omitempty"`   // Timestamp of this BeginFrame in Renderer TimeTicks (milliseconds of uptime). If not set, the current time will be used.
	Interval         float64           `json:"interval,omitempty"`         // The interval between BeginFrames that is reported to the compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e. about 16.666 milliseconds.
	NoDisplayUpdates bool              `json:"noDisplayUpdates,omitempty"` // Whether updates should not be committed and drawn onto the display. False by default. If true, only side effects of the BeginFrame will be run, such as layout and animations, but any visual updates may not be visible on the display or in screenshots.
	Screenshot       *ScreenshotParams `json:"screenshot,omitempty"`       // If set, a screenshot of the frame will be captured and returned in the response. Otherwise, no screenshot will be captured. Note that capturing a screenshot can fail, for example, during renderer initialization. In such a case, no screenshot data will be returned.
}

// BeginFrame sends a BeginFrame to the target and returns when the frame was
// completed. Optionally captures a screenshot from the resulting frame.
// Requires that the target was created with enabled BeginFrameControl. Designed
// for use with --run-all-compositor-stages-before-draw, see also
// https://goo.gl/3zHXhB for more background.
//
// parameters:
func BeginFrame() *BeginFrameParams {
	return &BeginFrameParams{}
}

// WithFrameTimeTicks timestamp of this BeginFrame in Renderer TimeTicks
// (milliseconds of uptime). If not set, the current time will be used.
func (p BeginFrameParams) WithFrameTimeTicks(frameTimeTicks float64) *BeginFrameParams {
	p.FrameTimeTicks = frameTimeTicks
	return &p
}

// WithInterval the interval between BeginFrames that is reported to the
// compositor, in milliseconds. Defaults to a 60 frames/second interval, i.e.
// about 16.666 milliseconds.
func (p BeginFrameParams) WithInterval(interval float64) *BeginFrameParams {
	p.Interval = interval
	return &p
}

// WithNoDisplayUpdates whether updates should not be committed and drawn
// onto the display. False by default. If true, only side effects of the
// BeginFrame will be run, such as layout and animations, but any visual updates
// may not be visible on the display or in screenshots.
func (p BeginFrameParams) WithNoDisplayUpdates(noDisplayUpdates bool) *BeginFrameParams {
	p.NoDisplayUpdates = noDisplayUpdates
	return &p
}

// WithScreenshot if set, a screenshot of the frame will be captured and
// returned in the response. Otherwise, no screenshot will be captured. Note
// that capturing a screenshot can fail, for example, during renderer
// initialization. In such a case, no screenshot data will be returned.
func (p BeginFrameParams) WithScreenshot(screenshot *ScreenshotParams) *BeginFrameParams {
	p.Screenshot = screenshot
	return &p
}

// BeginFrameReturns return values.
type BeginFrameReturns struct {
	HasDamage      bool   `json:"hasDamage,omitempty"`      // Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
	ScreenshotData string `json:"screenshotData,omitempty"` // Base64-encoded image data of the screenshot, if one was requested and successfully taken.
}

// Do executes HeadlessExperimental.beginFrame against the provided context.
//
// returns:
//   hasDamage - Whether the BeginFrame resulted in damage and, thus, a new frame was committed to the display. Reported for diagnostic uses, may be removed in the future.
//   screenshotData - Base64-encoded image data of the screenshot, if one was requested and successfully taken.
func (p *BeginFrameParams) Do(ctxt context.Context) (hasDamage bool, screenshotData []byte, err error) {
	// execute
	var res BeginFrameReturns
	err = cdp.Execute(ctxt, CommandBeginFrame, p, &res)
	if err != nil {
		return false, nil, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.ScreenshotData)
	if err != nil {
		return false, nil, err
	}
	return res.HasDamage, dec, nil
}

// DisableParams disables headless events for the target.
type DisableParams struct{}

// Disable disables headless events for the target.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes HeadlessExperimental.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables headless events for the target.
type EnableParams struct{}

// Enable enables headless events for the target.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes HeadlessExperimental.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandEnable, nil, nil)
}

// Command names.
const (
	CommandBeginFrame = "HeadlessExperimental.beginFrame"
	CommandDisable    = "HeadlessExperimental.disable"
	CommandEnable     = "HeadlessExperimental.enable"
)
