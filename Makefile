# GO=go
# GOCOVER=$(GO) tool cover

# .PHONY: test/cover
# test/cover:
#     $(GOTEST) -v -coverprofile=coverage.out ./...
#     $(GOCOVER) -func=coverage.out
#     $(GOCOVER) -html=coverage.out


simple_test:
	go test ./...

simple_cover:
	go test -cover ./...

coverprofile:
	go test -coverprofile=coverage.out ./...

func_cover:
	go tool cover -func=coverage.out

html_cover:
	go tool cover -html=coverage.out

test: coverprofile func_cover html_cover