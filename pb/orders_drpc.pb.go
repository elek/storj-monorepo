// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.31
// source: orders.proto

package pb

import (
	bytes "bytes"
	context "context"
	errors "errors"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_orders_proto struct{}

func (drpcEncoding_File_orders_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_orders_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_orders_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_orders_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCOrdersClient interface {
	DRPCConn() drpc.Conn

	SettlementWithWindow(ctx context.Context) (DRPCOrders_SettlementWithWindowClient, error)
}

type drpcOrdersClient struct {
	cc drpc.Conn
}

func NewDRPCOrdersClient(cc drpc.Conn) DRPCOrdersClient {
	return &drpcOrdersClient{cc}
}

func (c *drpcOrdersClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcOrdersClient) SettlementWithWindow(ctx context.Context) (DRPCOrders_SettlementWithWindowClient, error) {
	stream, err := c.cc.NewStream(ctx, "/orders.Orders/SettlementWithWindow", drpcEncoding_File_orders_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcOrders_SettlementWithWindowClient{stream}
	return x, nil
}

type DRPCOrders_SettlementWithWindowClient interface {
	drpc.Stream
	Send(*SettlementRequest) error
	CloseAndRecv() (*SettlementWithWindowResponse, error)
}

type drpcOrders_SettlementWithWindowClient struct {
	drpc.Stream
}

func (x *drpcOrders_SettlementWithWindowClient) Send(m *SettlementRequest) error {
	return x.MsgSend(m, drpcEncoding_File_orders_proto{})
}

func (x *drpcOrders_SettlementWithWindowClient) CloseAndRecv() (*SettlementWithWindowResponse, error) {
	if err := x.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SettlementWithWindowResponse)
	if err := x.MsgRecv(m, drpcEncoding_File_orders_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcOrders_SettlementWithWindowClient) CloseAndRecvMsg(m *SettlementWithWindowResponse) error {
	if err := x.CloseSend(); err != nil {
		return err
	}
	return x.MsgRecv(m, drpcEncoding_File_orders_proto{})
}

type DRPCOrdersServer interface {
	SettlementWithWindow(DRPCOrders_SettlementWithWindowStream) error
}

type DRPCOrdersUnimplementedServer struct{}

func (s *DRPCOrdersUnimplementedServer) SettlementWithWindow(DRPCOrders_SettlementWithWindowStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCOrdersDescription struct{}

func (DRPCOrdersDescription) NumMethods() int { return 1 }

func (DRPCOrdersDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/orders.Orders/SettlementWithWindow", drpcEncoding_File_orders_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCOrdersServer).
					SettlementWithWindow(
						&drpcOrders_SettlementWithWindowStream{in1.(drpc.Stream)},
					)
			}, DRPCOrdersServer.SettlementWithWindow, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterOrders(mux drpc.Mux, impl DRPCOrdersServer) error {
	return mux.Register(impl, DRPCOrdersDescription{})
}

type DRPCOrders_SettlementWithWindowStream interface {
	drpc.Stream
	SendAndClose(*SettlementWithWindowResponse) error
	Recv() (*SettlementRequest, error)
}

type drpcOrders_SettlementWithWindowStream struct {
	drpc.Stream
}

func (x *drpcOrders_SettlementWithWindowStream) SendAndClose(m *SettlementWithWindowResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_orders_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

func (x *drpcOrders_SettlementWithWindowStream) Recv() (*SettlementRequest, error) {
	m := new(SettlementRequest)
	if err := x.MsgRecv(m, drpcEncoding_File_orders_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcOrders_SettlementWithWindowStream) RecvMsg(m *SettlementRequest) error {
	return x.MsgRecv(m, drpcEncoding_File_orders_proto{})
}
