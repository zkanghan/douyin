// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageserivce

import (
	"context"
	message "douyin/kitex_gen/message"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Send(ctx context.Context, req *message.ActionRequest, callOptions ...callopt.Option) (r *message.ActionResponse, err error)
	Record(ctx context.Context, req *message.RecordRequest, callOptions ...callopt.Option) (r *message.RecordResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageSerivceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageSerivceClient struct {
	*kClient
}

func (p *kMessageSerivceClient) Send(ctx context.Context, req *message.ActionRequest, callOptions ...callopt.Option) (r *message.ActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Send(ctx, req)
}

func (p *kMessageSerivceClient) Record(ctx context.Context, req *message.RecordRequest, callOptions ...callopt.Option) (r *message.RecordResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Record(ctx, req)
}
