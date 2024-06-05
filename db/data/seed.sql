-------------------------------------------------------------------------------------------------------
INSERT INTO courses (
	code, name, underline, image, image_cover, excerpt, description
) VALUES ('GEO', 'Curso de Geografia', 'underline geo', 'image geo', 'image_cover geo', 'excerpt geo', 'description geo');


INSERT INTO courses  (
	code, name, underline, image, image_cover, excerpt, description
) VALUES ('CCO', 'Curso de Computacao', 'underline cco', 'image  cco', 'image_cover  cco', 'excerpt  cco', 'description  cco');


INSERT INTO courses (
	code, name, underline, image, image_cover, excerpt, description
) VALUES ('BIO', 'Curso de Biologia', 'underline bio', 'image bio', 'image_cover bio', 'excerpt bio', 'description bio');


INSERT INTO courses (
	code, name, underline, image, image_cover, excerpt, description
) VALUES ('MTM', 'Curso de Matemática', 'underline mtm', 'image mtm', 'image_cover mtm', 'excerpt mtm', 'description mtm');

-------------------------------------------------------------------------------------------------------

INSERT INTO matrices (
	code, "name", description, course_id
) VALUES('MGEO', 'Matriz de Geografia', 'description', 1);


INSERT INTO matrices (
	code, "name", description, course_id
) VALUES('MCCO', 'Matriz de Computação', 'description', 2);

-------------------------------------------------------------------------------------------------------

INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SGEO-1', 'Materia de Geografia 1', 'objective GEO 1', 11, 12, NOW() + INTERVAL '10 DAY');

INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SGEO-2', 'Materia de Geografia 2', 'objective GEO 2', 21, 22, NOW() + INTERVAL '20 DAY');

INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SGEO-3', 'Materia de Geografia 3', 'objective GEO 3', 31, 32, NOW() + INTERVAL '30 DAY');


INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SCCO-1', 'Materia de Computação 1', 'objective CCO 1', 11, 12, NOW() + INTERVAL '10 DAY');

INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SCCO-2', 'Materia de Computação 2', 'objective CCO 2', 21, 22, NOW() + INTERVAL '20 DAY');


INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SBIO-1', 'Materia de Biologia 1', 'objective BIO 1', 11, 12, NOW() + INTERVAL '10 DAY');

INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SBIO-2', 'Materia de Biologia 2', 'objective BIO 2', 21, 22, NOW() + INTERVAL '20 DAY');


INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SMTM-1', 'Materia de Matemática 1', 'objective MTM 1', 11, 12, NOW() + INTERVAL '10 DAY');

INSERT INTO subjects (
	code, "name", objective, credit, workload, published_at 
) VALUES('SMTM-2', 'Materia de Matemática 2', 'objective MTM 2', 21, 22, NOW() + INTERVAL '20 DAY');

-------------------------------------------------------------------------------------------------------

INSERT INTO public.matrix_subjects (
    matrix_id, subject_id, is_required, "group"
) VALUES(1, 1, true, '1SEM');

INSERT INTO public.matrix_subjects (
    matrix_id, subject_id, is_required, "group"
) VALUES(1, 2, false, '1SEM');

INSERT INTO public.matrix_subjects (
    matrix_id, subject_id, is_required, "group"
) VALUES(1, 3, true, '2SEM');

INSERT INTO public.matrix_subjects (
    matrix_id, subject_id, is_required, "group"
) VALUES(2, 4, true, '1SEM');

INSERT INTO public.matrix_subjects (
    matrix_id, subject_id, is_required, "group"
) VALUES(2, 5, false, '1SEM');

-------------------------------------------------------------------------------------------------------

INSERT INTO public.subscriptions (
	user_uuid, course_id, matrix_id, "role"
) VALUES('00000000-0000-0000-0001-000000000001', 1, 1, 'teacher');

INSERT INTO public.subscriptions (
	user_uuid, course_id, matrix_id, "role"
) VALUES('00000000-0000-0000-0001-000000000020', 1, 1, 'student');

INSERT INTO public.subscriptions (
	user_uuid, course_id, matrix_id, "role"
) VALUES('00000000-0000-0000-0001-000000000021', 1, 1, 'student');


INSERT INTO public.subscriptions (
	user_uuid, course_id, matrix_id, "role"
) VALUES('00000000-0000-0000-0002-000000000001', 2, 2, 'teacher');

INSERT INTO public.subscriptions (
	user_uuid, course_id, matrix_id, "role"
) VALUES('00000000-0000-0000-0002-000000000020', 2, 2, 'student');

INSERT INTO public.subscriptions (
	user_uuid, course_id, matrix_id, "role"
) VALUES('00000000-0000-0000-0002-000000000021', 2, 2, 'student');


INSERT INTO public.subscriptions (
	user_uuid, course_id, "role"
) VALUES('00000000-0000-0000-0003-000000000001', 3,'teacher');

INSERT INTO public.subscriptions (
	user_uuid, course_id, "role"
) VALUES('00000000-0000-0000-0003-000000000020', 3,'student');

INSERT INTO public.subscriptions (
	user_uuid, course_id, "role"
) VALUES('00000000-0000-0000-0003-000000000021', 3,'student');


INSERT INTO public.subscriptions (
	user_uuid, course_id, "role"
) VALUES('00000000-0000-0000-0004-000000000001', 4,'teacher');

INSERT INTO public.subscriptions (
	user_uuid, course_id, "role"
) VALUES('00000000-0000-0000-0004-000000000020', 4,'student');

INSERT INTO public.subscriptions (
	user_uuid, course_id, "role"
) VALUES('00000000-0000-0000-0004-000000000021', 4,'student');
-------------------------------------------------------------------------------------------------------