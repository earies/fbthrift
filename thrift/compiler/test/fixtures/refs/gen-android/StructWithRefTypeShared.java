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
public class StructWithRefTypeShared implements TBase, java.io.Serializable, Cloneable {
  private static final TStruct STRUCT_DESC = new TStruct("StructWithRefTypeShared");
  private static final TField DEF_FIELD_FIELD_DESC = new TField("def_field", TType.STRUCT, (short)1);
  private static final TField OPT_FIELD_FIELD_DESC = new TField("opt_field", TType.STRUCT, (short)2);
  private static final TField REQ_FIELD_FIELD_DESC = new TField("req_field", TType.STRUCT, (short)3);

  public final Empty def_field;
  public final Empty opt_field;
  public final Empty req_field;
  public static final int DEF_FIELD = 1;
  public static final int OPT_FIELD = 2;
  public static final int REQ_FIELD = 3;

  public StructWithRefTypeShared(
      Empty def_field,
      Empty opt_field,
      Empty req_field) {
    this.def_field = def_field;
    this.opt_field = opt_field;
    this.req_field = req_field;
  }

  /**
   * Performs a deep copy on <i>other</i>.
   */
  public StructWithRefTypeShared(StructWithRefTypeShared other) {
    if (other.isSetDef_field()) {
      this.def_field = TBaseHelper.deepCopy(other.def_field);
    } else {
      this.def_field = null;
    }
    if (other.isSetOpt_field()) {
      this.opt_field = TBaseHelper.deepCopy(other.opt_field);
    } else {
      this.opt_field = null;
    }
    if (other.isSetReq_field()) {
      this.req_field = TBaseHelper.deepCopy(other.req_field);
    } else {
      this.req_field = null;
    }
  }

  public StructWithRefTypeShared deepCopy() {
    return new StructWithRefTypeShared(this);
  }

  public Empty getDef_field() {
    return this.def_field;
  }

  // Returns true if field def_field is set (has been assigned a value) and false otherwise
  public boolean isSetDef_field() {
    return this.def_field != null;
  }

  public Empty getOpt_field() {
    return this.opt_field;
  }

  // Returns true if field opt_field is set (has been assigned a value) and false otherwise
  public boolean isSetOpt_field() {
    return this.opt_field != null;
  }

  public Empty getReq_field() {
    return this.req_field;
  }

  // Returns true if field req_field is set (has been assigned a value) and false otherwise
  public boolean isSetReq_field() {
    return this.req_field != null;
  }

  @Override
  public boolean equals(Object _that) {
    if (_that == null)
      return false;
    if (this == _that)
      return true;
    if (!(_that instanceof StructWithRefTypeShared))
      return false;
    StructWithRefTypeShared that = (StructWithRefTypeShared)_that;

    if (!TBaseHelper.equalsNobinary(this.isSetDef_field(), that.isSetDef_field(), this.def_field, that.def_field)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetOpt_field(), that.isSetOpt_field(), this.opt_field, that.opt_field)) { return false; }

    if (!TBaseHelper.equalsNobinary(this.isSetReq_field(), that.isSetReq_field(), this.req_field, that.req_field)) { return false; }

    return true;
  }

  @Override
  public int hashCode() {
    return Arrays.deepHashCode(new Object[] {def_field, opt_field, req_field});
  }

  // This is required to satisfy the TBase interface, but can't be implemented on immutable struture.
  public void read(TProtocol iprot) throws TException {
    throw new TException("unimplemented in android immutable structure");
  }

  public static StructWithRefTypeShared deserialize(TProtocol iprot) throws TException {
    Empty tmp_def_field = null;
    Empty tmp_opt_field = null;
    Empty tmp_req_field = null;
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
        case DEF_FIELD:
          if (__field.type == TType.STRUCT) {
            tmp_def_field = Empty.deserialize(iprot);
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case OPT_FIELD:
          if (__field.type == TType.STRUCT) {
            tmp_opt_field = Empty.deserialize(iprot);
          } else {
            TProtocolUtil.skip(iprot, __field.type);
          }
          break;
        case REQ_FIELD:
          if (__field.type == TType.STRUCT) {
            tmp_req_field = Empty.deserialize(iprot);
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

    StructWithRefTypeShared _that;
    _that = new StructWithRefTypeShared(
      tmp_def_field
      ,tmp_opt_field
      ,tmp_req_field
    );
    _that.validate();
    return _that;
  }

  public void write(TProtocol oprot) throws TException {
    validate();

    oprot.writeStructBegin(STRUCT_DESC);
    if (this.def_field != null) {
      oprot.writeFieldBegin(DEF_FIELD_FIELD_DESC);
      this.def_field.write(oprot);
      oprot.writeFieldEnd();
    }
    if (this.opt_field != null) {
      if (isSetOpt_field()) {
        oprot.writeFieldBegin(OPT_FIELD_FIELD_DESC);
        this.opt_field.write(oprot);
        oprot.writeFieldEnd();
      }
    }
    if (this.req_field != null) {
      oprot.writeFieldBegin(REQ_FIELD_FIELD_DESC);
      this.req_field.write(oprot);
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
    if (req_field == null) {
      throw new TProtocolException(TProtocolException.MISSING_REQUIRED_FIELD, "Required field 'req_field' was not present! Struct: " + toString());
    }
  }

}

