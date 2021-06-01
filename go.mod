module github.com/kardiachain/go-kardia

go 1.14

require (
	github.com/Workiva/go-datastructures v1.0.52
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/aristanetworks/goarista v0.0.0-20190712234253-ed1100a1c015
	github.com/binance-chain/tss-lib v1.3.2
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cosmos/cosmos-sdk v0.42.5
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.7.1
	github.com/ebuchman/fail-test v0.0.0-20170303061230-95f809107225
	github.com/ethereum/go-ethereum v1.10.2
	github.com/fortytw2/leaktest v1.3.0
	github.com/go-stack/stack v1.8.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.3.2
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/gtank/merlin v0.1.1
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/holiman/uint256 v1.1.1
	github.com/libp2p/go-buffer-pool v0.0.2
	github.com/minio/highwayhash v1.0.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/tsdb v0.10.0
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.7.0
	github.com/sasha-s/go-deadlock v0.2.1-0.20190427202633-1595213edefa
	github.com/shirou/gopsutil v2.20.5+incompatible
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/status-im/keycard-go v0.0.0-20190424133014-d95853db0f48 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954
	github.com/tyler-smith/go-bip39 v1.0.2 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20210220033124-5f55cee0dc0d
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/urfave/cli.v1 v1.20.0
	gopkg.in/yaml.v2 v2.4.0

)

// replace github.com/binance-chain/tss-lib => gitlab.com/thorchain/tss/tss-lib v0.0.0-20201118045712-70b2cb4bf916
replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
