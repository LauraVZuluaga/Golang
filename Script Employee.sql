CREATE TABLE "tbl_employee"(
"employee_id" SERIAL,
"full_name" STRING(100),
"department" STRING(50),
"designation" STRING(50),
"created_at" TIMESTAMPTZ,
"update_at" TIMESTAMPTZ,
PRIMARY KEY ("employee_id")
);