CREATE TABLE if NOT EXISTS pigs(
    pig_id bigserial PRIMARY KEY,
    room text NOT NULL REFERENCES rooms(name),
    pigsty text NOT NULL REFERENCES pigsty(name),
    breed text NOT NULL, 
    age text NOT NULL,
    dob date,
    weight text NOT NULL, 
    gender text NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW ()
  
); 