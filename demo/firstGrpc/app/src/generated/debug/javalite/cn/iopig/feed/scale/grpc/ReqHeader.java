// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: global_def.proto

package cn.iopig.feed.scale.grpc;

/**
 * Protobuf type {@code fsapi.ReqHeader}
 */
public  final class ReqHeader extends
    com.google.protobuf.GeneratedMessageLite<
        ReqHeader, ReqHeader.Builder> implements
    // @@protoc_insertion_point(message_implements:fsapi.ReqHeader)
    ReqHeaderOrBuilder {
  private ReqHeader() {
    devId_ = "";
  }
  public static final int VERSION_FIELD_NUMBER = 1;
  private int version_;
  /**
   * <pre>
   * 现在版本默认为 1
   * </pre>
   *
   * <code>optional int32 Version = 1;</code>
   */
  public int getVersion() {
    return version_;
  }
  /**
   * <pre>
   * 现在版本默认为 1
   * </pre>
   *
   * <code>optional int32 Version = 1;</code>
   */
  private void setVersion(int value) {
    
    version_ = value;
  }
  /**
   * <pre>
   * 现在版本默认为 1
   * </pre>
   *
   * <code>optional int32 Version = 1;</code>
   */
  private void clearVersion() {
    
    version_ = 0;
  }

  public static final int DEVID_FIELD_NUMBER = 2;
  private java.lang.String devId_;
  /**
   * <pre>
   * 设备ID
   * </pre>
   *
   * <code>optional string DevId = 2;</code>
   */
  public java.lang.String getDevId() {
    return devId_;
  }
  /**
   * <pre>
   * 设备ID
   * </pre>
   *
   * <code>optional string DevId = 2;</code>
   */
  public com.google.protobuf.ByteString
      getDevIdBytes() {
    return com.google.protobuf.ByteString.copyFromUtf8(devId_);
  }
  /**
   * <pre>
   * 设备ID
   * </pre>
   *
   * <code>optional string DevId = 2;</code>
   */
  private void setDevId(
      java.lang.String value) {
    if (value == null) {
    throw new NullPointerException();
  }
  
    devId_ = value;
  }
  /**
   * <pre>
   * 设备ID
   * </pre>
   *
   * <code>optional string DevId = 2;</code>
   */
  private void clearDevId() {
    
    devId_ = getDefaultInstance().getDevId();
  }
  /**
   * <pre>
   * 设备ID
   * </pre>
   *
   * <code>optional string DevId = 2;</code>
   */
  private void setDevIdBytes(
      com.google.protobuf.ByteString value) {
    if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
    
    devId_ = value.toStringUtf8();
  }

  public static final int TS_FIELD_NUMBER = 3;
  private int ts_;
  /**
   * <pre>
   *时间戳 unix_timestamp
   * </pre>
   *
   * <code>optional uint32 Ts = 3;</code>
   */
  public int getTs() {
    return ts_;
  }
  /**
   * <pre>
   *时间戳 unix_timestamp
   * </pre>
   *
   * <code>optional uint32 Ts = 3;</code>
   */
  private void setTs(int value) {
    
    ts_ = value;
  }
  /**
   * <pre>
   *时间戳 unix_timestamp
   * </pre>
   *
   * <code>optional uint32 Ts = 3;</code>
   */
  private void clearTs() {
    
    ts_ = 0;
  }

  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (version_ != 0) {
      output.writeInt32(1, version_);
    }
    if (!devId_.isEmpty()) {
      output.writeString(2, getDevId());
    }
    if (ts_ != 0) {
      output.writeUInt32(3, ts_);
    }
  }

  public int getSerializedSize() {
    int size = memoizedSerializedSize;
    if (size != -1) return size;

    size = 0;
    if (version_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, version_);
    }
    if (!devId_.isEmpty()) {
      size += com.google.protobuf.CodedOutputStream
        .computeStringSize(2, getDevId());
    }
    if (ts_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt32Size(3, ts_);
    }
    memoizedSerializedSize = size;
    return size;
  }

  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, data, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return parseDelimitedFrom(DEFAULT_INSTANCE, input, extensionRegistry);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input);
  }
  public static cn.iopig.feed.scale.grpc.ReqHeader parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageLite.parseFrom(
        DEFAULT_INSTANCE, input, extensionRegistry);
  }

  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(cn.iopig.feed.scale.grpc.ReqHeader prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }

  /**
   * Protobuf type {@code fsapi.ReqHeader}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageLite.Builder<
        cn.iopig.feed.scale.grpc.ReqHeader, Builder> implements
      // @@protoc_insertion_point(builder_implements:fsapi.ReqHeader)
      cn.iopig.feed.scale.grpc.ReqHeaderOrBuilder {
    // Construct using cn.iopig.feed.scale.grpc.ReqHeader.newBuilder()
    private Builder() {
      super(DEFAULT_INSTANCE);
    }


    /**
     * <pre>
     * 现在版本默认为 1
     * </pre>
     *
     * <code>optional int32 Version = 1;</code>
     */
    public int getVersion() {
      return instance.getVersion();
    }
    /**
     * <pre>
     * 现在版本默认为 1
     * </pre>
     *
     * <code>optional int32 Version = 1;</code>
     */
    public Builder setVersion(int value) {
      copyOnWrite();
      instance.setVersion(value);
      return this;
    }
    /**
     * <pre>
     * 现在版本默认为 1
     * </pre>
     *
     * <code>optional int32 Version = 1;</code>
     */
    public Builder clearVersion() {
      copyOnWrite();
      instance.clearVersion();
      return this;
    }

    /**
     * <pre>
     * 设备ID
     * </pre>
     *
     * <code>optional string DevId = 2;</code>
     */
    public java.lang.String getDevId() {
      return instance.getDevId();
    }
    /**
     * <pre>
     * 设备ID
     * </pre>
     *
     * <code>optional string DevId = 2;</code>
     */
    public com.google.protobuf.ByteString
        getDevIdBytes() {
      return instance.getDevIdBytes();
    }
    /**
     * <pre>
     * 设备ID
     * </pre>
     *
     * <code>optional string DevId = 2;</code>
     */
    public Builder setDevId(
        java.lang.String value) {
      copyOnWrite();
      instance.setDevId(value);
      return this;
    }
    /**
     * <pre>
     * 设备ID
     * </pre>
     *
     * <code>optional string DevId = 2;</code>
     */
    public Builder clearDevId() {
      copyOnWrite();
      instance.clearDevId();
      return this;
    }
    /**
     * <pre>
     * 设备ID
     * </pre>
     *
     * <code>optional string DevId = 2;</code>
     */
    public Builder setDevIdBytes(
        com.google.protobuf.ByteString value) {
      copyOnWrite();
      instance.setDevIdBytes(value);
      return this;
    }

    /**
     * <pre>
     *时间戳 unix_timestamp
     * </pre>
     *
     * <code>optional uint32 Ts = 3;</code>
     */
    public int getTs() {
      return instance.getTs();
    }
    /**
     * <pre>
     *时间戳 unix_timestamp
     * </pre>
     *
     * <code>optional uint32 Ts = 3;</code>
     */
    public Builder setTs(int value) {
      copyOnWrite();
      instance.setTs(value);
      return this;
    }
    /**
     * <pre>
     *时间戳 unix_timestamp
     * </pre>
     *
     * <code>optional uint32 Ts = 3;</code>
     */
    public Builder clearTs() {
      copyOnWrite();
      instance.clearTs();
      return this;
    }

    // @@protoc_insertion_point(builder_scope:fsapi.ReqHeader)
  }
  protected final Object dynamicMethod(
      com.google.protobuf.GeneratedMessageLite.MethodToInvoke method,
      Object arg0, Object arg1) {
    switch (method) {
      case NEW_MUTABLE_INSTANCE: {
        return new cn.iopig.feed.scale.grpc.ReqHeader();
      }
      case IS_INITIALIZED: {
        return DEFAULT_INSTANCE;
      }
      case MAKE_IMMUTABLE: {
        return null;
      }
      case NEW_BUILDER: {
        return new Builder();
      }
      case VISIT: {
        Visitor visitor = (Visitor) arg0;
        cn.iopig.feed.scale.grpc.ReqHeader other = (cn.iopig.feed.scale.grpc.ReqHeader) arg1;
        version_ = visitor.visitInt(version_ != 0, version_,
            other.version_ != 0, other.version_);
        devId_ = visitor.visitString(!devId_.isEmpty(), devId_,
            !other.devId_.isEmpty(), other.devId_);
        ts_ = visitor.visitInt(ts_ != 0, ts_,
            other.ts_ != 0, other.ts_);
        if (visitor == com.google.protobuf.GeneratedMessageLite.MergeFromVisitor
            .INSTANCE) {
        }
        return this;
      }
      case MERGE_FROM_STREAM: {
        com.google.protobuf.CodedInputStream input =
            (com.google.protobuf.CodedInputStream) arg0;
        com.google.protobuf.ExtensionRegistryLite extensionRegistry =
            (com.google.protobuf.ExtensionRegistryLite) arg1;
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              default: {
                if (!input.skipField(tag)) {
                  done = true;
                }
                break;
              }
              case 8: {

                version_ = input.readInt32();
                break;
              }
              case 18: {
                String s = input.readStringRequireUtf8();

                devId_ = s;
                break;
              }
              case 24: {

                ts_ = input.readUInt32();
                break;
              }
            }
          }
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw new RuntimeException(e.setUnfinishedMessage(this));
        } catch (java.io.IOException e) {
          throw new RuntimeException(
              new com.google.protobuf.InvalidProtocolBufferException(
                  e.getMessage()).setUnfinishedMessage(this));
        } finally {
        }
      }
      case GET_DEFAULT_INSTANCE: {
        return DEFAULT_INSTANCE;
      }
      case GET_PARSER: {
        if (PARSER == null) {    synchronized (cn.iopig.feed.scale.grpc.ReqHeader.class) {
            if (PARSER == null) {
              PARSER = new DefaultInstanceBasedParser(DEFAULT_INSTANCE);
            }
          }
        }
        return PARSER;
      }
    }
    throw new UnsupportedOperationException();
  }


  // @@protoc_insertion_point(class_scope:fsapi.ReqHeader)
  private static final cn.iopig.feed.scale.grpc.ReqHeader DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new ReqHeader();
    DEFAULT_INSTANCE.makeImmutable();
  }

  public static cn.iopig.feed.scale.grpc.ReqHeader getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static volatile com.google.protobuf.Parser<ReqHeader> PARSER;

  public static com.google.protobuf.Parser<ReqHeader> parser() {
    return DEFAULT_INSTANCE.getParserForType();
  }
}

