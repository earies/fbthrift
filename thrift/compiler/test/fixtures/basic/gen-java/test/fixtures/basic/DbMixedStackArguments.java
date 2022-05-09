/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.basic;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.service.*;
import com.facebook.thrift.client.*;
import com.google.common.util.concurrent.ListenableFuture;
import java.io.*;
import java.util.*;
import reactor.core.publisher.Mono;

@SwiftGenerated
@com.facebook.swift.service.ThriftService("DbMixedStackArguments")
public interface DbMixedStackArguments extends java.io.Closeable, com.facebook.thrift.util.BlockingService {
    static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final DbMixedStackArguments _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
        return new DbMixedStackArgumentsRpcServerHandler(_serviceImpl, _eventHandlers);
    }

    static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
        final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
        final com.facebook.thrift.util.TransportType _transportType,
        final DbMixedStackArguments _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

        final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

        return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
    }

    static com.facebook.thrift.client.ClientBuilder<DbMixedStackArguments> clientBuilder() {
        return new ClientBuilder<DbMixedStackArguments>() {
            @Override
            public DbMixedStackArguments build(Mono<RpcClient> rpcClientMono) {
                DbMixedStackArguments.Reactive _delegate =
                    new DbMixedStackArgumentsReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                return new DbMixedStackArgumentsReactiveBlockingWrapper(_delegate);
            }
        };
    }

    @com.facebook.swift.service.ThriftService("DbMixedStackArguments")
    public interface Async extends java.io.Closeable, com.facebook.thrift.util.AsyncService {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final DbMixedStackArguments.Async _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new DbMixedStackArgumentsRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final DbMixedStackArguments.Async _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<DbMixedStackArguments.Async> clientBuilder() {
            return new ClientBuilder<DbMixedStackArguments.Async>() {
                @Override
                public DbMixedStackArguments.Async build(Mono<RpcClient> rpcClientMono) {
                    DbMixedStackArguments.Reactive _delegate =
                        new DbMixedStackArgumentsReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                    return new DbMixedStackArgumentsReactiveAsyncWrapper(_delegate);
                }
            };
        }

        @java.lang.Override void close();

        @ThriftMethod(value = "getDataByKey0")
        ListenableFuture<byte[]> getDataByKey0(
            @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key);

        default ListenableFuture<byte[]> getDataByKey0(
            @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default ListenableFuture<ResponseWrapper<byte[]>> getDataByKey0Wrapper(
            @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        @ThriftMethod(value = "getDataByKey1")
        ListenableFuture<byte[]> getDataByKey1(
            @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key);

        default ListenableFuture<byte[]> getDataByKey1(
            @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default ListenableFuture<ResponseWrapper<byte[]>> getDataByKey1Wrapper(
            @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }
    }
    @java.lang.Override void close();

    @ThriftMethod(value = "getDataByKey0")
    byte[] getDataByKey0(
        @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key) throws org.apache.thrift.TException;

    default byte[] getDataByKey0(
        @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    default ResponseWrapper<byte[]> getDataByKey0Wrapper(
        @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    @ThriftMethod(value = "getDataByKey1")
    byte[] getDataByKey1(
        @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key) throws org.apache.thrift.TException;

    default byte[] getDataByKey1(
        @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    default ResponseWrapper<byte[]> getDataByKey1Wrapper(
        @com.facebook.swift.codec.ThriftField(value=1, name="key", requiredness=Requiredness.NONE) final String key,
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    @com.facebook.swift.service.ThriftService("DbMixedStackArguments")
    interface Reactive extends reactor.core.Disposable, com.facebook.thrift.util.ReactiveService {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final DbMixedStackArguments.Reactive _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new DbMixedStackArgumentsRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final DbMixedStackArguments.Reactive _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<DbMixedStackArguments.Reactive> clientBuilder() {
            return new ClientBuilder<DbMixedStackArguments.Reactive>() {
                @Override
                public DbMixedStackArguments.Reactive build(Mono<RpcClient> rpcClientMono) {
                    return new DbMixedStackArgumentsReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                }
            };
        }

        @ThriftMethod(value = "getDataByKey0")
        reactor.core.publisher.Mono<byte[]> getDataByKey0(final String key);

        default reactor.core.publisher.Mono<byte[]> getDataByKey0(final String key, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<byte[]>> getDataByKey0Wrapper(final String key, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        @ThriftMethod(value = "getDataByKey1")
        reactor.core.publisher.Mono<byte[]> getDataByKey1(final String key);

        default reactor.core.publisher.Mono<byte[]> getDataByKey1(final String key, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<byte[]>> getDataByKey1Wrapper(final String key, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

    }
}
