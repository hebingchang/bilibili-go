package bilibili

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/golang/protobuf/proto"
	"github.com/hebingchang/bilibili-go/proto/generated/app/archive/middleware"
	"github.com/hebingchang/bilibili-go/proto/generated/app/playurl"
	"github.com/hebingchang/bilibili-go/proto/generated/app/view"
	bilimeta "github.com/hebingchang/bilibili-go/proto/generated/metadata"
	"github.com/hebingchang/bilibili-go/proto/generated/metadata/device"
	"github.com/hebingchang/bilibili-go/proto/generated/metadata/locale"
	"github.com/hebingchang/bilibili-go/proto/generated/metadata/network"
	"github.com/hebingchang/bilibili-go/proto/generated/metadata/parabox"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

var (
	conn *grpc.ClientConn
)

type Client struct {
	AccessToken string
}

func init() {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	conn, err = grpc.Dial("grpc.biliapi.net:443", grpc.WithTransportCredentials(
		credentials.NewTLS(&tls.Config{RootCAs: systemRoots}),
	))
	if err != nil {
		log.Fatal(err)
	}
}

func buildContext(context context.Context) context.Context {
	localeBin, _ := proto.Marshal(&locale.Locale{
		SLocale: &locale.LocaleIds{
			Language: "zh",
			Script:   "Hans",
			Region:   "JP",
		},
	})
	metadataBin, _ := proto.Marshal(&bilimeta.Metadata{
		MobiApp:  "iphone",
		Device:   "phone",
		Build:    70700200,
		Channel:  "apple",
		Platform: "ios",
	})
	expsBin, _ := proto.Marshal(&parabox.Exps{Exps: []*parabox.Exp{{
		Id: 1,
	}}})
	networkBin, _ := proto.Marshal(&network.Network{Type: network.NetworkType_WIFI})
	deviceBin, _ := proto.Marshal(&device.Device{
		AppId:       1,
		Build:       70700200,
		MobiApp:     "iphone",
		Platform:    "ios",
		Device:      "phone",
		Channel:     "apple",
		Brand:       "Apple",
		Model:       "iPhone 12 Pro",
		Osver:       "16.1.1",
		VersionName: "7.7.0",
	})

	return metadata.NewOutgoingContext(context, metadata.New(map[string]string{
		"x-bili-locale-bin":   string(localeBin),
		"x-bili-metadata-bin": string(metadataBin),
		"x-bili-exps-bin":     string(expsBin),
		"x-bili-network-bin":  string(networkBin),
		"x-bili-device-bin":   string(deviceBin),
	}))
}

func (c *Client) GetViewByAvid(avid int64) (*view.ViewReply, error) {
	req := &view.ViewReq{
		Aid:       avid,
		From:      "64",
		Qn:        112,
		Fnval:     976,
		Fourk:     1,
		Spmid:     "main.ugc-video-detail.0.0",
		FromSpmid: "main.my-history.0.0",
		PlayerArgs: &middleware.PlayerArgs{
			Qn:           112,
			Fnval:        976,
			VoiceBalance: 1,
		},
	}
	return c.GetView(req)
}

func (c *Client) GetViewByBvid(bvid string) (*view.ViewReply, error) {
	req := &view.ViewReq{
		Bvid:      bvid,
		From:      "64",
		Qn:        112,
		Fnval:     976,
		Fourk:     1,
		Spmid:     "main.ugc-video-detail.0.0",
		FromSpmid: "main.my-history.0.0",
		PlayerArgs: &middleware.PlayerArgs{
			Qn:           112,
			Fnval:        976,
			VoiceBalance: 1,
		},
	}
	return c.GetView(req)
}

func (c *Client) GetView(req *view.ViewReq) (*view.ViewReply, error) {
	client := view.NewViewClient(conn)

	return client.View(
		buildContext(context.Background()),
		req,
	)
}

func (c *Client) GetPlayViewById(aid int64, cid int64) (*playurl.PlayViewReply, error) {
	req := &playurl.PlayViewReq{
		Aid:             aid,
		Cid:             cid,
		Qn:              64,
		Fnval:           976,
		Spmid:           "main.ugc-video-detail.0.0",
		FromSpmid:       "tm.recommend.0.0",
		PreferCodecType: playurl.CodeType_CODE265,
		VoiceBalance:    1,
	}
	return c.GetPlayView(req)
}

func (c *Client) GetPlayView(req *playurl.PlayViewReq) (*playurl.PlayViewReply, error) {
	client := playurl.NewPlayURLClient(conn)

	return client.PlayView(
		buildContext(context.Background()),
		req,
	)
}

func (c *Client) GetPlayUrlById(aid int64, cid int64) (*playurl.PlayURLReply, error) {
	req := &playurl.PlayURLReq{
		Aid:       aid,
		Cid:       cid,
		Qn:        64,
		Fnval:     976,
		Spmid:     "main.ugc-video-detail.0.0",
		FromSpmid: "tm.recommend.0.0",
	}
	return c.GetPlayUrl(req)
}

func (c *Client) GetPlayUrl(req *playurl.PlayURLReq) (*playurl.PlayURLReply, error) {
	client := playurl.NewPlayURLClient(conn)

	return client.PlayURL(
		buildContext(context.Background()),
		req,
	)
}
