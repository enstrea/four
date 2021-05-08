package data

import (
	"four/app/book/internal/conf"
	"four/app/book/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewData, NewBookRepo2)

type Data struct {
	db *ent.Client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	//helper := log.NewHelper("server-service/data", logger)
	//
	//client, err := ent.Open(c.Database.Driver, c.Database.Source)
	//if err != nil {
	//	helper.Errorf("failed opening connection to mysql: %v", err)
	//	return nil, nil, err
	//}
	//if err = client.Schema.Create(context.Background()); err != nil {
	//	helper.Errorf("failed creating schema resources: %v", err)
	//	return nil, nil, err
	//}
	d := &Data{
		//db: client,
	}
	//
	//return d, func() {
	//	if err := d.db.Close(); err != nil {
	//		helper.Error(err)
	//	}
	//}, nil
	return d, func() {

	}, nil
}
