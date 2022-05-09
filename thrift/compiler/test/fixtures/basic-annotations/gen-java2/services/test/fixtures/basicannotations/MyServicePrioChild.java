/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.basicannotations;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.service.*;
import com.facebook.thrift.client.*;
import com.google.common.util.concurrent.ListenableFuture;
import java.io.*;
import java.util.*;
import reactor.core.publisher.Mono;

@SwiftGenerated
@com.facebook.swift.service.ThriftService("MyServicePrioChild")
public interface MyServicePrioChild extends java.io.Closeable, com.facebook.thrift.util.BlockingService, test.fixtures.basicannotations.MyServicePrioParent {
    static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final MyServicePrioChild _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
        return new MyServicePrioChildRpcServerHandler(_serviceImpl, _eventHandlers);
    }

    static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
        final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
        final com.facebook.thrift.util.TransportType _transportType,
        final MyServicePrioChild _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

        final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

        return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
    }

    static com.facebook.thrift.client.ClientBuilder<MyServicePrioChild> clientBuilder() {
        return new ClientBuilder<MyServicePrioChild>() {
            @Override
            public MyServicePrioChild build(Mono<RpcClient> rpcClientMono) {
                MyServicePrioChild.Reactive _delegate =
                    new MyServicePrioChildReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                return new MyServicePrioChildReactiveBlockingWrapper(_delegate);
            }
        };
    }

    @com.facebook.swift.service.ThriftService("MyServicePrioChild")
    public interface Async extends java.io.Closeable, com.facebook.thrift.util.AsyncService, test.fixtures.basicannotations.MyServicePrioParent.Async {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final MyServicePrioChild.Async _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new MyServicePrioChildRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final MyServicePrioChild.Async _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<MyServicePrioChild.Async> clientBuilder() {
            return new ClientBuilder<MyServicePrioChild.Async>() {
                @Override
                public MyServicePrioChild.Async build(Mono<RpcClient> rpcClientMono) {
                    MyServicePrioChild.Reactive _delegate =
                        new MyServicePrioChildReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                    return new MyServicePrioChildReactiveAsyncWrapper(_delegate);
                }
            };
        }

        @java.lang.Override void close();

        @ThriftMethod(value = "pang")
        ListenableFuture<Void> pang();

        default ListenableFuture<Void> pang(
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default ListenableFuture<ResponseWrapper<Void>> pangWrapper(
            RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }
    }
    @java.lang.Override void close();

    @ThriftMethod(value = "pang")
    void pang() throws org.apache.thrift.TException;

    default void pang(
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    default ResponseWrapper<Void> pangWrapper(
        RpcOptions rpcOptions) throws org.apache.thrift.TException {
        throw new UnsupportedOperationException();
    }

    @com.facebook.swift.service.ThriftService("MyServicePrioChild")
    interface Reactive extends reactor.core.Disposable, com.facebook.thrift.util.ReactiveService, test.fixtures.basicannotations.MyServicePrioParent.Reactive {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final MyServicePrioChild.Reactive _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new MyServicePrioChildRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final MyServicePrioChild.Reactive _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<MyServicePrioChild.Reactive> clientBuilder() {
            return new ClientBuilder<MyServicePrioChild.Reactive>() {
                @Override
                public MyServicePrioChild.Reactive build(Mono<RpcClient> rpcClientMono) {
                    return new MyServicePrioChildReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                }
            };
        }

        @ThriftMethod(value = "pang")
        reactor.core.publisher.Mono<Void> pang();

        default reactor.core.publisher.Mono<Void> pang(RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<Void>> pangWrapper(RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

    }
}
