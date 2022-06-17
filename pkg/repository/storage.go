package repository

import (
	"context"
	"time"

	"github.com/colt005/TrackPixel/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	logger "github.com/sirupsen/logrus"
)

type Storage interface {
	CreatePixelTrackLink(request string, ctx context.Context) error
	HandleImageLink(request string, ctx *gin.Context) error
	FetchUriInfo(request string, ctx *gin.Context) (pkg.URIInfo, error)
}

type storage struct {
	db *pgx.Conn
}

func NewStorage(db *pgx.Conn) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) CreatePixelTrackLink(request string, ctx context.Context) error {
	newLinkStmt := `
		INSERT INTO "track_pixels" (uri, created_at) 
		VALUES ($1, $2);
		`

	_, err := s.db.Exec(ctx, newLinkStmt, request, time.Now())

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (s *storage) HandleImageLink(request string, ctx *gin.Context) error {

	newLinkStmt := `
		INSERT INTO "track_pixels_timeseries" (uri, ip_address,ts) 
		VALUES ($1, $2, $3);
		`

	_, err := s.db.Exec(ctx, newLinkStmt, request, ctx.RemoteIP(), time.Now())

	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (s *storage) FetchUriInfo(request string, ctx *gin.Context) (pkg.URIInfo, error) {
	var uriInfo pkg.URIInfo
	uriInfo.Uri = request
	q := `select created_at from track_pixels where uri ilike '` + request + `%'`

	err := s.db.QueryRow(ctx, q).Scan(&uriInfo.CreatedAt)

	if err != nil {
		logger.Error(err)
	}

	q = `select count(*) from track_pixels_timeseries where uri ilike '` + request + `%'`

	err = s.db.QueryRow(ctx, q).Scan(&uriInfo.TotalOpens)

	if err != nil {
		logger.Error(err)
	}

	q = `
		select uri,ip_address,ts from  "track_pixels_timeseries" where uri ilike '` + request + `%'`

	rows, err := s.db.Query(ctx, q)
	if err != nil {
		logger.Error(err)
	}
	defer rows.Close()

	for rows.Next() {
		var d pkg.TimeseriesDatum

		err := rows.Scan(&d.URI, &d.IPAddress, &d.TimeStamp)
		if err != nil {
			logger.Error(err)
		}
		uriInfo.TimeseriesData = append(uriInfo.TimeseriesData, d)
	}

	return uriInfo, nil
}
