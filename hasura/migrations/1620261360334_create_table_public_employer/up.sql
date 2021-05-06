CREATE TABLE "public"."employer"("uuid" uuid NOT NULL, "cnpj" text NOT NULL, "name" text NOT NULL, "password" text NOT NULL, PRIMARY KEY ("uuid") , UNIQUE ("name") , UNIQUE("cnpj"));
