package db

import (
	"fmt"
	"github.com/Elvilius/in-memory-store/internal/db/compute"
	"go.uber.org/zap"
)

type EngineInterface interface {
	Set(key string, value string)
	Get(key string) (string, bool)
	Del(key string)
}

type ComputeInterface interface {
	Parse(cmd string) (compute.PreparedCommand, error)
}

type DB struct {
	engine  EngineInterface
	compute ComputeInterface
	logger  *zap.Logger
}

func New(logger *zap.Logger, engine EngineInterface, compute ComputeInterface) *DB {
	return &DB{
		engine:  engine,
		compute: compute,
		logger:  logger,
	}
}

func (db *DB) CommandHandle(cmd string) string {
	db.logger.Info("Start parsing command:", zap.Any("cmd", cmd))

	preparedCmd, err := db.compute.Parse(cmd)
	if err != nil {
		db.logger.Error("Failed to parse command", zap.Error(err))
		return fmt.Sprintf("[ERROR] parsing command: %s", err.Error())
	}

	result := fmt.Sprintf("[UNKNOWN COMMAND: %s]", preparedCmd.Cmd)
	switch preparedCmd.Cmd {
	case compute.SET:
		result = db.ExecuteSet(preparedCmd)
	case compute.GET:
		result = db.ExecuteGet(preparedCmd)
	case compute.DEL:
		result = db.ExecuteDel(preparedCmd)
	}
	return result
}

func (db *DB) ExecuteSet(cmd compute.PreparedCommand) string {
	db.engine.Set(cmd.Key, cmd.Value)
	return "[OK]"
}

func (db *DB) ExecuteGet(cmd compute.PreparedCommand) string {
	data, ok := db.engine.Get(cmd.Key)
	if !ok {
		return "[NOT FOUND]"
	}

	return data

}

func (db *DB) ExecuteDel(cmd compute.PreparedCommand) string {
	db.engine.Del(cmd.Key)
	return "[OK]"

}
