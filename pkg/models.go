package pkg

import "time"

type URIInfo struct {
	Uri            string            `json:"uri"`
	TotalOpens     int64             `json:"total_opens"`
	CreatedAt      time.Time         `json:"created_at"`
	TimeseriesData []TimeseriesDatum `json:"timeseries_data"`
}

type TimeseriesDatum struct {
	URI       string    `json:"uri"`
	TimeStamp time.Time `json:"time_stamp"`
	IPAddress string    `json:"ip_address"`
}
