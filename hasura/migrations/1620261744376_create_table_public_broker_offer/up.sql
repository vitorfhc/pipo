CREATE TABLE "public"."broker_offer"("uuid" uuid NOT NULL, "broker_cnpj" text NOT NULL, "offer_uuid" uuid NOT NULL, PRIMARY KEY ("uuid") , FOREIGN KEY ("broker_cnpj") REFERENCES "public"."broker"("cnpj") ON UPDATE restrict ON DELETE restrict, FOREIGN KEY ("offer_uuid") REFERENCES "public"."offer"("uuid") ON UPDATE restrict ON DELETE restrict);
