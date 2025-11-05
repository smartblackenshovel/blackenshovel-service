CREATE TABLE spot_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    spot_id UUID NOT NULL REFERENCES spots(id),
    status VARCHAR(20) NOT NULL,
    session_id UUID REFERENCES sessions(id),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);