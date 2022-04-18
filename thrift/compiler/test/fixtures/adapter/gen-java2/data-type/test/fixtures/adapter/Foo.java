/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.adapter;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import com.google.common.collect.*;
import java.util.*;
import javax.annotation.Nullable;
import org.apache.thrift.*;
import org.apache.thrift.transport.*;
import org.apache.thrift.protocol.*;
import static com.google.common.base.MoreObjects.toStringHelper;
import static com.google.common.base.MoreObjects.ToStringHelper;

@SwiftGenerated
@com.facebook.swift.codec.ThriftStruct(value="Foo", builder=Foo.Builder.class)
public final class Foo implements com.facebook.thrift.payload.ThriftSerializable {

    @ThriftConstructor
    public Foo(
        @com.facebook.swift.codec.ThriftField(value=1, name="intField", requiredness=Requiredness.NONE) final int intField,
        @com.facebook.swift.codec.ThriftField(value=2, name="optionalIntField", requiredness=Requiredness.OPTIONAL) final Integer optionalIntField,
        @com.facebook.swift.codec.ThriftField(value=3, name="intFieldWithDefault", requiredness=Requiredness.NONE) final int intFieldWithDefault,
        @com.facebook.swift.codec.ThriftField(value=4, name="setField", requiredness=Requiredness.NONE) final Set<String> setField,
        @com.facebook.swift.codec.ThriftField(value=5, name="optionalSetField", requiredness=Requiredness.OPTIONAL) final Set<String> optionalSetField,
        @com.facebook.swift.codec.ThriftField(value=6, name="mapField", requiredness=Requiredness.NONE) final Map<String, List<String>> mapField,
        @com.facebook.swift.codec.ThriftField(value=7, name="optionalMapField", requiredness=Requiredness.OPTIONAL) final Map<String, List<String>> optionalMapField,
        @com.facebook.swift.codec.ThriftField(value=8, name="binaryField", requiredness=Requiredness.NONE) final byte[] binaryField,
        @com.facebook.swift.codec.ThriftField(value=9, name="longField", requiredness=Requiredness.NONE) final long longField,
        @com.facebook.swift.codec.ThriftField(value=10, name="adaptedLongField", requiredness=Requiredness.NONE) final long adaptedLongField
    ) {
        this.intField = intField;
        this.optionalIntField = optionalIntField;
        this.intFieldWithDefault = intFieldWithDefault;
        this.setField = setField;
        this.optionalSetField = optionalSetField;
        this.mapField = mapField;
        this.optionalMapField = optionalMapField;
        this.binaryField = binaryField;
        this.longField = longField;
        this.adaptedLongField = adaptedLongField;
    }
    
    @ThriftConstructor
    protected Foo() {
      this.intField = 0;
      this.optionalIntField = null;
      this.intFieldWithDefault = 13;
      this.setField = null;
      this.optionalSetField = null;
      this.mapField = null;
      this.optionalMapField = null;
      this.binaryField = null;
      this.longField = 0L;
      this.adaptedLongField = 0L;
    }
    
    public static class Builder {
    
        private int intField = 0;
        private Integer optionalIntField = null;
        private int intFieldWithDefault = 13;
        private Set<String> setField = null;
        private Set<String> optionalSetField = null;
        private Map<String, List<String>> mapField = null;
        private Map<String, List<String>> optionalMapField = null;
        private byte[] binaryField = null;
        private long longField = 0L;
        private long adaptedLongField = 0L;
    
        @com.facebook.swift.codec.ThriftField(value=1, name="intField", requiredness=Requiredness.NONE)
        public Builder setIntField(int intField) {
            this.intField = intField;
            return this;
        }
    
        public int getIntField() { return intField; }
    
            @com.facebook.swift.codec.ThriftField(value=2, name="optionalIntField", requiredness=Requiredness.OPTIONAL)
        public Builder setOptionalIntField(Integer optionalIntField) {
            this.optionalIntField = optionalIntField;
            return this;
        }
    
        public Integer getOptionalIntField() { return optionalIntField; }
    
            @com.facebook.swift.codec.ThriftField(value=3, name="intFieldWithDefault", requiredness=Requiredness.NONE)
        public Builder setIntFieldWithDefault(int intFieldWithDefault) {
            this.intFieldWithDefault = intFieldWithDefault;
            return this;
        }
    
        public int getIntFieldWithDefault() { return intFieldWithDefault; }
    
            @com.facebook.swift.codec.ThriftField(value=4, name="setField", requiredness=Requiredness.NONE)
        public Builder setSetField(Set<String> setField) {
            this.setField = setField;
            return this;
        }
    
        public Set<String> getSetField() { return setField; }
    
            @com.facebook.swift.codec.ThriftField(value=5, name="optionalSetField", requiredness=Requiredness.OPTIONAL)
        public Builder setOptionalSetField(Set<String> optionalSetField) {
            this.optionalSetField = optionalSetField;
            return this;
        }
    
        public Set<String> getOptionalSetField() { return optionalSetField; }
    
            @com.facebook.swift.codec.ThriftField(value=6, name="mapField", requiredness=Requiredness.NONE)
        public Builder setMapField(Map<String, List<String>> mapField) {
            this.mapField = mapField;
            return this;
        }
    
        public Map<String, List<String>> getMapField() { return mapField; }
    
            @com.facebook.swift.codec.ThriftField(value=7, name="optionalMapField", requiredness=Requiredness.OPTIONAL)
        public Builder setOptionalMapField(Map<String, List<String>> optionalMapField) {
            this.optionalMapField = optionalMapField;
            return this;
        }
    
        public Map<String, List<String>> getOptionalMapField() { return optionalMapField; }
    
            @com.facebook.swift.codec.ThriftField(value=8, name="binaryField", requiredness=Requiredness.NONE)
        public Builder setBinaryField(byte[] binaryField) {
            this.binaryField = binaryField;
            return this;
        }
    
        public byte[] getBinaryField() { return binaryField; }
    
            @com.facebook.swift.codec.ThriftField(value=9, name="longField", requiredness=Requiredness.NONE)
        public Builder setLongField(long longField) {
            this.longField = longField;
            return this;
        }
    
        public long getLongField() { return longField; }
    
            @com.facebook.swift.codec.ThriftField(value=10, name="adaptedLongField", requiredness=Requiredness.NONE)
        public Builder setAdaptedLongField(long adaptedLongField) {
            this.adaptedLongField = adaptedLongField;
            return this;
        }
    
        public long getAdaptedLongField() { return adaptedLongField; }
    
        public Builder() { }
        public Builder(Foo other) {
            this.intField = other.intField;
            this.optionalIntField = other.optionalIntField;
            this.intFieldWithDefault = other.intFieldWithDefault;
            this.setField = other.setField;
            this.optionalSetField = other.optionalSetField;
            this.mapField = other.mapField;
            this.optionalMapField = other.optionalMapField;
            this.binaryField = other.binaryField;
            this.longField = other.longField;
            this.adaptedLongField = other.adaptedLongField;
        }
    
        @ThriftConstructor
        public Foo build() {
            Foo result = new Foo (
                this.intField,
                this.optionalIntField,
                this.intFieldWithDefault,
                this.setField,
                this.optionalSetField,
                this.mapField,
                this.optionalMapField,
                this.binaryField,
                this.longField,
                this.adaptedLongField
            );
            return result;
        }
    }
                                            public static final Map<String, Integer> NAMES_TO_IDS = new HashMap();
    public static final Map<String, Integer> THRIFT_NAMES_TO_IDS = new HashMap();
    public static final Map<Integer, TField> FIELD_METADATA = new HashMap<>();
    private static final TStruct STRUCT_DESC = new TStruct("Foo");
    private final int intField;
    public static final int _INTFIELD = 1;
    private static final TField INT_FIELD_FIELD_DESC = new TField("intField", TType.I32, (short)1);
        private final Integer optionalIntField;
    public static final int _OPTIONALINTFIELD = 2;
    private static final TField OPTIONAL_INT_FIELD_FIELD_DESC = new TField("optionalIntField", TType.I32, (short)2);
        private final int intFieldWithDefault;
    public static final int _INTFIELDWITHDEFAULT = 3;
    private static final TField INT_FIELD_WITH_DEFAULT_FIELD_DESC = new TField("intFieldWithDefault", TType.I32, (short)3);
        private final Set<String> setField;
    public static final int _SETFIELD = 4;
    private static final TField SET_FIELD_FIELD_DESC = new TField("setField", TType.SET, (short)4);
        private final Set<String> optionalSetField;
    public static final int _OPTIONALSETFIELD = 5;
    private static final TField OPTIONAL_SET_FIELD_FIELD_DESC = new TField("optionalSetField", TType.SET, (short)5);
        private final Map<String, List<String>> mapField;
    public static final int _MAPFIELD = 6;
    private static final TField MAP_FIELD_FIELD_DESC = new TField("mapField", TType.MAP, (short)6);
        private final Map<String, List<String>> optionalMapField;
    public static final int _OPTIONALMAPFIELD = 7;
    private static final TField OPTIONAL_MAP_FIELD_FIELD_DESC = new TField("optionalMapField", TType.MAP, (short)7);
        private final byte[] binaryField;
    public static final int _BINARYFIELD = 8;
    private static final TField BINARY_FIELD_FIELD_DESC = new TField("binaryField", TType.STRING, (short)8);
        private final long longField;
    public static final int _LONGFIELD = 9;
    private static final TField LONG_FIELD_FIELD_DESC = new TField("longField", TType.I64, (short)9);
        private final long adaptedLongField;
    public static final int _ADAPTEDLONGFIELD = 10;
    private static final TField ADAPTED_LONG_FIELD_FIELD_DESC = new TField("adaptedLongField", TType.I64, (short)10);
    static {
      NAMES_TO_IDS.put("intField", 1);
      THRIFT_NAMES_TO_IDS.put("intField", 1);
      FIELD_METADATA.put(1, INT_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("optionalIntField", 2);
      THRIFT_NAMES_TO_IDS.put("optionalIntField", 2);
      FIELD_METADATA.put(2, OPTIONAL_INT_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("intFieldWithDefault", 3);
      THRIFT_NAMES_TO_IDS.put("intFieldWithDefault", 3);
      FIELD_METADATA.put(3, INT_FIELD_WITH_DEFAULT_FIELD_DESC);
      NAMES_TO_IDS.put("setField", 4);
      THRIFT_NAMES_TO_IDS.put("setField", 4);
      FIELD_METADATA.put(4, SET_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("optionalSetField", 5);
      THRIFT_NAMES_TO_IDS.put("optionalSetField", 5);
      FIELD_METADATA.put(5, OPTIONAL_SET_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("mapField", 6);
      THRIFT_NAMES_TO_IDS.put("mapField", 6);
      FIELD_METADATA.put(6, MAP_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("optionalMapField", 7);
      THRIFT_NAMES_TO_IDS.put("optionalMapField", 7);
      FIELD_METADATA.put(7, OPTIONAL_MAP_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("binaryField", 8);
      THRIFT_NAMES_TO_IDS.put("binaryField", 8);
      FIELD_METADATA.put(8, BINARY_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("longField", 9);
      THRIFT_NAMES_TO_IDS.put("longField", 9);
      FIELD_METADATA.put(9, LONG_FIELD_FIELD_DESC);
      NAMES_TO_IDS.put("adaptedLongField", 10);
      THRIFT_NAMES_TO_IDS.put("adaptedLongField", 10);
      FIELD_METADATA.put(10, ADAPTED_LONG_FIELD_FIELD_DESC);
    }
    
    
    @com.facebook.swift.codec.ThriftField(value=1, name="intField", requiredness=Requiredness.NONE)
    public int getIntField() { return intField; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=2, name="optionalIntField", requiredness=Requiredness.OPTIONAL)
    public Integer getOptionalIntField() { return optionalIntField; }
    
    
    
    @com.facebook.swift.codec.ThriftField(value=3, name="intFieldWithDefault", requiredness=Requiredness.NONE)
    public int getIntFieldWithDefault() { return intFieldWithDefault; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=4, name="setField", requiredness=Requiredness.NONE)
    public Set<String> getSetField() { return setField; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=5, name="optionalSetField", requiredness=Requiredness.OPTIONAL)
    public Set<String> getOptionalSetField() { return optionalSetField; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=6, name="mapField", requiredness=Requiredness.NONE)
    public Map<String, List<String>> getMapField() { return mapField; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=7, name="optionalMapField", requiredness=Requiredness.OPTIONAL)
    public Map<String, List<String>> getOptionalMapField() { return optionalMapField; }
    
    
    @Nullable
    @com.facebook.swift.codec.ThriftField(value=8, name="binaryField", requiredness=Requiredness.NONE)
    public byte[] getBinaryField() { return binaryField; }
    
    
    
    @com.facebook.swift.codec.ThriftField(value=9, name="longField", requiredness=Requiredness.NONE)
    public long getLongField() { return longField; }
    
    
    
    @com.facebook.swift.codec.ThriftField(value=10, name="adaptedLongField", requiredness=Requiredness.NONE)
    public long getAdaptedLongField() { return adaptedLongField; }
    
    @java.lang.Override
    public String toString() {
        ToStringHelper helper = toStringHelper(this);
        helper.add("intField", intField);
        helper.add("optionalIntField", optionalIntField);
        helper.add("intFieldWithDefault", intFieldWithDefault);
        helper.add("setField", setField);
        helper.add("optionalSetField", optionalSetField);
        helper.add("mapField", mapField);
        helper.add("optionalMapField", optionalMapField);
        helper.add("binaryField", binaryField);
        helper.add("longField", longField);
        helper.add("adaptedLongField", adaptedLongField);
        return helper.toString();
    }
    
    @java.lang.Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
    
        Foo other = (Foo)o;
    
        return
            Objects.equals(intField, other.intField) &&
    Objects.equals(optionalIntField, other.optionalIntField) &&
    Objects.equals(intFieldWithDefault, other.intFieldWithDefault) &&
    Objects.equals(setField, other.setField) &&
    Objects.equals(optionalSetField, other.optionalSetField) &&
    Objects.equals(mapField, other.mapField) &&
    Objects.equals(optionalMapField, other.optionalMapField) &&
    Arrays.equals(binaryField, other.binaryField) &&
    Objects.equals(longField, other.longField) &&
    Objects.equals(adaptedLongField, other.adaptedLongField) &&
            true;
    }
    
    @java.lang.Override
    public int hashCode() {
        return Arrays.deepHashCode(new Object[] {
            intField,
            optionalIntField,
            intFieldWithDefault,
            setField,
            optionalSetField,
            mapField,
            optionalMapField,
            binaryField,
            longField,
            adaptedLongField
        });
    }
    
    
    public static com.facebook.thrift.payload.Reader<Foo> asReader() {
      return Foo::read0;
    }
    
    public static Foo read0(TProtocol oprot) throws TException {
      TField __field;
      oprot.readStructBegin(Foo.NAMES_TO_IDS, Foo.THRIFT_NAMES_TO_IDS, Foo.FIELD_METADATA);
      Foo.Builder builder = new Foo.Builder();
      while (true) {
        __field = oprot.readFieldBegin();
        if (__field.type == TType.STOP) { break; }
        switch (__field.id) {
        case _INTFIELD:
          if (__field.type == TType.I32) {
            int intField = oprot.readI32();
            builder.setIntField(intField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _OPTIONALINTFIELD:
          if (__field.type == TType.I32) {
            Integer optionalIntField = oprot.readI32();
            builder.setOptionalIntField(optionalIntField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _INTFIELDWITHDEFAULT:
          if (__field.type == TType.I32) {
            int intFieldWithDefault = oprot.readI32();
            builder.setIntFieldWithDefault(intFieldWithDefault);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _SETFIELD:
          if (__field.type == TType.SET) {
            Set<String> setField;
            {
            TSet _set = oprot.readSetBegin();
            setField = new HashSet<String>(Math.max(0, _set.size));
            for (int _i = 0; (_set.size < 0) ? oprot.peekSet() : (_i < _set.size); _i++) {
                
                String _value1 = oprot.readString();
                setField.add(_value1);
            }
            oprot.readSetEnd();
            }
            builder.setSetField(setField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _OPTIONALSETFIELD:
          if (__field.type == TType.SET) {
            Set<String> optionalSetField;
            {
            TSet _set = oprot.readSetBegin();
            optionalSetField = new HashSet<String>(Math.max(0, _set.size));
            for (int _i = 0; (_set.size < 0) ? oprot.peekSet() : (_i < _set.size); _i++) {
                
                String _value1 = oprot.readString();
                optionalSetField.add(_value1);
            }
            oprot.readSetEnd();
            }
            builder.setOptionalSetField(optionalSetField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _MAPFIELD:
          if (__field.type == TType.MAP) {
            Map<String, List<String>> mapField;
            {
            TMap _map = oprot.readMapBegin();
            mapField = new HashMap<String, List<String>>(Math.max(0, _map.size));
            for (int _i = 0; (_map.size < 0) ? oprot.peekMap() : (_i < _map.size); _i++) {
                
                String _key1 = oprot.readString();
                List<String> _value1;
            {
            TList _list1 = oprot.readListBegin();
            _value1 = new ArrayList<String>(Math.max(0, _list1.size));
            for (int _i1 = 0; (_list1.size < 0) ? oprot.peekList() : (_i1 < _list1.size); _i1++) {
                
                
                String _value2 = oprot.readString();
                
                
                _value1.add(_value2);
                
            }
            oprot.readListEnd();
            }
                mapField.put(_key1, _value1);
            }
            }
            oprot.readMapEnd();
            builder.setMapField(mapField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _OPTIONALMAPFIELD:
          if (__field.type == TType.MAP) {
            Map<String, List<String>> optionalMapField;
            {
            TMap _map = oprot.readMapBegin();
            optionalMapField = new HashMap<String, List<String>>(Math.max(0, _map.size));
            for (int _i = 0; (_map.size < 0) ? oprot.peekMap() : (_i < _map.size); _i++) {
                
                String _key1 = oprot.readString();
                List<String> _value1;
            {
            TList _list1 = oprot.readListBegin();
            _value1 = new ArrayList<String>(Math.max(0, _list1.size));
            for (int _i1 = 0; (_list1.size < 0) ? oprot.peekList() : (_i1 < _list1.size); _i1++) {
                
                
                String _value2 = oprot.readString();
                
                
                _value1.add(_value2);
                
            }
            oprot.readListEnd();
            }
                optionalMapField.put(_key1, _value1);
            }
            }
            oprot.readMapEnd();
            builder.setOptionalMapField(optionalMapField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _BINARYFIELD:
          if (__field.type == TType.STRING) {
            byte[] binaryField = oprot.readBinary().array();
            builder.setBinaryField(binaryField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _LONGFIELD:
          if (__field.type == TType.I64) {
            long longField = oprot.readI64();
            builder.setLongField(longField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        case _ADAPTEDLONGFIELD:
          if (__field.type == TType.I64) {
            long adaptedLongField = oprot.readI64();
            builder.setAdaptedLongField(adaptedLongField);
          } else {
            TProtocolUtil.skip(oprot, __field.type);
          }
          break;
        default:
          TProtocolUtil.skip(oprot, __field.type);
          break;
        }
        oprot.readFieldEnd();
      }
      oprot.readStructEnd();
      return builder.build();
    }
    
    public void write0(TProtocol oprot) throws TException {
      oprot.writeStructBegin(STRUCT_DESC);
      oprot.writeFieldBegin(INT_FIELD_FIELD_DESC);
      oprot.writeI32(this.intField);
      oprot.writeFieldEnd();
      if (this.optionalIntField != null) {
        oprot.writeFieldBegin(OPTIONAL_INT_FIELD_FIELD_DESC);
        oprot.writeI32(this.optionalIntField);
        oprot.writeFieldEnd();
      }
      oprot.writeFieldBegin(INT_FIELD_WITH_DEFAULT_FIELD_DESC);
      oprot.writeI32(this.intFieldWithDefault);
      oprot.writeFieldEnd();
      if (this.setField != null) {
        oprot.writeFieldBegin(SET_FIELD_FIELD_DESC);
        Set<String> _iter0 = this.setField;
        oprot.writeSetBegin(new TSet(TType.STRING, _iter0.size()));
        for (String _iter1 : _iter0) {
          oprot.writeString(_iter1);
        }
        oprot.writeSetEnd();
        oprot.writeFieldEnd();
      }
      if (this.optionalSetField != null) {
        oprot.writeFieldBegin(OPTIONAL_SET_FIELD_FIELD_DESC);
        Set<String> _iter0 = this.optionalSetField;
        oprot.writeSetBegin(new TSet(TType.STRING, _iter0.size()));
        for (String _iter1 : _iter0) {
          oprot.writeString(_iter1);
        }
        oprot.writeSetEnd();
        oprot.writeFieldEnd();
      }
      if (this.mapField != null) {
        oprot.writeFieldBegin(MAP_FIELD_FIELD_DESC);
        Map<String, List<String>> _iter0 = this.mapField;
        oprot.writeMapBegin(new TMap(TType.STRING, TType.LIST, _iter0.size()));
        for (Map.Entry<String, List<String>> _iter1 : _iter0.entrySet()) {
          oprot.writeString(_iter1.getKey());
          
          oprot.writeListBegin(new TList(TType.STRING, _iter1.getValue().size()));
        for (String _iter2 : _iter1.getValue()) {
          oprot.writeString(_iter2);
        }
        oprot.writeListEnd();
        }
        oprot.writeMapEnd();
        oprot.writeFieldEnd();
      }
      if (this.optionalMapField != null) {
        oprot.writeFieldBegin(OPTIONAL_MAP_FIELD_FIELD_DESC);
        Map<String, List<String>> _iter0 = this.optionalMapField;
        oprot.writeMapBegin(new TMap(TType.STRING, TType.LIST, _iter0.size()));
        for (Map.Entry<String, List<String>> _iter1 : _iter0.entrySet()) {
          oprot.writeString(_iter1.getKey());
          
          oprot.writeListBegin(new TList(TType.STRING, _iter1.getValue().size()));
        for (String _iter2 : _iter1.getValue()) {
          oprot.writeString(_iter2);
        }
        oprot.writeListEnd();
        }
        oprot.writeMapEnd();
        oprot.writeFieldEnd();
      }
      if (this.binaryField != null) {
        oprot.writeFieldBegin(BINARY_FIELD_FIELD_DESC);
        oprot.writeBinary(java.nio.ByteBuffer.wrap(this.binaryField));
        oprot.writeFieldEnd();
      }
      oprot.writeFieldBegin(LONG_FIELD_FIELD_DESC);
      oprot.writeI64(this.longField);
      oprot.writeFieldEnd();
      oprot.writeFieldBegin(ADAPTED_LONG_FIELD_FIELD_DESC);
      oprot.writeI64(this.adaptedLongField);
      oprot.writeFieldEnd();
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }
    
    private static class _FooLazy {
        private static final Foo _DEFAULT = new Foo.Builder().build();
    }
    
    public static Foo defaultInstance() {
        return  _FooLazy._DEFAULT;
    }
}
