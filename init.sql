-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               PostgreSQL 17.4 on x86_64-windows, compiled by msvc-19.42.34436, 64-bit
-- Server OS:                    
-- HeidiSQL Version:             12.10.0.7000
-- --------------------------------------------------------
-- Dumping data for table public.leave_balances: 1 rows
/*!40000 ALTER TABLE "leave_balances" DISABLE KEYS */;
INSERT INTO "leave_balances" ("id", "created_at", "updated_at", "deleted_at", "user_id", "annual_leave", "sick_leave", "unpaid_leave") VALUES
	(1, '2025-04-11 23:36:37.487483+07', '2025-04-11 23:36:37.487483+07', NULL, 1, 15, 0, 0);
/*!40000 ALTER TABLE "leave_balances" ENABLE KEYS */;

-- Dumping data for table public.profiles: 3 rows
/*!40000 ALTER TABLE "profiles" DISABLE KEYS */;
INSERT INTO "profiles" ("id", "created_at", "updated_at", "deleted_at", "user_id", "national_id", "birth_place", "birth_date", "address", "civilization_level") VALUES
	(1, '2025-04-11 23:30:59.128686+07', '2025-04-11 23:36:21.002288+07', NULL, 1, '3713061208930011', 'Jakarta', '1993-01-01', 'Jl. Merdeka 10', NULL);
/*!40000 ALTER TABLE "profiles" ENABLE KEYS */;

-- Dumping data for table public.users: 3 rows
/*!40000 ALTER TABLE "users" DISABLE KEYS */;
INSERT INTO "users" ("id", "email", "name", "created_at", "updated_at", "deleted_at", "password") VALUES
	(1, 'irwan.rosyadi12@gmail.com', 'Mamarose', '2025-04-11 23:30:59.127469+07', '2025-04-11 23:36:21.001016+07', NULL, '$2a$12$qdsHiUi/.yNDw6P7sq8Jke5Iid4hS/j5RYZjwmpjKiPahqEfY1vE6');
/*!40000 ALTER TABLE "users" ENABLE KEYS */;