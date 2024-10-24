package database

import (
	"NBAPI/internal/config"
	"NBAPI/internal/sqlc"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close()
}

type Database = *pgxpool.Pool

var (
	DB      Database
	Queries *sqlc.Queries
)

type PgTracer struct {
	logger *logrus.Logger
}

func (pt *PgTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	pt.logger.Printf("Executing query: %s, args: %v", data.SQL, data.Args)
	return ctx
}

func (pt *PgTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		pt.logger.Printf("Query failed: %v", data.Err)
	} else {
		pt.logger.Printf("Query executed successfully.")
	}
}

func Init() {
	appConf := config.Config

	username := appConf.DBUsername
	password := appConf.DBPassword
	host := appConf.DBHost
	port := appConf.DBPort
	database := appConf.DBDatabase

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	logrus.Info("Connecting to database ", connStr)

	pgxLogger := logrus.New()
	tracer := &PgTracer{logger: pgxLogger}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logrus.Fatalf("Unable to parse config: %v\n", err)
	}

	config.ConnConfig.Tracer = tracer
	ctx := context.Background()
	conn, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Connected!")
	queries := sqlc.New(conn)
	DB = conn
	Queries = queries
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := DB.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		logrus.Fatalf("db down: %v", err) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := DB.Stat()
	stats["open_connections"] = strconv.Itoa(int(dbStats.TotalConns()))
	stats["in_use"] = strconv.Itoa(int(dbStats.AcquiredConns()))
	stats["idle"] = strconv.Itoa(int(dbStats.IdleConns()))
	stats["wait_count"] = strconv.FormatInt(dbStats.AcquireCount(), 10)
	stats["wait_duration"] = dbStats.AcquireDuration().String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleDestroyCount(), 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeDestroyCount(), 10)

	// Evaluate stats to provide a health message
	if dbStats.TotalConns() > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.AcquireCount() > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleDestroyCount() > int64(dbStats.TotalConns())/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeDestroyCount() > int64(dbStats.TotalConns())/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func Close() {
	logrus.Printf("Disconnected from database: %s", config.Config.DBDatabase)
	DB.Close()
}
