package rateengine

import (
	"time"

	"github.com/gobuffalo/pop"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/unit"
)

// RateEngine encapsulates the TSP rate engine process
type RateEngine struct {
	db     *pop.Connection
	logger *zap.Logger
}

func (re *RateEngine) determineCWT(weight int) (cwt int) {
	return weight / 100
}

func (re *RateEngine) computePPM(weight int, originZip string, destinationZip string, date time.Time, inverseDiscount float64) (unit.Cents, error) {
	cwt := re.determineCWT(weight)
	// Linehaul charges
	mileage, err := re.determineMileage(originZip, destinationZip)
	if err != nil {
		re.logger.Error("Failed to determine mileage", zap.Error(err))
		return 0, err
	}
	baseLinehaulChargeCents, err := re.baseLinehaul(mileage, cwt, date)
	if err != nil {
		re.logger.Error("Failed to determine base linehaul charge", zap.Error(err))
		return 0, err
	}
	originLinehaulFactorCents, err := re.linehaulFactors(cwt, originZip, date)
	if err != nil {
		re.logger.Error("Failed to determine origin linehaul factor", zap.Error(err))
		return 0, err
	}
	destinationLinehaulFactorCents, err := re.linehaulFactors(cwt, destinationZip, date)
	if err != nil {
		re.logger.Error("Failed to determine destination linehaul factor", zap.Error(err))
		return 0, err
	}
	shorthaulChargeCents, err := re.shorthaulCharge(mileage, cwt, date)
	if err != nil {
		re.logger.Error("Failed to determine shorthaul charge", zap.Error(err))
		return 0, err
	}
	// Non linehaul charges
	originServiceFee, err := re.serviceFeeCents(cwt, originZip)
	if err != nil {
		re.logger.Error("Failed to determine origin service fee", zap.Error(err))
		return 0, err
	}
	destinationServiceFee, err := re.serviceFeeCents(cwt, destinationZip)
	if err != nil {
		re.logger.Error("Failed to determine destination service fee", zap.Error(err))
		return 0, err
	}
	pack, err := re.fullPackCents(cwt, originZip)
	if err != nil {
		re.logger.Error("Failed to determine full pack cost", zap.Error(err))
		return 0, err
	}
	unpack, err := re.fullUnpackCents(cwt, destinationZip)
	if err != nil {
		re.logger.Error("Failed to determine full unpack cost", zap.Error(err))
		return 0, err
	}
	ppmSubtotal := baseLinehaulChargeCents + originLinehaulFactorCents + destinationLinehaulFactorCents +
		shorthaulChargeCents + originServiceFee + destinationServiceFee + pack + unpack
	ppmBestValue := ppmSubtotal.MultiplyFloat64(inverseDiscount)

	// PPMs only pay 95% of the best value
	ppmPayback := ppmBestValue.MultiplyFloat64(.95)

	re.logger.Info("PPM compensation total calculated",
		zap.Int("PPM compensation total", ppmPayback.Int()),
		zap.Int("PPM subtotal", ppmSubtotal.Int()),
		zap.Float64("inverse discount", inverseDiscount),
		zap.Int("base linehaul", baseLinehaulChargeCents.Int()),
		zap.Int("origin lh factor", originLinehaulFactorCents.Int()),
		zap.Int("destination lh factor", destinationLinehaulFactorCents.Int()),
		zap.Int("shorthaul", shorthaulChargeCents.Int()),
		zap.Int("origin service fee", originServiceFee.Int()),
		zap.Int("destination service fee", destinationServiceFee.Int()),
		zap.Int("pack fee", pack.Int()),
		zap.Int("unpack fee", unpack.Int()),
	)

	return ppmPayback, nil
}

// NewRateEngine creates a new RateEngine
func NewRateEngine(db *pop.Connection, logger *zap.Logger) *RateEngine {
	return &RateEngine{db: db, logger: logger}
}