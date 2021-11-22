-- +goose Up
-- SQL in this section is executed when the migration is applied.
INSERT INTO `user` (`id`, `email`, `password`, `role`, `id_merchant`, `id_outlet`, `created_at`, `updated_at`) VALUES (555,'a@a.id', '$2a$08$cwJLLr.LfnKjUdpW6C3kE.KnEhdcQiVXGXcMd3iAXJ9IgMgLDnOci', 'admin', NULL, NULL, '2021-11-21 00:00:00', '2021-11-21 00:00:00');

INSERT INTO `merchant` (`id`, `name`, `created_at`, `updated_at`) VALUES (888, 'INDOMART', '2021-11-21 00:00:00', '2021-11-21 00:00:00'), (999, 'ALFAMART', '2021-11-21 00:00:00', '2021-11-21 00:00:00');

INSERT INTO `outlet` (`id`, `id_merchant`, `name`, `created_at`, `updated_at`) VALUES (888001, 888, 'INDOMART PAKIS', '2021-11-21 00:00:00', '2021-11-21 00:00:00'), (999001, 999, 'ALFAMART TUMPANG', '2021-11-21 00:00:00', '2021-11-21 00:00:00');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DELETE FROM `outlet` WHERE `id` IN (999001, 888001);
DELETE FROM `merchant` WHERE `id` IN (999, 888);
DELETE FROM `user` WHERE `id` = 555;
