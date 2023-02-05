package auth

import (
	"bytes"
	"encoding/json"
	"sync"
)

type UbisoftSession struct {
	Ticket        string `json:"ticket"`
	SessionId     string `json:"sessionId"`
	Expiration    string `json:"expiration"`
	TicketNew     string `json:"ticketNew"`
	SessionIdNew  string `json:"sessionIdNew"`
	ExpirationNew string `json:"expirationNew"`
}

type AuthStore struct {
	m *sync.RWMutex
	UbisoftSession
}

func New() *AuthStore {
	return &AuthStore{
		m:              &sync.RWMutex{},
		UbisoftSession: UbisoftSession{},
	}
}

func (as *AuthStore) Read() *UbisoftSession {
	as.m.RLock()
	defer as.m.RUnlock()
	session := as.UbisoftSession
	return &session
}

func (as *AuthStore) Write(b []byte) error {
	var packet UbisoftSession
	err := json.NewDecoder(bytes.NewReader(b)).Decode(&packet)
	if err != nil {
		return err
	}

	as.m.Lock()
	as.UbisoftSession = packet
	as.m.Unlock()

	return nil
}
