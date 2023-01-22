package wal

import (
	"betbright-management-console/domain"
	"context"
	"encoding/binary"
	"github.com/golover/wal-listener/config"
	"github.com/golover/wal-listener/listener"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
)

type PGWalDelivery struct {
	postgresWalListener *listener.Listener
	su                  domain.SearchUsecase
}

func (pgwd *PGWalDelivery) Publish(subject string, event listener.Event) error {
	switch listener.ActionKind(event.Action) {
	case listener.ActionKindInsert:
		err := pgwd.su.Index(context.Background(), event.Table, event.Data)
		return err
	case listener.ActionKindUpdate:
		//err := pgwd.su.DeleteIndex()
		//err = pgwd.su.Index()
		//return err
	case listener.ActionKindDelete:
		//err := pgwd.su.DeleteIndex()
		//return err
	}
	return nil
}
func (pgwd *PGWalDelivery) BindListener(logger *logrus.Entry, pgxConn *pgx.Conn, pgxReplicationConnection *pgx.ReplicationConn) {
	pgwd.postgresWalListener = listener.NewWalListener(&config.Config{
		Listener: config.ListenerCfg{
			SlotName:          "betbright",
			RefreshConnection: 30,
			HeartbeatInterval: 10,
			Filter: config.FilterStruct{Tables: map[string][]string{"events": []string{"insert", "update", "delete"},
				"markets":    []string{"insert", "update", "delete"},
				"selections": []string{"insert", "update", "delete"},
				"sports":     []string{"insert", "update", "delete"}}},
			TopicsMap: nil,
		},
		Database: config.DatabaseCfg{},
		Nats:     config.NatsCfg{},
		Logger: config.LoggerCfg{
			Caller: false,
			Level:  "info",
			Format: "json",
		},
		Monitoring: config.MonitoringCfg{},
	}, logger, listener.NewRepository(pgxConn), pgxReplicationConnection, pgwd, listener.NewBinaryParser(binary.BigEndian))
	err := pgwd.postgresWalListener.Process(context.Background())
	if err != nil {
		panic(err)
	}
}
func New(searchUsecase domain.SearchUsecase) *PGWalDelivery {
	pgwd := &PGWalDelivery{nil, searchUsecase}
	return pgwd
}
