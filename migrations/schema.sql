
CREATE TABLE limits (
    user_id BIGINT,
    tenor INT,
    limit_amount DECIMAL(18,2),
    used_amount DECIMAL(18,2) DEFAULT 0,
    PRIMARY KEY(user_id, tenor)
);
