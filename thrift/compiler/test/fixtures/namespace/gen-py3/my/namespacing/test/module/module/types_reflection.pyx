#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#


import folly.iobuf as _fbthrift_iobuf

from thrift.py3.reflection cimport (
    NumberType as __NumberType,
    StructType as __StructType,
    Qualifier as __Qualifier,
)


cimport my.namespacing.test.module.module.types as _my_namespacing_test_module_module_types

from thrift.py3.types cimport (
    constant_shared_ptr,
    default_inst,
)


cdef __StructSpec get_reflection__Foo():
    cdef _my_namespacing_test_module_module_types.Foo defaults = _my_namespacing_test_module_module_types.Foo._fbthrift_create(
        constant_shared_ptr[_my_namespacing_test_module_module_types.cFoo](
            default_inst[_my_namespacing_test_module_module_types.cFoo]()
        )
    )
    cdef __StructSpec spec = __StructSpec._fbthrift_create(
        name="Foo",
        kind=__StructType.STRUCT,
        annotations={
        },
    )
    spec.add_field(
        __FieldSpec._fbthrift_create(
            id=1,
            name="MyInt",
            py_name="MyInt",
            type=int,
            kind=__NumberType.I64,
            qualifier=__Qualifier.UNQUALIFIED,
            default=None,
            annotations={
            },
        ),
    )
    return spec
