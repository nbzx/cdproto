package domsnapshot

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/nbzx/cdproto/cdp"
	"github.com/nbzx/cdproto/dom"
	"github.com/nbzx/cdproto/domdebugger"
)

// DOMNode a Node in the DOM tree.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-DOMNode
type DOMNode struct {
	NodeType             cdp.NodeType                 `json:"nodeType"`                       // Node's nodeType.
	NodeName             string                       `json:"nodeName"`                       // Node's nodeName.
	NodeValue            string                       `json:"nodeValue"`                      // Node's nodeValue.
	TextValue            string                       `json:"textValue,omitempty"`            // Only set for textarea elements, contains the text value.
	InputValue           string                       `json:"inputValue,omitempty"`           // Only set for input elements, contains the input's associated text value.
	InputChecked         bool                         `json:"inputChecked,omitempty"`         // Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected       bool                         `json:"optionSelected,omitempty"`       // Only set for option elements, indicates if the element has been selected
	BackendNodeID        cdp.BackendNodeID            `json:"backendNodeId"`                  // Node's id, corresponds to DOM.Node.backendNodeId.
	ChildNodeIndexes     []int64                      `json:"childNodeIndexes,omitempty"`     // The indexes of the node's child nodes in the domNodes array returned by getSnapshot, if any.
	Attributes           []*NameValue                 `json:"attributes,omitempty"`           // Attributes of an Element node.
	PseudoElementIndexes []int64                      `json:"pseudoElementIndexes,omitempty"` // Indexes of pseudo elements associated with this node in the domNodes array returned by getSnapshot, if any.
	LayoutNodeIndex      int64                        `json:"layoutNodeIndex,omitempty"`      // The index of the node's related layout tree node in the layoutTreeNodes array returned by getSnapshot, if any.
	DocumentURL          string                       `json:"documentURL,omitempty"`          // Document URL that Document or FrameOwner node points to.
	BaseURL              string                       `json:"baseURL,omitempty"`              // Base URL that Document or FrameOwner node uses for URL completion.
	ContentLanguage      string                       `json:"contentLanguage,omitempty"`      // Only set for documents, contains the document's content language.
	DocumentEncoding     string                       `json:"documentEncoding,omitempty"`     // Only set for documents, contains the document's character set encoding.
	PublicID             string                       `json:"publicId,omitempty"`             // DocumentType node's publicId.
	SystemID             string                       `json:"systemId,omitempty"`             // DocumentType node's systemId.
	FrameID              cdp.FrameID                  `json:"frameId,omitempty"`              // Frame ID for frame owner elements and also for the document node.
	ContentDocumentIndex int64                        `json:"contentDocumentIndex,omitempty"` // The index of a frame owner element's content document in the domNodes array returned by getSnapshot, if any.
	PseudoType           cdp.PseudoType               `json:"pseudoType,omitempty"`           // Type of a pseudo element node.
	ShadowRootType       cdp.ShadowRootType           `json:"shadowRootType,omitempty"`       // Shadow root type.
	IsClickable          bool                         `json:"isClickable,omitempty"`          // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	EventListeners       []*domdebugger.EventListener `json:"eventListeners,omitempty"`       // Details of the node's event listeners, if any.
	CurrentSourceURL     string                       `json:"currentSourceURL,omitempty"`     // The selected url for nodes with a srcset attribute.
	OriginURL            string                       `json:"originURL,omitempty"`            // The url of the script (if any) that generates this node.
	ScrollOffsetX        float64                      `json:"scrollOffsetX,omitempty"`        // Scroll offsets, set when this node is a Document.
	ScrollOffsetY        float64                      `json:"scrollOffsetY,omitempty"`
}

// InlineTextBox details of post layout rendered text positions. The exact
// layout should not be regarded as stable and may change between versions.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-InlineTextBox
type InlineTextBox struct {
	BoundingBox         *dom.Rect `json:"boundingBox"`         // The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	StartCharacterIndex int64     `json:"startCharacterIndex"` // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	NumCharacters       int64     `json:"numCharacters"`       // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}

// LayoutTreeNode details of an element in the DOM tree with a LayoutObject.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-LayoutTreeNode
type LayoutTreeNode struct {
	DomNodeIndex      int64            `json:"domNodeIndex"`                // The index of the related DOM node in the domNodes array returned by getSnapshot.
	BoundingBox       *dom.Rect        `json:"boundingBox"`                 // The bounding box in document coordinates. Note that scroll offset of the document is ignored.
	LayoutText        string           `json:"layoutText,omitempty"`        // Contents of the LayoutText, if any.
	InlineTextNodes   []*InlineTextBox `json:"inlineTextNodes,omitempty"`   // The post-layout inline text nodes, if any.
	StyleIndex        int64            `json:"styleIndex,omitempty"`        // Index into the computedStyles array returned by getSnapshot.
	PaintOrder        int64            `json:"paintOrder,omitempty"`        // Global paint order index, which is determined by the stacking order of the nodes. Nodes that are painted together will have the same index. Only provided if includePaintOrder in getSnapshot was true.
	IsStackingContext bool             `json:"isStackingContext,omitempty"` // Set to true to indicate the element begins a new stacking context.
}

// ComputedStyle a subset of the full ComputedStyle as defined by the request
// whitelist.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-ComputedStyle
type ComputedStyle struct {
	Properties []*NameValue `json:"properties"` // Name/value pairs of computed style properties.
}

// NameValue a name/value pair.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-NameValue
type NameValue struct {
	Name  string `json:"name"`  // Attribute/property name.
	Value string `json:"value"` // Attribute/property value.
}

// StringIndex index of the string in the strings table.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-StringIndex
type StringIndex int64

// Int64 returns the StringIndex as int64 value.
func (t StringIndex) Int64() int64 {
	return int64(t)
}

// ArrayOfStrings index of the string in the strings table.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-ArrayOfStrings
type ArrayOfStrings []int64

// RareStringData data that is only present on rare nodes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-RareStringData
type RareStringData struct {
	Index []int64       `json:"index"`
	Value []StringIndex `json:"value"`
}

// RareBooleanData [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-RareBooleanData
type RareBooleanData struct {
	Index []int64 `json:"index"`
}

// RareIntegerData [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-RareIntegerData
type RareIntegerData struct {
	Index []int64 `json:"index"`
	Value []int64 `json:"value"`
}

// Rectangle [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-Rectangle
type Rectangle []float64

// DocumentSnapshot document snapshot.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-DocumentSnapshot
type DocumentSnapshot struct {
	DocumentURL     StringIndex         `json:"documentURL"`             // Document URL that Document or FrameOwner node points to.
	BaseURL         StringIndex         `json:"baseURL"`                 // Base URL that Document or FrameOwner node uses for URL completion.
	ContentLanguage StringIndex         `json:"contentLanguage"`         // Contains the document's content language.
	EncodingName    StringIndex         `json:"encodingName"`            // Contains the document's character set encoding.
	PublicID        StringIndex         `json:"publicId"`                // DocumentType node's publicId.
	SystemID        StringIndex         `json:"systemId"`                // DocumentType node's systemId.
	FrameID         StringIndex         `json:"frameId"`                 // Frame ID for frame owner elements and also for the document node.
	Nodes           *NodeTreeSnapshot   `json:"nodes"`                   // A table with dom nodes.
	Layout          *LayoutTreeSnapshot `json:"layout"`                  // The nodes in the layout tree.
	TextBoxes       *TextBoxSnapshot    `json:"textBoxes"`               // The post-layout inline text nodes.
	ScrollOffsetX   float64             `json:"scrollOffsetX,omitempty"` // Scroll offsets.
	ScrollOffsetY   float64             `json:"scrollOffsetY,omitempty"`
}

// NodeTreeSnapshot table containing nodes.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-NodeTreeSnapshot
type NodeTreeSnapshot struct {
	ParentIndex          []int64             `json:"parentIndex,omitempty"`          // Parent node index.
	NodeType             []int64             `json:"nodeType,omitempty"`             // Node's nodeType.
	NodeName             []StringIndex       `json:"nodeName,omitempty"`             // Node's nodeName.
	NodeValue            []StringIndex       `json:"nodeValue,omitempty"`            // Node's nodeValue.
	BackendNodeID        []cdp.BackendNodeID `json:"backendNodeId,omitempty"`        // Node's id, corresponds to DOM.Node.backendNodeId.
	Attributes           []ArrayOfStrings    `json:"attributes,omitempty"`           // Attributes of an Element node. Flatten name, value pairs.
	TextValue            *RareStringData     `json:"textValue,omitempty"`            // Only set for textarea elements, contains the text value.
	InputValue           *RareStringData     `json:"inputValue,omitempty"`           // Only set for input elements, contains the input's associated text value.
	InputChecked         *RareBooleanData    `json:"inputChecked,omitempty"`         // Only set for radio and checkbox input elements, indicates if the element has been checked
	OptionSelected       *RareBooleanData    `json:"optionSelected,omitempty"`       // Only set for option elements, indicates if the element has been selected
	ContentDocumentIndex *RareIntegerData    `json:"contentDocumentIndex,omitempty"` // The index of the document in the list of the snapshot documents.
	PseudoType           *RareStringData     `json:"pseudoType,omitempty"`           // Type of a pseudo element node.
	IsClickable          *RareBooleanData    `json:"isClickable,omitempty"`          // Whether this DOM node responds to mouse clicks. This includes nodes that have had click event listeners attached via JavaScript as well as anchor tags that naturally navigate when clicked.
	CurrentSourceURL     *RareStringData     `json:"currentSourceURL,omitempty"`     // The selected url for nodes with a srcset attribute.
	OriginURL            *RareStringData     `json:"originURL,omitempty"`            // The url of the script (if any) that generates this node.
}

// LayoutTreeSnapshot details of an element in the DOM tree with a
// LayoutObject.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-LayoutTreeSnapshot
type LayoutTreeSnapshot struct {
	NodeIndex        []int64          `json:"nodeIndex"`        // The index of the related DOM node in the domNodes array returned by getSnapshot.
	Styles           []ArrayOfStrings `json:"styles"`           // Index into the computedStyles array returned by captureSnapshot.
	Bounds           []Rectangle      `json:"bounds"`           // The absolute position bounding box.
	Text             []StringIndex    `json:"text"`             // Contents of the LayoutText, if any.
	StackingContexts *RareBooleanData `json:"stackingContexts"` // Stacking context information.
}

// TextBoxSnapshot details of post layout rendered text positions. The exact
// layout should not be regarded as stable and may change between versions.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/DOMSnapshot#type-TextBoxSnapshot
type TextBoxSnapshot struct {
	LayoutIndex []int64     `json:"layoutIndex"` // Intex of th elayout tree node that owns this box collection.
	Bounds      []Rectangle `json:"bounds"`      // The absolute position bounding box.
	Start       []int64     `json:"start"`       // The starting index in characters, for this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
	Length      []int64     `json:"length"`      // The number of characters in this post layout textbox substring. Characters that would be represented as a surrogate pair in UTF-16 have length 2.
}
