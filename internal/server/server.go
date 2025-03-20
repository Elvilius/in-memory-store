package server

import (
	"context"
	"net"
	"sync"

	"github.com/Elvilius/in-memory-store/internal/config"
	"github.com/Elvilius/in-memory-store/internal/db"
	"go.uber.org/zap"
)

type TCPServer struct {
	cfg             *config.Config
	db              *db.DB
	logger          *zap.Logger
	listener        net.Listener
	connectionCount chan struct{}
}

func NewTCPServer(cfg *config.Config, db *db.DB, logger *zap.Logger) (*TCPServer, error) {
	listener, err := net.Listen("tcp", cfg.Network.Address)
	if err != nil {
		return nil, err
	}

	return &TCPServer{
		cfg:             cfg,
		db:              db,
		logger:          logger,
		listener:        listener,
		connectionCount: make(chan struct{}, cfg.Network.MaxConnections),
	}, nil
}

func (s *TCPServer) Run(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)
	s.logger.Info("tcp server start address")

	go func() {
		defer wg.Done()
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.logger.Sugar().Errorln(err)
				continue
			}

			s.connectionCount <- struct{}{}

			go func(c net.Conn) {

				defer func() {
					<-s.connectionCount
				}()


				request := make([]byte, 4<<10)
				count, err := conn.Read(request)


				if err != nil {
					s.logger.Sugar().Errorln(err)
				}

				res := s.db.CommandHandle(string(request[:count]))
				if _, err := conn.Write([]byte(res)); err != nil {
					s.logger.Warn(
						"failed to write data",
						zap.String("address", conn.RemoteAddr().String()),
						zap.Error(err),
					)
				}

			}(conn)

		}
	}()

	<-ctx.Done()
	close(s.connectionCount)
	wg.Wait()


}
