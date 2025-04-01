for i in `seq 1 100`; do go clean -testcache; go test ./... -run=TestRaceCondition; done
