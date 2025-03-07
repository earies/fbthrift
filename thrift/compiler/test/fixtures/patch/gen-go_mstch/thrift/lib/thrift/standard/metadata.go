// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package standard // [[[ program thrift source path ]]]

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
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
            )
    premadeThriftType_standard_Uri = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("standard.Uri").
            SetUnderlyingType(premadeThriftType_string),
            )
    premadeThriftType_binary = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BINARY_TYPE.Ptr(),
            )
    premadeThriftType_standard_ByteString = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("standard.ByteString").
            SetUnderlyingType(premadeThriftType_binary),
            )
    premadeThriftType_standard_Void = metadata.NewThriftType().SetTEnum(
        metadata.NewThriftEnumType().
            SetName("standard.Void"),
            )
    premadeThriftType_standard_TypeUri = metadata.NewThriftType().SetTUnion(
        metadata.NewThriftUnionType().
            SetName("standard.TypeUri"),
            )
)

var structMetadatas = []*metadata.ThriftStruct{
    metadata.NewThriftStruct().
    SetName("standard.TypeUri").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("uri").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Uri),
            metadata.NewThriftField().
    SetId(2).
    SetName("typeHashPrefixSha2_256").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_ByteString),
            metadata.NewThriftField().
    SetId(3).
    SetName("scopedName").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
    metadata.NewThriftStruct().
    SetName("standard.TypeName").
    SetIsUnion(true).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("boolType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(2).
    SetName("byteType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(3).
    SetName("i16Type").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(4).
    SetName("i32Type").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(5).
    SetName("i64Type").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(6).
    SetName("floatType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(7).
    SetName("doubleType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(8).
    SetName("stringType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(9).
    SetName("binaryType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(10).
    SetName("enumType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_TypeUri),
            metadata.NewThriftField().
    SetId(11).
    SetName("structType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_TypeUri),
            metadata.NewThriftField().
    SetId(12).
    SetName("unionType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_TypeUri),
            metadata.NewThriftField().
    SetId(13).
    SetName("exceptionType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_TypeUri),
            metadata.NewThriftField().
    SetId(14).
    SetName("listType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(15).
    SetName("setType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(16).
    SetName("mapType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_Void),
            metadata.NewThriftField().
    SetId(17).
    SetName("typedefType").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_TypeUri),
        },
    ),
}

var exceptionMetadatas = []*metadata.ThriftException{
}

var enumMetadatas = []*metadata.ThriftEnum{
    metadata.NewThriftEnum().
    SetName("standard.Void").
    SetElements(
        map[int32]string{
            0: "Unused",
        },
    ),
    metadata.NewThriftEnum().
    SetName("standard.StandardProtocol").
    SetElements(
        map[int32]string{
            0: "Custom",
            1: "Binary",
            2: "Compact",
            3: "Json",
            4: "SimpleJson",
        },
    ),
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
