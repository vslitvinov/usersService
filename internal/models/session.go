package models

import "time"

type Session struct {
	UUIDUser string    
	expiry   time.Time 
}

