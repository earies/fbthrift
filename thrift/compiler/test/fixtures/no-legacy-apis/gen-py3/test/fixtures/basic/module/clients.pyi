#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

import folly.iobuf as _fbthrift_iobuf
import thrift.py3.types
import thrift.py3.client
import thrift.py3.common
import typing as _typing
from types import TracebackType

import test.fixtures.basic.module.types as _test_fixtures_basic_module_types


_MyServiceT = _typing.TypeVar('_MyServiceT', bound='MyService')


class MyService(thrift.py3.client.Client):

    async def query(
        self,
        u: _test_fixtures_basic_module_types.MyUnion,
        rpc_options: _typing.Optional[thrift.py3.common.RpcOptions]=None
    ) -> _test_fixtures_basic_module_types.MyStruct: ...

