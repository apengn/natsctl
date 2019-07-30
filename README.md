# natsctl
nats natsctl 

Used to connect NATs

#### Support: 

##### publish、subscribe 、queue subscribe 

### help
```bash
natsctl connect -h
natsctl gitCommit: c5ab24b 
Usage:
  natsctl connect

Examples:


publish message : pub subj value

subscribe message through async :subasync subj

subscribe to messages through queues : subqueue subj queue



Flags:
      --ca stringArray              ca.pem file path
      --cert string                 cert.pem file path
  -h, --help                        help for connect
      --key string                  key.pem file path
  -m, --max-reconnect int           MaxReconnect sets the number of reconnect attempts that will be tried before giving up. If negative, then it will never give up trying to reconnect.
      --name string                 Name is an optional name label which will be sent to the server on CONNECT to identify the client.
      --password string             Password sets the password to be used when connecting to a server.
      --ping-interval duration      PingInterval is the period at which the client will be sending ping commands to the server, disabled if 0 or negative.
  -b, --reconnect-buffer-size int   ReconnectBufSize is the size of the backing bufio during reconnect. Once this has been exhausted publish operations will return an error.
  -r, --reconnect-wait duration     ReconnectWait sets the time to back off after attempting a reconnect to a server that we were already connected to previously.
  -s, --servers stringArray         Servers is a configured set of servers which this client will use when attempting to connect.
      --token string                Token sets the token to be used when connecting to a server.
      --user string                 User sets the username to be used when connecting to the server.

```
### connect nats server example
```
natsctl connect --name test-cluster -s localhost
natsctl gitCommit: c5ab24b 
connect nats server:nats://localhost:4222,status:Connected

``` 

### connect nats server success publish message
```bash
pub subj value
```

### subscribe subj
```bash
subasync subj
```
### subqueue subj
```bash
subqueue subj
```