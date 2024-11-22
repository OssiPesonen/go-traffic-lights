// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "traffic_lights_service.proto" (package "proto", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { TrafficLightsService } from "./traffic_lights_service";
import type { RefreshTokenRequest } from "./traffic_lights_service";
import type { RegistrationRequest } from "./traffic_lights_service";
import type { TokenResponse } from "./traffic_lights_service";
import type { LoginRequest } from "./traffic_lights_service";
import type { CreateProviderResponse } from "./traffic_lights_service";
import type { CreateProviderRequest } from "./traffic_lights_service";
import type { ReadProviderResponse } from "./traffic_lights_service";
import type { ReadProviderRequest } from "./traffic_lights_service";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ListProviderResponse } from "./traffic_lights_service";
import type { Empty } from "./google/protobuf/empty";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service proto.TrafficLightsService
 */
export interface ITrafficLightsServiceClient {
    /**
     * Providers
     *
     * @generated from protobuf rpc: ListProviders(google.protobuf.Empty) returns (proto.ListProviderResponse);
     */
    listProviders(input: Empty, options?: RpcOptions): UnaryCall<Empty, ListProviderResponse>;
    /**
     * @generated from protobuf rpc: ReadProvider(proto.ReadProviderRequest) returns (proto.ReadProviderResponse);
     */
    readProvider(input: ReadProviderRequest, options?: RpcOptions): UnaryCall<ReadProviderRequest, ReadProviderResponse>;
    /**
     * @generated from protobuf rpc: CreateProvider(proto.CreateProviderRequest) returns (proto.CreateProviderResponse);
     */
    createProvider(input: CreateProviderRequest, options?: RpcOptions): UnaryCall<CreateProviderRequest, CreateProviderResponse>;
    /**
     * Users
     *
     * @generated from protobuf rpc: GetToken(proto.LoginRequest) returns (proto.TokenResponse);
     */
    getToken(input: LoginRequest, options?: RpcOptions): UnaryCall<LoginRequest, TokenResponse>;
    /**
     * @generated from protobuf rpc: RegisterUser(proto.RegistrationRequest) returns (google.protobuf.Empty);
     */
    registerUser(input: RegistrationRequest, options?: RpcOptions): UnaryCall<RegistrationRequest, Empty>;
    /**
     * @generated from protobuf rpc: RefreshToken(proto.RefreshTokenRequest) returns (proto.TokenResponse);
     */
    refreshToken(input: RefreshTokenRequest, options?: RpcOptions): UnaryCall<RefreshTokenRequest, TokenResponse>;
    /**
     * @generated from protobuf rpc: RevokeRefreshToken(proto.RefreshTokenRequest) returns (google.protobuf.Empty);
     */
    revokeRefreshToken(input: RefreshTokenRequest, options?: RpcOptions): UnaryCall<RefreshTokenRequest, Empty>;
    /**
     * @generated from protobuf rpc: RevokeAllRefreshTokens(google.protobuf.Empty) returns (google.protobuf.Empty);
     */
    revokeAllRefreshTokens(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty>;
}
/**
 * @generated from protobuf service proto.TrafficLightsService
 */
export class TrafficLightsServiceClient implements ITrafficLightsServiceClient, ServiceInfo {
    typeName = TrafficLightsService.typeName;
    methods = TrafficLightsService.methods;
    options = TrafficLightsService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * Providers
     *
     * @generated from protobuf rpc: ListProviders(google.protobuf.Empty) returns (proto.ListProviderResponse);
     */
    listProviders(input: Empty, options?: RpcOptions): UnaryCall<Empty, ListProviderResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<Empty, ListProviderResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ReadProvider(proto.ReadProviderRequest) returns (proto.ReadProviderResponse);
     */
    readProvider(input: ReadProviderRequest, options?: RpcOptions): UnaryCall<ReadProviderRequest, ReadProviderResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<ReadProviderRequest, ReadProviderResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: CreateProvider(proto.CreateProviderRequest) returns (proto.CreateProviderResponse);
     */
    createProvider(input: CreateProviderRequest, options?: RpcOptions): UnaryCall<CreateProviderRequest, CreateProviderResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateProviderRequest, CreateProviderResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Users
     *
     * @generated from protobuf rpc: GetToken(proto.LoginRequest) returns (proto.TokenResponse);
     */
    getToken(input: LoginRequest, options?: RpcOptions): UnaryCall<LoginRequest, TokenResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<LoginRequest, TokenResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: RegisterUser(proto.RegistrationRequest) returns (google.protobuf.Empty);
     */
    registerUser(input: RegistrationRequest, options?: RpcOptions): UnaryCall<RegistrationRequest, Empty> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<RegistrationRequest, Empty>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: RefreshToken(proto.RefreshTokenRequest) returns (proto.TokenResponse);
     */
    refreshToken(input: RefreshTokenRequest, options?: RpcOptions): UnaryCall<RefreshTokenRequest, TokenResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<RefreshTokenRequest, TokenResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: RevokeRefreshToken(proto.RefreshTokenRequest) returns (google.protobuf.Empty);
     */
    revokeRefreshToken(input: RefreshTokenRequest, options?: RpcOptions): UnaryCall<RefreshTokenRequest, Empty> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<RefreshTokenRequest, Empty>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: RevokeAllRefreshTokens(google.protobuf.Empty) returns (google.protobuf.Empty);
     */
    revokeAllRefreshTokens(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<Empty, Empty>("unary", this._transport, method, opt, input);
    }
}
