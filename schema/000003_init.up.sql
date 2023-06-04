CREATE TABLE sessions (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE
);

--Inserting data for testing (need to be uncommented for inserting)
INSERT INTO "public"."reqs" ("title") VALUES ('Потреба в психоемоційній підтримці');
INSERT INTO "public"."reqs" ("title") VALUES ('Проблеми з емоційною регуляцією');
INSERT INTO "public"."reqs" ("title") VALUES ('Переживання/страх майбутнього');
INSERT INTO "public"."reqs" ("title") VALUES ('Труднощі в батьківсько-дитячих стосунках');
INSERT INTO "public"."reqs" ("title") VALUES ('Конфлікти, труднощі з комунікацією');
INSERT INTO "public"."reqs" ("title") VALUES ('Переживання втрати, горя');
INSERT INTO "public"."reqs" ("title") VALUES ('Пригнічений настрій, депресивні стани');
INSERT INTO "public"."reqs" ("title") VALUES ('Почуття самотності');
INSERT INTO "public"."reqs" ("title") VALUES ('Інша проблема (вказати у примітках)');
INSERT INTO "public"."reqs" ("title") VALUES ('Кризова інтервенція, стабілізація, ППД');
INSERT INTO "public"."organisations" ("title", "code") VALUES ('Карітас Хмельницький', 'KML');
INSERT INTO "public"."users" ("login", "pass", "first_name", "last_name", "email", "phone", "org_id", "role") VALUES ('mariia','mariia','Марія','Іванько','mivanko@gmail.com','+380671234566','2','2');
INSERT INTO "public"."users" ("login", "pass", "first_name", "last_name", "email", "phone", "org_id", "role") VALUES ('ivan','ivan','Іван','Похилько','ipohylko@gmail.com','+380501234456','2','2');
INSERT INTO "public"."users" ("login", "pass", "first_name", "last_name", "email", "phone", "org_id", "role") VALUES ('petro','petro','Петро','Хоньків','phonkiv@gmail.com','+380931224565','2','2');
INSERT INTO "public"."projects" ("short_title", "code") VALUES ('Bridging Divides','BD');
INSERT INTO "public"."projects" ("short_title", "code") VALUES ('Caritas Norwey','CN');
INSERT INTO "public"."beneficiaries" ("first_name", "middle_name", "last_name", "phone", "birthday", "prj_id", "org_id", "user_id") VALUES ('Олексій','Іванович','Завгородній','+380501231129','19.05.1989','1','2','2');
INSERT INTO "public"."beneficiaries" ("first_name", "middle_name", "last_name", "phone", "birthday", "prj_id", "org_id", "user_id") VALUES ('Олексій','Андрійович','Потапенко','+380501231129','08.05.1981','1','2','2');
INSERT INTO "public"."beneficiaries" ("first_name", "middle_name", "last_name", "phone", "birthday", "prj_id", "org_id", "user_id") VALUES ('Вадим','Олексійович','Федоров','+380631231125','03.02.1980','1','2','2');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('23.03.2023','2','1','2','Проблеми з музикальним слухом');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('24.03.2023','3','1','2','Сумніви у вдалості обраного музикального напряму');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('25.03.2023','2','1','2','Сумує, що розпалася музикальна група');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('26.03.2023','1','2','2','Проблеми із самооцінкою');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('27.03.2023','4','2','2','Почуває себе зайвим');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('28.03.2023','2','2','2','Вважає свою творчість безцінним надбанням українського нароуд');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('29.03.2023','3','3','2','Хоче зрозуміти хто він');
INSERT INTO "public"."activities" ("time", "req", "bnf_id", "user_id", "description") VALUES ('30.03.2023','4','3','2','Мені надоїло придумувати ци тексти');