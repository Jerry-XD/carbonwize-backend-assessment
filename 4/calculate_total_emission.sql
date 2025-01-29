UPDATE emission_data
SET total_emission = COALESCE(input_value, 0) * COALESCE(emission_factor, 0);