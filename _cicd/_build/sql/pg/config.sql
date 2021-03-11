BEGIN;

INSERT INTO Promotion(promotion_name, description, start_date, end_date)
VALUES ('Promo1', 'Rate < 10', '2020-01-01', '2020-03-31');
INSERT INTO Promotion(promotion_name, description, start_date, end_date)
VALUES ('Promo2', 'Rate > 10 < 2', '22020-04-01', '2020-06-30');
INSERT INTO Promotion(promotion_name, description, start_date, end_date)
VALUES ('Promo3', 'Rate > 20 < 28', '2020-07-01', '2020-12-30');

END;