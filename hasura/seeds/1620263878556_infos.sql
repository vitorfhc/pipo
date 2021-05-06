INSERT INTO public.info (uuid, name, regex) VALUES ('80d4670b-866f-44d9-8584-908e4ccbbc4a', 'Nome', NULL);
INSERT INTO public.info (uuid, name, regex) VALUES ('657e5605-bc40-41e1-924f-6461d79aa93e', 'CPF', '^\d{3}\.\d{3}\.\d{3}-\d{2}$');
INSERT INTO public.info (uuid, name, regex) VALUES ('7c37929e-d43d-4202-954d-b687366a2bae', 'Data de Admissão', '^(([0-2][0-9])|(3[01]))\/((0[1-9])|(1[0-2]))\/[12]\d{3}$');
INSERT INTO public.info (uuid, name, regex) VALUES ('1b127588-08a5-4328-b3b9-b6b63c492f46', 'Endereço', NULL);
INSERT INTO public.info (uuid, name, regex) VALUES ('18515458-a6a2-4e43-a22f-7203219a0782', 'Peso (kg)', '^\d{1,3}\.\d{1}$');
INSERT INTO public.info (uuid, name, regex) VALUES ('7d5be03a-0e5d-4feb-8d29-76169f3e3c39', 'Altura (cm)', '^\d{2,3}$');
INSERT INTO public.info (uuid, name, regex) VALUES ('755fb76e-7a7d-4267-b7c6-d7af26b7c814', 'Horas meditadas nos últimos 7 dias', '^\d{1,4}$');
INSERT INTO public.info (uuid, name, regex) VALUES ('70d0476b-f365-4c4a-a3d3-541f9f289688', 'Email', '^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$');
