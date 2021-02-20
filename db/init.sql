GRANT ALL PRIVILEGES
ON DATABASE simple_db TO demo;

CREATE SEQUENCE IF NOT EXISTS hubs_id_seq;
CREATE TABLE "hubs"
(
    "id"           int4         NOT NULL DEFAULT nextval('hubs_id_seq'::regclass),
    "name"         varchar(100) NOT NULL,
    "geo_location" point,
    PRIMARY KEY ("id")
);
ALTER SEQUENCE hubs_id_seq OWNED BY hubs.id;

CREATE SEQUENCE IF NOT EXISTS teams_id_seq;
CREATE TABLE "teams"
(
    "id"     int4         NOT NULL DEFAULT nextval('teams_id_seq'::regclass),
    "name"   varchar(200) NOT NULL,
    "type"   varchar(30)  NOT NULL,
    "hub_id" int4         NOT NULL,
    CONSTRAINT "teams_hub_id_fkey" FOREIGN KEY ("hub_id") REFERENCES "hubs" ("id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);
ALTER SEQUENCE teams_id_seq OWNED BY teams.id;

CREATE SEQUENCE IF NOT EXISTS users_id_seq;
CREATE TABLE "users"
(
    "id"    int4         NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "role"  varchar(30)  NOT NULL,
    "email" varchar(100) NOT NULL,
    PRIMARY KEY ("id")
);
ALTER SEQUENCE users_id_seq OWNED BY users.id;
CREATE UNIQUE INDEX "email_uniq" ON "users" USING BTREE ("email");

CREATE SEQUENCE IF NOT EXISTS teams_users_id_seq;
CREATE TABLE "teams_users"
(
    "id"      int4 NOT NULL DEFAULT nextval('teams_users_id_seq'::regclass),
    "team_id" int4 NOT NULL,
    "user_id" int4 NOT NULL,
    CONSTRAINT "teams_users_team_id_fkey" FOREIGN KEY ("team_id") REFERENCES "teams" ("id") ON DELETE CASCADE,
    CONSTRAINT "teams_users_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    PRIMARY KEY ("id")
);
ALTER SEQUENCE teams_users_id_seq OWNED BY teams_users.id;
CREATE UNIQUE INDEX "teams_users_uniq" ON "teams_users" USING BTREE ("team_id","user_id");


-- INSERT INTO "hubs" ("id", "name", "geo_location")
-- VALUES ('1', 'Hub A', '(21.349835,105.881325)'),
--        ('2', 'Hub B', '(21.349835,105.881325)');
--
-- INSERT INTO "teams" ("id", "name", "type", "hub_id")
-- VALUES ('1', 'Team Red', 'cool', '2'),
--        ('2', 'Team Blue', 'pro', '1');
--
-- INSERT INTO "users" ("id", "role", "email")
-- VALUES ('1', 'hub_admin', 'demo@email.com'),
--        ('2', 'hub_staff', 'staff@email.com');
--
-- INSERT INTO "teams_users" ("id", "team_id", "user_id")
-- VALUES ('1', '1', '1'),
--        ('2', '1', '2'),
--        ('3', '2', '1');

