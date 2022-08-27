/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

#pragma once

#include <functional>
#include <folly/Range.h>

#include <thrift/lib/py3/enums.h>
#include "thrift/lib/thrift/gen-cpp2/standard_data.h"
#include "thrift/lib/thrift/gen-cpp2/standard_types.h"
#include "thrift/lib/thrift/gen-cpp2/standard_metadata.h"
namespace thrift {
namespace py3 {


template<>
const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::apache::thrift::type::StandardProtocol>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}


template<>
const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::apache::thrift::type::Void>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}


template<>
const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::apache::thrift::type::TypeUri::Type>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}

template<>
const std::vector<std::pair<std::string_view, std::string_view>>& PyEnumTraits<
    ::apache::thrift::type::TypeName::Type>::namesmap() {
  static const folly::Indestructible<NamesMap> pairs {
    {
    }
  };
  return *pairs;
}


template<>
void reset_field<::apache::thrift::type::UriStruct>(
    ::apache::thrift::type::UriStruct& obj, uint16_t index) {
  switch (index) {
    case 0:
      obj.scheme_ref().copy_from(default_inst<::apache::thrift::type::UriStruct>().scheme_ref());
      return;
    case 1:
      obj.domain_ref().copy_from(default_inst<::apache::thrift::type::UriStruct>().domain_ref());
      return;
    case 2:
      obj.path_ref().copy_from(default_inst<::apache::thrift::type::UriStruct>().path_ref());
      return;
    case 3:
      obj.query_ref().copy_from(default_inst<::apache::thrift::type::UriStruct>().query_ref());
      return;
    case 4:
      obj.fragment_ref().copy_from(default_inst<::apache::thrift::type::UriStruct>().fragment_ref());
      return;
  }
}

template<>
const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::apache::thrift::type::UriStruct>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::apache::thrift::type::TypeUri>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}

template<>
const std::unordered_map<std::string_view, std::string_view>& PyStructTraits<
    ::apache::thrift::type::TypeName>::namesmap() {
  static const folly::Indestructible<NamesMap> map {
    {
    }
  };
  return *map;
}
} // namespace py3
} // namespace thrift