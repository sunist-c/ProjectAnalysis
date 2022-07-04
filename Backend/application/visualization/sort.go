package visualization

import "ProjectAnalysis/domain/visualization"

type byTotalRecovered []visualization.Record

func (r byTotalRecovered) Len() int {
	return len(r)
}

func (r byTotalRecovered) Less(i, j int) bool {
	return r[i].TotalRecoveredCase < r[j].TotalRecoveredCase
}

func (r byTotalRecovered) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byMonthlyRecovered []visualization.Record

func (r byMonthlyRecovered) Len() int {
	return len(r)
}

func (r byMonthlyRecovered) Less(i, j int) bool {
	return r[i].MonthlyRecoveredCase < r[j].MonthlyRecoveredCase
}

func (r byMonthlyRecovered) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byWeeklyRecovered []visualization.Record

func (r byWeeklyRecovered) Len() int {
	return len(r)
}

func (r byWeeklyRecovered) Less(i, j int) bool {
	return r[i].WeeklyRecoveredCase < r[j].WeeklyRecoveredCase
}

func (r byWeeklyRecovered) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byDailyRecovered []visualization.Record

func (r byDailyRecovered) Len() int {
	return len(r)
}

func (r byDailyRecovered) Less(i, j int) bool {
	return r[i].DailyRecoveredCase < r[j].DailyRecoveredCase
}

func (r byDailyRecovered) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byTotalDeath []visualization.Record

func (r byTotalDeath) Len() int {
	return len(r)
}

func (r byTotalDeath) Less(i, j int) bool {
	return r[i].TotalDeathCase < r[j].TotalDeathCase
}

func (r byTotalDeath) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byMonthlyDeath []visualization.Record

func (r byMonthlyDeath) Len() int {
	return len(r)
}

func (r byMonthlyDeath) Less(i, j int) bool {
	return r[i].MonthlyDeathCase < r[j].MonthlyDeathCase
}

func (r byMonthlyDeath) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byWeeklyDeath []visualization.Record

func (r byWeeklyDeath) Len() int {
	return len(r)
}

func (r byWeeklyDeath) Less(i, j int) bool {
	return r[i].WeeklyDeathCase < r[j].WeeklyDeathCase
}

func (r byWeeklyDeath) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byDailyDeath []visualization.Record

func (r byDailyDeath) Len() int {
	return len(r)
}

func (r byDailyDeath) Less(i, j int) bool {
	return r[i].DailyDeathCase < r[j].DailyDeathCase
}

func (r byDailyDeath) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byTotalConfirm []visualization.Record

func (r byTotalConfirm) Len() int {
	return len(r)
}

func (r byTotalConfirm) Less(i, j int) bool {
	return r[i].TotalConfirmCase < r[j].TotalConfirmCase
}

func (r byTotalConfirm) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byMonthlyConfirm []visualization.Record

func (r byMonthlyConfirm) Len() int {
	return len(r)
}

func (r byMonthlyConfirm) Less(i, j int) bool {
	return r[i].MonthlyConfirmCase < r[j].MonthlyConfirmCase
}

func (r byMonthlyConfirm) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byWeeklyConfirm []visualization.Record

func (r byWeeklyConfirm) Len() int {
	return len(r)
}

func (r byWeeklyConfirm) Less(i, j int) bool {
	return r[i].WeeklyConfirmCase < r[j].WeeklyConfirmCase
}

func (r byWeeklyConfirm) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type byDailyConfirm []visualization.Record

func (r byDailyConfirm) Len() int {
	return len(r)
}

func (r byDailyConfirm) Less(i, j int) bool {
	return r[i].DailyConfirmCase < r[j].DailyConfirmCase
}

func (r byDailyConfirm) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
