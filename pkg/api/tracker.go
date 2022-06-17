package api

import (
	"context"

	"github.com/colt005/TrackPixel/pkg"
	"github.com/gin-gonic/gin"
)

// TrackerService contains the methods of the tracker service
type TrackerService interface {
	New(uri string, ctx context.Context) error
	HandleImageLink(uri string, ctx *gin.Context) error
	FetchUriInfo(uri string, ctx *gin.Context) (pkg.URIInfo, error)
}

// TrackerRepository is what lets our service do db operations without knowing anything about the implementation
type TrackerRepository interface {
	CreatePixelTrackLink(string, context.Context) error
	HandleImageLink(string, *gin.Context) error
	FetchUriInfo(uri string, ctx *gin.Context) (pkg.URIInfo, error)
}

type trackerService struct {
	storage TrackerRepository
}

func NewTrackerService(trackerRepo TrackerRepository) TrackerService {
	return &trackerService{
		storage: trackerRepo,
	}
}

func (t *trackerService) New(uri string, ctx context.Context) error {

	err := t.storage.CreatePixelTrackLink(uri, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (t *trackerService) HandleImageLink(uri string, ctx *gin.Context) error {

	err := t.storage.HandleImageLink(uri, ctx)

	if err != nil {
		return err
	}

	return nil
}

func (t *trackerService) FetchUriInfo(uri string, ctx *gin.Context) (pkg.URIInfo, error) {

	uriInfo, err := t.storage.FetchUriInfo(uri, ctx)

	if err != nil {
		return uriInfo, err
	}

	return uriInfo, nil
}
