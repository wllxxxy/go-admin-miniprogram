module adminapi

go 1.17

require github.com/beego/beego/v2 v2.0.5

require (
	common v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/smartystreets/goconvey v1.7.2
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20181017120253-0766667cb4d1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/smartystreets/assertions v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace common => ../common
