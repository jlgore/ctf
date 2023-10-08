-- Customers table data
INSERT INTO customers (first_name, last_name, email, phone_number, address, city, state, zip_code)
VALUES 
('John', 'Doe', 'john.doe@example.com', '123-456-7890', '123 Elm Street', 'Anytown', 'CA', '12345'),
('Jane', 'Smith', 'jane.smith@example.com', '987-654-3210', '456 Oak Road', 'Sometown', 'NY', '67890'),
('Alice', 'Johnson', 'alice.johnson@example.com', '555-555-5555', '789 Pine Lane', 'Othertown', 'TX', '11223');

-- Panels table data
INSERT INTO panels (model, manufacturer, capacity)
VALUES 
('Model A', 'SolarCorp', 100.5),
('Model B', 'BrightSun', 150.75),
('Model C', 'EcoEnergy', 200.0);

-- Installations table data (assuming the IDs from the inserts above)
INSERT INTO installations (customer_id, panel_id, installation_date, warranty_expiry_date, status)
VALUES 
(1, 1, '2023-01-15', '2033-01-15', 'Completed'),
(2, 2, '2023-02-20', '2033-02-20', 'Pending'),
(3, 3, '2023-03-10', '2038-03-10', 'Completed');

-- Service records table data
INSERT INTO service_records (installation_id, service_date, service_notes)
VALUES 
(1, '2023-06-10', 'Routine check. All panels functioning optimally.'),
(3, '2023-04-05', 'Initial service after installation.');

-- ... You can continue to add more mock data as required for development purposes.

