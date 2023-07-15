CREATE DATABASE IF NOT EXISTS `data` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `data`;

CREATE TABLE IF NOT EXISTS `menu` (
  `id` int PRIMARY KEY,
  `parent_id` int,
  `title` varchar(255) CHARACTER SET utf8mb4,
  `name` varchar(255) CHARACTER SET utf8mb4,
  `route` varchar(255) CHARACTER SET utf8mb4,
  `icon` varchar(255) CHARACTER SET utf8mb4,
  `is_children` boolean
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `menu` (`id`, `parent_id`, `title`, `name`, `route`, `icon`, `is_children`) VALUES
(1, 0, '', 'ข้อมูลสรุป', '/dashboards', 'fi-rr-chart-pie-alt', false),
(2, 0, '', 'ข้อมูลสายทาง', '/roads', 'fi-rr-map', true),
(3, 0, '', 'อนุมัติข้อมูล', '/approves', 'fi-rr-list-check', false),
(4, 0, '', 'ติดตามการซ่อมบำรุง', '/maintenances', 'fi fi-rr-pulse', true),
(7, 0, '', 'ตั้งค่าระบบ', '/settings', 'fi-rr-settings', false),
(8, 0, '', 'ผู้ใช้งาน', '/users', 'fi fi-rr-users', false),
(16, 1, '', 'สินทรัพย์', '/dashboards/road-asset', '', true),
(17, 1, '', 'ผิวทาง', '/dashboards/road-surface', '', true),
(18, 1, '', 'สภาพทาง', '/dashboards/road-condition', '', true),
(20, 3, '', 'สินทรัพย์ในเขตทาง', '/approves/in-asset', '', true),
(21, 3, '', 'สินทรัพย์นอกเขตทาง', '/approves/out-asset', '', true),
(22, 3, '', 'ผิวทาง', '/approves/surface', '', true),
(23, 3, '', 'สภาพทาง', '/approves/condition', '', true),
(24, 3, '', 'ความเสียหาย', '/approves/damage', '', true),
(34, 3, '', 'ปริมาณจราจร', '/approves/traffic', '', true),
(35, 3, '', 'อุบัติเหตุ', '/approves/accident', '', true),
(25, 7, '', 'แผนก', '/settings/departments', '', true),
(26, 7, '', 'กลุ่มสินทรัพย์', '/settings/asset-groups', '', true),
(27, 7, '', 'สินทรัพย์ในเขตทาง', '/settings/in-assets', '', true),
(28, 7, '', 'สินทรัพย์นอกเขตทาง', '/settings/out-assets', '', true),
(29, 7, '', 'เครื่องหมายจราจร', '/settings/traffic-signs', '', true),
(30, 7, '', 'เกณฑ์การจำแนกสภาพทาง', '/settings/survey-rules', '', true),
(36, 7, '', 'ชนิดผิวทาง', '/settings/surfaces', '', true),
(37, 7, '', 'ประเภทงบประมาณ', '/settings/budgets', '', true),
(38, 7, '', 'ตั้งค่า Model', '/settings/models', '', true),
(32, 8, '', 'ผู้ใช้งานภายในระบบ', '/users/users', '', true),
(33, 8, '', 'ตั้งค่าสิทธิ์ผู้ใช้งาน', '/users/roles', '', true);
