-- Create the alerts table
CREATE TABLE IF NOT EXISTS alerts (
    alert_id UUID PRIMARY KEY,
    service_id VARCHAR(50) NOT NULL,
    service_name VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    alert_type VARCHAR(50) NOT NULL,
    alert_ts TIMESTAMP WITH TIME ZONE NOT NULL,
    severity VARCHAR(20) NOT NULL,
    team_slack VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- Create an index on service_id for faster retrieval
CREATE INDEX idx_service_id ON alerts (service_id);
