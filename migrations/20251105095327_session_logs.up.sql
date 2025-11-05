CREATE TABLE session_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL REFERENCES sessions(id),
    spot_id UUID REFERENCES spots(id),
    timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    altitude DOUBLE PRECISION,
    movement_type VARCHAR(20),
    movement_probability DOUBLE PRECISION,
    employee_confirmed BOOLEAN,
    sensor_data JSONB,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);