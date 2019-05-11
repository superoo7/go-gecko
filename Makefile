test:
	@go test -count=1 v3/v3_test.go v3/v3.go

test_v3:
	@go test -count=1 v3/v3_test.go v3/v3.go

test_example:
	@./script/test_example.sh
	