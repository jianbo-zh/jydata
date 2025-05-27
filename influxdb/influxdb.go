package influxdb

import (
	configV1 "github.com/jianbo-zh/jypb/config/v1"

	"github.com/go-kratos/kratos/v2/log"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type InfluxDB struct {
	org    string
	bucket string
	client influxdb2.Client
}

func (i *InfluxDB) WritePoint(point *write.Point) {
	i.client.WriteAPI(i.org, i.bucket).WritePoint(point)
}

func NewInfluxDB(conf *configV1.Infra, logger log.Logger) (*InfluxDB, func(), error) {
	if conf.Influxdb == nil {
		panic("config infra.influxdb not found")
	}

	client := influxdb2.NewClient(conf.Influxdb.Url, conf.Influxdb.Token)

	return &InfluxDB{
			client: influxdb2.NewClient(conf.Influxdb.Url, conf.Influxdb.Token),
			org:    conf.Influxdb.Org,
			bucket: conf.Influxdb.Bucket,
		}, func() {
			client.Close()
		}, nil
}
