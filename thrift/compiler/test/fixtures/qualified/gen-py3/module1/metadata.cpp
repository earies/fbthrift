/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#include "gen-py3/module1/metadata.h"

#include <thrift/lib/py3/metadata.h>

namespace module1 {
::apache::thrift::metadata::ThriftMetadata module1_getThriftModuleMetadata() {
  ::apache::thrift::metadata::ThriftMetadata metadata;
  ::apache::thrift::metadata::ThriftServiceContext serviceContext;
  ::apache::thrift::detail::md::EnumMetadata<Enum>::gen(metadata);
  ::apache::thrift::detail::md::StructMetadata<Struct>::gen(metadata);
  return metadata;
}
} // namespace module1
