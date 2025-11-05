CREATE TABLE shovel_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    shovel_id UUID NOT NULL REFERENCES shovels(id),
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);