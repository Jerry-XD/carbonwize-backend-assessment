SELECT
	name,
	activity_type,
	vehicle_type,
	fuel_type,
	emission_factor,
	unit_type,
	detail,
	source
FROM
	emission_factors
WHERE
	emission_factor IS NOT NULL
ORDER BY
	emission_factor DESC