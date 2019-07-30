# natsctl
nats natsctl 

Used to connect NATs

#### Support: 

##### publish、subscribe 、queue subscribe 

### connect nats server example

```
natsctl connect --name test-cluster -s localhost
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