package mysql

import (
	"fmt"
	"testing"

	. "github.com/freeusd/solebtc/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"github.com/freeusd/solebtc/errors"
	"github.com/freeusd/solebtc/models"
)

func TestGetLatestConfig(t *testing.T) {
	Convey("Given mysql storage with default config", t, func() {
		s := prepareDatabaseForTesting()

		Convey("When get latest config", func() {
			result, err := s.GetLatestConfig()

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Result should be zero value", func() {
				So(result, func(actual interface{}, expected ...interface{}) string {
					config := actual.(models.Config)
					if config.TotalRewardThreshold == 100000000 &&
						config.RefererRewardRate == 10 {
						return ""
					}
					return fmt.Sprintf("Config %v is not expected", config)
				})
			})
		})
	})

	withClosedConn(t, "When get latest config", func(s Storage) *errors.Error {
		_, err := s.GetLatestConfig()
		return err
	})
}

func TestUpdateLatestBitcoinPrice(t *testing.T) {
	Convey("Given mysql storage", t, func() {
		s := prepareDatabaseForTesting()

		Convey("When update bitcoin price", func() {
			err := s.UpdateLatestBitcoinPrice(10000000)

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	withClosedConn(t, "When update bitcoin price", func(s Storage) *errors.Error {
		return s.UpdateLatestBitcoinPrice(100)
	})
}
