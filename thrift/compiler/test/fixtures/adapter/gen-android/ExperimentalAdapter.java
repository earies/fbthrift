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
public class ExperimentalAdapter implements TBase, java.io.Serializable, Cloneable {
  private static final TStruct STRUCT_DESC = new TStruct("ExperimentalAdapter");
  private static final TField NAME_FIELD_DESC = new TField("name", TType.STRING, (short)1);

  public final String name;
  public static final int NAME = 1;

  public ExperimentalAdapter(
      String name) {
    this.name = name;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public ExperimentalAdapter(ExperimentalAdapter other) {
    if (other.isSetName()) {
      this.name = TBaseHelper.deepCopy(other.name);
    } else {
      this.name = null;
    }
  }

  public ExperimentalAdapter deepCopy() {
    return new ExperimentalAdapter(this);
  }

  public String getName() {
    return this.name;
  }

  // Returns true if field name is set (has been assigned a value) and false otherwise
  public boolean isSetName() {
    return this.name != null;
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof ExperimentalAdapter))
      return false;
    ExperimentalAdapter that = (ExperimentalAdapter)_that;

    if (!TBaseHelper.equalsNobinary(this.isSetName(), that.isSetName(), this.name, that.name)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {name});
  }

  // This is required to satisfy the TBase interface, but can't be implemented on immutable struture.
  public void read(TProtocol iprot) throws TException {
    throw new TException("unimplemented in android immutable structure");
  }

  public static ExperimentalAdapter deserialize(TProtocol iprot) throws TException {
    String tmp_name = null;
    TField __field;
    iprot.readStructBegin();
    while (true)
    {
      __field = iprot.readFieldBegin();
      if (__field.type == TType.STOP) {
        break;
      }
      switch (__field.id)
      {
        case NAME:
          if (__field.type == TType.STRING) {
            tmp_name = iprot.readString();
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

    ExperimentalAdapter _that;
    _that = new ExperimentalAdapter(
      tmp_name
    );
    _that.validate();
    return _that;
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    if (this.name != null) {
      oprot.writeFieldBegin(NAME_FIELD_DESC);
      oprot.writeString(this.name);
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
    return TBaseHelper.toStringHelper(this, indent, prettyPrint);
  }

  public void validate() throws TException {
    // check for required fields
  }

}

