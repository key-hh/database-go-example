.PHONY: test
test:
	go test ./...

.PHONY: testmock
testmock:
	SBTEST=mock go test ./...

.PHONY: watch
watch:
	reflex -r "\.go" -d fancy -s -- go run cmd/server/main.go

.PHONY: gen
gen:
	go generate ./ent

.PHONY: ent
ent:
	go run entgo.io/ent/cmd/ent init ${ENT}
