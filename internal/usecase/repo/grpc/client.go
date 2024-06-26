package grpcclient

import (
	"context"
	"encoding/json"
	"github.com/Miroshinsv/wcharge_back/config"
	"github.com/Miroshinsv/wcharge_back/internal/entity"
	pb "github.com/Miroshinsv/wcharge_back/internal/usecase/repo/grpc/gen/v1" // Замените на путь к вашему сгенерированному gRPC коду
	"strconv"

	//"github.com/Miroshinsv/wcharge_back/pkg/logger"
	"github.com/slukits/graylog"
	"google.golang.org/grpc"
)

type MqttV1Client struct {
	logger  *graylog.Logger
	conn    *grpc.ClientConn
	service pb.MqttMiddlewareV1Client // Замените YourServiceClient на название вашего сервиса
}

func NewMqttV1Client(cfg *config.Config) (*MqttV1Client, error) {
	conn, err := grpc.Dial(cfg.GRPC.URL, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	service := pb.NewMqttMiddlewareV1Client(conn) // Замените NewYourServiceClient на функцию создания клиента вашего сервиса

	//test()

	return &MqttV1Client{
		//logger:  l,
		conn:    conn,
		service: service,
	}, nil
}

func (c *MqttV1Client) Close() {
	c.conn.Close()
}

func (c *MqttV1Client) PushPowerBank(ctx context.Context, st *entity.Station, pw *entity.Powerbank) (bool, error) {
	device := pb.Device{
		Cabinet:      st.SerialNumber,
		DeviceNumber: st.SerialNumber,
	}
	pr := pb.RequestPush{
		RlSlot: uint32(pw.Position),
		//RlSeq:  1,
	}
	command := pb.CommandPush{
		Device: &device,
		Push:   &pr,
	}
	req, err := c.service.PushPowerBank(ctx, &command)
	if err != nil {
		c.logger.Info("MqttV1Client - PushPowerBank - c.service.PushPowerBank - err")
		return false, err
	} else {
		c.logger.Debug(strconv.Itoa(int(req.RlResult)))
	}
	return true, nil
}

func (c *MqttV1Client) ForcePushPowerBank(ctx context.Context, st *entity.Station, pw *entity.Powerbank) error {
	return nil
}

func (c *MqttV1Client) QueryInventory(ctx context.Context, st *entity.Station) (*pb.ResponseInventory, error) {
	device := pb.Device{
		Cabinet:      st.SerialNumber,
		DeviceNumber: st.SerialNumber,
	}
	pr := pb.RequestInventory{
		//RlSlot: 1,
		RlSeq: 1,
	}
	command := pb.CommandInventory{
		Device: &device,
		Invent: &pr,
	}
	req, err := c.service.QueryInventory(ctx, &command)
	if err != nil {
		c.logger.Info("MqttV1Client - QueryInventory - c.service.QueryInventory - err")
		return req, err
	} else {
		r, _ := json.Marshal(req)
		c.logger.Debug(string(r))
	}
	return req, nil
}

func (c *MqttV1Client) QueryServerInformation(ctx context.Context) error {
	return nil
}

func (c *MqttV1Client) QueryCabinetAPN(ctx context.Context) error {
	return nil
}

func (c *MqttV1Client) QuerySIMCardICCID(ctx context.Context) error {
	return nil
}

func (c *MqttV1Client) QueryNetworkInformation(ctx context.Context) error {
	return nil
}

func (c *MqttV1Client) ResetCabinet(ctx context.Context) error {
	return nil
}

func (c *MqttV1Client) Subscribe(ctx context.Context) error {
	return nil
}
