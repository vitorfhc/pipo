alter table "public"."employee_info" add constraint "employee_info_employee_cpf_info_uuid_key" unique ("employee_cpf", "info_uuid");
