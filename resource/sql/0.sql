CREATE TABLE "fronius"."current_data_meter" (
  "id"                                       SERIAL8,
  "current_ac_phase_1"                       DOUBLE PRECISION NOT NULL,
  "current_ac_phase_2"                       DOUBLE PRECISION,
  "current_ac_phase_3"                       DOUBLE PRECISION,
  "current_ac_sum"                           DOUBLE PRECISION NOT NULL,
  "enable"                                   BOOLEAN          NOT NULL DEFAULT FALSE,
  "energy_reactive_v_ar_ac_phase_1_consumed" INTEGER          NOT NULL,
  "energy_reactive_v_ar_ac_phase_1_produced" INTEGER          NOT NULL,
  "energy_reactive_v_ar_ac_phase_2_consumed" INTEGER,
  "energy_reactive_v_ar_ac_phase_2_produced" INTEGER,
  "energy_reactive_v_ar_ac_phase_3_consumed" INTEGER,
  "energy_reactive_v_ar_ac_phase_3_produced" INTEGER,
  "energy_reactive_v_ar_ac_sum_consumed"     INTEGER          NOT NULL,
  "energy_reactive_v_ar_ac_sum_produced"     INTEGER          NOT NULL,
  "energy_real_w_ac_minus_absolute"          INTEGER          NOT NULL,
  "energy_real_w_ac_phase_1_consumed"        INTEGER          NOT NULL,
  "energy_real_w_ac_phase_1_produced"        INTEGER          NOT NULL,
  "energy_real_w_ac_phase_2_consumed"        INTEGER,
  "energy_real_w_ac_phase_2_produced"        INTEGER,
  "energy_real_w_ac_phase_3_consumed"        INTEGER,
  "energy_real_w_ac_phase_3_produced"        INTEGER,
  "energy_real_w_ac_plus_absolute"           INTEGER          NOT NULL,
  "energy_real_w_ac_sum_consumed"            INTEGER          NOT NULL,
  "energy_real_w_ac_sum_produced"            INTEGER          NOT NULL,
  "frequency_phase_average"                  DOUBLE PRECISION NOT NULL,
  "meter_location_current"                   INTEGER          NOT NULL,
  "power_apparent_s_phase_1"                 DOUBLE PRECISION NOT NULL,
  "power_apparent_s_phase_2"                 DOUBLE PRECISION,
  "power_apparent_s_phase_3"                 DOUBLE PRECISION,
  "power_apparent_s_sum"                     DOUBLE PRECISION NOT NULL,
  "power_factor_phase_1"                     DOUBLE PRECISION NOT NULL,
  "power_factor_phase_2"                     DOUBLE PRECISION,
  "power_factor_phase_3"                     DOUBLE PRECISION,
  "power_factor_sum"                         DOUBLE PRECISION NOT NULL,
  "power_reactive_q_phase_1"                 DOUBLE PRECISION NOT NULL,
  "power_reactive_q_phase_2"                 DOUBLE PRECISION,
  "power_reactive_q_phase_3"                 DOUBLE PRECISION,
  "power_reactive_q_sum"                     DOUBLE PRECISION NOT NULL,
  "power_real_p_phase_1"                     DOUBLE PRECISION NOT NULL,
  "power_real_p_phase_2"                     DOUBLE PRECISION,
  "power_real_p_phase_3"                     DOUBLE PRECISION,
  "power_real_p_sum"                         DOUBLE PRECISION NOT NULL,
  "timestamp"                                TIMESTAMP        NOT NULL,
  "visible"                                  BOOLEAN          NOT NULL DEFAULT FALSE,
  "voltage_ac_phase_1"                       DOUBLE PRECISION NOT NULL,
  "voltage_ac_phase_2"                       DOUBLE PRECISION,
  "voltage_ac_phase_3"                       DOUBLE PRECISION,
  PRIMARY KEY ("id")
);

CREATE TABLE "fronius"."current_powerflow_site" (
  "id"                        SERIAL8,
  "battery_standby"           BOOLEAN,
  "backup_mode"               BOOLEAN,
  "power_from_grid"           DOUBLE PRECISION,
  "power_load"                DOUBLE PRECISION,
  "power_akku"                DOUBLE PRECISION,
  "power_from_pv"             DOUBLE PRECISION,
  "relative_self_consumption" INTEGER,
  "relative_autonomy"         INTEGER,
  "meter_location"            VARCHAR(16) NOT NULL,
  "energy_day"                INTEGER,
  "energy_year"               INTEGER,
  "energy_total"              INTEGER,
  "created_at"                TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "pk_current_powerflow_site" PRIMARY KEY ("id")
);

CREATE TABLE "fronius"."current_powerflow_inverter" (
  "id"            SERIAL8,
  "id_site"       INT8      NOT NULL,
  "battery_mode"  VARCHAR(32),
  "device_type"   INTEGER   NOT NULL,
  "energy_day"    INTEGER,
  "energy_year"   INTEGER,
  "energy_total"  INTEGER,
  "current_power" INTEGER,
  "soc"           INTEGER,
  "created_at"    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "pk_current_powerflow_inverter" PRIMARY KEY ("id"),
  CONSTRAINT "fk_id_site" FOREIGN KEY ("id_site") REFERENCES "fronius"."current_powerflow_site" ("id")
  ON DELETE CASCADE
  ON UPDATE CASCADE
);

CREATE TABLE "fronius"."current_powerflow_ohmpilot" (
  "id"             SERIAL8,
  "id_site"        INT8             NOT NULL,
  "power_ac_total" INTEGER          NOT NULL,
  "state"          VARCHAR(32)      NOT NULL,
  "temperature"    DOUBLE PRECISION NOT NULL,
  "created_at"     TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "pk_current_powerflow_ohmpilot" PRIMARY KEY ("id"),
  CONSTRAINT "fk_id_site" FOREIGN KEY ("id_site") REFERENCES "fronius"."current_powerflow_site" ("id")
  ON DELETE CASCADE
  ON UPDATE CASCADE
);

CREATE TABLE "fronius"."current_data_inverter" (
  "id"           SERIAL8,
  "day_energy"   INTEGER   NOT NULL,
  "pac"          INTEGER   NOT NULL,
  "total_energy" INTEGER   NOT NULL,
  "year_energy"  INTEGER   NOT NULL,
  "timestamp"    timestamp NOT NULL,
  "created_at"   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "pk_current_data_inverter" PRIMARY KEY ("id")
);
