CREATE TABLE version (
  id varchar(64) NOT NULL,
  service_id uuid NOT NULL,  
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  PRIMARY KEY (service_id, id),
  FOREIGN KEY (service_id) REFERENCES service (id) ON DELETE CASCADE
);

INSERT INTO version (service_id, id) VALUES 
  ('14f777b2-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f777b2-997e-11ec-b909-0242ac120002', '0.2'),
  ('14f777b2-997e-11ec-b909-0242ac120002', '0.3'),
  ('14f777b2-997e-11ec-b909-0242ac120002', '1.0'),
  ('14f79454-997e-11ec-b909-0242ac120002', '1.1'),
  ('14f79454-997e-11ec-b909-0242ac120002', '1.2'),
  ('14f79454-997e-11ec-b909-0242ac120002', '2.0'),
  ('14f79454-997e-11ec-b909-0242ac120002', '2.1'),
  ('14f795f8-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f795f8-997e-11ec-b909-0242ac120002', '0.2'),
  ('14f795f8-997e-11ec-b909-0242ac120002', '0.3'),
  ('14f79828-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f799a4-997e-11ec-b909-0242ac120002', '1.0'),
  ('14f79b20-997e-11ec-b909-0242ac120002', '1.0'),
  ('14f79c92-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f79f58-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f79f58-997e-11ec-b909-0242ac120002', '1.0'),
  ('14f79f58-997e-11ec-b909-0242ac120002', '2.0'),
  ('14f7a106-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f7a2b4-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f7a700-997e-11ec-b909-0242ac120002', '0.1'),
  ('14f7a912-997e-11ec-b909-0242ac120002', '0.1')
;