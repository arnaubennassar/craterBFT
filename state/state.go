package state

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/cometbft/cometbft-db"
)

const (
	dbName        = "craterDB"
	craterBaseKey = "craterNum"
	cometBaseKey  = "cometNum"
)

var (
	lastCometNumKey  = []byte("lastComet")
	lastCraterNumKey = []byte("lastCrater")
)

type State struct {
	db *db.GoLevelDB
}

func NewState(path string) *State {
	var state State
	db, err := db.NewGoLevelDB(dbName, path)
	if err != nil {
		log.Fatalf("failed to create persistent state at %s: %v", path, err)
	}
	lastCraterNumBytes, err := db.Get(lastCraterNumKey)
	if err != nil {
		log.Fatalf("failed to load state: %v", err)
	}
	var lastCraterNum uint64
	if len(lastCraterNumBytes) > 0 {
		lastCraterNum = binary.BigEndian.Uint64(lastCraterNumBytes)
	}
	log.Printf("State loaded. Last crater num is %d", lastCraterNum)
	state.db = db
	return &state
}

func (s *State) AddCrater(crater Crater) error {
	lastCraterNumBytes, err := s.db.Get(lastCraterNumKey)
	if err != nil {
		return err
	}
	var lastCraterNum uint64
	if len(lastCraterNumBytes) > 0 {
		lastCraterNum = binary.BigEndian.Uint64(lastCraterNumBytes)
	}
	b := s.db.NewBatch()
	var newLastCraterNumBytes []byte
	lastCraterNum += 1
	binary.BigEndian.PutUint64(newLastCraterNumBytes, lastCraterNum)
	if err := b.Set(lastCraterNumKey, newLastCraterNumBytes); err != nil {
		if closeErr := b.Close(); closeErr != nil {
			return fmt.Errorf("error closing the db batch: %e. Error stroing lastCraterNumKey: %e", closeErr, err)
		}
		return err
	}
	key := []byte(fmt.Sprintf(craterBaseKey+"%d", lastCraterNum))
	craterBytes, err := json.Marshal(crater)
	if err != nil {
		return err
	}
	if err := b.Set(key, craterBytes); err != nil {
		if closeErr := b.Close(); closeErr != nil {
			return fmt.Errorf("error closing the db batch: %e. Error stroing crater: %e", closeErr, err)
		}
		return err
	}
	if err := b.Write(); err != nil {
		if closeErr := b.Close(); closeErr != nil {
			return fmt.Errorf("error closing the db batch: %e. Error writting batch: %e", closeErr, err)
		}
		return err
	}

	return b.Close()
}

func (s *State) AddComet(comet Comet) error {
	lastCometNumBytes, err := s.db.Get(lastCometNumKey)
	if err != nil {
		return err
	}
	var lastCometNum uint64
	if len(lastCometNumBytes) > 0 {
		lastCometNum = binary.BigEndian.Uint64(lastCometNumBytes)
	}
	b := s.db.NewBatch()
	var newLastCometNumBytes []byte
	lastCometNum += 1
	binary.BigEndian.PutUint64(newLastCometNumBytes, lastCometNum)
	if err := b.Set(lastCometNumKey, newLastCometNumBytes); err != nil {
		if closeErr := b.Close(); closeErr != nil {
			return fmt.Errorf("error closing the db batch: %e. Error stroing lastCometNumKey: %e", closeErr, err)
		}
		return err
	}
	cometBytes, err := json.Marshal(comet)
	if err != nil {
		return err
	}
	key := []byte(fmt.Sprintf(cometBaseKey+"%d", lastCometNum))
	if err := b.Set(key, cometBytes); err != nil {
		if closeErr := b.Close(); closeErr != nil {
			return fmt.Errorf("error closing the db batch: %e. Error stroing comet: %e", closeErr, err)
		}
		return err
	}
	if err := b.Write(); err != nil {
		if closeErr := b.Close(); closeErr != nil {
			return fmt.Errorf("error closing the db batch: %e. Error writting batch: %e", closeErr, err)
		}
		return err
	}

	return b.Close()
}

func (s *State) Close() {
	if err := s.db.Close(); err != nil {
		log.Printf("Closing state database: %v", err)
	}
}
