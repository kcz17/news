package news

// service.go contains the definition and implementation (business logic) of the
// news service. Everything here is agnostic to the transport (HTTP).

import (
	"errors"
	"golang.org/x/exp/rand"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"gonum.org/v1/gonum/stat/distuv"
)

// Service is the news service, providing read operations on a saleable
// news of sock products.
type Service interface {
	List() ([]NewsItem, error) // GET /news
	Health() []Health          // GET /health
}

// Middleware decorates a Service.
type Middleware func(Service) Service

// NewsItem describes a news item.
type NewsItem struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Contents string `json:"contents" db:"contents"`
}

// Health describes the health of a service
type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

// ErrNotFound is returned when there is no sock for a given ID.
var ErrNotFound = errors.New("not found")

// ErrDBConnection is returned when connection with the database fails.
var ErrDBConnection = errors.New("database connection error")

// NewNewsService returns an implementation of the Service interface,
// with connection to an SQL database.
func NewNewsService(db *sqlx.DB, logger log.Logger) Service {
	return &newsService{
		db:     db,
		logger: logger,
	}
}

type newsService struct {
	db     *sqlx.DB
	logger log.Logger
}

func (s *newsService) List() ([]NewsItem, error) {
	var newsItems []NewsItem
	query := "SELECT news_id AS id, title, contents FROM news ORDER BY id DESC;"

	err := s.db.Select(&newsItems, query)
	if err != nil {
		s.logger.Log("database error", err)
		return []NewsItem{}, ErrDBConnection
	}

	// Set the random seed to the current time for sufficient uniqueness.
	randSeed := uint64(time.Now().UTC().UnixNano())
	delay := distuv.Normal{
		Mu:    1,
		Sigma: 1,
		Src:   rand.NewSource(randSeed),
	}.Rand()
	if delay > 0 {
		time.Sleep(time.Duration(delay) * time.Second)
	}

	return newsItems, nil
}

func (s *newsService) Health() []Health {
	var health []Health
	dbstatus := "OK"

	err := s.db.Ping()
	if err != nil {
		dbstatus = "err"
	}

	app := Health{"news", "OK", time.Now().String()}
	db := Health{"news-db", dbstatus, time.Now().String()}

	health = append(health, app)
	health = append(health, db)

	return health
}
