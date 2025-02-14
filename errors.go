package main

import "errors"

var (
	ErrAPIKeyDecode      = errors.New("error decoding api key")
	ErrTXTRecordNotFound = errors.New("TXT record not found")
	ErrTXTRecordCreate   = errors.New("TXT record could not be created")
	ErrTXTRecordFetch    = errors.New("TXT record fetch failed")
	ErrTXTRecordDelete   = errors.New("TXT record delete failed")
)
