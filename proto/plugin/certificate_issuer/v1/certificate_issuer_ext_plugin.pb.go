// Code generated by protoc-gen-go-extension. DO NOT EDIT.

package certificate_issuerv1

import (
	api "github.com/openkcm/plugin-sdk/api"
	grpc "google.golang.org/grpc"
)

const (
	Type                = "CertificateIssuerService"
	GRPCServiceFullName = "plugin.certificate_issuer.v1.CertificateIssuerService"
)

func CertificateIssuerServicePluginServer(server CertificateIssuerServiceServer) api.PluginServer {
	return certificateIssuerServicePluginServer{CertificateIssuerServiceServer: server}
}

type certificateIssuerServicePluginServer struct {
	CertificateIssuerServiceServer
}

func (s certificateIssuerServicePluginServer) Type() string {
	return Type
}

func (s certificateIssuerServicePluginServer) GRPCServiceName() string {
	return GRPCServiceFullName
}

func (s certificateIssuerServicePluginServer) RegisterServer(server *grpc.Server) any {
	RegisterCertificateIssuerServiceServer(server, s.CertificateIssuerServiceServer)
	return s.CertificateIssuerServiceServer
}

type CertificateIssuerServicePluginClient struct {
	CertificateIssuerServiceClient
}

func (s CertificateIssuerServicePluginClient) Type() string {
	return Type
}

func (c *CertificateIssuerServicePluginClient) IsInitialized() bool {
	return c.CertificateIssuerServiceClient != nil
}

func (c *CertificateIssuerServicePluginClient) GRPCServiceName() string {
	return GRPCServiceFullName
}

func (c *CertificateIssuerServicePluginClient) InitClient(conn grpc.ClientConnInterface) any {
	c.CertificateIssuerServiceClient = NewCertificateIssuerServiceClient(conn)
	return c.CertificateIssuerServiceClient
}
