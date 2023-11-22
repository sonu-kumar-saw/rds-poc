package model

// Ingest represents the 'ingests' table in the database.
type Ingest struct {
	ID          int    `gorm:"column:id;primaryKey;"`
	IngestData  string `gorm:"column:ingest_data;index"`
	IngestLabel string `gorm:"column:ingest_label"`
	SourceLabel string `gorm:"column:source_label"`
	PlacementId string `gorm:"column:placement_id;index"`
}

func (db *Ingest) TableName() string {
	return "ingests"
}

func (db *Database) GetAllIngests() ([]Ingest, error) {
	var ingests []Ingest
	var ingest Ingest
	err := db.mysqlConn.Table(ingest.TableName()).Find(&ingests).Error
	if err != nil {
		return nil, err
	}
	return ingests, err
}

func (db *Database) CreateIngestData(data Ingest) error {
	var ingest Ingest
	err := db.mysqlConn.Table(ingest.TableName()).Create(&data).Error
	if err != nil {
		db.logger.Error("err", err)
	}
	return err
}
