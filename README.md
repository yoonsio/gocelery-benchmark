# gocelery-benchmark

## run python worker
```
celery -A worker worker --loglevel=debug --without-heartbeat --without-mingle
```

## run go worker
```
go run worker/main.go
```

## run all python client tests
```
python test_client.py
```

## run single python client test
```
python test_client.py TestCeleryExecution.test_add
```
