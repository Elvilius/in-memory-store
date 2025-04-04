package server

import (
	"context"
	"io"
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

	tcpServer := &TCPServer{
		cfg:             cfg,
		db:              db,
		logger:          logger,
		listener:        listener,
		connectionCount: nil,
	}

	if cfg.Network.MaxConnections != 0 {
		tcpServer.connectionCount = make(chan struct{}, cfg.Network.MaxConnections)
	}

	return tcpServer, nil
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

			s.Wait()
			go func() {
				defer func() {
					s.Signal()
				}()

				s.queryHandler(conn)
			}()

		}
	}()

	<-ctx.Done()

	if s.connectionCount != nil {
		close(s.connectionCount)
	}
	wg.Wait()
}

func (s *TCPServer) queryHandler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, s.cfg.Network.BufferSize)

	for {
		count, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				s.logger.Info("client disconnected", zap.String("address", conn.RemoteAddr().String()))
			} else {
				s.logger.Error("failed to read from connection", zap.Error(err))
			}
			break
		}

		request := string(buf[:count])
		response := s.db.CommandHandle(request)

		_, err = conn.Write([]byte(response))
		if err != nil {
			s.logger.Warn("failed to write response", zap.String("address", conn.RemoteAddr().String()), zap.Error(err))
			break
		}
	}
}

func (s *TCPServer) Wait() {
	if s.connectionCount != nil {
		s.connectionCount <- struct{}{}
	}
}

func (s *TCPServer) Signal() {
	if s.connectionCount != nil {
		<-s.connectionCount
	}
}
