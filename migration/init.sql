DROP TABLE IF EXISTS correspondre;
DROP TABLE IF EXISTS reaction;
DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS fil_discussion;
DROP TABLE IF EXISTS categorie_jeu;
DROP TABLE IF EXISTS utilisateur;

CREATE TABLE utilisateur (
    id_user TEXT PRIMARY KEY,
    nom_utilisateur TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    mot_passe_hashe TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    est_banni INTEGER NOT NULL DEFAULT 0, -- 0 = Faux, 1 = Vrai -> kyky
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categorie_jeu (
    id_jeu TEXT PRIMARY KEY,
    nom_jeu TEXT NOT NULL,
    genre TEXT NOT NULL
);

CREATE TABLE fil_discussion (
    id_fil TEXT PRIMARY KEY,
    titre TEXT NOT NULL,
    description TEXT NOT NULL,
    etat TEXT NOT NULL DEFAULT 'ouvert', 
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    id_user TEXT NOT NULL,
    FOREIGN KEY (id_user) REFERENCES utilisateur(id_user) ON DELETE CASCADE
);

CREATE TABLE message (
    id_message TEXT PRIMARY KEY,
    contenu TEXT NOT NULL,
    date_envoi DATETIME DEFAULT CURRENT_TIMESTAMP,
    id_fil TEXT NOT NULL,
    id_user TEXT NOT NULL,
    score INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (id_fil) REFERENCES fil_discussion(id_fil) ON DELETE CASCADE,
    FOREIGN KEY (id_user) REFERENCES utilisateur(id_user) ON DELETE CASCADE
);

CREATE TABLE reaction (
    id_reaction TEXT PRIMARY KEY,
    type_reaction TEXT NOT NULL, -- 'like' ou 'dislike' en gros
    id_user TEXT NOT NULL,
    id_message TEXT NOT NULL,
    FOREIGN KEY (id_user) REFERENCES utilisateur(id_user) ON DELETE CASCADE,
    FOREIGN KEY (id_message) REFERENCES message(id_message) ON DELETE CASCADE,
    UNIQUE(id_user, id_message) -- ca empêche un utilisateur de réagir plusieurs fois au même message
);

CREATE TABLE correspondre (
    id_fil TEXT NOT NULL,
    id_jeu TEXT NOT NULL,
    PRIMARY KEY (id_fil, id_jeu),
    FOREIGN KEY (id_fil) REFERENCES fil_discussion(id_fil) ON DELETE CASCADE,
    FOREIGN KEY (id_jeu) REFERENCES categorie_jeu(id_jeu) ON DELETE CASCADE
);

-- qulques jeux de base pour tester le matchmaking 
INSERT INTO categorie_jeu (id_jeu, nom_jeu, genre) VALUES 
('1', 'Apex Legends', 'FPS / Battle Royale'),
('2', 'League of Legends', 'MOBA'),
('3', 'Valorant', 'FPS / Tactique'),
('4', 'Rocket League', 'Sport / Arcade');