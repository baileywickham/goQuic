# go_server

## Goals
I want to better understading networking, so I want to write a server to handle http requests. Following this I want to imliment http/3, QUIC

## Design
This design will copy the basics of the nginx server. This is because go offers nice ways to handle concurrency through goroutines.

Processes:
- Listen socket
- Worker proc, should be indepenent. 




State machines? 
