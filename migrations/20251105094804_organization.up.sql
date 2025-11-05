CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    legal_form VARCHAR(50),
    registration_number VARCHAR(20),
    street VARCHAR(255),
    postal_code VARCHAR(20),
    city VARCHAR(100),
    country CHAR(2) DEFAULT 'CH',
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT now()
);
