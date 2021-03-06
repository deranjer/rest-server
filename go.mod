module github.com/deranjer/rest-server

go 1.14

require (
	github.com/asdine/storm/v3 v3.2.1
	github.com/beorn7/perks v0.0.0-20160804104726-4c0e84591b9a // indirect
	github.com/gorilla/handlers v1.3.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.0 // indirect
	github.com/miolini/datacounter v0.0.0-20171104152933-fd4e42a1d5e0
	github.com/prometheus/client_golang v0.8.0
	github.com/prometheus/client_model v0.0.0-20171117100541-99fa1f4be8e5 // indirect
	github.com/prometheus/common v0.0.0-20180110214958-89604d197083 // indirect
	github.com/prometheus/procfs v0.0.0-20180212145926-282c8707aa21 // indirect
	github.com/restic/rest-server v0.10.0
	github.com/spf13/cobra v0.0.1
	github.com/spf13/pflag v1.0.0 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	goji.io v2.0.2+incompatible
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
)

replace goji.io v2.0.0+incompatible => github.com/goji/goji v2.0.0+incompatible

replace github.com/gorilla/handlers v1.3.0 => github.com/gorilla/handlers v1.3.0

replace github.com/deranjer/rest-server => ../rest-server
