package utils

import "math"

type StatisticsCalculator struct {
	resultSet []ResultSet
}

func NewStatisticsCalculator() *StatisticsCalculator {
	return &StatisticsCalculator{}
}

func (sc *StatisticsCalculator) AddResultSet(set ResultSet) {
	sc.resultSet = append(sc.resultSet, set)
}

func (sc StatisticsCalculator) GetResultSet() []ResultSet {
	return sc.resultSet
}

func (sc StatisticsCalculator) GetAverageTimeInMillis() int64 {
	var sumTime int64
	for _, resultSet := range sc.resultSet {
		sumTime += resultSet.Duration.Milliseconds()
	}
	return sumTime / int64(len(sc.resultSet))
}

func (sc StatisticsCalculator) GetStandardDeviation () float64 {
	var execTime float64
	for _, rs := range sc.resultSet {
		aux := rs.Duration.Milliseconds() - sc.GetAverageTimeInMillis()
		execTime += math.Pow(float64(aux), 2) / float64(len(sc.resultSet))
	}
	return math.Sqrt(execTime)
}

func (sc StatisticsCalculator) GetCoefficientOfVariation () float64 {
	return (sc.GetStandardDeviation() / float64(sc.GetAverageTimeInMillis())) * 100
}