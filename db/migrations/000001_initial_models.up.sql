CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE chan_width AS ENUM ('W', 'N');

CREATE TYPE tone_type AS ENUM ('CSQ', 'CTCSS', 'DCS');

CREATE TYPE chan_mode AS ENUM ('A', 'D', 'M');

CREATE TABLE users
(
    id         SERIAL                  NOT NULL,
    created_at timestamp DEFAULT NOW() NOT NULL,
    updated_at timestamp DEFAULT NOW() NOT NULL,
    deleted_at timestamp,
    username   varchar(255)            NOT NULL UNIQUE,
    auth0_id   varchar(255)            NOT NULL,
    PRIMARY KEY (id)
);
COMMENT ON TABLE users IS 'A system user.';
COMMENT ON COLUMN users.username IS 'Unique identifier to be used publicly.';
COMMENT ON COLUMN users.auth0_id IS 'The ID by which the user will be referred to in the Auth0 system.';

CREATE TABLE channels
(
    id           SERIAL                                  NOT NULL,
    public_id    uuid         DEFAULT UUID_GENERATE_V4() NOT NULL UNIQUE,
    created_at   timestamp    DEFAULT NOW()              NOT NULL,
    updated_at   timestamp    DEFAULT NOW()              NOT NULL,
    deleted_at   timestamp,
    name         varchar(255)                            NOT NULL,
    rx_frequency numeric(8, 4)                           NOT NULL,
    rx_width     chan_width                              NOT NULL,
    rx_tone_type tone_type    DEFAULT 'CSQ'              NOT NULL,
    rx_tone      numeric(4, 1),
    tx_frequency numeric(8, 4),
    tx_width     chan_width,
    tx_tone_type tone_type,
    tx_tone      numeric(4, 1),
    mode         chan_mode    DEFAULT 'A'                NOT NULL,
    remarks      varchar(255) DEFAULT ''                 NOT NULL,
    is_public    bool         DEFAULT 'false'            NOT NULL,
    user_id      int4                                    NOT NULL,
    band_id      int4                                    NOT NULL,
    PRIMARY KEY (id)
);
COMMENT ON COLUMN channels.name IS 'User-defined channel name.';
COMMENT ON COLUMN channels.rx_frequency IS 'The receive frequency of this channel.';
COMMENT ON COLUMN channels.rx_width IS 'Channel width enum with values (W, N). ';
COMMENT ON COLUMN channels.rx_tone_type IS 'Enum of tone types with values: (CSQ, CTCSS, DCS).';
COMMENT ON COLUMN channels.rx_tone IS 'The PL tone needed for receiving this signal. If rx_tone_type is ''CSQ'', this should be null. If rx_tone_type is NOT CSQ, then this should not be null.';
COMMENT ON COLUMN channels.tx_frequency IS 'The frequency to use when transmitting on this channel. If this is a receive-only frequency, this will be null. For simplex operation, this field should equal the rq_frequency field.';
COMMENT ON COLUMN channels.tx_width IS 'Channel width enum with values (W, N).';
COMMENT ON COLUMN channels.tx_tone_type IS 'Enum of tone types with values: (CSQ, CTCSS, DCS).';
COMMENT ON COLUMN channels.tx_tone IS 'The PL tone to be used when transmitting.';
COMMENT ON COLUMN channels.mode IS 'Enum with values (A, D, M) representing Analog, Digital, or Mixed mode transmissions.';
COMMENT ON COLUMN channels.remarks IS 'User-provided remarks about this channel.';
COMMENT ON COLUMN channels.is_public IS 'Flag indicating whether this channel should be visible to users other than the owner.';

CREATE TABLE bands
(
    id              SERIAL                               NOT NULL,
    public_id       uuid      DEFAULT UUID_GENERATE_V4() NOT NULL UNIQUE,
    created_at      timestamp DEFAULT NOW()              NOT NULL,
    updated_at      timestamp DEFAULT NOW()              NOT NULL,
    deleted_at      timestamp,
    name            varchar(255)                         NOT NULL,
    lower_frequency numeric(8, 4)                        NOT NULL,
    upper_frequency numeric(8, 4)                        NOT NULL,
    service_id      int4                                 NOT NULL,
    PRIMARY KEY (id)
);
COMMENT ON TABLE bands IS 'Grouping of similar frequencies by service.';
COMMENT ON COLUMN bands.name IS 'The name of this frequency band. This is usually the rough wavelength of these frequencies in meters.';
COMMENT ON COLUMN bands.lower_frequency IS 'The lowest frequency in this band.';
COMMENT ON COLUMN bands.upper_frequency IS 'The highest frequency in this band.';

CREATE TABLE services
(
    id         SERIAL                               NOT NULL,
    public_id  uuid      DEFAULT UUID_GENERATE_V4() NOT NULL UNIQUE,
    created_at timestamp DEFAULT NOW()              NOT NULL,
    updated_at timestamp DEFAULT NOW()              NOT NULL,
    deleted_at timestamp,
    name       varchar(255)                         NOT NULL,
    cfr_part   varchar(6)                           NOT NULL,
    license    bool      DEFAULT 'false'            NOT NULL,
    PRIMARY KEY (id)
);
COMMENT ON TABLE services IS 'Radio service defined by Title 47 of the Code of Federal Regulations (CFR).';
COMMENT ON COLUMN services.name IS 'The name of this radio service.';
COMMENT ON COLUMN services.cfr_part IS 'The part and subpart of 47 CFR where this service is defined.';
COMMENT ON COLUMN services.license IS 'Flag indicates if a user must hold a license to operate with this service.';

CREATE TABLE zones
(
    id          SERIAL                               NOT NULL,
    public_id   uuid      DEFAULT UUID_GENERATE_V4() NOT NULL UNIQUE,
    created_at  timestamp DEFAULT NOW()              NOT NULL,
    updated_at  timestamp DEFAULT NOW()              NOT NULL,
    deleted_at  timestamp,
    name        varchar(255)                         NOT NULL,
    description text                                 NOT NULL,
    is_public   bytea     DEFAULT 'false'            NOT NULL,
    user_id     int4                                 NOT NULL,
    PRIMARY KEY (id)
);
COMMENT ON TABLE zones IS 'Geographic area under which channels can be grouped.';
COMMENT ON COLUMN zones.name IS 'User defined name for this zone.';
COMMENT ON COLUMN zones.description IS 'User-provided description of this zone.';
COMMENT ON COLUMN zones.is_public IS 'Flag indicating whether other users can see this Zone.';

CREATE TABLE zone_channels
(
    zone_id    int4 NOT NULL,
    channel_id int4 NOT NULL,
    PRIMARY KEY (zone_id,
                 channel_id)
);
ALTER TABLE bands
    ADD CONSTRAINT fk_bands_services FOREIGN KEY (service_id) REFERENCES services (id);
ALTER TABLE zones
    ADD CONSTRAINT fk_zones_users FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE channels
    ADD CONSTRAINT fk_channels_users FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE channels
    ADD CONSTRAINT fk_channels_bands FOREIGN KEY (band_id) REFERENCES bands (id);
ALTER TABLE zone_channels
    ADD CONSTRAINT fk_zone_channels_zones FOREIGN KEY (zone_id) REFERENCES zones (id);
ALTER TABLE zone_channels
    ADD CONSTRAINT fk_zone_channels_channels FOREIGN KEY (channel_id) REFERENCES channels (id);
