#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
from cpython cimport bool as pbool, int as pint, float as pfloat

cimport folly.iobuf as _fbthrift_iobuf

cimport thrift.py3.builder


cimport foo.types as _foo_types

cdef class Fields_Builder(thrift.py3.builder.StructBuilder):
    cdef public str injected_field
    cdef public str injected_structured_annotation_field
    cdef public str injected_unstructured_annotation_field


