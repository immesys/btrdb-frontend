package main

import (
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"

	btrdb "gopkg.in/btrdb.v4"

	"github.com/SoftwareDefinedBuildings/btrdb/version"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	logging "github.com/op/go-logging"
	"github.com/pborman/uuid"
	"google.golang.org/grpc"
	pb "gopkg.in/btrdb.v4/grpcinterface"
)

const MajorVersion = version.Major
const MinorVersion = version.Minor

var VersionString = version.VersionString

type apiProvider struct {
	s          *grpc.Server
	downstream *btrdb.BTrDB
}

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("log")
}

//go:generate ./genswag.py
//go:generate go-bindata -pkg main swag/...
func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "swag",
	})
	prefix := "/swag/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

const (
	P_NULL   = iota
	P_READ   = iota
	P_MODIFY = iota
	P_CREATE = iota
	P_ADMIN  = iota
)

type GRPCInterface interface {
	InitiateShutdown() chan struct{}
}

func (a *apiProvider) writeEndpoint(ctx context.Context, uu uuid.UUID) (*btrdb.Endpoint, error) {
	return a.downstream.EndpointFor(ctx, uu)
}
func (a *apiProvider) readEndpoint(ctx context.Context, uu uuid.UUID) (*btrdb.Endpoint, error) {
	return a.downstream.ReadEndpointFor(ctx, uu)
}
func (a *apiProvider) checkPermissionsByUUID(ctx context.Context, uu uuid.UUID, perms int) error {
	return nil
}
func (a *apiProvider) checkPermissionsByCollection(ctx context.Context, collection string, perms int) error {
	return nil
}
func ProxyGRPC(laddr string) GRPCInterface {
	go func() {
		fmt.Println("==== PROFILING ENABLED ==========")
		err := http.ListenAndServe("0.0.0.0:6060", nil)
		panic(err)
	}()

	l, err := net.Listen("tcp", laddr)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	downstream, err := btrdb.Connect(ctx, btrdb.EndpointsFromEnv()...)
	cancel()
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	api := &apiProvider{s: grpcServer,
		downstream: downstream,
	}
	pb.RegisterBTrDBServer(grpcServer, api)
	go grpcServer.Serve(l)
	return api
}

func main() {
	ProxyGRPC("0.0.0.0:4410")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/v4/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(SwaggerJSON))
	})

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterBTrDBHandlerFromEndpoint(ctx, gwmux, "127.0.0.1:4410", opts)
	if err != nil {
		panic(err)
	}
	mux.Handle("/", gwmux)
	serveSwagger(mux)
	err = http.ListenAndServe(":9000", mux)
	panic(err)

}

func (a *apiProvider) InitiateShutdown() chan struct{} {
	done := make(chan struct{})
	go func() {
		a.s.GracefulStop()
		close(done)
	}()
	return done
}

func (a *apiProvider) RawValues(p *pb.RawValuesParams, r pb.BTrDB_RawValuesServer) error {
	ctx := r.Context()
	uu := p.Uuid
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_READ)
	if err != nil {
		return err
	}
	ds, err := a.readEndpoint(ctx, uu)
	if err != nil {
		return err
	}
	client, err := ds.GetGRPC().RawValues(ctx, p)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = r.Send(resp)
		if err != nil {
			return err
		}
	}
}
func (a *apiProvider) AlignedWindows(p *pb.AlignedWindowsParams, r pb.BTrDB_AlignedWindowsServer) error {
	ctx := r.Context()
	uu := p.Uuid
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_READ)
	if err != nil {
		return err
	}
	ds, err := a.readEndpoint(ctx, uu)
	if err != nil {
		return err
	}
	client, err := ds.GetGRPC().AlignedWindows(ctx, p)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = r.Send(resp)
		if err != nil {
			return err
		}
	}
}
func (a *apiProvider) Windows(p *pb.WindowsParams, r pb.BTrDB_WindowsServer) error {
	ctx := r.Context()
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_READ)
	if err != nil {
		return err
	}
	ds, err := a.readEndpoint(ctx, p.GetUuid())
	if err != nil {
		return err
	}
	client, err := ds.GetGRPC().Windows(ctx, p)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = r.Send(resp)
		if err != nil {
			return err
		}
	}
}
func (a *apiProvider) StreamInfo(ctx context.Context, p *pb.StreamInfoParams) (*pb.StreamInfoResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_READ)
	if err != nil {
		return nil, err
	}
	ds, err := a.readEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	return ds.GetGRPC().StreamInfo(ctx, p)
}

func (a *apiProvider) SetStreamAnnotations(ctx context.Context, p *pb.SetStreamAnnotationsParams) (*pb.SetStreamAnnotationsResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_MODIFY)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	return ds.GetGRPC().SetStreamAnnotations(ctx, p)
}
func (a *apiProvider) Changes(p *pb.ChangesParams, r pb.BTrDB_ChangesServer) error {
	ctx := r.Context()
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_READ)
	if err != nil {
		return err
	}
	ds, err := a.readEndpoint(ctx, p.GetUuid())
	if err != nil {
		return err
	}
	client, err := ds.GetGRPC().Changes(ctx, p)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = r.Send(resp)
		if err != nil {
			return err
		}
	}
}
func (a *apiProvider) Create(ctx context.Context, p *pb.CreateParams) (*pb.CreateResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_CREATE)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	return ds.GetGRPC().Create(ctx, p)
}
func (a *apiProvider) ListCollections(ctx context.Context, p *pb.ListCollectionsParams) (*pb.ListCollectionsResponse, error) {
	ds, err := a.readEndpoint(ctx, uuid.NewRandom())
	if err != nil {
		return nil, err
	}
	lcr, err := ds.GetGRPC().ListCollections(ctx, p)
	if err != nil {
		return nil, err
	}
	filtCollections := make([]string, 0, len(lcr.Collections))
	for _, col := range lcr.Collections {
		cerr := a.checkPermissionsByCollection(ctx, col, P_READ)
		if cerr != nil {
			continue
		}
		filtCollections = append(filtCollections, col)
	}
	lcr.Collections = filtCollections
	return lcr, err
}
func (a *apiProvider) LookupStreams(p *pb.LookupStreamsParams, r pb.BTrDB_LookupStreamsServer) error {
	ctx := r.Context()
	ds, err := a.readEndpoint(ctx, uuid.NewRandom())
	if err != nil {
		return err
	}
	client, err := ds.GetGRPC().LookupStreams(ctx, p)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if resp.Stat != nil {
			cerr := r.Send(resp)
			if cerr != nil {
				return cerr
			}
		}
		//Filter the results by permitted ones
		nr := make([]*pb.StreamDescriptor, 0, len(resp.Results))
		for _, res := range resp.Results {
			sterr := a.checkPermissionsByCollection(ctx, res.Collection, P_READ)
			if sterr != nil {
				continue
			}
			nr = append(nr, res)
		}
		resp.Results = nr
		err = r.Send(resp)
		if err != nil {
			return err
		}
	}
}
func (a *apiProvider) Nearest(ctx context.Context, p *pb.NearestParams) (*pb.NearestResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_READ)
	if err != nil {
		return nil, err
	}
	ds, err := a.readEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	return ds.GetGRPC().Nearest(ctx, p)
}

func (a *apiProvider) Insert(ctx context.Context, p *pb.InsertParams) (*pb.InsertResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_MODIFY)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	rv, e := ds.GetGRPC().Insert(ctx, p)
	return rv, e
}
func (a *apiProvider) Delete(ctx context.Context, p *pb.DeleteParams) (*pb.DeleteResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_MODIFY)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	return ds.GetGRPC().Delete(ctx, p)
}

func (a *apiProvider) Flush(ctx context.Context, p *pb.FlushParams) (*pb.FlushResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_MODIFY)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	rv, e := ds.GetGRPC().Flush(ctx, p)
	return rv, e
}

func (a *apiProvider) Obliterate(ctx context.Context, p *pb.ObliterateParams) (*pb.ObliterateResponse, error) {
	err := a.checkPermissionsByUUID(ctx, p.GetUuid(), P_ADMIN)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, p.GetUuid())
	if err != nil {
		return nil, err
	}
	rv, e := ds.GetGRPC().Obliterate(ctx, p)
	return rv, e
}

func (a *apiProvider) FaultInject(ctx context.Context, p *pb.FaultInjectParams) (*pb.FaultInjectResponse, error) {
	err := a.checkPermissionsByUUID(ctx, uuid.NewRandom(), P_ADMIN)
	if err != nil {
		return nil, err
	}
	ds, err := a.writeEndpoint(ctx, uuid.NewRandom())
	if err != nil {
		return nil, err
	}
	rv, e := ds.GetGRPC().FaultInject(ctx, p)
	return rv, e
}

func (a *apiProvider) Info(ctx context.Context, params *pb.InfoParams) (*pb.InfoResponse, error) {
	//We do not forward the info call, as we want the client to always contact us
	// nevertheless tihs PARTICULAR repsonse is a hack
	ProxyInfo := &pb.ProxyInfo{
		ProxyEndpoints: []string{"127.0.0.1:4410"},
	}
	return &pb.InfoResponse{
		MajorVersion: MajorVersion,
		MinorVersion: MinorVersion,
		Build:        VersionString,
		Proxy:        ProxyInfo,
	}, nil
}
