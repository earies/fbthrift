// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package scope // [[[ program thrift source path ]]]

import (
    "maps"

    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = maps.Copy[map[int]int, map[int]int]
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
)

var structMetadatas = []*metadata.ThriftStruct{
    metadata.NewThriftStruct().
    SetName("scope.Transitive").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Program").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Struct").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Union").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Exception").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Field").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Typedef").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Service").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Interaction").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Function").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.EnumValue").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Const").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Enum").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Structured").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Interface").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.RootDefinition").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("scope.Definition").
    SetIsUnion(false),
}

var exceptionMetadatas = []*metadata.ThriftException{
}

var enumMetadatas = []*metadata.ThriftEnum{
}

var serviceMetadatas = []*metadata.ThriftService{
}

// GetThriftMetadata returns complete Thrift metadata for current and imported packages.
func GetThriftMetadata() *metadata.ThriftMetadata {
    allEnums := GetEnumsMetadata()
    allStructs := GetStructsMetadata()
    allExceptions := GetExceptionsMetadata()
    allServices := GetServicesMetadata()

    return metadata.NewThriftMetadata().
        SetEnums(allEnums).
        SetStructs(allStructs).
        SetExceptions(allExceptions).
        SetServices(allServices)
}

// GetEnumsMetadata returns Thrift metadata for enums in the current and recursively included packages.
func GetEnumsMetadata() map[string]*metadata.ThriftEnum {
    allEnumsMap := make(map[string]*metadata.ThriftEnum)

    // Add enum metadatas from the current program...
    for _, enumMetadata := range enumMetadatas {
        allEnumsMap[enumMetadata.GetName()] = enumMetadata
    }

    // ...now add enum metadatas from recursively included programs.

    return allEnumsMap
}

// GetStructsMetadata returns Thrift metadata for structs in the current and recursively included packages.
func GetStructsMetadata() map[string]*metadata.ThriftStruct {
    allStructsMap := make(map[string]*metadata.ThriftStruct)

    // Add struct metadatas from the current program...
    for _, structMetadata := range structMetadatas {
        allStructsMap[structMetadata.GetName()] = structMetadata
    }

    // ...now add struct metadatas from recursively included programs.

    return allStructsMap
}

// GetExceptionsMetadata returns Thrift metadata for exceptions in the current and recursively included packages.
func GetExceptionsMetadata() map[string]*metadata.ThriftException {
    allExceptionsMap := make(map[string]*metadata.ThriftException)

    // Add exception metadatas from the current program...
    for _, exceptionMetadata := range exceptionMetadatas {
        allExceptionsMap[exceptionMetadata.GetName()] = exceptionMetadata
    }

    // ...now add exception metadatas from recursively included programs.

    return allExceptionsMap
}

// GetServicesMetadata returns Thrift metadata for services in the current and recursively included packages.
func GetServicesMetadata() map[string]*metadata.ThriftService {
    allServicesMap := make(map[string]*metadata.ThriftService)

    // Add service metadatas from the current program...
    for _, serviceMetadata := range serviceMetadatas {
        allServicesMap[serviceMetadata.GetName()] = serviceMetadata
    }

    // ...now add service metadatas from recursively included programs.

    return allServicesMap
}
