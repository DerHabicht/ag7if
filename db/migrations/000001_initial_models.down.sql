ALTER TABLE zone_channels
    DROP CONSTRAINT fk_zone_channels_channels;
ALTER TABLE channels
    DROP CONSTRAINT fk_channels_bands;
ALTER TABLE channels
    DROP CONSTRAINT fk_channels_users;
ALTER TABLE zones
    DROP CONSTRAINT fk_zones_users;
ALTER TABLE bands
    DROP CONSTRAINT fk_bands_services;

DROP TABLE zone_channels;
DROP TABLE zones;
DROP TABLE services;
DROP TABLE bands;
DROP TABLE channels;
DROP TABLE users;

DROP TYPE chan_mode;
DROP TYPE tone_type;
DROP TYPE chan_width;

