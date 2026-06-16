const API_URL = "http://localhost:8080/api";

document.addEventListener("DOMContentLoaded", () => {

    checkAuthStatus();
    
    if (document.getElementById("categories-list")) loadCategories();
    if (document.getElementById("threads-list")) loadThreads();

    const formLogin = document.getElementById("form-login");
    if (formLogin) {
        formLogin.addEventListener("submit", async (e) => {
            e.preventDefault();
            const email = document.getElementById("login-email").value;
            const password = document.getElementById("login-password").value;
            await handleLogin(email, password);
        });
    }

    const formRegister = document.getElementById("form-register");
    if (formRegister) {
        formRegister.addEventListener("submit", async (e) => {
            e.preventDefault();
            const pseudo = document.getElementById("reg-pseudo").value;
            const email = document.getElementById("reg-email").value;
            const password = document.getElementById("reg-password").value;
            await handleRegister(pseudo, email, password);
        });
    }
});


// Gérer la connexion
async function handleLogin(email, password) {
    try {
        const response = await fetch(`${API_URL}/login`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ identifiant: email, mot_passe: password })
        });

        const data = await response.json();
        if (!response.ok) throw new Error(data.error || "Identifiants invalides");

        localStorage.setItem("token", data.token);
        localStorage.setItem("user_pseudo", data.user.nom_utilisateur);
        localStorage.setItem("user_id", data.user.id);

        alert(` Content de te revoir, ${data.user.nom_utilisateur} !`);
        window.location.href = "index.html"; 
    } catch (error) {
        alert(`Erreur : ${error.message}`);
    }
}

// Gérer l'inscription
async function handleRegister(pseudo, email, password) {
    try {
        const response = await fetch(`${API_URL}/register`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ nom_utilisateur: pseudo, email: email, mot_passe: password })
        });

        const data = await response.json();
        if (!response.ok) throw new Error(data.error || "Erreur lors de l'inscription");

        alert("Compte créé avec succès ! Connecte-toi maintenant.");
        window.location.reload(); 
    } catch (error) {
        alert(`Erreur : ${error.message}`);
    }
}

function checkAuthStatus() {
    const token = localStorage.getItem("token");
    const pseudo = localStorage.getItem("user_pseudo");
    const navAuth = document.getElementById("nav-auth");
    const btnNewThread = document.getElementById("btn-new-thread");

    if (token && navAuth) {
        navAuth.innerHTML = `
            <span style="margin-right: 1rem; color: #949ba4;">👤 ${pseudo}</span>
            <button onclick="handleLogout()" class="btn" style="background-color: #ff4757;">Déconnexion</button>
        `;
        if (btnNewThread) btnNewThread.classList.remove("hidden");
    }
}

function handleLogout() {
    localStorage.clear();
    window.location.reload();
}


// Charger les catégories de jeux depuis l'API
async function loadCategories() {
    const listContainer = document.getElementById("categories-list");
    try {
        const response = await fetch(`${API_URL}/categories`);
        const categories = await response.json();
        listContainer.innerHTML = "";

        categories.forEach(cat => {
            const li = document.createElement("li");
            li.innerHTML = `<strong>${cat.nom_jeu}</strong> <br><small style="color: #949ba4;">${cat.genre}</small>`;
            listContainer.appendChild(li);
        });
    } catch (error) {
        listContainer.innerHTML = "<li>Erreur de chargement</li>";
    }
}

// Charger toutes les annonces depuis l'API
async function loadThreads() {
    const threadsContainer = document.getElementById("threads-list");
    try {
        const response = await fetch(`${API_URL}/threads`);
        const threads = await response.json();
        threadsContainer.innerHTML = "";

        if (!threads || threads.length === 0) {
            threadsContainer.innerHTML = "<p style='color: #949ba4;'>Aucune annonce publiée pour le moment.</p>";
            return;
        }

        threads.forEach(thread => {
            const card = document.createElement("div");
            card.className = "thread-card";
            card.innerHTML = `
                <h3>${thread.titre}</h3>
                <p>${thread.description}</p>
                <div style="margin-top: 1rem; display:flex; justify-content:space-between; font-size:0.85rem; color:#949ba4;">
                    <span>Statut : 🟢 ${thread.etat}</span>
                    <button class="btn" style="padding: 0.3rem 0.6rem; font-size:0.8rem;" onclick="loadMessages('${thread.id}')">Voir les réponses</button>
                </div>
                <div id="messages-container-${thread.id}" style="margin-top:1rem; padding-top:1rem; border-top:1px solid #2f3342;" class="hidden">
                    <div id="messages-list-${thread.id}"></div>
                    ${localStorage.getItem("token") ? `
                        <div style="margin-top:1rem; display:flex; gap:0.5rem;">
                            <input type="text" id="input-msg-${thread.id}" placeholder="Propose ton Discord..." style="flex:1; padding:0.5rem; background:#1a1c23; border:1px solid #2f3342; border-radius:4px; color:white;">
                            <button class="btn" style="padding:0.5rem;" onclick="postMessage('${thread.id}')">Répondre</button>
                        </div>
                    ` : ''}
                </div>
            `;
            threadsContainer.appendChild(card);
        });
    } catch (error) {
        threadsContainer.innerHTML = "<p>Erreur lors du chargement des annonces</p>";
    }
}


// Charger les messages d'un thread la
async function loadMessages(threadId) {
    const container = document.getElementById(`messages-container-${threadId}`);
    const list = document.getElementById(`messages-list-${threadId}`);
    
    if (!container.classList.contains("hidden")) {
        container.classList.add("hidden");
        return;
    }
    container.classList.remove("hidden");

    try {
        const response = await fetch(`${API_URL}/threads/${threadId}/messages`);
        const messages = await response.json();
        list.innerHTML = "";

        if (!messages || messages.length === 0) {
            list.innerHTML = "<p style='font-size:0.9rem; color:#949ba4;'>Aucune réponse pour le moment.</p>";
            return;
        }

        messages.forEach(msg => {
            const div = document.createElement("div");
            div.style = "background:#1f212a; padding:0.75rem; border-radius:6px; margin-bottom:0.5rem; display:flex; justify-content:space-between; align-items:center;";
            div.innerHTML = `
                <div>
                    <p style="font-size:0.95rem;">${msg.contenu}</p>
                    <small style="color:#949ba4;">Par Joueur #${msg.user_id.substring(0,5)}</small>
                </div>
                <div style="display:flex; align-items:center; gap:0.5rem;">
                    <button onclick="voteMessage('${msg.id}', 'like', '${threadId}')" style="background:none; border:none; cursor:pointer;">👍</button>
                    <span style="font-weight:bold; color:${msg.score >= 0 ? '#2ed573' : '#ff4757'}">${msg.score}</span>
                    <button onclick="voteMessage('${msg.id}', 'dislike', '${threadId}')" style="background:none; border:none; cursor:pointer;">👎</button>
                </div>
            `;
            list.appendChild(div);
        });
    } catch (error) {
        list.innerHTML = "<p>Impossible de charger les messages</p>";
    }
}

// Envoyer un nouveau message de réponse 
async function postMessage(threadId) {
    const input = document.getElementById(`input-msg-${threadId}`);
    const contenu = input.value.trim();
    if (!contenu) return;

    try {
        const response = await fetch(`${API_URL}/messages`, {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem("token")}`
            },
            body: JSON.stringify({ contenu: contenu, thread_id: threadId })
        });

        if (!response.ok) throw new Error("Échec de l'envoi");
        
        input.value = "";
    
        container = document.getElementById(`messages-container-${threadId}`);
        container.classList.add("hidden");
        await loadMessages(threadId);
    } catch (error) {
        alert("Action impossible (vérifie ton authentification)");
    }
}

// Voter pour un message 
async function voteMessage(messageId, typeVote, threadId) {
    if (!localStorage.getItem("token")) {
        alert("Tu dois être connecté pour voter");
        return;
    }

    try {
        const response = await fetch(`${API_URL}/reactions`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem("token")}`
            },
            body: JSON.stringify({ type: typeVote, message_id: messageId })
        });

        if (!response.ok) throw new Error();
        
        container = document.getElementById(`messages-container-${threadId}`);
        container.classList.add("hidden");
        await loadMessages(threadId);
    } catch (error) {
        alert("Impossible d'enregistrer le vote");
    }
}

// Gérer l'affichage du formulaire et l'envoi de l'annonce
document.addEventListener("DOMContentLoaded", () => {
    const btnNew = document.getElementById("btn-new-thread");
    const modal = document.getElementById("modal-new-thread");
    const formThread = document.getElementById("form-thread");

    if (btnNew && modal) {
        btnNew.addEventListener("click", () => modal.classList.remove("hidden"));
    }

    if (formThread) {
        formThread.addEventListener("submit", async (e) => {
            e.preventDefault();
            const idJeu = document.getElementById("thread-jeu").value;
            const titre = document.getElementById("thread-titre").value;
            const desc = document.getElementById("thread-desc").value;

            try {
                const response = await fetch(`${API_URL}/threads`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": `Bearer ${localStorage.getItem("token")}`
                    },
                    body: JSON.stringify({ titre: titre, description: desc, id_jeu: idJeu })
                });

                if (!response.ok) throw new Error("Erreur de publication");
                
                alert("Annonce publiée !");
                window.location.reload();
            } catch (error) {
                alert("Impossible de publier. Es-tu bien connecté ?");
            }
        });
    }
});
