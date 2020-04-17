package utils

import (
	"context"

	sharedpb "github.com/tron-us/go-btfs-common/protos/shared"

	"github.com/tron-us/go-common/v2/constant"
	"github.com/tron-us/go-common/v2/db"
	"github.com/tron-us/go-common/v2/db/postgres"
	"github.com/tron-us/go-common/v2/db/redis"
	"github.com/tron-us/go-common/v2/log"

	"go.uber.org/zap"
)

const DBURLDNE = "DB URL does not exist !!"
const RDURLDNE = "RD URL does not exist !!"

func CheckRuntime(ctx context.Context, runtime *sharedpb.SignedRuntimeInfoRequest, connection db.ConnectionUrls) (*sharedpb.RuntimeInfoReport, error) {
	// db runtime
	report := new(sharedpb.RuntimeInfoReport)
	report.Status = sharedpb.RuntimeInfoReport_RUNNING
	report.DbStatusExtra = make(map[string]string)

	for key, url := range connection.PgURL {
		// Assume the database connection is healthy
		report.DbStatusExtra[key] = DBURLDNE

		if url != "" {
			// Log the connection string
			pgdb := postgres.CreateTGPGDB(url)
			log.Info("Postgres",
				zap.String("name", key),
				zap.String("user", pgdb.Options().User),
				zap.String("host", pgdb.Options().Addr),
				zap.String("db", pgdb.Options().Database),
			)

			// Check postgres dbWrite
			if err := pgdb.Ping(); err != nil {
				report.DbStatusExtra[key] = constant.DBWriteConnectionError
				report.Status = sharedpb.RuntimeInfoReport_SICK
				log.Error(constant.DBWriteConnectionError, zap.Error(err))
			}

			// Set the database connection is healthy
			report.DbStatusExtra[key] = constant.DBConnectionHealthy
		}

		// Log status
		log.Info(key + ":" + report.DbStatusExtra[key])
	}

	// Assume the redis connection is not present
	report.RdStatusExtra = RDURLDNE

	// Check redis environment variable
	if connection.RdURL != "" {
		// Parse redis url
		opts, errParse := redis.ParseRedisURL(connection.RdURL)
		if errParse != nil {
			log.Error(constant.RDURLParseError, zap.Error(errParse))
		}

		// Log connection string
		log.Info("Redis URL",
			zap.String("host", opts.Addr),
			zap.String("db", string(opts.DB)),
		)

		// Check redis connection
		errConn := redis.CheckRedisConnection(redis.NewRedisConn(opts))
		if errConn != nil {
			report.RdStatusExtra = constant.RDConnectionError
			report.Status = sharedpb.RuntimeInfoReport_SICK
			log.Error(constant.RDConnectionError, zap.Error(errConn))
		}

		// Set redis connection to healthy
		report.RdStatusExtra = constant.RDConnectionHealthy
	}

	// Log status
	log.Info(report.RdStatusExtra)

	// Remaining fields will be populated by the calling service
	// Reserve: only pass fatal error to higher level
	return report, nil
}
