package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

var qualityBands = []int{1, 2, 3, 4}

// awardsPerQualityBand struct contains the number of shipments awarded to each tsp according to quality band
var awardsPerQualityBand = map[int]int{
	1: 5,
	2: 3,
	3: 2,
	4: 1,
}

// TransportationServiceProviderPerformance is a combination of all TSP
// performance metrics (BVS, Quality Band) for a performance period.
type TransportationServiceProviderPerformance struct {
	ID                              uuid.UUID `json:"id" db:"id"`
	CreatedAt                       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at" db:"updated_at"`
	PerformancePeriodStart          time.Time `json:"performance_period_start" db:"performance_period_start"`
	PerformancePeriodEnd            time.Time `json:"performance_period_end" db:"performance_period_end"`
	TrafficDistributionListID       uuid.UUID `json:"traffic_distribution_list_id" db:"traffic_distribution_list_id"`
	TransportationServiceProviderID uuid.UUID `json:"transportation_service_provider_id" db:"transportation_service_provider_id"`
	QualityBand                     *int      `json:"quality_band" db:"quality_band"`
	BestValueScore                  int       `json:"best_value_score" db:"best_value_score"`
	AwardCount                      int       `json:"award_count" db:"award_count"`
}

// String is not required by pop and may be deleted
func (t TransportationServiceProviderPerformance) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// TransportationServiceProviderPerformances is a handy type for multiple TransportationServiceProviderPerformance structs
type TransportationServiceProviderPerformances []TransportationServiceProviderPerformance

// String is not required by pop and may be deleted
func (t TransportationServiceProviderPerformances) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
func (t *TransportationServiceProviderPerformance) Validate(tx *pop.Connection) (*validate.Errors, error) {
	// Pop can't validate pointers to ints, so turn the pointer into an integer.
	// Our valid values are [nil, 1, 2, 3, 4]
	qualityBand := 1
	if t.QualityBand != nil {
		qualityBand = *t.QualityBand
	}

	return validate.Validate(
		// Quality Bands can have a range from 1 - 4 as defined in DTR 402. See page 67 of
		// https://www.ustranscom.mil/dtr/part-iv/dtr-part-4-402.pdf
		&validators.IntIsGreaterThan{Field: qualityBand, Name: "QualityBand", Compared: 0},
		&validators.IntIsLessThan{Field: qualityBand, Name: "QualityBand", Compared: 5},

		// Best Value Scores can range from 0 - 100, as defined in DTR403. See page 7
		// of https://www.ustranscom.mil/dtr/part-iv/dtr-part-4-403.pdf
		&validators.IntIsGreaterThan{Field: t.BestValueScore, Name: "BestValueScore", Compared: -1},
		&validators.IntIsLessThan{Field: t.BestValueScore, Name: "BestValueScore", Compared: 101},
	), nil
}

// FetchNextQualityBandTSPPerformance returns TSP performance records in a given TDL
// in the order that they should be awarded new shipments.
func FetchNextQualityBandTSPPerformance(tx *pop.Connection, tdlID uuid.UUID, qualityBand int) (
	TransportationServiceProviderPerformance, error) {

	sql := `SELECT
			*
		FROM
			transportation_service_provider_performances
		WHERE
			traffic_distribution_list_id = $1
			AND
			quality_band = $2
		ORDER BY
			award_count ASC,
			best_value_score DESC
		`

	tsp := TransportationServiceProviderPerformance{}
	err := tx.RawQuery(sql, tdlID, qualityBand).First(&tsp)

	return tsp, err
}

// GatherNextEligibleTSPPerformances returns a map of QualityBands to their next eligible TSPPerformance.
func GatherNextEligibleTSPPerformances(tx *pop.Connection, tdlID uuid.UUID) (TransportationServiceProviderPerformances, error) {
	tspPerformances := TransportationServiceProviderPerformances{}
	for _, qualityBand := range qualityBands {
		tspPerformance, err := FetchNextQualityBandTSPPerformance(tx, tdlID, qualityBand)
		if err != nil {
			fmt.Printf("\tNo TSP returned for Quality Band: %d\n; See error: %s", qualityBand, err)
			return nil, err
		}
		tspPerformances = append(tspPerformances, tspPerformance)
	}
	return tspPerformances, nil
}

// DetermineNextTSPPerformance returns the tspPerformance that is next to receive a shipment.
func DetermineNextTSPPerformance(tspPerformances TransportationServiceProviderPerformances) (TransportationServiceProviderPerformance, error) {
	// First time through, no rounds have yet occurred so set to 0.
	var tspPerformance TransportationServiceProviderPerformance
	previousRounds := 0
	for qualityBand, tspPerformance := range tspPerformances {
		rounds := tspPerformance.AwardCount / awardsPerQualityBand[qualityBand]

		if rounds <= previousRounds {
			return tspPerformance, nil
		}
		previousRounds = rounds
	}
	return tspPerformance, errors.New("there was an issue determining which TSP should next be awarded a shipment")
}

// FetchTSPPerformanceForQualityBandAssignment returns TSPs in a given TDL in the
// order that they should be assigned quality bands.
func FetchTSPPerformanceForQualityBandAssignment(tx *pop.Connection, tdlID uuid.UUID, mps int) (TransportationServiceProviderPerformances, error) {

	sql := `SELECT
			*
		FROM
			transportation_service_provider_performances
		WHERE
			traffic_distribution_list_id = $1
			AND
			best_value_score > $2
		ORDER BY
			best_value_score DESC
		`

	tsps := TransportationServiceProviderPerformances{}
	err := tx.RawQuery(sql, tdlID, mps).All(&tsps)

	return tsps, err
}

// AssignQualityBandToTSPPerformance sets the QualityBand value for a TransportationServiceProviderPerformance.
func AssignQualityBandToTSPPerformance(db *pop.Connection, band int, id uuid.UUID) error {
	performance := TransportationServiceProviderPerformance{}
	if err := db.Find(&performance, id); err != nil {
		return err
	}
	performance.QualityBand = &band
	verrs, err := db.ValidateAndUpdate(&performance)
	if err != nil {
		return err
	} else if verrs.Count() > 0 {
		return errors.New("could not update quality band")
	}
	return nil
}
