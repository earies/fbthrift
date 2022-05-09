/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.sink;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.service.*;
import com.facebook.thrift.client.*;
import com.google.common.util.concurrent.ListenableFuture;
import java.io.*;
import java.util.*;
import reactor.core.publisher.Mono;

@SwiftGenerated
@com.facebook.swift.service.ThriftService("SinkService")
public interface SinkService extends java.io.Closeable, com.facebook.thrift.util.BlockingService {
    static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final SinkService _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
        return new SinkServiceRpcServerHandler(_serviceImpl, _eventHandlers);
    }

    static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
        final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
        final com.facebook.thrift.util.TransportType _transportType,
        final SinkService _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

        final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

        return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
    }

    static com.facebook.thrift.client.ClientBuilder<SinkService> clientBuilder() {
        return new ClientBuilder<SinkService>() {
            @Override
            public SinkService build(Mono<RpcClient> rpcClientMono) {
                SinkService.Reactive _delegate =
                    new SinkServiceReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                return new SinkServiceReactiveBlockingWrapper(_delegate);
            }
        };
    }

    @com.facebook.swift.service.ThriftService("SinkService")
    public interface Async extends java.io.Closeable, com.facebook.thrift.util.AsyncService {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final SinkService.Async _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new SinkServiceRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final SinkService.Async _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<SinkService.Async> clientBuilder() {
            return new ClientBuilder<SinkService.Async>() {
                @Override
                public SinkService.Async build(Mono<RpcClient> rpcClientMono) {
                    SinkService.Reactive _delegate =
                        new SinkServiceReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                    return new SinkServiceReactiveAsyncWrapper(_delegate);
                }
            };
        }

        @java.lang.Override void close();

    }
    @java.lang.Override void close();

    @com.facebook.swift.service.ThriftService("SinkService")
    interface Reactive extends reactor.core.Disposable, com.facebook.thrift.util.ReactiveService {
        static com.facebook.thrift.server.RpcServerHandler createRpcServerHandler(
        final SinkService.Reactive _serviceImpl,
        final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {
            return new SinkServiceRpcServerHandler(_serviceImpl, _eventHandlers);
        }

        static reactor.core.publisher.Mono<? extends com.facebook.thrift.server.ServerTransport> createServer(
            final com.facebook.swift.service.ThriftServerConfig _thriftServerConfig,
            final com.facebook.thrift.util.TransportType _transportType,
            final SinkService.Reactive _serviceImpl,
            final List<com.facebook.swift.service.ThriftEventHandler> _eventHandlers) {

            final com.facebook.thrift.server.RpcServerHandler _serverHandler = createRpcServerHandler(_serviceImpl, _eventHandlers);

            return com.facebook.thrift.util.RpcServerUtils.createServerTransport(_thriftServerConfig, _transportType, _serverHandler);
        }

        static com.facebook.thrift.client.ClientBuilder<SinkService.Reactive> clientBuilder() {
            return new ClientBuilder<SinkService.Reactive>() {
                @Override
                public SinkService.Reactive build(Mono<RpcClient> rpcClientMono) {
                    return new SinkServiceReactiveClient(protocolId, rpcClientMono, headers, persistentHeaders);
                }
            };
        }

        reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> method( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> method( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.sink.FinalResponse>> methodWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        reactor.core.publisher.Flux<com.facebook.thrift.model.StreamResponse<test.fixtures.sink.InitialResponse,test.fixtures.sink.FinalResponse>> methodAndReponse( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Flux<com.facebook.thrift.model.StreamResponse<test.fixtures.sink.InitialResponse,test.fixtures.sink.FinalResponse>> methodAndReponse( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Flux<ResponseWrapper<com.facebook.thrift.model.StreamResponse<test.fixtures.sink.InitialResponse,test.fixtures.sink.FinalResponse>>> methodAndReponseWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.sink.FinalResponse>> methodThrowWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodSinkThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodSinkThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.sink.FinalResponse>> methodSinkThrowWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodFinalThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodFinalThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.sink.FinalResponse>> methodFinalThrowWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodBothThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodBothThrow( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.sink.FinalResponse>> methodBothThrowWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodFast( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads);

        default reactor.core.publisher.Mono<test.fixtures.sink.FinalResponse> methodFast( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

        default reactor.core.publisher.Mono<ResponseWrapper<test.fixtures.sink.FinalResponse>> methodFastWrapper( org.reactivestreams.Publisher<test.fixtures.sink.SinkPayload> payloads, RpcOptions rpcOptions) {
            throw new UnsupportedOperationException();
        }

    }
}
