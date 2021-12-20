/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.util.Set;
import java.util.HashSet;
import java.util.Collections;
import java.util.BitSet;
import java.util.Arrays;
import com.facebook.thrift.*;
import com.facebook.thrift.annotations.*;
import com.facebook.thrift.async.*;
import com.facebook.thrift.meta_data.*;
import com.facebook.thrift.server.*;
import com.facebook.thrift.transport.*;
import com.facebook.thrift.protocol.*;

@SuppressWarnings({ "unused", "serial" })
public class MyStructTypeDef implements TBase, java.io.Serializable, Cloneable, Comparable<MyStructTypeDef> {
  private static final TStruct STRUCT_DESC = new TStruct("MyStructTypeDef");
  private static final TField MY_LONG_FIELD_FIELD_DESC = new TField("myLongField", TType.I64, (short)1);
  private static final TField MY_LONG_TYPE_DEF_FIELD_DESC = new TField("myLongTypeDef", TType.I64, (short)2);
  private static final TField MY_STRING_FIELD_FIELD_DESC = new TField("myStringField", TType.STRING, (short)3);
  private static final TField MY_STRING_TYPEDEF_FIELD_DESC = new TField("myStringTypedef", TType.STRING, (short)4);
  private static final TField MY_MAP_FIELD_FIELD_DESC = new TField("myMapField", TType.MAP, (short)5);
  private static final TField MY_MAP_TYPEDEF_FIELD_DESC = new TField("myMapTypedef", TType.MAP, (short)6);
  private static final TField MY_LIST_FIELD_FIELD_DESC = new TField("myListField", TType.LIST, (short)7);
  private static final TField MY_LIST_TYPEDEF_FIELD_DESC = new TField("myListTypedef", TType.LIST, (short)8);
  private static final TField MY_MAP_LIST_OF_TYPE_DEF_FIELD_DESC = new TField("myMapListOfTypeDef", TType.MAP, (short)9);

  public long myLongField;
  public long myLongTypeDef;
  public String myStringField;
  public String myStringTypedef;
  public Map<Short,String> myMapField;
  public Map<Short,String> myMapTypedef;
  public List<Double> myListField;
  public List<Double> myListTypedef;
  public Map<Short,List<List<Double>>> myMapListOfTypeDef;
  public static final int MYLONGFIELD = 1;
  public static final int MYLONGTYPEDEF = 2;
  public static final int MYSTRINGFIELD = 3;
  public static final int MYSTRINGTYPEDEF = 4;
  public static final int MYMAPFIELD = 5;
  public static final int MYMAPTYPEDEF = 6;
  public static final int MYLISTFIELD = 7;
  public static final int MYLISTTYPEDEF = 8;
  public static final int MYMAPLISTOFTYPEDEF = 9;

  // isset id assignments
  private static final int __MYLONGFIELD_ISSET_ID = 0;
  private static final int __MYLONGTYPEDEF_ISSET_ID = 1;
  private BitSet __isset_bit_vector = new BitSet(2);

  public static final Map<Integer, FieldMetaData> metaDataMap;

  static {
    Map<Integer, FieldMetaData> tmpMetaDataMap = new HashMap<Integer, FieldMetaData>();
    tmpMetaDataMap.put(MYLONGFIELD, new FieldMetaData("myLongField", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.I64)));
    tmpMetaDataMap.put(MYLONGTYPEDEF, new FieldMetaData("myLongTypeDef", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.I64)));
    tmpMetaDataMap.put(MYSTRINGFIELD, new FieldMetaData("myStringField", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.STRING)));
    tmpMetaDataMap.put(MYSTRINGTYPEDEF, new FieldMetaData("myStringTypedef", TFieldRequirementType.DEFAULT, 
        new FieldValueMetaData(TType.STRING)));
    tmpMetaDataMap.put(MYMAPFIELD, new FieldMetaData("myMapField", TFieldRequirementType.DEFAULT, 
        new MapMetaData(TType.MAP, 
            new FieldValueMetaData(TType.I16), 
            new FieldValueMetaData(TType.STRING))));
    tmpMetaDataMap.put(MYMAPTYPEDEF, new FieldMetaData("myMapTypedef", TFieldRequirementType.DEFAULT, 
        new MapMetaData(TType.MAP, 
            new FieldValueMetaData(TType.I16), 
            new FieldValueMetaData(TType.STRING))));
    tmpMetaDataMap.put(MYLISTFIELD, new FieldMetaData("myListField", TFieldRequirementType.DEFAULT, 
        new ListMetaData(TType.LIST, 
            new FieldValueMetaData(TType.DOUBLE))));
    tmpMetaDataMap.put(MYLISTTYPEDEF, new FieldMetaData("myListTypedef", TFieldRequirementType.DEFAULT, 
        new ListMetaData(TType.LIST, 
            new FieldValueMetaData(TType.DOUBLE))));
    tmpMetaDataMap.put(MYMAPLISTOFTYPEDEF, new FieldMetaData("myMapListOfTypeDef", TFieldRequirementType.DEFAULT, 
        new MapMetaData(TType.MAP, 
            new FieldValueMetaData(TType.I16), 
            new ListMetaData(TType.LIST, 
                new ListMetaData(TType.LIST, 
                    new FieldValueMetaData(TType.DOUBLE))))));
    metaDataMap = Collections.unmodifiableMap(tmpMetaDataMap);
  }

  static {
    FieldMetaData.addStructMetaDataMap(MyStructTypeDef.class, metaDataMap);
  }

  public MyStructTypeDef() {
  }

  public MyStructTypeDef(
      long myLongField,
      long myLongTypeDef,
      String myStringField,
      String myStringTypedef,
      Map<Short,String> myMapField,
      Map<Short,String> myMapTypedef,
      List<Double> myListField,
      List<Double> myListTypedef,
      Map<Short,List<List<Double>>> myMapListOfTypeDef) {
    this();
    this.myLongField = myLongField;
    setMyLongFieldIsSet(true);
    this.myLongTypeDef = myLongTypeDef;
    setMyLongTypeDefIsSet(true);
    this.myStringField = myStringField;
    this.myStringTypedef = myStringTypedef;
    this.myMapField = myMapField;
    this.myMapTypedef = myMapTypedef;
    this.myListField = myListField;
    this.myListTypedef = myListTypedef;
    this.myMapListOfTypeDef = myMapListOfTypeDef;
  }

  public static class Builder {
    private long myLongField;
    private long myLongTypeDef;
    private String myStringField;
    private String myStringTypedef;
    private Map<Short,String> myMapField;
    private Map<Short,String> myMapTypedef;
    private List<Double> myListField;
    private List<Double> myListTypedef;
    private Map<Short,List<List<Double>>> myMapListOfTypeDef;

    BitSet __optional_isset = new BitSet(2);

    public Builder() {
    }

    public Builder setMyLongField(final long myLongField) {
      this.myLongField = myLongField;
      __optional_isset.set(__MYLONGFIELD_ISSET_ID, true);
      return this;
    }

    public Builder setMyLongTypeDef(final long myLongTypeDef) {
      this.myLongTypeDef = myLongTypeDef;
      __optional_isset.set(__MYLONGTYPEDEF_ISSET_ID, true);
      return this;
    }

    public Builder setMyStringField(final String myStringField) {
      this.myStringField = myStringField;
      return this;
    }

    public Builder setMyStringTypedef(final String myStringTypedef) {
      this.myStringTypedef = myStringTypedef;
      return this;
    }

    public Builder setMyMapField(final Map<Short,String> myMapField) {
      this.myMapField = myMapField;
      return this;
    }

    public Builder setMyMapTypedef(final Map<Short,String> myMapTypedef) {
      this.myMapTypedef = myMapTypedef;
      return this;
    }

    public Builder setMyListField(final List<Double> myListField) {
      this.myListField = myListField;
      return this;
    }

    public Builder setMyListTypedef(final List<Double> myListTypedef) {
      this.myListTypedef = myListTypedef;
      return this;
    }

    public Builder setMyMapListOfTypeDef(final Map<Short,List<List<Double>>> myMapListOfTypeDef) {
      this.myMapListOfTypeDef = myMapListOfTypeDef;
      return this;
    }

    public MyStructTypeDef build() {
      MyStructTypeDef result = new MyStructTypeDef();
      if (__optional_isset.get(__MYLONGFIELD_ISSET_ID)) {
        result.setMyLongField(this.myLongField);
      }
      if (__optional_isset.get(__MYLONGTYPEDEF_ISSET_ID)) {
        result.setMyLongTypeDef(this.myLongTypeDef);
      }
      result.setMyStringField(this.myStringField);
      result.setMyStringTypedef(this.myStringTypedef);
      result.setMyMapField(this.myMapField);
      result.setMyMapTypedef(this.myMapTypedef);
      result.setMyListField(this.myListField);
      result.setMyListTypedef(this.myListTypedef);
      result.setMyMapListOfTypeDef(this.myMapListOfTypeDef);
      return result;
    }
  }

  public static Builder builder() {
    return new Builder();
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public MyStructTypeDef(MyStructTypeDef other) {
    __isset_bit_vector.clear();
    __isset_bit_vector.or(other.__isset_bit_vector);
    this.myLongField = TBaseHelper.deepCopy(other.myLongField);
    this.myLongTypeDef = TBaseHelper.deepCopy(other.myLongTypeDef);
    if (other.isSetMyStringField()) {
      this.myStringField = TBaseHelper.deepCopy(other.myStringField);
    }
    if (other.isSetMyStringTypedef()) {
      this.myStringTypedef = TBaseHelper.deepCopy(other.myStringTypedef);
    }
    if (other.isSetMyMapField()) {
      this.myMapField = TBaseHelper.deepCopy(other.myMapField);
    }
    if (other.isSetMyMapTypedef()) {
      this.myMapTypedef = TBaseHelper.deepCopy(other.myMapTypedef);
    }
    if (other.isSetMyListField()) {
      this.myListField = TBaseHelper.deepCopy(other.myListField);
    }
    if (other.isSetMyListTypedef()) {
      this.myListTypedef = TBaseHelper.deepCopy(other.myListTypedef);
    }
    if (other.isSetMyMapListOfTypeDef()) {
      this.myMapListOfTypeDef = TBaseHelper.deepCopy(other.myMapListOfTypeDef);
    }
  }

  public MyStructTypeDef deepCopy() {
    return new MyStructTypeDef(this);
  }

  public long getMyLongField() {
    return this.myLongField;
  }

  public MyStructTypeDef setMyLongField(long myLongField) {
    this.myLongField = myLongField;
    setMyLongFieldIsSet(true);
    return this;
  }

  public void unsetMyLongField() {
    __isset_bit_vector.clear(__MYLONGFIELD_ISSET_ID);
  }

  // Returns true if field myLongField is set (has been assigned a value) and false otherwise
  public boolean isSetMyLongField() {
    return __isset_bit_vector.get(__MYLONGFIELD_ISSET_ID);
  }

  public void setMyLongFieldIsSet(boolean __value) {
    __isset_bit_vector.set(__MYLONGFIELD_ISSET_ID, __value);
  }

  public long getMyLongTypeDef() {
    return this.myLongTypeDef;
  }

  public MyStructTypeDef setMyLongTypeDef(long myLongTypeDef) {
    this.myLongTypeDef = myLongTypeDef;
    setMyLongTypeDefIsSet(true);
    return this;
  }

  public void unsetMyLongTypeDef() {
    __isset_bit_vector.clear(__MYLONGTYPEDEF_ISSET_ID);
  }

  // Returns true if field myLongTypeDef is set (has been assigned a value) and false otherwise
  public boolean isSetMyLongTypeDef() {
    return __isset_bit_vector.get(__MYLONGTYPEDEF_ISSET_ID);
  }

  public void setMyLongTypeDefIsSet(boolean __value) {
    __isset_bit_vector.set(__MYLONGTYPEDEF_ISSET_ID, __value);
  }

  public String getMyStringField() {
    return this.myStringField;
  }

  public MyStructTypeDef setMyStringField(String myStringField) {
    this.myStringField = myStringField;
    return this;
  }

  public void unsetMyStringField() {
    this.myStringField = null;
  }

  // Returns true if field myStringField is set (has been assigned a value) and false otherwise
  public boolean isSetMyStringField() {
    return this.myStringField != null;
  }

  public void setMyStringFieldIsSet(boolean __value) {
    if (!__value) {
      this.myStringField = null;
    }
  }

  public String getMyStringTypedef() {
    return this.myStringTypedef;
  }

  public MyStructTypeDef setMyStringTypedef(String myStringTypedef) {
    this.myStringTypedef = myStringTypedef;
    return this;
  }

  public void unsetMyStringTypedef() {
    this.myStringTypedef = null;
  }

  // Returns true if field myStringTypedef is set (has been assigned a value) and false otherwise
  public boolean isSetMyStringTypedef() {
    return this.myStringTypedef != null;
  }

  public void setMyStringTypedefIsSet(boolean __value) {
    if (!__value) {
      this.myStringTypedef = null;
    }
  }

  public Map<Short,String> getMyMapField() {
    return this.myMapField;
  }

  public MyStructTypeDef setMyMapField(Map<Short,String> myMapField) {
    this.myMapField = myMapField;
    return this;
  }

  public void unsetMyMapField() {
    this.myMapField = null;
  }

  // Returns true if field myMapField is set (has been assigned a value) and false otherwise
  public boolean isSetMyMapField() {
    return this.myMapField != null;
  }

  public void setMyMapFieldIsSet(boolean __value) {
    if (!__value) {
      this.myMapField = null;
    }
  }

  public Map<Short,String> getMyMapTypedef() {
    return this.myMapTypedef;
  }

  public MyStructTypeDef setMyMapTypedef(Map<Short,String> myMapTypedef) {
    this.myMapTypedef = myMapTypedef;
    return this;
  }

  public void unsetMyMapTypedef() {
    this.myMapTypedef = null;
  }

  // Returns true if field myMapTypedef is set (has been assigned a value) and false otherwise
  public boolean isSetMyMapTypedef() {
    return this.myMapTypedef != null;
  }

  public void setMyMapTypedefIsSet(boolean __value) {
    if (!__value) {
      this.myMapTypedef = null;
    }
  }

  public List<Double> getMyListField() {
    return this.myListField;
  }

  public MyStructTypeDef setMyListField(List<Double> myListField) {
    this.myListField = myListField;
    return this;
  }

  public void unsetMyListField() {
    this.myListField = null;
  }

  // Returns true if field myListField is set (has been assigned a value) and false otherwise
  public boolean isSetMyListField() {
    return this.myListField != null;
  }

  public void setMyListFieldIsSet(boolean __value) {
    if (!__value) {
      this.myListField = null;
    }
  }

  public List<Double> getMyListTypedef() {
    return this.myListTypedef;
  }

  public MyStructTypeDef setMyListTypedef(List<Double> myListTypedef) {
    this.myListTypedef = myListTypedef;
    return this;
  }

  public void unsetMyListTypedef() {
    this.myListTypedef = null;
  }

  // Returns true if field myListTypedef is set (has been assigned a value) and false otherwise
  public boolean isSetMyListTypedef() {
    return this.myListTypedef != null;
  }

  public void setMyListTypedefIsSet(boolean __value) {
    if (!__value) {
      this.myListTypedef = null;
    }
  }

  public Map<Short,List<List<Double>>> getMyMapListOfTypeDef() {
    return this.myMapListOfTypeDef;
  }

  public MyStructTypeDef setMyMapListOfTypeDef(Map<Short,List<List<Double>>> myMapListOfTypeDef) {
    this.myMapListOfTypeDef = myMapListOfTypeDef;
    return this;
  }

  public void unsetMyMapListOfTypeDef() {
    this.myMapListOfTypeDef = null;
  }

  // Returns true if field myMapListOfTypeDef is set (has been assigned a value) and false otherwise
  public boolean isSetMyMapListOfTypeDef() {
    return this.myMapListOfTypeDef != null;
  }

  public void setMyMapListOfTypeDefIsSet(boolean __value) {
    if (!__value) {
      this.myMapListOfTypeDef = null;
    }
  }

  @SuppressWarnings("unchecked")
  public void setFieldValue(int fieldID, Object __value) {
    switch (fieldID) {
    case MYLONGFIELD:
      if (__value == null) {
        unsetMyLongField();
      } else {
        setMyLongField((Long)__value);
      }
      break;

    case MYLONGTYPEDEF:
      if (__value == null) {
        unsetMyLongTypeDef();
      } else {
        setMyLongTypeDef((Long)__value);
      }
      break;

    case MYSTRINGFIELD:
      if (__value == null) {
        unsetMyStringField();
      } else {
        setMyStringField((String)__value);
      }
      break;

    case MYSTRINGTYPEDEF:
      if (__value == null) {
        unsetMyStringTypedef();
      } else {
        setMyStringTypedef((String)__value);
      }
      break;

    case MYMAPFIELD:
      if (__value == null) {
        unsetMyMapField();
      } else {
        setMyMapField((Map<Short,String>)__value);
      }
      break;

    case MYMAPTYPEDEF:
      if (__value == null) {
        unsetMyMapTypedef();
      } else {
        setMyMapTypedef((Map<Short,String>)__value);
      }
      break;

    case MYLISTFIELD:
      if (__value == null) {
        unsetMyListField();
      } else {
        setMyListField((List<Double>)__value);
      }
      break;

    case MYLISTTYPEDEF:
      if (__value == null) {
        unsetMyListTypedef();
      } else {
        setMyListTypedef((List<Double>)__value);
      }
      break;

    case MYMAPLISTOFTYPEDEF:
      if (__value == null) {
        unsetMyMapListOfTypeDef();
      } else {
        setMyMapListOfTypeDef((Map<Short,List<List<Double>>>)__value);
      }
      break;

    default:
      throw new IllegalArgumentException("Field " + fieldID + " doesn't exist!");
    }
  }

  public Object getFieldValue(int fieldID) {
    switch (fieldID) {
    case MYLONGFIELD:
      return new Long(getMyLongField());

    case MYLONGTYPEDEF:
      return new Long(getMyLongTypeDef());

    case MYSTRINGFIELD:
      return getMyStringField();

    case MYSTRINGTYPEDEF:
      return getMyStringTypedef();

    case MYMAPFIELD:
      return getMyMapField();

    case MYMAPTYPEDEF:
      return getMyMapTypedef();

    case MYLISTFIELD:
      return getMyListField();

    case MYLISTTYPEDEF:
      return getMyListTypedef();

    case MYMAPLISTOFTYPEDEF:
      return getMyMapListOfTypeDef();

    default:
      throw new IllegalArgumentException("Field " + fieldID + " doesn't exist!");
    }
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof MyStructTypeDef))
      return false;
    MyStructTypeDef that = (MyStructTypeDef)_that;

    if (!TBaseHelper.equalsNobinary(this.myLongField, that.myLongField)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.myLongTypeDef, that.myLongTypeDef)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyStringField(), that.isSetMyStringField(), this.myStringField, that.myStringField)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyStringTypedef(), that.isSetMyStringTypedef(), this.myStringTypedef, that.myStringTypedef)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyMapField(), that.isSetMyMapField(), this.myMapField, that.myMapField)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyMapTypedef(), that.isSetMyMapTypedef(), this.myMapTypedef, that.myMapTypedef)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyListField(), that.isSetMyListField(), this.myListField, that.myListField)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyListTypedef(), that.isSetMyListTypedef(), this.myListTypedef, that.myListTypedef)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetMyMapListOfTypeDef(), that.isSetMyMapListOfTypeDef(), this.myMapListOfTypeDef, that.myMapListOfTypeDef)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {myLongField, myLongTypeDef, myStringField, myStringTypedef, myMapField, myMapTypedef, myListField, myListTypedef, myMapListOfTypeDef});
  }

  @Override
  public int compareTo(MyStructTypeDef other) {
    if (other == null) {
      // See java.lang.Comparable docs
      throw new NullPointerException();
    }

    if (other == this) {
      return 0;
    }
    int lastComparison = 0;

    lastComparison = Boolean.valueOf(isSetMyLongField()).compareTo(other.isSetMyLongField());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myLongField, other.myLongField);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyLongTypeDef()).compareTo(other.isSetMyLongTypeDef());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myLongTypeDef, other.myLongTypeDef);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyStringField()).compareTo(other.isSetMyStringField());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myStringField, other.myStringField);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyStringTypedef()).compareTo(other.isSetMyStringTypedef());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myStringTypedef, other.myStringTypedef);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyMapField()).compareTo(other.isSetMyMapField());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myMapField, other.myMapField);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyMapTypedef()).compareTo(other.isSetMyMapTypedef());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myMapTypedef, other.myMapTypedef);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyListField()).compareTo(other.isSetMyListField());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myListField, other.myListField);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyListTypedef()).compareTo(other.isSetMyListTypedef());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myListTypedef, other.myListTypedef);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    lastComparison = Boolean.valueOf(isSetMyMapListOfTypeDef()).compareTo(other.isSetMyMapListOfTypeDef());
    if (lastComparison != 0) {
      return lastComparison;
    }
    lastComparison = TBaseHelper.compareTo(myMapListOfTypeDef, other.myMapListOfTypeDef);
    if (lastComparison != 0) { 
      return lastComparison;
    }
    return 0;
  }

  public void read(TProtocol iprot) throws TException {
    TField __field;
    iprot.readStructBegin(metaDataMap);
    while (true)
    {
      __field = iprot.readFieldBegin();
      if (__field.type == TType.STOP) {
        break;
      }
      switch (__field.id)
      {
        case MYLONGFIELD:
          if (__field.type == TType.I64) {
            this.myLongField = iprot.readI64();
            setMyLongFieldIsSet(true);
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYLONGTYPEDEF:
          if (__field.type == TType.I64) {
            this.myLongTypeDef = iprot.readI64();
            setMyLongTypeDefIsSet(true);
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYSTRINGFIELD:
          if (__field.type == TType.STRING) {
            this.myStringField = iprot.readString();
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYSTRINGTYPEDEF:
          if (__field.type == TType.STRING) {
            this.myStringTypedef = iprot.readString();
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYMAPFIELD:
          if (__field.type == TType.MAP) {
            {
              TMap _map170 = iprot.readMapBegin();
              this.myMapField = new HashMap<Short,String>(Math.max(0, 2*_map170.size));
              for (int _i171 = 0; 
                   (_map170.size < 0) ? iprot.peekMap() : (_i171 < _map170.size); 
                   ++_i171)
              {
                short _key172;
                String _val173;
                _key172 = iprot.readI16();
                _val173 = iprot.readString();
                this.myMapField.put(_key172, _val173);
              }
              iprot.readMapEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYMAPTYPEDEF:
          if (__field.type == TType.MAP) {
            {
              TMap _map174 = iprot.readMapBegin();
              this.myMapTypedef = new HashMap<Short,String>(Math.max(0, 2*_map174.size));
              for (int _i175 = 0; 
                   (_map174.size < 0) ? iprot.peekMap() : (_i175 < _map174.size); 
                   ++_i175)
              {
                short _key176;
                String _val177;
                _key176 = iprot.readI16();
                _val177 = iprot.readString();
                this.myMapTypedef.put(_key176, _val177);
              }
              iprot.readMapEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYLISTFIELD:
          if (__field.type == TType.LIST) {
            {
              TList _list178 = iprot.readListBegin();
              this.myListField = new ArrayList<Double>(Math.max(0, _list178.size));
              for (int _i179 = 0; 
                   (_list178.size < 0) ? iprot.peekList() : (_i179 < _list178.size); 
                   ++_i179)
              {
                double _elem180;
                _elem180 = iprot.readDouble();
                this.myListField.add(_elem180);
              }
              iprot.readListEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYLISTTYPEDEF:
          if (__field.type == TType.LIST) {
            {
              TList _list181 = iprot.readListBegin();
              this.myListTypedef = new ArrayList<Double>(Math.max(0, _list181.size));
              for (int _i182 = 0; 
                   (_list181.size < 0) ? iprot.peekList() : (_i182 < _list181.size); 
                   ++_i182)
              {
                double _elem183;
                _elem183 = iprot.readDouble();
                this.myListTypedef.add(_elem183);
              }
              iprot.readListEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case MYMAPLISTOFTYPEDEF:
          if (__field.type == TType.MAP) {
            {
              TMap _map184 = iprot.readMapBegin();
              this.myMapListOfTypeDef = new HashMap<Short,List<List<Double>>>(Math.max(0, 2*_map184.size));
              for (int _i185 = 0; 
                   (_map184.size < 0) ? iprot.peekMap() : (_i185 < _map184.size); 
                   ++_i185)
              {
                short _key186;
                List<List<Double>> _val187;
                _key186 = iprot.readI16();
                {
                  TList _list188 = iprot.readListBegin();
                  _val187 = new ArrayList<List<Double>>(Math.max(0, _list188.size));
                  for (int _i189 = 0; 
                       (_list188.size < 0) ? iprot.peekList() : (_i189 < _list188.size); 
                       ++_i189)
                  {
                    List<Double> _elem190;
                    {
                      TList _list191 = iprot.readListBegin();
                      _elem190 = new ArrayList<Double>(Math.max(0, _list191.size));
                      for (int _i192 = 0; 
                           (_list191.size < 0) ? iprot.peekList() : (_i192 < _list191.size); 
                           ++_i192)
                      {
                        double _elem193;
                        _elem193 = iprot.readDouble();
                        _elem190.add(_elem193);
                      }
                      iprot.readListEnd();
                    }
                    _val187.add(_elem190);
                  }
                  iprot.readListEnd();
                }
                this.myMapListOfTypeDef.put(_key186, _val187);
              }
              iprot.readMapEnd();
            }
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        default:
          TProtocolUtil.skip(iprot, __field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();


    // check for required fields of primitive type, which can't be checked in the validate method
    validate();
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    oprot.writeFieldBegin(MY_LONG_FIELD_FIELD_DESC);
    oprot.writeI64(this.myLongField);
    oprot.writeFieldEnd();
    oprot.writeFieldBegin(MY_LONG_TYPE_DEF_FIELD_DESC);
    oprot.writeI64(this.myLongTypeDef);
    oprot.writeFieldEnd();
    if (this.myStringField != null) {
      oprot.writeFieldBegin(MY_STRING_FIELD_FIELD_DESC);
      oprot.writeString(this.myStringField);
      oprot.writeFieldEnd();
    }
    if (this.myStringTypedef != null) {
      oprot.writeFieldBegin(MY_STRING_TYPEDEF_FIELD_DESC);
      oprot.writeString(this.myStringTypedef);
      oprot.writeFieldEnd();
    }
    if (this.myMapField != null) {
      oprot.writeFieldBegin(MY_MAP_FIELD_FIELD_DESC);
      {
        oprot.writeMapBegin(new TMap(TType.I16, TType.STRING, this.myMapField.size()));
        for (Map.Entry<Short, String> _iter194 : this.myMapField.entrySet())        {
          oprot.writeI16(_iter194.getKey());
          oprot.writeString(_iter194.getValue());
        }
        oprot.writeMapEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.myMapTypedef != null) {
      oprot.writeFieldBegin(MY_MAP_TYPEDEF_FIELD_DESC);
      {
        oprot.writeMapBegin(new TMap(TType.I16, TType.STRING, this.myMapTypedef.size()));
        for (Map.Entry<Short, String> _iter195 : this.myMapTypedef.entrySet())        {
          oprot.writeI16(_iter195.getKey());
          oprot.writeString(_iter195.getValue());
        }
        oprot.writeMapEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.myListField != null) {
      oprot.writeFieldBegin(MY_LIST_FIELD_FIELD_DESC);
      {
        oprot.writeListBegin(new TList(TType.DOUBLE, this.myListField.size()));
        for (double _iter196 : this.myListField)        {
          oprot.writeDouble(_iter196);
        }
        oprot.writeListEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.myListTypedef != null) {
      oprot.writeFieldBegin(MY_LIST_TYPEDEF_FIELD_DESC);
      {
        oprot.writeListBegin(new TList(TType.DOUBLE, this.myListTypedef.size()));
        for (double _iter197 : this.myListTypedef)        {
          oprot.writeDouble(_iter197);
        }
        oprot.writeListEnd();
      }
      oprot.writeFieldEnd();
    }
    if (this.myMapListOfTypeDef != null) {
      oprot.writeFieldBegin(MY_MAP_LIST_OF_TYPE_DEF_FIELD_DESC);
      {
        oprot.writeMapBegin(new TMap(TType.I16, TType.LIST, this.myMapListOfTypeDef.size()));
        for (Map.Entry<Short, List<List<Double>>> _iter198 : this.myMapListOfTypeDef.entrySet())        {
          oprot.writeI16(_iter198.getKey());
          {
            oprot.writeListBegin(new TList(TType.LIST, _iter198.getValue().size()));
            for (List<Double> _iter199 : _iter198.getValue())            {
              {
                oprot.writeListBegin(new TList(TType.DOUBLE, _iter199.size()));
                for (double _iter200 : _iter199)                {
                  oprot.writeDouble(_iter200);
                }
                oprot.writeListEnd();
              }
            }
            oprot.writeListEnd();
          }
        }
        oprot.writeMapEnd();
      }
      oprot.writeFieldEnd();
    }
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  @Override
  public String toString() {
    return toString(1, true);
  }

  @Override
  public String toString(int indent, boolean prettyPrint) {
    String indentStr = prettyPrint ? TBaseHelper.getIndentedString(indent) : "";
    String newLine = prettyPrint ? "\n" : "";
    String space = prettyPrint ? " " : "";
    StringBuilder sb = new StringBuilder("MyStructTypeDef");
    sb.append(space);
    sb.append("(");
    sb.append(newLine);
    boolean first = true;

    sb.append(indentStr);
    sb.append("myLongField");
    sb.append(space);
    sb.append(":").append(space);
    sb.append(TBaseHelper.toString(this.getMyLongField(), indent + 1, prettyPrint));
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myLongTypeDef");
    sb.append(space);
    sb.append(":").append(space);
    sb.append(TBaseHelper.toString(this.getMyLongTypeDef(), indent + 1, prettyPrint));
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myStringField");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyStringField() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyStringField(), indent + 1, prettyPrint));
    }
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myStringTypedef");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyStringTypedef() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyStringTypedef(), indent + 1, prettyPrint));
    }
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myMapField");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyMapField() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyMapField(), indent + 1, prettyPrint));
    }
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myMapTypedef");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyMapTypedef() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyMapTypedef(), indent + 1, prettyPrint));
    }
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myListField");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyListField() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyListField(), indent + 1, prettyPrint));
    }
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myListTypedef");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyListTypedef() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyListTypedef(), indent + 1, prettyPrint));
    }
    first = false;
    if (!first) sb.append("," + newLine);
    sb.append(indentStr);
    sb.append("myMapListOfTypeDef");
    sb.append(space);
    sb.append(":").append(space);
    if (this.getMyMapListOfTypeDef() == null) {
      sb.append("null");
    } else {
      sb.append(TBaseHelper.toString(this.getMyMapListOfTypeDef(), indent + 1, prettyPrint));
    }
    first = false;
    sb.append(newLine + TBaseHelper.reduceIndent(indentStr));
    sb.append(")");
    return sb.toString();
  }

  public void validate() throws TException {
    // check for required fields
  }

}

