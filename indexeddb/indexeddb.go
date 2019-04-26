// Package indexeddb provides the Chrome DevTools Protocol
// commands, types, and events for the IndexedDB domain.
//
// Generated by the cdproto-gen command.
package indexeddb

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/nbzx/cdproto/cdp"
)

// ClearObjectStoreParams clears all entries from an object store.
type ClearObjectStoreParams struct {
	SecurityOrigin  string `json:"securityOrigin"`  // Security origin.
	DatabaseName    string `json:"databaseName"`    // Database name.
	ObjectStoreName string `json:"objectStoreName"` // Object store name.
}

// ClearObjectStore clears all entries from an object store.
//
// parameters:
//   securityOrigin - Security origin.
//   databaseName - Database name.
//   objectStoreName - Object store name.
func ClearObjectStore(securityOrigin string, databaseName string, objectStoreName string) *ClearObjectStoreParams {
	return &ClearObjectStoreParams{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
	}
}

// Do executes IndexedDB.clearObjectStore against the provided context.
func (p *ClearObjectStoreParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandClearObjectStore, p, nil)
}

// DeleteDatabaseParams deletes a database.
type DeleteDatabaseParams struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
	DatabaseName   string `json:"databaseName"`   // Database name.
}

// DeleteDatabase deletes a database.
//
// parameters:
//   securityOrigin - Security origin.
//   databaseName - Database name.
func DeleteDatabase(securityOrigin string, databaseName string) *DeleteDatabaseParams {
	return &DeleteDatabaseParams{
		SecurityOrigin: securityOrigin,
		DatabaseName:   databaseName,
	}
}

// Do executes IndexedDB.deleteDatabase against the provided context.
func (p *DeleteDatabaseParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDeleteDatabase, p, nil)
}

// DeleteObjectStoreEntriesParams delete a range of entries from an object
// store.
type DeleteObjectStoreEntriesParams struct {
	SecurityOrigin  string    `json:"securityOrigin"`
	DatabaseName    string    `json:"databaseName"`
	ObjectStoreName string    `json:"objectStoreName"`
	KeyRange        *KeyRange `json:"keyRange"` // Range of entry keys to delete
}

// DeleteObjectStoreEntries delete a range of entries from an object store.
//
// parameters:
//   securityOrigin
//   databaseName
//   objectStoreName
//   keyRange - Range of entry keys to delete
func DeleteObjectStoreEntries(securityOrigin string, databaseName string, objectStoreName string, keyRange *KeyRange) *DeleteObjectStoreEntriesParams {
	return &DeleteObjectStoreEntriesParams{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
		KeyRange:        keyRange,
	}
}

// Do executes IndexedDB.deleteObjectStoreEntries against the provided context.
func (p *DeleteObjectStoreEntriesParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDeleteObjectStoreEntries, p, nil)
}

// DisableParams disables events from backend.
type DisableParams struct{}

// Disable disables events from backend.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes IndexedDB.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables events from backend.
type EnableParams struct{}

// Enable enables events from backend.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes IndexedDB.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context) (err error) {
	return cdp.Execute(ctxt, CommandEnable, nil, nil)
}

// RequestDataParams requests data from object store or index.
type RequestDataParams struct {
	SecurityOrigin  string    `json:"securityOrigin"`     // Security origin.
	DatabaseName    string    `json:"databaseName"`       // Database name.
	ObjectStoreName string    `json:"objectStoreName"`    // Object store name.
	IndexName       string    `json:"indexName"`          // Index name, empty string for object store data requests.
	SkipCount       int64     `json:"skipCount"`          // Number of records to skip.
	PageSize        int64     `json:"pageSize"`           // Number of records to fetch.
	KeyRange        *KeyRange `json:"keyRange,omitempty"` // Key range.
}

// RequestData requests data from object store or index.
//
// parameters:
//   securityOrigin - Security origin.
//   databaseName - Database name.
//   objectStoreName - Object store name.
//   indexName - Index name, empty string for object store data requests.
//   skipCount - Number of records to skip.
//   pageSize - Number of records to fetch.
func RequestData(securityOrigin string, databaseName string, objectStoreName string, indexName string, skipCount int64, pageSize int64) *RequestDataParams {
	return &RequestDataParams{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
		IndexName:       indexName,
		SkipCount:       skipCount,
		PageSize:        pageSize,
	}
}

// WithKeyRange key range.
func (p RequestDataParams) WithKeyRange(keyRange *KeyRange) *RequestDataParams {
	p.KeyRange = keyRange
	return &p
}

// RequestDataReturns return values.
type RequestDataReturns struct {
	ObjectStoreDataEntries []*DataEntry `json:"objectStoreDataEntries,omitempty"` // Array of object store data entries.
	HasMore                bool         `json:"hasMore,omitempty"`                // If true, there are more entries to fetch in the given range.
}

// Do executes IndexedDB.requestData against the provided context.
//
// returns:
//   objectStoreDataEntries - Array of object store data entries.
//   hasMore - If true, there are more entries to fetch in the given range.
func (p *RequestDataParams) Do(ctxt context.Context) (objectStoreDataEntries []*DataEntry, hasMore bool, err error) {
	// execute
	var res RequestDataReturns
	err = cdp.Execute(ctxt, CommandRequestData, p, &res)
	if err != nil {
		return nil, false, err
	}

	return res.ObjectStoreDataEntries, res.HasMore, nil
}

// GetMetadataParams gets metadata of an object store.
type GetMetadataParams struct {
	SecurityOrigin  string `json:"securityOrigin"`  // Security origin.
	DatabaseName    string `json:"databaseName"`    // Database name.
	ObjectStoreName string `json:"objectStoreName"` // Object store name.
}

// GetMetadata gets metadata of an object store.
//
// parameters:
//   securityOrigin - Security origin.
//   databaseName - Database name.
//   objectStoreName - Object store name.
func GetMetadata(securityOrigin string, databaseName string, objectStoreName string) *GetMetadataParams {
	return &GetMetadataParams{
		SecurityOrigin:  securityOrigin,
		DatabaseName:    databaseName,
		ObjectStoreName: objectStoreName,
	}
}

// GetMetadataReturns return values.
type GetMetadataReturns struct {
	EntriesCount      float64 `json:"entriesCount,omitempty"`      // the entries count
	KeyGeneratorValue float64 `json:"keyGeneratorValue,omitempty"` // the current value of key generator, to become the next inserted key into the object store. Valid if objectStore.autoIncrement is true.
}

// Do executes IndexedDB.getMetadata against the provided context.
//
// returns:
//   entriesCount - the entries count
//   keyGeneratorValue - the current value of key generator, to become the next inserted key into the object store. Valid if objectStore.autoIncrement is true.
func (p *GetMetadataParams) Do(ctxt context.Context) (entriesCount float64, keyGeneratorValue float64, err error) {
	// execute
	var res GetMetadataReturns
	err = cdp.Execute(ctxt, CommandGetMetadata, p, &res)
	if err != nil {
		return 0, 0, err
	}

	return res.EntriesCount, res.KeyGeneratorValue, nil
}

// RequestDatabaseParams requests database with given name in given frame.
type RequestDatabaseParams struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
	DatabaseName   string `json:"databaseName"`   // Database name.
}

// RequestDatabase requests database with given name in given frame.
//
// parameters:
//   securityOrigin - Security origin.
//   databaseName - Database name.
func RequestDatabase(securityOrigin string, databaseName string) *RequestDatabaseParams {
	return &RequestDatabaseParams{
		SecurityOrigin: securityOrigin,
		DatabaseName:   databaseName,
	}
}

// RequestDatabaseReturns return values.
type RequestDatabaseReturns struct {
	DatabaseWithObjectStores *DatabaseWithObjectStores `json:"databaseWithObjectStores,omitempty"` // Database with an array of object stores.
}

// Do executes IndexedDB.requestDatabase against the provided context.
//
// returns:
//   databaseWithObjectStores - Database with an array of object stores.
func (p *RequestDatabaseParams) Do(ctxt context.Context) (databaseWithObjectStores *DatabaseWithObjectStores, err error) {
	// execute
	var res RequestDatabaseReturns
	err = cdp.Execute(ctxt, CommandRequestDatabase, p, &res)
	if err != nil {
		return nil, err
	}

	return res.DatabaseWithObjectStores, nil
}

// RequestDatabaseNamesParams requests database names for given security
// origin.
type RequestDatabaseNamesParams struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
}

// RequestDatabaseNames requests database names for given security origin.
//
// parameters:
//   securityOrigin - Security origin.
func RequestDatabaseNames(securityOrigin string) *RequestDatabaseNamesParams {
	return &RequestDatabaseNamesParams{
		SecurityOrigin: securityOrigin,
	}
}

// RequestDatabaseNamesReturns return values.
type RequestDatabaseNamesReturns struct {
	DatabaseNames []string `json:"databaseNames,omitempty"` // Database names for origin.
}

// Do executes IndexedDB.requestDatabaseNames against the provided context.
//
// returns:
//   databaseNames - Database names for origin.
func (p *RequestDatabaseNamesParams) Do(ctxt context.Context) (databaseNames []string, err error) {
	// execute
	var res RequestDatabaseNamesReturns
	err = cdp.Execute(ctxt, CommandRequestDatabaseNames, p, &res)
	if err != nil {
		return nil, err
	}

	return res.DatabaseNames, nil
}

// Command names.
const (
	CommandClearObjectStore         = "IndexedDB.clearObjectStore"
	CommandDeleteDatabase           = "IndexedDB.deleteDatabase"
	CommandDeleteObjectStoreEntries = "IndexedDB.deleteObjectStoreEntries"
	CommandDisable                  = "IndexedDB.disable"
	CommandEnable                   = "IndexedDB.enable"
	CommandRequestData              = "IndexedDB.requestData"
	CommandGetMetadata              = "IndexedDB.getMetadata"
	CommandRequestDatabase          = "IndexedDB.requestDatabase"
	CommandRequestDatabaseNames     = "IndexedDB.requestDatabaseNames"
)
