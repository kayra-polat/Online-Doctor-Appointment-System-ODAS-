-- PATIENTS --
insert into users(name, surname, email, username, gender, phone, address, date_of_birth, password_hash, role)
VALUES ('Joe', 'R.', 'joe@gmail.com', 'joer', 'Male', '05544567572', 'Ankara',
        date '2002-10-05', '$2a$16$geSKx/HuTySDuNQF7jEMh.1NHyJI20ZShTvg4SNujbWBWSUleAUe2', 'Patient');

insert into users(name, surname, email, username, gender, phone, address, date_of_birth, password_hash, role)
VALUES ('George', 'B.', 'george@gmail.com', 'georgeb', 'Male', '05544567572', 'Damascus',
        date '1945-10-05', '$2a$16$geSKx/HuTySDuNQF7jEMh.1NHyJI20ZShTvg4SNujbWBWSUleAUe2', 'Patient');


-- DOCTORS --
insert into users(name, surname, email, username, gender, phone, address, date_of_birth, password_hash, role,
                  title, expertise, hospital, room)
VALUES ('Mary', 'K.', 'mary@gmail.com', 'maryk', 'Female', '+90 5314554655', 'Istanbul',
        date '1932-10-05', '$2a$16$geSKx/HuTySDuNQF7jEMh.1NHyJI20ZShTvg4SNujbWBWSUleAUe2', 'Doctor', 'Prof. Dr.', 'Pediatrics', 'Hospital A', 'A02');

insert into users(name, surname, email, username, gender, phone, address, date_of_birth, password_hash, role,
                  title, expertise, hospital, room)
VALUES ('Angela', 'L.', 'angela@gmail.com', 'angelal', 'Female', '+61 5753844655', 'Canberra',
        date '1976-10-05', '$2a$16$geSKx/HuTySDuNQF7jEMh.1NHyJI20ZShTvg4SNujbWBWSUleAUe2', 'Doctor', 'Dr.', 'Dermatology', 'Hospital B', 'B01');

insert into available_times( doctor_id, date, hospital, room)
VALUES (1,'2021-12-07T13:42:31+03:00', 'City Hospital', 'O870');