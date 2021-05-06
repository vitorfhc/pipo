
ALTER TABLE ONLY "public"."offer" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."employee_info" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."employer_broker_offer" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."registration" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."info" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."broker_offer" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."offer_info" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();

ALTER TABLE ONLY "public"."employer" ALTER COLUMN "uuid" SET DEFAULT gen_random_uuid();
