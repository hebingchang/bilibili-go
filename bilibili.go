package bilibili

import (
	"bilibili/proto/generated/app/archive/middleware"
	"bilibili/proto/generated/app/view"
	bilimeta "bilibili/proto/generated/metadata"
	"bilibili/proto/generated/metadata/device"
	"bilibili/proto/generated/metadata/locale"
	"bilibili/proto/generated/metadata/network"
	"bilibili/proto/generated/metadata/parabox"
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

var (
	conn *grpc.ClientConn
)

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

func GetArchiveInfo[T int64 | string](archive T) (*view.ViewReply, error) {
	client := view.NewViewClient(conn)
	req := &view.ViewReq{
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

	switch v := ((any)(archive)).(type) {
	case string:
		req.Bvid = v
		break
	case int64:
		req.Aid = v
	}

	return client.View(
		buildContext(context.Background()),
		req,
	)
}
