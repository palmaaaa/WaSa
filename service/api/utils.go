package api

import (
	//"fmt"
)

func valid_identifier(identifier string) bool{
	return len(identifier)>=3 && len(identifier)<=16
}