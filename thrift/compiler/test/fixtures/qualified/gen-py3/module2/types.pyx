#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libcpp.memory cimport shared_ptr, make_shared, unique_ptr, make_unique
from libcpp.string cimport string
from libcpp cimport bool as cbool
from libcpp.iterator cimport inserter as cinserter
from cpython cimport bool as pbool
from libc.stdint cimport int8_t, int16_t, int32_t, int64_t, uint32_t
from cython.operator cimport dereference as deref, preincrement as inc, address as ptr_address
import thrift.py3.types
cimport thrift.py3.types
cimport thrift.py3.exceptions
from thrift.py3.types import NOTSET as __NOTSET
from thrift.py3.types cimport translate_cpp_enum_to_python
cimport thrift.py3.std_libcpp as std_libcpp
from thrift.py3.serializer import Protocol as __Protocol
cimport thrift.py3.serializer as serializer
from thrift.py3.serializer import deserialize, serialize
import folly.iobuf as __iobuf
from folly.optional cimport cOptional

import sys
import itertools
from collections import Sequence, Set, Mapping, Iterable
import enum as __enum
import warnings
import builtins as _builtins
cimport module0.types as _module0_types
import module0.types as _module0_types
cimport module1.types as _module1_types
import module1.types as _module1_types




cdef cStruct _Struct_defaults = cStruct()

cdef class Struct(thrift.py3.types.Struct):

    def __init__(
        Struct self, *,
        _module0_types.Struct first=None,
        _module1_types.Struct second=None
    ):
        self._cpp_obj = move(Struct._make_instance(
          NULL,
          first,
          second,
        ))

    def __call__(
        Struct self,
        first=__NOTSET,
        second=__NOTSET
    ):
        changes = any((
            first is not __NOTSET,

            second is not __NOTSET,
        ))

        if not changes:
            return self

        if None is not first is not __NOTSET:
            if not isinstance(first, _module0_types.Struct):
                raise TypeError(f'first is not a { _module0_types.Struct !r}.')

        if None is not second is not __NOTSET:
            if not isinstance(second, _module1_types.Struct):
                raise TypeError(f'second is not a { _module1_types.Struct !r}.')

        inst = <Struct>Struct.__new__(Struct)
        inst._cpp_obj = move(Struct._make_instance(
          self._cpp_obj.get(),
          first,
          second,
        ))
        return inst

    @staticmethod
    cdef unique_ptr[cStruct] _make_instance(
        cStruct* base_instance,
        object first,
        object second
    ) except *:
        cdef unique_ptr[cStruct] c_inst
        if base_instance:
            c_inst = make_unique[cStruct](deref(base_instance))
        else:
            c_inst = make_unique[cStruct]()

        if base_instance:
            # Convert None's to default value. (or unset)
            if first is None:
                deref(c_inst).first = _Struct_defaults.first
                deref(c_inst).__isset.first = False
                pass
            elif first is __NOTSET:
                first = None

            if second is None:
                deref(c_inst).second = _Struct_defaults.second
                deref(c_inst).__isset.second = False
                pass
            elif second is __NOTSET:
                second = None

        if first is not None:
            deref(c_inst).first = deref((<_module0_types.Struct?> first)._cpp_obj)
            deref(c_inst).__isset.first = True
        if second is not None:
            deref(c_inst).second = deref((<_module1_types.Struct?> second)._cpp_obj)
            deref(c_inst).__isset.second = True
        # in C++ you don't have to call move(), but this doesn't translate
        # into a C++ return statement, so you do here
        return move_unique(c_inst)

    def __iter__(self):
        yield 'first', self.first
        yield 'second', self.second

    def __bool__(self):
        return True or True

    @staticmethod
    cdef create(shared_ptr[cStruct] cpp_obj):
        inst = <Struct>Struct.__new__(Struct)
        inst._cpp_obj = cpp_obj
        return inst

    @property
    def first(self):

        if self.__first is None:
            self.__first = _module0_types.Struct.create(make_shared[_module0_types.cStruct](deref(self._cpp_obj).first))
        return self.__first

    @property
    def second(self):

        if self.__second is None:
            self.__second = _module1_types.Struct.create(make_shared[_module1_types.cStruct](deref(self._cpp_obj).second))
        return self.__second


    def __hash__(Struct self):
        if not self.__hash:
            self.__hash = hash((
            self.first,
            self.second,
            ))
        return self.__hash

    def __repr__(Struct self):
        return f'Struct(first={repr(self.first)}, second={repr(self.second)})'
    def __richcmp__(self, other, op):
        cdef int cop = op
        if cop not in (2, 3):
            raise TypeError("unorderable types: {}, {}".format(self, other))
        if not (
                isinstance(self, Struct) and
                isinstance(other, Struct)):
            if cop == 2:  # different types are never equal
                return False
            else:         # different types are always notequal
                return True

        cdef cStruct cself = deref((<Struct>self)._cpp_obj)
        cdef cStruct cother = deref((<Struct>other)._cpp_obj)
        cdef cbool cmp = cself == cother
        if cop == 2:
            return cmp
        return not cmp

    cdef __iobuf.IOBuf _serialize(Struct self, proto):
        cdef __iobuf.cIOBufQueue queue = __iobuf.cIOBufQueue(__iobuf.cacheChainLength())
        cdef cStruct* cpp_obj = self._cpp_obj.get()
        if proto is __Protocol.COMPACT:
            with nogil:
                serializer.CompactSerialize[cStruct](deref(cpp_obj), &queue, serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.BINARY:
            with nogil:
                serializer.BinarySerialize[cStruct](deref(cpp_obj), &queue, serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.JSON:
            with nogil:
                serializer.JSONSerialize[cStruct](deref(cpp_obj), &queue, serializer.SHARE_EXTERNAL_BUFFER)
        return __iobuf.from_unique_ptr(queue.move())

    cdef uint32_t _deserialize(Struct self, const __iobuf.cIOBuf* buf, proto) except? 0:
        cdef uint32_t needed
        self._cpp_obj = make_shared[cStruct]()
        cdef cStruct* cpp_obj = self._cpp_obj.get()
        if proto is __Protocol.COMPACT:
            with nogil:
                needed = serializer.CompactDeserialize[cStruct](buf, deref(cpp_obj), serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.BINARY:
            with nogil:
                needed = serializer.BinaryDeserialize[cStruct](buf, deref(cpp_obj), serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.JSON:
            with nogil:
                needed = serializer.JSONDeserialize[cStruct](buf, deref(cpp_obj), serializer.SHARE_EXTERNAL_BUFFER)
        return needed

    def __reduce__(self):
        return (deserialize, (Struct, serialize(self)))


cdef cBigStruct _BigStruct_defaults = cBigStruct()

cdef class BigStruct(thrift.py3.types.Struct):

    def __init__(
        BigStruct self, *,
        Struct s=None,
        id=None
    ):
        if id is not None:
            if not isinstance(id, int):
                raise TypeError(f'id is not a { int !r}.')
            id = <int32_t> id

        self._cpp_obj = move(BigStruct._make_instance(
          NULL,
          s,
          id,
        ))

    def __call__(
        BigStruct self,
        s=__NOTSET,
        id=__NOTSET
    ):
        changes = any((
            s is not __NOTSET,

            id is not __NOTSET,
        ))

        if not changes:
            return self

        if None is not s is not __NOTSET:
            if not isinstance(s, Struct):
                raise TypeError(f's is not a { Struct !r}.')

        if None is not id is not __NOTSET:
            if not isinstance(id, int):
                raise TypeError(f'id is not a { int !r}.')
            id = <int32_t> id

        inst = <BigStruct>BigStruct.__new__(BigStruct)
        inst._cpp_obj = move(BigStruct._make_instance(
          self._cpp_obj.get(),
          s,
          id,
        ))
        return inst

    @staticmethod
    cdef unique_ptr[cBigStruct] _make_instance(
        cBigStruct* base_instance,
        object s,
        object id
    ) except *:
        cdef unique_ptr[cBigStruct] c_inst
        if base_instance:
            c_inst = make_unique[cBigStruct](deref(base_instance))
        else:
            c_inst = make_unique[cBigStruct]()

        if base_instance:
            # Convert None's to default value. (or unset)
            if s is None:
                deref(c_inst).s = _BigStruct_defaults.s
                deref(c_inst).__isset.s = False
                pass
            elif s is __NOTSET:
                s = None

            if id is None:
                deref(c_inst).id = _BigStruct_defaults.id
                deref(c_inst).__isset.id = False
                pass
            elif id is __NOTSET:
                id = None

        if s is not None:
            deref(c_inst).s = deref((<Struct?> s)._cpp_obj)
            deref(c_inst).__isset.s = True
        if id is not None:
            deref(c_inst).id = id
            deref(c_inst).__isset.id = True
        # in C++ you don't have to call move(), but this doesn't translate
        # into a C++ return statement, so you do here
        return move_unique(c_inst)

    def __iter__(self):
        yield 's', self.s
        yield 'id', self.id

    def __bool__(self):
        return True or True

    @staticmethod
    cdef create(shared_ptr[cBigStruct] cpp_obj):
        inst = <BigStruct>BigStruct.__new__(BigStruct)
        inst._cpp_obj = cpp_obj
        return inst

    @property
    def s(self):

        if self.__s is None:
            self.__s = Struct.create(make_shared[cStruct](deref(self._cpp_obj).s))
        return self.__s

    @property
    def id(self):

        return self._cpp_obj.get().id


    def __hash__(BigStruct self):
        if not self.__hash:
            self.__hash = hash((
            self.s,
            self.id,
            ))
        return self.__hash

    def __repr__(BigStruct self):
        return f'BigStruct(s={repr(self.s)}, id={repr(self.id)})'
    def __richcmp__(self, other, op):
        cdef int cop = op
        if cop not in (2, 3):
            raise TypeError("unorderable types: {}, {}".format(self, other))
        if not (
                isinstance(self, BigStruct) and
                isinstance(other, BigStruct)):
            if cop == 2:  # different types are never equal
                return False
            else:         # different types are always notequal
                return True

        cdef cBigStruct cself = deref((<BigStruct>self)._cpp_obj)
        cdef cBigStruct cother = deref((<BigStruct>other)._cpp_obj)
        cdef cbool cmp = cself == cother
        if cop == 2:
            return cmp
        return not cmp

    cdef __iobuf.IOBuf _serialize(BigStruct self, proto):
        cdef __iobuf.cIOBufQueue queue = __iobuf.cIOBufQueue(__iobuf.cacheChainLength())
        cdef cBigStruct* cpp_obj = self._cpp_obj.get()
        if proto is __Protocol.COMPACT:
            with nogil:
                serializer.CompactSerialize[cBigStruct](deref(cpp_obj), &queue, serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.BINARY:
            with nogil:
                serializer.BinarySerialize[cBigStruct](deref(cpp_obj), &queue, serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.JSON:
            with nogil:
                serializer.JSONSerialize[cBigStruct](deref(cpp_obj), &queue, serializer.SHARE_EXTERNAL_BUFFER)
        return __iobuf.from_unique_ptr(queue.move())

    cdef uint32_t _deserialize(BigStruct self, const __iobuf.cIOBuf* buf, proto) except? 0:
        cdef uint32_t needed
        self._cpp_obj = make_shared[cBigStruct]()
        cdef cBigStruct* cpp_obj = self._cpp_obj.get()
        if proto is __Protocol.COMPACT:
            with nogil:
                needed = serializer.CompactDeserialize[cBigStruct](buf, deref(cpp_obj), serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.BINARY:
            with nogil:
                needed = serializer.BinaryDeserialize[cBigStruct](buf, deref(cpp_obj), serializer.SHARE_EXTERNAL_BUFFER)
        elif proto is __Protocol.JSON:
            with nogil:
                needed = serializer.JSONDeserialize[cBigStruct](buf, deref(cpp_obj), serializer.SHARE_EXTERNAL_BUFFER)
        return needed

    def __reduce__(self):
        return (deserialize, (BigStruct, serialize(self)))


c2 = Struct.create(make_shared[cStruct](cc2()))
c3 = Struct.create(make_shared[cStruct](cc3()))
c4 = Struct.create(make_shared[cStruct](cc4()))
