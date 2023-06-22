#
# Autogenerated by Thrift
#
# DO NOT EDIT
#  @generated
#

from __future__ import annotations


import typing as _typing

import enum

import folly.iobuf as _fbthrift_iobuf
import thrift.python.types as _fbthrift_python_types
import thrift.python.exceptions as _fbthrift_python_exceptions


class MyEnum(_fbthrift_python_types.Enum, int):
    MyValue1: MyEnum = ...
    MyValue2: MyEnum = ...
    def _to_python(self) -> MyEnum: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.MyEnum": ...  # type: ignore
    def _to_py_deprecated(self) -> int: ...


class MyStruct(_fbthrift_python_types.Struct):
    MyIntField: _typing.Final[int] = ...
    MyStringField: _typing.Final[str] = ...
    MyDataField: _typing.Final[MyDataItem] = ...
    myEnum: _typing.Final[MyEnum] = ...
    oneway: _typing.Final[bool] = ...
    floatList: _typing.Final[_typing.Sequence[float]] = ...
    strMap: _typing.Final[_typing.Mapping[bytes, str]] = ...
    floatSet: _typing.Final[_typing.AbstractSet[int]] = ...
    def __init__(
        self, *,
        MyIntField: _typing.Optional[int]=...,
        MyStringField: _typing.Optional[str]=...,
        MyDataField: _typing.Optional[MyDataItem]=...,
        myEnum: _typing.Optional[MyEnum]=...,
        oneway: _typing.Optional[bool]=...,
        floatList: _typing.Optional[_typing.Sequence[float]]=...,
        strMap: _typing.Optional[_typing.Mapping[bytes, str]]=...,
        floatSet: _typing.Optional[_typing.AbstractSet[int]]=...
    ) -> None: ...

    def __call__(
        self, *,
        MyIntField: _typing.Optional[int]=...,
        MyStringField: _typing.Optional[str]=...,
        MyDataField: _typing.Optional[MyDataItem]=...,
        myEnum: _typing.Optional[MyEnum]=...,
        oneway: _typing.Optional[bool]=...,
        floatList: _typing.Optional[_typing.Sequence[float]]=...,
        strMap: _typing.Optional[_typing.Mapping[bytes, str]]=...,
        floatSet: _typing.Optional[_typing.AbstractSet[int]]=...
    ) -> MyStruct: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[int, str, MyDataItem, MyEnum, bool, _typing.Sequence[float], _typing.Mapping[bytes, str], _typing.AbstractSet[int]]]]: ...
    def _to_python(self) -> MyStruct: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.MyStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.MyStruct": ...  # type: ignore


class MyDataItem(_fbthrift_python_types.Struct):
    def __init__(
        self,
    ) -> None: ...

    def __call__(
        self,
    ) -> MyDataItem: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[None]]]: ...
    def _to_python(self) -> MyDataItem: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.MyDataItem": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.MyDataItem": ...  # type: ignore


class TransitiveDoubler(_fbthrift_python_types.Struct):
    def __init__(
        self,
    ) -> None: ...

    def __call__(
        self,
    ) -> TransitiveDoubler: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[None]]]: ...
    def _to_python(self) -> TransitiveDoubler: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.TransitiveDoubler": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.TransitiveDoubler": ...  # type: ignore


class DoubledPair(_fbthrift_python_types.Struct):
    s: _typing.Final[str] = ...
    x: _typing.Final[int] = ...
    def __init__(
        self, *,
        s: _typing.Optional[str]=...,
        x: _typing.Optional[int]=...
    ) -> None: ...

    def __call__(
        self, *,
        s: _typing.Optional[str]=...,
        x: _typing.Optional[int]=...
    ) -> DoubledPair: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str, int]]]: ...
    def _to_python(self) -> DoubledPair: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.DoubledPair": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.DoubledPair": ...  # type: ignore


class StringPair(_fbthrift_python_types.Struct):
    normal: _typing.Final[str] = ...
    doubled: _typing.Final[str] = ...
    def __init__(
        self, *,
        normal: _typing.Optional[str]=...,
        doubled: _typing.Optional[str]=...
    ) -> None: ...

    def __call__(
        self, *,
        normal: _typing.Optional[str]=...,
        doubled: _typing.Optional[str]=...
    ) -> StringPair: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[str, str]]]: ...
    def _to_python(self) -> StringPair: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.StringPair": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.StringPair": ...  # type: ignore


class EmptyStruct(_fbthrift_python_types.Struct):
    def __init__(
        self,
    ) -> None: ...

    def __call__(
        self,
    ) -> EmptyStruct: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[None]]]: ...
    def _to_python(self) -> EmptyStruct: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.EmptyStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.EmptyStruct": ...  # type: ignore


class PrimitiveStruct(_fbthrift_python_types.Struct):
    booly: _typing.Final[bool] = ...
    charry: _typing.Final[int] = ...
    shorty: _typing.Final[int] = ...
    inty: _typing.Final[int] = ...
    longy: _typing.Final[int] = ...
    floaty: _typing.Final[_typing.Optional[float]] = ...
    dubby: _typing.Final[_typing.Optional[float]] = ...
    stringy: _typing.Final[_typing.Optional[str]] = ...
    bytey: _typing.Final[_typing.Optional[bytes]] = ...
    def __init__(
        self, *,
        booly: _typing.Optional[bool]=...,
        charry: _typing.Optional[int]=...,
        shorty: _typing.Optional[int]=...,
        inty: _typing.Optional[int]=...,
        longy: _typing.Optional[int]=...,
        floaty: _typing.Optional[float]=...,
        dubby: _typing.Optional[float]=...,
        stringy: _typing.Optional[str]=...,
        bytey: _typing.Optional[bytes]=...
    ) -> None: ...

    def __call__(
        self, *,
        booly: _typing.Optional[bool]=...,
        charry: _typing.Optional[int]=...,
        shorty: _typing.Optional[int]=...,
        inty: _typing.Optional[int]=...,
        longy: _typing.Optional[int]=...,
        floaty: _typing.Optional[float]=...,
        dubby: _typing.Optional[float]=...,
        stringy: _typing.Optional[str]=...,
        bytey: _typing.Optional[bytes]=...
    ) -> PrimitiveStruct: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[bool, int, int, int, int, float, float, str, bytes]]]: ...
    def _to_python(self) -> PrimitiveStruct: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.PrimitiveStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.PrimitiveStruct": ...  # type: ignore


class ListStruct(_fbthrift_python_types.Struct):
    boolz: _typing.Final[_typing.Sequence[bool]] = ...
    intz: _typing.Final[_typing.Optional[_typing.Sequence[int]]] = ...
    stringz: _typing.Final[_typing.Optional[_typing.Sequence[str]]] = ...
    encoded: _typing.Final[_typing.Sequence[bytes]] = ...
    uidz: _typing.Final[_typing.Sequence[int]] = ...
    matrix: _typing.Final[_typing.Sequence[_typing.Sequence[float]]] = ...
    ucharz: _typing.Final[_typing.Sequence[_typing.Sequence[int]]] = ...
    voxels: _typing.Final[_typing.Sequence[_typing.Sequence[_typing.Sequence[int]]]] = ...
    def __init__(
        self, *,
        boolz: _typing.Optional[_typing.Sequence[bool]]=...,
        intz: _typing.Optional[_typing.Sequence[int]]=...,
        stringz: _typing.Optional[_typing.Sequence[str]]=...,
        encoded: _typing.Optional[_typing.Sequence[bytes]]=...,
        uidz: _typing.Optional[_typing.Sequence[int]]=...,
        matrix: _typing.Optional[_typing.Sequence[_typing.Sequence[float]]]=...,
        ucharz: _typing.Optional[_typing.Sequence[_typing.Sequence[int]]]=...,
        voxels: _typing.Optional[_typing.Sequence[_typing.Sequence[_typing.Sequence[int]]]]=...
    ) -> None: ...

    def __call__(
        self, *,
        boolz: _typing.Optional[_typing.Sequence[bool]]=...,
        intz: _typing.Optional[_typing.Sequence[int]]=...,
        stringz: _typing.Optional[_typing.Sequence[str]]=...,
        encoded: _typing.Optional[_typing.Sequence[bytes]]=...,
        uidz: _typing.Optional[_typing.Sequence[int]]=...,
        matrix: _typing.Optional[_typing.Sequence[_typing.Sequence[float]]]=...,
        ucharz: _typing.Optional[_typing.Sequence[_typing.Sequence[int]]]=...,
        voxels: _typing.Optional[_typing.Sequence[_typing.Sequence[_typing.Sequence[int]]]]=...
    ) -> ListStruct: ...
    def __iter__(self) -> _typing.Iterator[_typing.Tuple[str, _typing.Union[_typing.Sequence[bool], _typing.Sequence[int], _typing.Sequence[str], _typing.Sequence[bytes], _typing.Sequence[int], _typing.Sequence[_typing.Sequence[float]], _typing.Sequence[_typing.Sequence[int]], _typing.Sequence[_typing.Sequence[_typing.Sequence[int]]]]]]: ...
    def _to_python(self) -> ListStruct: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.ListStruct": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.ListStruct": ...  # type: ignore


class MyUnion(_fbthrift_python_types.Union):
    myEnum: _typing.Final[MyEnum] = ...
    myStruct: _typing.Final[MyStruct] = ...
    myDataItem: _typing.Final[MyDataItem] = ...
    doubleSet: _typing.Final[_typing.AbstractSet[int]] = ...
    doubleList: _typing.Final[_typing.Sequence[float]] = ...
    strMap: _typing.Final[_typing.Mapping[bytes, str]] = ...
    def __init__(
        self, *,
        myEnum: _typing.Optional[MyEnum]=...,
        myStruct: _typing.Optional[MyStruct]=...,
        myDataItem: _typing.Optional[MyDataItem]=...,
        doubleSet: _typing.Optional[_typing.AbstractSet[int]]=...,
        doubleList: _typing.Optional[_typing.Sequence[float]]=...,
        strMap: _typing.Optional[_typing.Mapping[bytes, str]]=...
    ) -> None: ...


    class Type(enum.Enum):
        EMPTY: MyUnion.Type = ...
        myEnum: MyUnion.Type = ...
        myStruct: MyUnion.Type = ...
        myDataItem: MyUnion.Type = ...
        doubleSet: MyUnion.Type = ...
        doubleList: MyUnion.Type = ...
        strMap: MyUnion.Type = ...

    @classmethod
    def fromValue(cls, value: _typing.Union[None, MyEnum, MyStruct, MyDataItem, _typing.AbstractSet[int], _typing.Sequence[float], _typing.Mapping[bytes, str]]) -> MyUnion: ...
    value: _typing.Final[_typing.Union[None, MyEnum, MyStruct, MyDataItem, _typing.AbstractSet[int], _typing.Sequence[float], _typing.Mapping[bytes, str]]]
    type: Type
    def get_type(self) -> Type:...
    def _to_python(self) -> MyUnion: ...
    def _to_py3(self) -> "test.fixtures.basic-python-capi.module.types.MyUnion": ...  # type: ignore
    def _to_py_deprecated(self) -> "module.ttypes.MyUnion": ...  # type: ignore

signed_byte = int