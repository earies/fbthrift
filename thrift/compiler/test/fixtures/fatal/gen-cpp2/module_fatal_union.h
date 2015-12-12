/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#pragma once

#include "module_fatal.h"

#include <fatal/type/enum.h>
#include <fatal/type/variant_traits.h>

#include <type_traits>

namespace test_cpp2 { namespace cpp_reflection {

namespace detail {

struct union1_Type_enum_traits {
  using type = ::test_cpp2::cpp_reflection::union1::Type;
  using name = detail::test_cpp2_cpp_reflection_module__unique_strings_list::Type;

  struct str {
    using ui = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ui;
    using ud = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ud;
    using us = detail::test_cpp2_cpp_reflection_module__unique_strings_list::us;
    using ue = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ue;
  };

  using name_to_value = ::fatal::type_map<
    ::fatal::type_pair<
      str::ui,
      std::integral_constant<type, type::ui>
    >,
    ::fatal::type_pair<
      str::ud,
      std::integral_constant<type, type::ud>
    >,
    ::fatal::type_pair<
      str::us,
      std::integral_constant<type, type::us>
    >,
    ::fatal::type_pair<
      str::ue,
      std::integral_constant<type, type::ue>
    >
  >;

  static char const *to_string(type e, char const *fallback) {
    switch (e) {
      case type::ui: return "ui";
      case type::ud: return "ud";
      case type::us: return "us";
      case type::ue: return "ue";
      default: return fallback;
    }
  }
};

} // detail

FATAL_REGISTER_ENUM_TRAITS(
  ::test_cpp2::cpp_reflection::detail::union1_Type_enum_traits,
  ::fatal::type_pair<
    module_tags::module,
    ::fatal::type_map<
    >
  >
);

namespace detail {

class union1_variant_traits {
  struct get {
    struct ui {
      int32_t const &operator ()(union1 const &variant) const {
        return variant.get_ui();
      }

      int32_t &operator ()(union1 &variant) const {
        return variant.mutable_ui();
      }

      int32_t operator ()(union1 &&variant) const {
        return std::move(variant).move_ui();
      }
    };

    struct ud {
      double const &operator ()(union1 const &variant) const {
        return variant.get_ud();
      }

      double &operator ()(union1 &variant) const {
        return variant.mutable_ud();
      }

      double operator ()(union1 &&variant) const {
        return std::move(variant).move_ud();
      }
    };

    struct us {
      std::string const &operator ()(union1 const &variant) const {
        return variant.get_us();
      }

      std::string &operator ()(union1 &variant) const {
        return variant.mutable_us();
      }

      std::string operator ()(union1 &&variant) const {
        return std::move(variant).move_us();
      }
    };

    struct ue {
       ::test_cpp2::cpp_reflection::enum1 const &operator ()(union1 const &variant) const {
        return variant.get_ue();
      }

       ::test_cpp2::cpp_reflection::enum1 &operator ()(union1 &variant) const {
        return variant.mutable_ue();
      }

       ::test_cpp2::cpp_reflection::enum1 operator ()(union1 &&variant) const {
        return std::move(variant).move_ue();
      }
    };
  };

  struct set {
    struct ui {
      template <typename... Args>
      void operator ()(union1 &variant, Args &&...args) const {
        return variant.set_ui(std::forward<Args>(args)...);
      }
    };

    struct ud {
      template <typename... Args>
      void operator ()(union1 &variant, Args &&...args) const {
        return variant.set_ud(std::forward<Args>(args)...);
      }
    };

    struct us {
      template <typename... Args>
      void operator ()(union1 &variant, Args &&...args) const {
        return variant.set_us(std::forward<Args>(args)...);
      }
    };

    struct ue {
      template <typename... Args>
      void operator ()(union1 &variant, Args &&...args) const {
        return variant.set_ue(std::forward<Args>(args)...);
      }
    };
  };

  public:
  using type = ::test_cpp2::cpp_reflection::union1;
  using name = detail::test_cpp2_cpp_reflection_module__unique_strings_list::union1;
  using id = type::Type;

  struct ids {
    using ui = std::integral_constant<id, id::ui>;
    using ud = std::integral_constant<id, id::ud>;
    using us = std::integral_constant<id, id::us>;
    using ue = std::integral_constant<id, id::ue>;
  };

  using descriptors = ::fatal::type_list<
    ::fatal::variant_type_descriptor<
      int32_t,
      ids::ui,
      get::ui,
      set::ui
    >,
    ::fatal::variant_type_descriptor<
      double,
      ids::ud,
      get::ud,
      set::ud
    >,
    ::fatal::variant_type_descriptor<
      std::string,
      ids::us,
      get::us,
      set::us
    >,
    ::fatal::variant_type_descriptor<
       ::test_cpp2::cpp_reflection::enum1,
      ids::ue,
      get::ue,
      set::ue
    >
  >;

  static id get_id(type const &variant) {
    return variant.getType();
  }

  static bool empty(type const &variant) {
    return variant.getType() == id::__EMPTY__;
  }
};

} // detail
namespace detail {

struct union2_Type_enum_traits {
  using type = ::test_cpp2::cpp_reflection::union2::Type;
  using name = detail::test_cpp2_cpp_reflection_module__unique_strings_list::Type;

  struct str {
    using ui_2 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ui_2;
    using ud_2 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ud_2;
    using us_2 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::us_2;
    using ue_2 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ue_2;
  };

  using name_to_value = ::fatal::type_map<
    ::fatal::type_pair<
      str::ui_2,
      std::integral_constant<type, type::ui_2>
    >,
    ::fatal::type_pair<
      str::ud_2,
      std::integral_constant<type, type::ud_2>
    >,
    ::fatal::type_pair<
      str::us_2,
      std::integral_constant<type, type::us_2>
    >,
    ::fatal::type_pair<
      str::ue_2,
      std::integral_constant<type, type::ue_2>
    >
  >;

  static char const *to_string(type e, char const *fallback) {
    switch (e) {
      case type::ui_2: return "ui_2";
      case type::ud_2: return "ud_2";
      case type::us_2: return "us_2";
      case type::ue_2: return "ue_2";
      default: return fallback;
    }
  }
};

} // detail

FATAL_REGISTER_ENUM_TRAITS(
  ::test_cpp2::cpp_reflection::detail::union2_Type_enum_traits,
  ::fatal::type_pair<
    module_tags::module,
    ::fatal::type_map<
    >
  >
);

namespace detail {

class union2_variant_traits {
  struct get {
    struct ui_2 {
      int32_t const &operator ()(union2 const &variant) const {
        return variant.get_ui_2();
      }

      int32_t &operator ()(union2 &variant) const {
        return variant.mutable_ui_2();
      }

      int32_t operator ()(union2 &&variant) const {
        return std::move(variant).move_ui_2();
      }
    };

    struct ud_2 {
      double const &operator ()(union2 const &variant) const {
        return variant.get_ud_2();
      }

      double &operator ()(union2 &variant) const {
        return variant.mutable_ud_2();
      }

      double operator ()(union2 &&variant) const {
        return std::move(variant).move_ud_2();
      }
    };

    struct us_2 {
      std::string const &operator ()(union2 const &variant) const {
        return variant.get_us_2();
      }

      std::string &operator ()(union2 &variant) const {
        return variant.mutable_us_2();
      }

      std::string operator ()(union2 &&variant) const {
        return std::move(variant).move_us_2();
      }
    };

    struct ue_2 {
       ::test_cpp2::cpp_reflection::enum1 const &operator ()(union2 const &variant) const {
        return variant.get_ue_2();
      }

       ::test_cpp2::cpp_reflection::enum1 &operator ()(union2 &variant) const {
        return variant.mutable_ue_2();
      }

       ::test_cpp2::cpp_reflection::enum1 operator ()(union2 &&variant) const {
        return std::move(variant).move_ue_2();
      }
    };
  };

  struct set {
    struct ui_2 {
      template <typename... Args>
      void operator ()(union2 &variant, Args &&...args) const {
        return variant.set_ui_2(std::forward<Args>(args)...);
      }
    };

    struct ud_2 {
      template <typename... Args>
      void operator ()(union2 &variant, Args &&...args) const {
        return variant.set_ud_2(std::forward<Args>(args)...);
      }
    };

    struct us_2 {
      template <typename... Args>
      void operator ()(union2 &variant, Args &&...args) const {
        return variant.set_us_2(std::forward<Args>(args)...);
      }
    };

    struct ue_2 {
      template <typename... Args>
      void operator ()(union2 &variant, Args &&...args) const {
        return variant.set_ue_2(std::forward<Args>(args)...);
      }
    };
  };

  public:
  using type = ::test_cpp2::cpp_reflection::union2;
  using name = detail::test_cpp2_cpp_reflection_module__unique_strings_list::union2;
  using id = type::Type;

  struct ids {
    using ui_2 = std::integral_constant<id, id::ui_2>;
    using ud_2 = std::integral_constant<id, id::ud_2>;
    using us_2 = std::integral_constant<id, id::us_2>;
    using ue_2 = std::integral_constant<id, id::ue_2>;
  };

  using descriptors = ::fatal::type_list<
    ::fatal::variant_type_descriptor<
      int32_t,
      ids::ui_2,
      get::ui_2,
      set::ui_2
    >,
    ::fatal::variant_type_descriptor<
      double,
      ids::ud_2,
      get::ud_2,
      set::ud_2
    >,
    ::fatal::variant_type_descriptor<
      std::string,
      ids::us_2,
      get::us_2,
      set::us_2
    >,
    ::fatal::variant_type_descriptor<
       ::test_cpp2::cpp_reflection::enum1,
      ids::ue_2,
      get::ue_2,
      set::ue_2
    >
  >;

  static id get_id(type const &variant) {
    return variant.getType();
  }

  static bool empty(type const &variant) {
    return variant.getType() == id::__EMPTY__;
  }
};

} // detail
namespace detail {

struct union3_Type_enum_traits {
  using type = ::test_cpp2::cpp_reflection::union3::Type;
  using name = detail::test_cpp2_cpp_reflection_module__unique_strings_list::Type;

  struct str {
    using ui_3 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ui_3;
    using ud_3 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ud_3;
    using us_3 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::us_3;
    using ue_3 = detail::test_cpp2_cpp_reflection_module__unique_strings_list::ue_3;
  };

  using name_to_value = ::fatal::type_map<
    ::fatal::type_pair<
      str::ui_3,
      std::integral_constant<type, type::ui_3>
    >,
    ::fatal::type_pair<
      str::ud_3,
      std::integral_constant<type, type::ud_3>
    >,
    ::fatal::type_pair<
      str::us_3,
      std::integral_constant<type, type::us_3>
    >,
    ::fatal::type_pair<
      str::ue_3,
      std::integral_constant<type, type::ue_3>
    >
  >;

  static char const *to_string(type e, char const *fallback) {
    switch (e) {
      case type::ui_3: return "ui_3";
      case type::ud_3: return "ud_3";
      case type::us_3: return "us_3";
      case type::ue_3: return "ue_3";
      default: return fallback;
    }
  }
};

} // detail

FATAL_REGISTER_ENUM_TRAITS(
  ::test_cpp2::cpp_reflection::detail::union3_Type_enum_traits,
  ::fatal::type_pair<
    module_tags::module,
    ::fatal::type_map<
    >
  >
);

namespace detail {

class union3_variant_traits {
  struct get {
    struct ui_3 {
      int32_t const &operator ()(union3 const &variant) const {
        return variant.get_ui_3();
      }

      int32_t &operator ()(union3 &variant) const {
        return variant.mutable_ui_3();
      }

      int32_t operator ()(union3 &&variant) const {
        return std::move(variant).move_ui_3();
      }
    };

    struct ud_3 {
      double const &operator ()(union3 const &variant) const {
        return variant.get_ud_3();
      }

      double &operator ()(union3 &variant) const {
        return variant.mutable_ud_3();
      }

      double operator ()(union3 &&variant) const {
        return std::move(variant).move_ud_3();
      }
    };

    struct us_3 {
      std::string const &operator ()(union3 const &variant) const {
        return variant.get_us_3();
      }

      std::string &operator ()(union3 &variant) const {
        return variant.mutable_us_3();
      }

      std::string operator ()(union3 &&variant) const {
        return std::move(variant).move_us_3();
      }
    };

    struct ue_3 {
       ::test_cpp2::cpp_reflection::enum1 const &operator ()(union3 const &variant) const {
        return variant.get_ue_3();
      }

       ::test_cpp2::cpp_reflection::enum1 &operator ()(union3 &variant) const {
        return variant.mutable_ue_3();
      }

       ::test_cpp2::cpp_reflection::enum1 operator ()(union3 &&variant) const {
        return std::move(variant).move_ue_3();
      }
    };
  };

  struct set {
    struct ui_3 {
      template <typename... Args>
      void operator ()(union3 &variant, Args &&...args) const {
        return variant.set_ui_3(std::forward<Args>(args)...);
      }
    };

    struct ud_3 {
      template <typename... Args>
      void operator ()(union3 &variant, Args &&...args) const {
        return variant.set_ud_3(std::forward<Args>(args)...);
      }
    };

    struct us_3 {
      template <typename... Args>
      void operator ()(union3 &variant, Args &&...args) const {
        return variant.set_us_3(std::forward<Args>(args)...);
      }
    };

    struct ue_3 {
      template <typename... Args>
      void operator ()(union3 &variant, Args &&...args) const {
        return variant.set_ue_3(std::forward<Args>(args)...);
      }
    };
  };

  public:
  using type = ::test_cpp2::cpp_reflection::union3;
  using name = detail::test_cpp2_cpp_reflection_module__unique_strings_list::union3;
  using id = type::Type;

  struct ids {
    using ui_3 = std::integral_constant<id, id::ui_3>;
    using ud_3 = std::integral_constant<id, id::ud_3>;
    using us_3 = std::integral_constant<id, id::us_3>;
    using ue_3 = std::integral_constant<id, id::ue_3>;
  };

  using descriptors = ::fatal::type_list<
    ::fatal::variant_type_descriptor<
      int32_t,
      ids::ui_3,
      get::ui_3,
      set::ui_3
    >,
    ::fatal::variant_type_descriptor<
      double,
      ids::ud_3,
      get::ud_3,
      set::ud_3
    >,
    ::fatal::variant_type_descriptor<
      std::string,
      ids::us_3,
      get::us_3,
      set::us_3
    >,
    ::fatal::variant_type_descriptor<
       ::test_cpp2::cpp_reflection::enum1,
      ids::ue_3,
      get::ue_3,
      set::ue_3
    >
  >;

  static id get_id(type const &variant) {
    return variant.getType();
  }

  static bool empty(type const &variant) {
    return variant.getType() == id::__EMPTY__;
  }
};

} // detail

FATAL_REGISTER_VARIANT_TRAITS(
  ::test_cpp2::cpp_reflection::detail::union1_variant_traits,
  ::fatal::type_pair<
    module_tags::module,
    ::fatal::type_map<
    >
  >
);
FATAL_REGISTER_VARIANT_TRAITS(
  ::test_cpp2::cpp_reflection::detail::union2_variant_traits,
  ::fatal::type_pair<
    module_tags::module,
    ::fatal::type_map<
    >
  >
);
FATAL_REGISTER_VARIANT_TRAITS(
  ::test_cpp2::cpp_reflection::detail::union3_variant_traits,
  ::fatal::type_pair<
    module_tags::module,
    ::fatal::type_map<
    >
  >
);

}} // test_cpp2::cpp_reflection