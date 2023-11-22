package main

import (
	"fmt"

	"github.com/sonu-kumar-saw/rds-poc/pkg/logger"
	"github.com/sonu-kumar-saw/rds-poc/pkg/model"
)

func main() {
	db := model.Database{}
	logger := logger.NewSugaredLogger(false, "servo")
	db.Init(logger)

	db.CreateIngestData(model.Ingest{
		IngestData:  `data:"some-important-data-0"`,
		IngestLabel: "ip01",
		SourceLabel: "source-01",
		PlacementId: "s1-ip1-az-1",
	})

	db.CreateIngestData(model.Ingest{
		IngestData:  `data:"some-important-data-1"`,
		IngestLabel: "ip02",
		SourceLabel: "source-02",
		PlacementId: "s1-ip1-az-2",
	})

	db.CreateIngestData(model.Ingest{
		IngestData:  `data:"some-important-data-2"`,
		IngestLabel: "ip02",
		SourceLabel: "source-03",
		PlacementId: "s1-ip1-az-3",
	})

	allIngests, err := db.GetAllIngests()
	if err != nil {
		logger.Infow("Error occurred while getting all the ingests from the database", "error", err)
		return
	}
	fmt.Println(allIngests)
}
