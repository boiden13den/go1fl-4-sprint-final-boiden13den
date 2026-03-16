module github.com/Yandex-Practicum/tracker

go 1.25.0

require (
	daysteps v0.0.0-00010101000000-000000000000
	spentcalories v0.0.0-00010101000000-000000000000
)

replace spentcalories => ./internal/spentcalories

replace daysteps => ./internal/daysteps
