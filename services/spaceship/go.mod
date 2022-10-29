module imperial-fleet-inventory/services/spaceship

go 1.17

replace imperial-fleet-inventory/common => ../../common

replace imperial-fleet-inventory/api => ../../api

require (
	github.com/cristalhq/aconfig v0.18.3
	github.com/cristalhq/aconfig/aconfigdotenv v0.17.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.23.0
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.50.1
	gorm.io/driver/mysql v1.4.3
	gorm.io/gorm v1.24.0
	imperial-fleet-inventory/api v0.0.0-00010101000000-000000000000
	imperial-fleet-inventory/common v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220317061510-51cd9980dadf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
