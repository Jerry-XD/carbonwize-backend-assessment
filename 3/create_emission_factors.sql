CREATE SEQUENCE IF NOT EXISTS emission_factors_id_seq;

CREATE TABLE "public"."emission_factors" (
    "id" int8 NOT NULL DEFAULT nextval('emission_factors_id_seq'::regclass),
    "name" varchar(255),
    "detail" varchar(1000),
    "activity_type" varchar(255),
    "unit_type" varchar(255),
    "vehicle_type" varchar(255),
    "fuel_type" varchar(255),
    "emission_factor" numeric,
    "source" varchar(255),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    PRIMARY KEY ("id")
);

INSERT INTO
	emission_factors (id, name, detail, activity_type, unit_type, vehicle_type, fuel_type, emission_factor, source, created_at, updated_at)
VALUES
	(
		1,
		'Gasoline Passenger Car',
		'Emission factor for gasoline-powered passenger cars.',
		'transportation',
		'kgCO2e',
		'car',
		'gasoline',
		0.254,
		'https://co2.myclimate.org/en/car_calculators/new',
		NOW(),
		NOW()
	),
	(2, 'Diesel Truck', 'Emission factor for diesel-powered trucks.', 'transportation', 'kgCO2e', 'truck', 'diesel', 1.25, 'https://co2.myclimate.org/en/car_calculators/new', NOW(), NOW()),
	(3, 'Electric Bus', 'Emission factor for electric buses.', 'transportation', 'kgCO2e', 'bus', 'electric', 0.168, 'https://co2.myclimate.org/en/car_calculators/new', NOW(), NOW()),
	(4, 'Biodiesel Van', 'Emission factor for vans powered by biodiesel.', 'energy_consumption', 'kgCO2e', 'van', 'biodiesel', 1.04, 'https://co2.myclimate.org/en/car_calculators/new', NOW(), NOW()),
	(
		5,
		'Natural Gas Forklift',
		'Emission factor for forklifts powered by natural gas.',
		'energy_consumption',
		'kgCO2e',
		'forklift',
		'natural gas',
		0.132,
		'https://co2.myclimate.org/en/car_calculators/new',
		NOW(),
		NOW()
	);