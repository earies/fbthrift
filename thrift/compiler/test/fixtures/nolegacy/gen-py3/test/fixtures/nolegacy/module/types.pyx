#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#
cimport cython as __cython
from cpython.object cimport PyTypeObject, Py_LT, Py_LE, Py_EQ, Py_NE, Py_GT, Py_GE
from libcpp.memory cimport shared_ptr, make_shared, unique_ptr, make_unique
from libcpp.string cimport string
from libcpp cimport bool as cbool
from libcpp.iterator cimport inserter as cinserter
from cpython cimport bool as pbool
from cython.operator cimport dereference as deref, preincrement as inc, address as ptr_address
import thrift.py3.types
from thrift.py3.types import _IsSet as _fbthrift_IsSet
cimport thrift.py3.types
cimport thrift.py3.exceptions
from thrift.py3.std_libcpp cimport sv_to_str as __sv_to_str, string_view as __cstring_view
from thrift.py3.types cimport (
    cSetOp as __cSetOp,
    richcmp as __richcmp,
    set_op as __set_op,
    setcmp as __setcmp,
    list_index as __list_index,
    list_count as __list_count,
    list_slice as __list_slice,
    list_getitem as __list_getitem,
    set_iter as __set_iter,
    map_iter as __map_iter,
    map_contains as __map_contains,
    map_getitem as __map_getitem,
    reference_shared_ptr as __reference_shared_ptr,
    get_field_name_by_index as __get_field_name_by_index,
    reset_field as __reset_field,
    translate_cpp_enum_to_python,
    SetMetaClass as __SetMetaClass,
    const_pointer_cast,
    constant_shared_ptr,
    NOTSET as __NOTSET,
    EnumData as __EnumData,
    EnumFlagsData as __EnumFlagsData,
    UnionTypeEnumData as __UnionTypeEnumData,
    createEnumDataForUnionType as __createEnumDataForUnionType,
)
cimport thrift.py3.std_libcpp as std_libcpp
cimport thrift.py3.serializer as serializer
import folly.iobuf as _fbthrift_iobuf
from folly.optional cimport cOptional
from folly.memory cimport to_shared_ptr as __to_shared_ptr
from folly.range cimport Range as __cRange

import sys
from collections.abc import Sequence, Set, Mapping, Iterable
import weakref as __weakref
import builtins as _builtins

cimport test.fixtures.nolegacy.module.types_reflection as _types_reflection


cdef __EnumData __TestEnum_enum_data  = __EnumData._fbthrift_create(thrift.py3.types.createEnumData[cTestEnum](), TestEnum)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __TestEnumMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __TestEnum_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __TestEnum_enum_data.get_all_names()

    def __len__(cls):
        return __TestEnum_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __TestEnum_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class TestEnum(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __TestEnum_enum_data.get_by_name(name)


    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        EnumMetadata[cTestEnum].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.TestEnum"

    def _to_python(self):
        import importlib
        python_types = importlib.import_module(
            "test.fixtures.nolegacy.module.thrift_types"
        )
        return python_types.TestEnum(self.value)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        return self.value


__SetMetaClass(<PyTypeObject*> TestEnum, <PyTypeObject*> __TestEnumMeta)



cdef __UnionTypeEnumData __TestUnion_union_type_enum_data  = __UnionTypeEnumData._fbthrift_create(
    __createEnumDataForUnionType[cTestUnion](),
    __TestUnionType,
)


@__cython.internal
@__cython.auto_pickle(False)
cdef class __TestUnion_Union_TypeMeta(thrift.py3.types.EnumMeta):
    def _fbthrift_get_by_value(cls, int value):
        return __TestUnion_union_type_enum_data.get_by_value(value)

    def _fbthrift_get_all_names(cls):
        return __TestUnion_union_type_enum_data.get_all_names()

    def __len__(cls):
        return __TestUnion_union_type_enum_data.size()

    def __getattribute__(cls, str name not None):
        if name.startswith("__") or name.startswith("_fbthrift_") or name == "mro":
            return super().__getattribute__(name)
        return __TestUnion_union_type_enum_data.get_by_name(name)


@__cython.final
@__cython.auto_pickle(False)
cdef class __TestUnionType(thrift.py3.types.CompiledEnum):
    cdef get_by_name(self, str name):
        return __TestUnion_union_type_enum_data.get_by_name(name)


__SetMetaClass(<PyTypeObject*> __TestUnionType, <PyTypeObject*> __TestUnion_Union_TypeMeta)


@__cython.auto_pickle(False)
cdef class TestError(thrift.py3.exceptions.GeneratedError):
    def __init__(TestError self, *args, **kwargs):
        self._cpp_obj = make_shared[cTestError]()
        self._fields_setter = _fbthrift_types_fields.__TestError_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__( *args, **kwargs)

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("TestError", {
          "test_enum": deref(self._cpp_obj).test_enum_ref().has_value(),
          "code": deref(self._cpp_obj).code_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cTestError] cpp_obj):
        __fbthrift_inst = <TestError>TestError.__new__(TestError, (<bytes>deref(cpp_obj).what()).decode('utf-8'))
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        _builtins.Exception.__init__(__fbthrift_inst, *(v for _, v in __fbthrift_inst))
        return __fbthrift_inst

    cdef inline test_enum_impl(self):

        if self.__fbthrift_cached_test_enum is None:
            self.__fbthrift_cached_test_enum = translate_cpp_enum_to_python(TestEnum, <int>(deref(self._cpp_obj).test_enum_ref().value()))
        return self.__fbthrift_cached_test_enum

    @property
    def test_enum(self):
        return self.test_enum_impl()

    cdef inline code_impl(self):

        return deref(self._cpp_obj).code_ref().value()

    @property
    def code(self):
        return self.code_impl()


    def __hash__(TestError self):
        return super().__hash__()

    def __repr__(TestError self):
        return super().__repr__()

    def __str__(TestError self):
        return super().__str__()


    def __copy__(TestError self):
        cdef shared_ptr[cTestError] cpp_obj = make_shared[cTestError](
            deref(self._cpp_obj)
        )
        return TestError._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cTestError](
            self._cpp_obj,
            (<TestError>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__TestError()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        ExceptionMetadata[cTestError].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.TestError"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cTestError](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 2

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(TestError self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cTestError](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(TestError self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cTestError]()
        with nogil:
            needed = serializer.cdeserialize[cTestError](buf, self._cpp_obj.get(), proto)
        return needed

    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "test.fixtures.nolegacy.module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.TestError, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.TestError, self)
@__cython.auto_pickle(False)
cdef class TestMixin(thrift.py3.types.Struct):
    def __init__(TestMixin self, **kwargs):
        self._cpp_obj = make_shared[cTestMixin]()
        self._fields_setter = _fbthrift_types_fields.__TestMixin_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(TestMixin self, **kwargs):
        if not kwargs:
            return self
        cdef TestMixin __fbthrift_inst = TestMixin.__new__(TestMixin)
        __fbthrift_inst._cpp_obj = make_shared[cTestMixin](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__TestMixin_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("TestMixin", {
          "field1": deref(self._cpp_obj).field1_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cTestMixin] cpp_obj):
        __fbthrift_inst = <TestMixin>TestMixin.__new__(TestMixin)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline field1_impl(self):

        return (<bytes>deref(self._cpp_obj).field1_ref().value()).decode('UTF-8')

    @property
    def field1(self):
        return self.field1_impl()


    def __hash__(TestMixin self):
        return super().__hash__()

    def __repr__(TestMixin self):
        return super().__repr__()

    def __str__(TestMixin self):
        return super().__str__()


    def __copy__(TestMixin self):
        cdef shared_ptr[cTestMixin] cpp_obj = make_shared[cTestMixin](
            deref(self._cpp_obj)
        )
        return TestMixin._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cTestMixin](
            self._cpp_obj,
            (<TestMixin>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__TestMixin()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cTestMixin].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.TestMixin"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cTestMixin](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 1

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(TestMixin self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cTestMixin](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(TestMixin self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cTestMixin]()
        with nogil:
            needed = serializer.cdeserialize[cTestMixin](buf, self._cpp_obj.get(), proto)
        return needed

    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "test.fixtures.nolegacy.module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.TestMixin, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.TestMixin, self)
@__cython.auto_pickle(False)
cdef class TestStruct(thrift.py3.types.Struct):
    def __init__(TestStruct self, **kwargs):
        self._cpp_obj = make_shared[cTestStruct]()
        self._fields_setter = _fbthrift_types_fields.__TestStruct_FieldsSetter._fbthrift_create(self._cpp_obj.get())
        super().__init__(**kwargs)

    def __call__(TestStruct self, **kwargs):
        if not kwargs:
            return self
        cdef TestStruct __fbthrift_inst = TestStruct.__new__(TestStruct)
        __fbthrift_inst._cpp_obj = make_shared[cTestStruct](deref(self._cpp_obj))
        __fbthrift_inst._fields_setter = _fbthrift_types_fields.__TestStruct_FieldsSetter._fbthrift_create(__fbthrift_inst._cpp_obj.get())
        for __fbthrift_name, _fbthrift_value in kwargs.items():
            __fbthrift_inst._fbthrift_set_field(__fbthrift_name, _fbthrift_value)
        return __fbthrift_inst

    cdef void _fbthrift_set_field(self, str name, object value) except *:
        self._fields_setter.set_field(name.encode("utf-8"), value)

    cdef object _fbthrift_isset(self):
        return _fbthrift_IsSet("TestStruct", {
          "bar": deref(self._cpp_obj).bar_ref().has_value(),
          "baropt": deref(self._cpp_obj).baropt_ref().has_value(),
          "test_error": deref(self._cpp_obj).test_error_ref().has_value(),
          "test_mixin": deref(self._cpp_obj).test_mixin_ref().has_value(),
        })

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cTestStruct] cpp_obj):
        __fbthrift_inst = <TestStruct>TestStruct.__new__(TestStruct)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        return __fbthrift_inst

    cdef inline bar_impl(self):

        return (<bytes>deref(self._cpp_obj).bar_ref().value()).decode('UTF-8')

    @property
    def bar(self):
        return self.bar_impl()

    cdef inline baropt_impl(self):
        if not deref(self._cpp_obj).baropt_ref().has_value():
            return None

        return (<bytes>deref(self._cpp_obj).baropt_ref().value_unchecked()).decode('UTF-8')

    @property
    def baropt(self):
        return self.baropt_impl()

    cdef inline test_error_impl(self):

        if self.__fbthrift_cached_test_error is None:
            self.__fbthrift_cached_test_error = TestError._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).test_error_ref().ref(), self._cpp_obj))
        return self.__fbthrift_cached_test_error

    @property
    def test_error(self):
        return self.test_error_impl()

    cdef inline test_mixin_impl(self):

        if self.__fbthrift_cached_test_mixin is None:
            self.__fbthrift_cached_test_mixin = TestMixin._fbthrift_create(__reference_shared_ptr(deref(self._cpp_obj).test_mixin_ref().ref(), self._cpp_obj))
        return self.__fbthrift_cached_test_mixin

    @property
    def test_mixin(self):
        return self.test_mixin_impl()

    cdef inline field1_impl(self):

        return (<bytes>deref(self._cpp_obj).field1_ref().value()).decode('UTF-8')

    @property
    def field1(self):
        return self.field1_impl()


    def __hash__(TestStruct self):
        return super().__hash__()

    def __repr__(TestStruct self):
        return super().__repr__()

    def __str__(TestStruct self):
        return super().__str__()


    def __copy__(TestStruct self):
        cdef shared_ptr[cTestStruct] cpp_obj = make_shared[cTestStruct](
            deref(self._cpp_obj)
        )
        return TestStruct._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cTestStruct](
            self._cpp_obj,
            (<TestStruct>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__TestStruct()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cTestStruct].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.TestStruct"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cTestStruct](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 4

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(TestStruct self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cTestStruct](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(TestStruct self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cTestStruct]()
        with nogil:
            needed = serializer.cdeserialize[cTestStruct](buf, self._cpp_obj.get(), proto)
        return needed

    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "test.fixtures.nolegacy.module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.TestStruct, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.TestStruct, self)


@__cython.auto_pickle(False)
cdef class TestUnion(thrift.py3.types.Union):
    Type = __TestUnionType

    def __init__(
        self, *,
        TestEnum enumVal=None,
        TestStruct structVal=None
    ):
        self._cpp_obj = __to_shared_ptr(cmove(TestUnion._make_instance(
          NULL,
          enumVal,
          structVal,
        )))
        self._load_cache()

    @staticmethod
    def fromValue(value):
        if value is None:
            return TestUnion()
        if isinstance(value, TestEnum):
            return TestUnion(enumVal=value)
        if isinstance(value, TestStruct):
            return TestUnion(structVal=value)
        raise ValueError(f"Unable to derive correct union field for value: {value}")

    @staticmethod
    cdef unique_ptr[cTestUnion] _make_instance(
        cTestUnion* base_instance,
        TestEnum enumVal,
        TestStruct structVal
    ) except *:
        cdef unique_ptr[cTestUnion] c_inst = make_unique[cTestUnion]()
        cdef bint any_set = False
        if enumVal is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).enumVal_ref().assign(<cTestEnum><int>enumVal)
            any_set = True
        if structVal is not None:
            if any_set:
                raise TypeError("At most one field may be set when initializing a union")
            deref(c_inst).structVal_ref().assign(deref((<TestStruct?> structVal)._cpp_obj))
            any_set = True
        # in C++ you don't have to call move(), but this doesn't translate
        # into a C++ return statement, so you do here
        return cmove(c_inst)

    @staticmethod
    cdef _fbthrift_create(shared_ptr[cTestUnion] cpp_obj):
        __fbthrift_inst = <TestUnion>TestUnion.__new__(TestUnion)
        __fbthrift_inst._cpp_obj = cmove(cpp_obj)
        __fbthrift_inst._load_cache()
        return __fbthrift_inst

    @property
    def enumVal(self):
        if self.type.value != 1:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not enumVal')
        return self.value

    @property
    def structVal(self):
        if self.type.value != 2:
            raise AttributeError(f'Union contains a value of type {self.type.name}, not structVal')
        return self.value


    def __hash__(TestUnion self):
        return  super().__hash__()

    cdef _load_cache(TestUnion self):
        self.type = TestUnion.Type(<int>(deref(self._cpp_obj).getType()))
        cdef int type = self.type.value
        if type == 0:    # Empty
            self.value = None
        elif type == 1:
            self.value = translate_cpp_enum_to_python(TestEnum, <int>deref(self._cpp_obj).enumVal_ref().value())
        elif type == 2:
            self.value = TestStruct._fbthrift_create(make_shared[cTestStruct](deref(self._cpp_obj).structVal_ref().value()))

    def __copy__(TestUnion self):
        cdef shared_ptr[cTestUnion] cpp_obj = make_shared[cTestUnion](
            deref(self._cpp_obj)
        )
        return TestUnion._fbthrift_create(cmove(cpp_obj))

    def __richcmp__(self, other, int op):
        r = self._fbthrift_cmp_sametype(other, op)
        return __richcmp[cTestUnion](
            self._cpp_obj,
            (<TestUnion>other)._cpp_obj,
            op,
        ) if r is None else r

    @staticmethod
    def __get_reflection__():
        return _types_reflection.get_reflection__TestUnion()

    @staticmethod
    def __get_metadata__():
        cdef __fbthrift_cThriftMetadata meta
        StructMetadata[cTestUnion].gen(meta)
        return __MetadataBox.box(cmove(meta))

    @staticmethod
    def __get_thrift_name__():
        return "module.TestUnion"

    @classmethod
    def _fbthrift_get_field_name_by_index(cls, idx):
        return __sv_to_str(__get_field_name_by_index[cTestUnion](idx))

    @classmethod
    def _fbthrift_get_struct_size(cls):
        return 2

    cdef _fbthrift_iobuf.IOBuf _fbthrift_serialize(TestUnion self, __Protocol proto):
        cdef unique_ptr[_fbthrift_iobuf.cIOBuf] data
        with nogil:
            data = cmove(serializer.cserialize[cTestUnion](self._cpp_obj.get(), proto))
        return _fbthrift_iobuf.from_unique_ptr(cmove(data))

    cdef cuint32_t _fbthrift_deserialize(TestUnion self, const _fbthrift_iobuf.cIOBuf* buf, __Protocol proto) except? 0:
        cdef cuint32_t needed
        self._cpp_obj = make_shared[cTestUnion]()
        with nogil:
            needed = serializer.cdeserialize[cTestUnion](buf, self._cpp_obj.get(), proto)
        # force a cache reload since the underlying data's changed
        self._load_cache()
        return needed

    def _to_python(self):
        import importlib
        import thrift.python.converter
        python_types = importlib.import_module(
            "test.fixtures.nolegacy.module.thrift_types"
        )
        return thrift.python.converter.to_python_struct(python_types.TestUnion, self)

    def _to_py3(self):
        return self

    def _to_py_deprecated(self):
        import importlib
        import thrift.util.converter
        py_deprecated_types = importlib.import_module("module.ttypes")
        return thrift.util.converter.to_py_struct(py_deprecated_types.TestUnion, self)
