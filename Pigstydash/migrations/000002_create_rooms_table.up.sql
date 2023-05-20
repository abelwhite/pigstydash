CREATE TABLE if NOT EXISTS rooms(
    room_id bigserial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    num_of_pigsty bigserial,
    temperature text NOT NULL,
    humidity text NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW ()
);