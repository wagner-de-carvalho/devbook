insert into usuarios (nome, nick , email, senha) 
values 
("Paulo Santos", "Paulin", "paulin@mail.com", "$2a$10$mQSR1wSaGlCM89oU9jvDuOlHACtOtXC16Ey4BGL5ZkhyX65CdumSW" ),
("Maria Prado", "mamaia", "pradomaia@mail.com", "$2a$10$mQSR1wSaGlCM89oU9jvDuOlHACtOtXC16Ey4BGL5ZkhyX65CdumSW"  ),
("Sirlene Tavares", "silinha", "lenetavares@mail.com", "$2a$10$mQSR1wSaGlCM89oU9jvDuOlHACtOtXC16Ey4BGL5ZkhyX65CdumSW" ),
("Rodrigo Araújo", "diguin", "diguin_ara@mail.com", "$2a$10$mQSR1wSaGlCM89oU9jvDuOlHACtOtXC16Ey4BGL5ZkhyX65CdumSW" ),
("Carla Morgado", "Mora", "carla.mor@mail.com", "$2a$10$mQSR1wSaGlCM89oU9jvDuOlHACtOtXC16Ey4BGL5ZkhyX65CdumSW" );

insert into seguidores (usuario_id, seguidor_id) 
values
(1, 2),
(3, 1),
(1, 3),
(2, 4),
(4, 2),
(5, 1),
(5, 3),
(3, 5),
(4, 1),
(5, 2), 
(5, 4), 
(4, 5), 
(1, 5), 
(3, 2);

insert into publicacoes (titulo, conteudo, autor_id)
values
("Publicação do autor 1", "Esta é a publicação do autor 1", 1),
("outra Publicação do autor 1", "Esta é outra publicação do autor 1", 1),
("Publicação do autor 2", "Esta é a publicação do autor 2", 2),
("Outra Publicação do autor 2", "Esta é outra publicação do autor 2", 2),
("Publicação do autor 3", "Esta é a publicação do autor 3", 3),
("OutraPublicação do autor 3", "Esta é outra publicação do autor 3", 3);