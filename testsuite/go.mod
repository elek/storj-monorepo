module storj.io/gateway/testsuite

go 1.14

replace storj.io/gateway => ../

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/minio/minio v0.0.0-20201125204248-91130e884b5d
	github.com/stretchr/testify v1.7.0
	github.com/zeebo/errs v1.2.2
	go.uber.org/zap v1.16.0
	storj.io/common v0.0.0-20210408063549-b10b43cba35a
	storj.io/gateway v0.0.0-00010101000000-000000000000
	storj.io/storj v0.12.1-0.20210409141935-51d667a65e15
	storj.io/uplink v1.4.6-0.20210409140224-ad6bc2e5b2ee
)
