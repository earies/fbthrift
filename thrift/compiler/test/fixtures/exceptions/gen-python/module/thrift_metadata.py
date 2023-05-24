#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations

import apache.thrift.metadata.thrift_types as _fbthrift_metadata


# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_Fiery(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.Fiery"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
        _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="message", is_optional=False, structured_annotations=[
        ]),
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
        ])
    new_struct = metadata_struct(exceptions=struct_dict)

     # message

    return new_struct
def gen_metadata_exception_Fiery() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_Fiery(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_Serious(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.Serious"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
        _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="sonnet", is_optional=False, structured_annotations=[
        ]),
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
        ])
    new_struct = metadata_struct(exceptions=struct_dict)

     # sonnet

    return new_struct
def gen_metadata_exception_Serious() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_Serious(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_ComplexFieldNames(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.ComplexFieldNames"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
        _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="error_message", is_optional=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftField(id=2, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="internal_error_message", is_optional=False, structured_annotations=[
        ]),
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
        ])
    new_struct = metadata_struct(exceptions=struct_dict)

     # error_message
     # internal_error_message

    return new_struct
def gen_metadata_exception_ComplexFieldNames() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_ComplexFieldNames(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_CustomFieldNames(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.CustomFieldNames"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
        _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="error_message", is_optional=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftField(id=2, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="internal_error_message", is_optional=False, structured_annotations=[
        ]),
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
        ])
    new_struct = metadata_struct(exceptions=struct_dict)

     # error_message
     # internal_error_message

    return new_struct
def gen_metadata_exception_CustomFieldNames() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_CustomFieldNames(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_ExceptionWithPrimitiveField(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.ExceptionWithPrimitiveField"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
        _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="message", is_optional=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftField(id=2, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_I32_TYPE), name="error_code", is_optional=False, structured_annotations=[
        ]),
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
        ])
    new_struct = metadata_struct(exceptions=struct_dict)

     # message
     # error_code

    return new_struct
def gen_metadata_exception_ExceptionWithPrimitiveField() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_ExceptionWithPrimitiveField(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_ExceptionWithStructuredAnnotation(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.ExceptionWithStructuredAnnotation"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
        _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), name="message_field", is_optional=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftField(id=2, type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_I32_TYPE), name="error_code", is_optional=False, structured_annotations=[
        ]),
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
            _fbthrift_metadata.ThriftConstStruct(type=_fbthrift_metadata.ThriftStructType(name="thrift.ExceptionMessage"), fields= { "field": _fbthrift_metadata.ThriftConstValue(cv_string="message_field"),  }),
        ])
    new_struct = metadata_struct(exceptions=struct_dict)

     # message_field
     # error_code

    return new_struct
def gen_metadata_exception_ExceptionWithStructuredAnnotation() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_ExceptionWithStructuredAnnotation(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

# TODO (ffrancet): This general pattern can be optimized by using tuples and dicts
# instead of re-generating thrift structs
def _fbthrift_gen_metadata_exception_Banal(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.Banal"

    if qualified_name in metadata_struct.exceptions:
        return metadata_struct
    fields = [
    ]
    struct_dict = dict(metadata_struct.exceptions)
    struct_dict[qualified_name] = _fbthrift_metadata.ThriftException(name=qualified_name, fields=fields,
        structured_annotations=[
        ])
    new_struct = metadata_struct(exceptions=struct_dict)


    return new_struct
def gen_metadata_exception_Banal() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_exception_Banal(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))


def gen_metadata_service_Raiser() -> _fbthrift_metadata.ThriftMetadata:
    return _fbthrift_gen_metadata_service_Raiser(_fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={}))

def _fbthrift_gen_metadata_service_Raiser(metadata_struct: _fbthrift_metadata.ThriftMetadata) -> _fbthrift_metadata.ThriftMetadata:
    qualified_name = "module.Raiser"
    
    if qualified_name in metadata_struct.services:
        return metadata_struct
    
    functions = [
        _fbthrift_metadata.ThriftFunction(name="doBland", return_type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_VOID_TYPE), arguments=[
        ], exceptions = [
        ], is_oneway=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftFunction(name="doRaise", return_type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_VOID_TYPE), arguments=[
        ], exceptions = [
            _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_struct=_fbthrift_metadata.ThriftStructType(name="module.Banal")), name="b", is_optional=False, structured_annotations=[
            ]),
            _fbthrift_metadata.ThriftField(id=2, type=_fbthrift_metadata.ThriftType(t_struct=_fbthrift_metadata.ThriftStructType(name="module.Fiery")), name="f", is_optional=False, structured_annotations=[
            ]),
            _fbthrift_metadata.ThriftField(id=3, type=_fbthrift_metadata.ThriftType(t_struct=_fbthrift_metadata.ThriftStructType(name="module.Serious")), name="s", is_optional=False, structured_annotations=[
            ]),
        ], is_oneway=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftFunction(name="get200", return_type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), arguments=[
        ], exceptions = [
        ], is_oneway=False, structured_annotations=[
        ]),
        _fbthrift_metadata.ThriftFunction(name="get500", return_type=_fbthrift_metadata.ThriftType(t_primitive=_fbthrift_metadata.ThriftPrimitiveType.THRIFT_STRING_TYPE), arguments=[
        ], exceptions = [
            _fbthrift_metadata.ThriftField(id=1, type=_fbthrift_metadata.ThriftType(t_struct=_fbthrift_metadata.ThriftStructType(name="module.Fiery")), name="f", is_optional=False, structured_annotations=[
            ]),
            _fbthrift_metadata.ThriftField(id=2, type=_fbthrift_metadata.ThriftType(t_struct=_fbthrift_metadata.ThriftStructType(name="module.Banal")), name="b", is_optional=False, structured_annotations=[
            ]),
            _fbthrift_metadata.ThriftField(id=3, type=_fbthrift_metadata.ThriftType(t_struct=_fbthrift_metadata.ThriftStructType(name="module.Serious")), name="s", is_optional=False, structured_annotations=[
            ]),
        ], is_oneway=False, structured_annotations=[
        ]),
    ]
    
    service_dict = dict(metadata_struct.services)
    service_dict[qualified_name] = _fbthrift_metadata.ThriftService(name=qualified_name, functions=functions,  structured_annotations=[
    ])
    new_struct = metadata_struct(services=service_dict)
    
    
    
     # return value
    
    
    
    new_struct = _fbthrift_gen_metadata_exception_Banal(new_struct) # b
    new_struct = _fbthrift_gen_metadata_exception_Fiery(new_struct) # f
    new_struct = _fbthrift_gen_metadata_exception_Serious(new_struct) # s
    
     # return value
    
    
    
    
     # return value
    
    
    
    new_struct = _fbthrift_gen_metadata_exception_Fiery(new_struct) # f
    new_struct = _fbthrift_gen_metadata_exception_Banal(new_struct) # b
    new_struct = _fbthrift_gen_metadata_exception_Serious(new_struct) # s
    
     # return value
    
    
    return new_struct

def _fbthrift_metadata_service_response_Raiser() -> _fbthrift_metadata.ThriftServiceMetadataResponse:
    metadata = gen_metadata_service_Raiser()
    context = _fbthrift_metadata.ThriftServiceContext(service_info=metadata.services["module.Raiser"], module=_fbthrift_metadata.ThriftModuleContext(name="module"))
    services = [_fbthrift_metadata.ThriftServiceContextRef(module=_fbthrift_metadata.ThriftModuleContext(name=name.split('.')[0]), service_name=name) for name in metadata.services]
    return _fbthrift_metadata.ThriftServiceMetadataResponse(metadata=metadata,context=context,services=services)




def getThriftModuleMetadata() -> _fbthrift_metadata.ThriftMetadata:
    meta = _fbthrift_metadata.ThriftMetadata(structs={}, enums={}, exceptions={}, services={})
    meta = _fbthrift_gen_metadata_exception_Fiery(meta)
    meta = _fbthrift_gen_metadata_exception_Serious(meta)
    meta = _fbthrift_gen_metadata_exception_ComplexFieldNames(meta)
    meta = _fbthrift_gen_metadata_exception_CustomFieldNames(meta)
    meta = _fbthrift_gen_metadata_exception_ExceptionWithPrimitiveField(meta)
    meta = _fbthrift_gen_metadata_exception_ExceptionWithStructuredAnnotation(meta)
    meta = _fbthrift_gen_metadata_exception_Banal(meta)
    meta = _fbthrift_gen_metadata_service_Raiser(meta)
    return meta
