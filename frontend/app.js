const API_URL = "http://localhost:8080/api";

document.addEventListener("DOMContentLoaded", () => {
    checkAuthStatus();
    
    if (document.getElementById("categories-list")) loadCategories();
    if (document.getElementById("threads-list")) loadThreads();

    const formLogin = document.getElementById("form-login");
    if (formLogin) {
        formLogin.addEventListener("submit", async (e) => {
            e.preventDefault();
            await handleLogin(document.getElementById("login-email").value, document.getElementById("login-password").value);
        });
    }

    const formRegister = document.getElementById("form-register");
    if (formRegister) {
        formRegister.addEventListener("submit", async (e) => {
            e.preventDefault();
            await handleRegister(document.getElementById("reg-pseudo").value, document.getElementById("reg-email").value, document.getElementById("reg-password").value);
        });
    }

    const formThread = document.getElementById("form-thread");
    if (formThread) {
        formThread.addEventListener("submit", async (e) => {
            e.preventDefault();
        
            const idJeu = document.getElementById("thread-jeu").value;
            await createThread(document.getElementById("thread-titre").value, document.getElementById("thread-desc").value, idJeu);
        });
    }
});

async function handleLogin(email, password) {
    try {
        const res = await fetch(`${API_URL}/login`, { method: "POST", headers: {"Content-Type": "application/json"}, body: JSON.stringify({identifiant: email, mot_passe: password}) });
        const data = await res.json();
        if (!res.ok) throw new Error(data.error || "Identifiants invalides");
        localStorage.setItem("token", data.token);
        localStorage.setItem("user_pseudo", data.user.nom_utilisateur);
        window.location.href = "index.html";
    } catch (err) { alert(`err ${err.message}`); }
}

async function handleRegister(pseudo, email, password) {
    try {
        const res = await fetch(`${API_URL}/register`, { method: "POST", headers: {"Content-Type": "application/json"}, body: JSON.stringify({nom_utilisateur: pseudo, email: email, mot_passe: password}) });
        if (!res.ok) throw new Error("Erreur d'inscription. Pseudo ou email déjà pris.");
        alert("Compte créé avec succès ! Connecte-toi.");
        window.location.reload();
    } catch (err) { alert(`err ${err.message}`); }
}

function checkAuthStatus() {
    const token = localStorage.getItem("token");
    if (token && document.getElementById("nav-auth")) {
        document.getElementById("nav-auth").innerHTML = `
            <span style="color:white; margin-right:1rem; font-weight:bold;"> ${localStorage.getItem("user_pseudo")}</span>
            <button onclick="localStorage.clear(); window.location.reload();" class="btn" style="background:#ff4757;">Déconnexion</button>
        `;
        if (document.getElementById("btn-new-thread")) document.getElementById("btn-new-thread").classList.remove("hidden");
    }
}

async function loadCategories() {
    try {
        const res = await fetch(`${API_URL}/categories`);
        const cats = await res.json();
        const listContainer = document.getElementById("categories-list");
        const selectMenu = document.getElementById("thread-jeu");
        
        // Bouton pour réinitialiser le filtre
        listContainer.innerHTML = `<li style="cursor:pointer; color:var(--accent);" onclick="loadThreads()">Voir toutes les annonces</li>`;
        if (selectMenu) selectMenu.innerHTML = '<option value="">Choisis un jeu...</option>';

        cats.forEach(c => {
          
            const li = document.createElement("li");
            li.style.cursor = "pointer";
            li.innerHTML = `<strong>${c.nom_jeu}</strong><br><small style="color:#949ba4">${c.genre}</small>`;
            li.onclick = () => loadThreads(c.id);
            listContainer.appendChild(li);

            if (selectMenu) {
                selectMenu.innerHTML += `<option value="${c.id}">${c.nom_jeu}</option>`;
            }
        });
    } catch (err) { document.getElementById("categories-list").innerHTML = "<li>Erreur chargement</li>"; }
}

async function loadThreads(categoryId = null) {
    try {
        const res = await fetch(`${API_URL}/threads`);
        let threads = await res.json();
        
        if (categoryId) threads = threads.filter(t => t.id_jeu == categoryId);

        const container = document.getElementById("threads-list");
        if (!threads || threads.length === 0) {
            return container.innerHTML = "<p style='color:#949ba4;'>Aucune annonce pour ce salon.</p>";
        }
        
        container.innerHTML = threads.map(t => `
            <div class="thread-card">
                <h3>${t.titre}</h3>
                <p>${t.description}</p>
                <div style="margin-top:1rem; display:flex; justify-content:space-between; align-items:center;">
                    <span style="font-size:0.85rem; color:#949ba4;">Statut :  Ouvert</span>
                    <button class="btn" style="padding: 0.4rem 0.8rem;" onclick="loadMessages('${t.id}')">Voir les réponses</button>
                </div>
                <div id="msg-box-${t.id}" class="hidden" style="margin-top:1rem; border-top:1px solid var(--border); padding-top:1rem;">
                    <div id="msg-list-${t.id}"></div>
                    ${localStorage.getItem("token") ? `
                        <div style="display:flex; gap:0.5rem; margin-top:1rem;">
                            <input type="text" id="input-${t.id}" placeholder="Ton Discord ou ta réponse..." style="flex:1; padding:0.5rem; background:var(--bg-main); border:1px solid var(--border); color:white; border-radius:4px;">
                            <button class="btn" onclick="postMessage('${t.id}')">Envoyer</button>
                        </div>
                    ` : '<p style="font-size:0.85rem; color:#ff4757; margin-top:1rem;">Connecte-toi pour répondre.</p>'}
                </div>
            </div>
        `).join('');
    } catch (err) { console.log(err); }
}

async function createThread(titre, desc, idJeu) {
    try {
        const res = await fetch(`${API_URL}/threads`, { 
            method: "POST", 
            headers: {"Content-Type": "application/json", "Authorization": `Bearer ${localStorage.getItem("token")}`}, 
            body: JSON.stringify({titre: titre, description: desc, id_jeu: idJeu}) 
        });
        if (!res.ok) throw new Error();
        window.location.reload();
    } catch (err) { alert("Erreur de publication."); }
}

async function loadMessages(id) {
    const box = document.getElementById(`msg-box-${id}`);
    box.classList.toggle("hidden");
    
    if (!box.classList.contains("hidden")) {
        const res = await fetch(`${API_URL}/threads/${id}/messages`);
        const msgs = await res.json();
        const list = document.getElementById(`msg-list-${id}`);
        
        if (!msgs || msgs.length === 0) {
            list.innerHTML = '<p style="font-size:0.9rem; color:#949ba4;">Pas encore de réponses.</p>';
            return;
        }

        list.innerHTML = msgs.map(m => `
            <div style="background:#1f212a; padding:0.75rem; border-radius:6px; margin-bottom:0.5rem; display:flex; justify-content:space-between; align-items:center;">
                <div>
                    <p style="font-size:0.95rem;">${m.contenu}</p>
                </div>
                <div style="display:flex; align-items:center; gap:0.5rem;">
                    <button onclick="voteMessage('${m.id}', 'like', '${id}')" style="background:none; border:none; cursor:pointer;">👍</button>
                    <span style="font-weight:bold; color:${m.score >= 0 ? '#2ed573' : '#ff4757'}">${m.score || 0}</span>
                    <button onclick="voteMessage('${m.id}', 'dislike', '${id}')" style="background:none; border:none; cursor:pointer;">👎</button>
                </div>
            </div>
        `).join('');
    }
}

async function postMessage(threadId) {
    const contenu = document.getElementById(`input-${threadId}`).value;
    if (!contenu) return;
    await fetch(`${API_URL}/messages`, { 
        method: "POST", 
        headers: {"Content-Type": "application/json", "Authorization": `Bearer ${localStorage.getItem("token")}`}, 
        body: JSON.stringify({contenu: contenu, thread_id: threadId}) 
    });
    
    document.getElementById(`input-${threadId}`).value = "";
    document.getElementById(`msg-box-${threadId}`).classList.add("hidden");
    loadMessages(threadId);
}

async function voteMessage(messageId, typeVote, threadId) {
    if (!localStorage.getItem("token")) return alert("Connecte-toi pour voter !");
    try {
        await fetch(`${API_URL}/reactions`, {
            method: "POST",
            headers: { "Content-Type": "application/json", "Authorization": `Bearer ${localStorage.getItem("token")}` },
            body: JSON.stringify({ type: typeVote, message_id: messageId })
        });
        
        document.getElementById(`msg-box-${threadId}`).classList.add("hidden");
        loadMessages(threadId);
    } catch (err) { alert("Impossible d'enregistrer le vote"); }
}