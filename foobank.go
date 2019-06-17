package foobank

import (
	"github.com/cstrong1122/moonlight"
)

type FooBankRateProvider struct{}

func (b FooBankRateProvider) GetRateOptions(parameters *moonlight.RateOptionParameters) ([]moonlight.RateOption, error) {

	options := []moonlight.RateOption{}
	o1 := moonlight.RateOption{
		InterestRate:      6.75,
		DiscountPoints:    1.125,
		OriginationPoints: .5}
	options = append(options, o1)

	o2 := moonlight.RateOption{
		InterestRate:      6.5,
		DiscountPoints:    1.25,
		OriginationPoints: .75}
	options = append(options, o2)

	return options, nil
}
