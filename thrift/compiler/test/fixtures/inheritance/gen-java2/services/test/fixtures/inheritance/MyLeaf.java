/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.inheritance;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.service.*;
import com.facebook.thrift.client.*;
import com.google.common.util.concurrent.ListenableFuture;
import java.io.*;
import java.util.*;
import reactor.core.publisher.Mono;

@SwiftGenerated
@com.facebook.swift.service.ThriftService("MyLeaf")
public interface MyLeaf extends java.io.Closeable, com.facebook.thrift.util.BlockingService, test.fixtures.inheritance.MyNode {
    static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final MyLeaf _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
        return new MyLeafRpcServerHandler(_serviceImpl, _eventHandlers);
    }

    static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
        final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
        final com.facebook.thrift.util.TransportType _transportType,
        final MyLeaf _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

        final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

        return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
    }

    static com.facebook.thrift.client.ClientBuilder<MyLeaf> clientBuilder() {
        return new ClientBuilder<MyLeaf>() {
            @Override
            public MyLeaf build(Mono<RpcClient> rpcClientMono) {
                MyLeaf.Reactive _delegate =
                    new MyLeafReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                return new MyLeafReactiveBlockingWrapper(_delegate);
            }
        };
    }

    @com.facebook.swift.service.ThriftService("MyLeaf")
    public interface Async extends java.io.Closeable, com.facebook.thrift.util.AsyncService, test.fixtures.inheritance.MyNode.Async {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final MyLeaf.Async _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new MyLeafRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final MyLeaf.Async _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<MyLeaf.Async> clientBuilder() {
            return new ClientBuilder<MyLeaf.Async>() {
                @Override
                public MyLeaf.Async build(Mono<RpcClient> rpcClientMono) {
                    MyLeaf.Reactive _delegate =
                        new MyLeafReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                    return new MyLeafReactiveAsyncWrapper(_delegate);
                }
            };
        }

        @java.lang.Override void close();

        @ThriftMethod(value = "do_leaf")
        ListenableFuture<Void> doLeaf();

        default ListenableFuture<Void> doLeaf(
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default ListenableFuture<ResponseWrapper<Void>> doLeafWrapper(
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }
    }
    @java.lang.Override void close();

    @ThriftMethod(value = "do_leaf")
    void doLeaf() throws org.apache.thrift.TException;

    default void doLeaf(
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    default ResponseWrapper<Void> doLeafWrapper(
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    @com.facebook.swift.service.ThriftService("MyLeaf")
    interface Reactive extends reactor.core.Disposable, com.facebook.thrift.util.ReactiveService, test.fixtures.inheritance.MyNode.Reactive {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final MyLeaf.Reactive _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new MyLeafRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final MyLeaf.Reactive _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<MyLeaf.Reactive> clientBuilder() {
            return new ClientBuilder<MyLeaf.Reactive>() {
                @Override
                public MyLeaf.Reactive build(Mono<RpcClient> rpcClientMono) {
                    return new MyLeafReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                }
            };
        }

        @ThriftMethod(value = "do_leaf")
        reactor.core.publisher.Mono<Void> doLeaf();

        default reactor.core.publisher.Mono<Void> doLeaf(RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<Void>> doLeafWrapper(RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

    }
}
