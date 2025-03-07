/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/mcpp2-compare/src/enums.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include <thrift/lib/cpp2/gen/module_types_h.h>



namespace apache {
namespace thrift {
namespace ident {
struct fieldA;
} // namespace ident
namespace detail {
#ifndef APACHE_THRIFT_ACCESSOR_fieldA
#define APACHE_THRIFT_ACCESSOR_fieldA
APACHE_THRIFT_DEFINE_ACCESSOR(fieldA);
#endif
} // namespace detail
} // namespace thrift
} // namespace apache

// BEGIN declare_enums
namespace facebook { namespace ns { namespace qwerty {

enum class AnEnumA {
  FIELDA = 0,
};



enum class AnEnumB {
  FIELDA = 0,
  FIELDB = 2,
};



enum class AnEnumC {
  FIELDC = 0,
};



enum class AnEnumD {
  FIELDD = 0,
};



enum class AnEnumE {
  FIELDA = 0,
};



}}} // facebook::ns::qwerty

namespace std {
template<> struct hash<::facebook::ns::qwerty::AnEnumA> :
  ::apache::thrift::detail::enum_hash<::facebook::ns::qwerty::AnEnumA> {};
template<> struct hash<::facebook::ns::qwerty::AnEnumB> :
  ::apache::thrift::detail::enum_hash<::facebook::ns::qwerty::AnEnumB> {};
template<> struct hash<::facebook::ns::qwerty::AnEnumC> :
  ::apache::thrift::detail::enum_hash<::facebook::ns::qwerty::AnEnumC> {};
template<> struct hash<::facebook::ns::qwerty::AnEnumD> :
  ::apache::thrift::detail::enum_hash<::facebook::ns::qwerty::AnEnumD> {};
template<> struct hash<::facebook::ns::qwerty::AnEnumE> :
  ::apache::thrift::detail::enum_hash<::facebook::ns::qwerty::AnEnumE> {};
} // std

namespace apache { namespace thrift {


template <> struct TEnumDataStorage<::facebook::ns::qwerty::AnEnumA>;

template <> struct TEnumTraits<::facebook::ns::qwerty::AnEnumA> {
  using type = ::facebook::ns::qwerty::AnEnumA;

  static constexpr std::size_t const size = 1;
  static folly::Range<type const*> const values;
  static folly::Range<folly::StringPiece const*> const names;

  static bool findName(type value, folly::StringPiece* out) noexcept;
  static bool findValue(folly::StringPiece name, type* out) noexcept;

#if FOLLY_HAS_STRING_VIEW
  static bool findName(type value, std::string_view* out) noexcept {
    folly::StringPiece outp;
    return findName(value, &outp) && ((*out = outp), true);
  }
#endif
  static char const* findName(type value) noexcept {
    folly::StringPiece ret;
    (void)findName(value, &ret);
    return ret.data();
  }
  static constexpr type min() { return type::FIELDA; }
  static constexpr type max() { return type::FIELDA; }
};


template <> struct TEnumDataStorage<::facebook::ns::qwerty::AnEnumB>;

template <> struct TEnumTraits<::facebook::ns::qwerty::AnEnumB> {
  using type = ::facebook::ns::qwerty::AnEnumB;

  static constexpr std::size_t const size = 2;
  static folly::Range<type const*> const values;
  static folly::Range<folly::StringPiece const*> const names;

  static bool findName(type value, folly::StringPiece* out) noexcept;
  static bool findValue(folly::StringPiece name, type* out) noexcept;

#if FOLLY_HAS_STRING_VIEW
  static bool findName(type value, std::string_view* out) noexcept {
    folly::StringPiece outp;
    return findName(value, &outp) && ((*out = outp), true);
  }
#endif
  static char const* findName(type value) noexcept {
    folly::StringPiece ret;
    (void)findName(value, &ret);
    return ret.data();
  }
  static constexpr type min() { return type::FIELDA; }
  static constexpr type max() { return type::FIELDB; }
};


template <> struct TEnumDataStorage<::facebook::ns::qwerty::AnEnumC>;

template <> struct TEnumTraits<::facebook::ns::qwerty::AnEnumC> {
  using type = ::facebook::ns::qwerty::AnEnumC;

  static constexpr std::size_t const size = 1;
  static folly::Range<type const*> const values;
  static folly::Range<folly::StringPiece const*> const names;

  static bool findName(type value, folly::StringPiece* out) noexcept;
  static bool findValue(folly::StringPiece name, type* out) noexcept;

#if FOLLY_HAS_STRING_VIEW
  static bool findName(type value, std::string_view* out) noexcept {
    folly::StringPiece outp;
    return findName(value, &outp) && ((*out = outp), true);
  }
#endif
  static char const* findName(type value) noexcept {
    folly::StringPiece ret;
    (void)findName(value, &ret);
    return ret.data();
  }
  static constexpr type min() { return type::FIELDC; }
  static constexpr type max() { return type::FIELDC; }
};


template <> struct TEnumDataStorage<::facebook::ns::qwerty::AnEnumD>;

template <> struct TEnumTraits<::facebook::ns::qwerty::AnEnumD> {
  using type = ::facebook::ns::qwerty::AnEnumD;

  static constexpr std::size_t const size = 1;
  static folly::Range<type const*> const values;
  static folly::Range<folly::StringPiece const*> const names;

  static bool findName(type value, folly::StringPiece* out) noexcept;
  static bool findValue(folly::StringPiece name, type* out) noexcept;

#if FOLLY_HAS_STRING_VIEW
  static bool findName(type value, std::string_view* out) noexcept {
    folly::StringPiece outp;
    return findName(value, &outp) && ((*out = outp), true);
  }
#endif
  static char const* findName(type value) noexcept {
    folly::StringPiece ret;
    (void)findName(value, &ret);
    return ret.data();
  }
  static constexpr type min() { return type::FIELDD; }
  static constexpr type max() { return type::FIELDD; }
};


template <> struct TEnumDataStorage<::facebook::ns::qwerty::AnEnumE>;

template <> struct TEnumTraits<::facebook::ns::qwerty::AnEnumE> {
  using type = ::facebook::ns::qwerty::AnEnumE;

  static constexpr std::size_t const size = 1;
  static folly::Range<type const*> const values;
  static folly::Range<folly::StringPiece const*> const names;

  static bool findName(type value, folly::StringPiece* out) noexcept;
  static bool findValue(folly::StringPiece name, type* out) noexcept;

#if FOLLY_HAS_STRING_VIEW
  static bool findName(type value, std::string_view* out) noexcept {
    folly::StringPiece outp;
    return findName(value, &outp) && ((*out = outp), true);
  }
#endif
  static char const* findName(type value) noexcept {
    folly::StringPiece ret;
    (void)findName(value, &ret);
    return ret.data();
  }
  static constexpr type min() { return type::FIELDA; }
  static constexpr type max() { return type::FIELDA; }
};


}} // apache::thrift


// END declare_enums
// BEGIN forward_declare
namespace facebook { namespace ns { namespace qwerty {
class SomeStruct;
}}} // facebook::ns::qwerty
// END forward_declare
namespace apache::thrift::detail::annotation {
} // namespace apache::thrift::detail::annotation

namespace apache::thrift::detail::qualifier {
} // namespace apache::thrift::detail::qualifier

// BEGIN hash_and_equal_to
// END hash_and_equal_to
namespace facebook { namespace ns { namespace qwerty {
using ::apache::thrift::detail::operator!=;
using ::apache::thrift::detail::operator>;
using ::apache::thrift::detail::operator<=;
using ::apache::thrift::detail::operator>=;


class SomeStruct final  {
 private:
  friend struct ::apache::thrift::detail::st::struct_private_access;
  template<class> friend struct ::apache::thrift::detail::invoke_reffer;

  //  used by a static_assert in the corresponding source
  static constexpr bool __fbthrift_cpp2_gen_json = true;
  static constexpr bool __fbthrift_cpp2_is_runtime_annotation = false;
  static const folly::StringPiece __fbthrift_get_field_name(::apache::thrift::FieldOrdinal ord);
  static const folly::StringPiece __fbthrift_get_class_name();
  using __fbthrift_reflection_ident_list = folly::tag_t<
    ::apache::thrift::ident::fieldA
  >;

  static constexpr std::int16_t __fbthrift_reflection_field_id_list[] = {0,1};
  using __fbthrift_reflection_type_tags = folly::tag_t<
    ::apache::thrift::type::i32_t
  >;

  static constexpr std::size_t __fbthrift_field_size_v = 1;

  template<class T>
  using __fbthrift_id = ::apache::thrift::type::field_id<__fbthrift_reflection_field_id_list[folly::to_underlying(T::value)]>;

  template<class T>
  using __fbthrift_type_tag = ::apache::thrift::detail::at<__fbthrift_reflection_type_tags, T::value>;

  template<class T>
  using __fbthrift_ident = ::apache::thrift::detail::at<__fbthrift_reflection_ident_list, T::value>;

  template<class T> using __fbthrift_ordinal = ::apache::thrift::type::ordinal_tag<
    ::apache::thrift::detail::getFieldOrdinal<T,
                                              __fbthrift_reflection_ident_list,
                                              __fbthrift_reflection_type_tags>(
      __fbthrift_reflection_field_id_list
    )
  >;
  void __fbthrift_clear();
  void __fbthrift_clear_terse_fields();
  bool __fbthrift_is_empty() const;

 public:
  using __fbthrift_cpp2_type = SomeStruct;
  static constexpr bool __fbthrift_cpp2_is_union =
    false;


 public:

  SomeStruct() :
      __fbthrift_field_fieldA() {
  }
  // FragileConstructor for use in initialization lists only.
  [[deprecated("This constructor is deprecated")]]
  SomeStruct(apache::thrift::FragileConstructor, ::std::int32_t fieldA__arg);

  SomeStruct(SomeStruct&&) = default;

  SomeStruct(const SomeStruct&) = default;


  SomeStruct& operator=(SomeStruct&&) = default;

  SomeStruct& operator=(const SomeStruct&) = default;
 private:
  ::std::int32_t __fbthrift_field_fieldA;
 private:
  apache::thrift::detail::isset_bitset<1, apache::thrift::detail::IssetBitsetOption::Unpacked> __isset;

 public:

  bool operator==(const SomeStruct&) const;
  bool operator<(const SomeStruct&) const;

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&> fieldA_ref() const& {
    return {this->__fbthrift_field_fieldA, __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&&> fieldA_ref() const&& {
    return {static_cast<const T&&>(this->__fbthrift_field_fieldA), __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&> fieldA_ref() & {
    return {this->__fbthrift_field_fieldA, __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&&> fieldA_ref() && {
    return {static_cast<T&&>(this->__fbthrift_field_fieldA), __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&> fieldA() const& {
    return {this->__fbthrift_field_fieldA, __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<const T&&> fieldA() const&& {
    return {static_cast<const T&&>(this->__fbthrift_field_fieldA), __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&> fieldA() & {
    return {this->__fbthrift_field_fieldA, __isset.at(0), __isset.bit(0)};
  }

  template <typename..., typename T = ::std::int32_t>
  FOLLY_ERASE ::apache::thrift::field_ref<T&&> fieldA() && {
    return {static_cast<T&&>(this->__fbthrift_field_fieldA), __isset.at(0), __isset.bit(0)};
  }

  ::std::int32_t get_fieldA() const {
    return __fbthrift_field_fieldA;
  }

  [[deprecated("Use `FOO.fieldA_ref() = BAR;` instead of `FOO.set_fieldA(BAR);`")]]
  ::std::int32_t& set_fieldA(::std::int32_t fieldA_) {
    fieldA_ref() = fieldA_;
    return __fbthrift_field_fieldA;
  }

  template <class Protocol_>
  unsigned long read(Protocol_* iprot);
  template <class Protocol_>
  uint32_t serializedSize(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t serializedSizeZC(Protocol_ const* prot_) const;
  template <class Protocol_>
  uint32_t write(Protocol_* prot_) const;

 private:
  template <class Protocol_>
  void readNoXfer(Protocol_* iprot);

  friend class ::apache::thrift::Cpp2Ops<SomeStruct>;
  friend void swap(SomeStruct& a, SomeStruct& b);
};

template <class Protocol_>
unsigned long SomeStruct::read(Protocol_* iprot) {
  auto _xferStart = iprot->getCursorPosition();
  readNoXfer(iprot);
  return iprot->getCursorPosition() - _xferStart;
}


}}} // facebook::ns::qwerty
