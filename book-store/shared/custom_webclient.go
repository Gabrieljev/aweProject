package shared

import (
	"time"

	"github.com/geb/aweproj/book-store/shared/config"
	"github.com/pkg/errors"

	"github.com/geb/aweproj/book-store/shared/webclient"
	customHttpClient "github.com/geb/aweproj/book-store/shared/webclient/custom_http_client"
	"go.uber.org/dig"
)

type (
	WebClientHolder struct {
		dig.In
		Outbound Outbound
	}

	Outbound interface {
		GetToken() webclient.Client
		FindBookByPubId() webclient.Client
		Path() Path
	}

	Path struct {
		Token           string
		FindBookByPubId string
	}

	outbound struct {
		getToken        webclient.Client
		findBookByPubId webclient.Client
		path            Path
	}
)

func (o outbound) GetToken() webclient.Client {
	return o.getToken
}

func (o outbound) FindBookByPubId() webclient.Client {
	return o.findBookByPubId
}

func (o outbound) Path() Path {
	return o.path
}

func (ch *WebClientHolder) Close() {
	// - define closer here
	// - example
	/*

	   var i interface{} = nil

	   i = ch.SomeField
	   if closer, ok := i.(io.Closer); ok {
	   	_ = closer.Close()
	   }

	*/
}

func NewOutbound(config *config.EnvConfiguration) (Outbound, error) {
	outboundData := Path{
		Token:           "http://localhost:9000/book-store/member/user/login",
		FindBookByPubId: "http://localhost:9000/book-store/inventory/book/find/",
	}

	duration := time.Second * 3

	customHClient := customHttpClient.NewClientFactory()

	return &outbound{
		getToken:        customHClient.Create(duration),
		findBookByPubId: customHClient.Create(duration),
		path:            outboundData,
	}, nil
}

func RegisterCustomHolder(container *dig.Container) error {
	// - register your custom holder here
	if err := container.Provide(NewOutbound); err != nil {
		return errors.Wrap(err, "failed to provide Outbound")
	}
	return nil
}
