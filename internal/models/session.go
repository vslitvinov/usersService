package models

import "time"

type Session struct {
	UUIDUser string    
	Expiry   time.Time 
}

