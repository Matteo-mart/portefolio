        const menu   = document.getElementById('menu');
        const burger = document.getElementById('burger-btn');
        burger.addEventListener('click', e => { e.stopPropagation(); menu.classList.toggle('active'); });
        document.addEventListener('click', e => { if (!menu.contains(e.target) && e.target !== burger) menu.classList.remove('active'); });

        function showToast(msg, type = 'success') {
            const t = document.getElementById('toast');
            t.textContent = msg;
            t.className = type;
            t.style.display = 'block';
            t.style.animation = 'slideUp 0.2s ease';
            setTimeout(() => { t.style.display = 'none'; }, 3000);
        }

        function formatDate(iso) {
            if (!iso) return '—';
            return new Date(iso).toLocaleDateString('fr-FR', {
                day: '2-digit', month: '2-digit', year: 'numeric',
                hour: '2-digit', minute: '2-digit'
            });
        }

        function renderFiles(images) {
            if (!images || images.length === 0) return '';
            const IMAGE_EXT = ['jpg','jpeg','png','gif','webp'];
            const items = images.map(filename => {
                const ext = filename.split('.').pop().toLowerCase();
                if (IMAGE_EXT.includes(ext)) {
                    return `<a href="/uploads/${filename}" target="_blank">
                                <img src="/uploads/${filename}" alt="${filename}" title="${filename}">
                            </a>`;
                } else if (ext === 'pdf') {
                    return `<a href="/uploads/${filename}" target="_blank" class="doc-badge">${filename}</a>`;
                } else {
                    return `<a href="/uploads/${filename}" download class="doc-badge">${filename}</a>`;
                }
            }).join('');
            return `<div class="card-files">${items}</div>`;
        }

        function checkEmpty(listId, emptyId) {
            const list = document.getElementById(listId);
            if (list.children.length === 0) {
                document.getElementById(emptyId).style.display = 'block';
            }
        }


        async function loadCorbeille() {
            try {
                const res = await fetch('/corbeille-list');
                if (!res.ok) throw new Error('Erreur serveur');
                const data = await res.json();

                const list  = document.getElementById('project-list');
                const empty = document.getElementById('empty-projects');
                list.innerHTML = '';

                if (!data || data.length === 0) {
                    empty.style.display = 'block';
                    return;
                }

                empty.style.display = 'none';
                data.forEach(p => {
                    const card = document.createElement('div');
                    card.className = 'item-card';
                    card.id = `card-${p.id}`;
                    card.innerHTML = `
                        <div class="item-info">
                            <div class="titre">${p.titre}</div>
                            <div class="meta">
                                <span>ID projet : ${p.project_id}</span>
                                <span>Supprimé le : ${formatDate(p.date_suppression)}</span>
                            </div>
                            ${renderFiles(p.images)}
                        </div>
                        <div class="card-actions">
                            <button class="btn-restore" onclick="restorer(${p.id})">↩ Restaurer</button>
                            <button class="btn-del"     onclick="supprimerDefinitif(${p.id})">Supprimer</button>
                        </div>
                    `;
                    list.appendChild(card);
                });
            } catch(e) {
                showToast('Impossible de charger les projets : ' + e.message, 'error');
            }
        }

        async function restorer(id) {
            try {
                const res = await fetch(`/corbeille-restore?id=${id}`, { method: 'POST' });
                if (res.ok) {
                    document.getElementById(`card-${id}`)?.remove();
                    showToast('Projet restauré.');
                    checkEmpty('project-list', 'empty-projects');
                } else {
                    showToast('Erreur : ' + res.status, 'error');
                }
            } catch(e) {
                showToast('Erreur de connexion : ' + e.message, 'error');
            }
        }

        async function supprimerDefinitif(id) {
            try {
                const res = await fetch(`/corbeille-delete?id=${id}`, { method: 'DELETE' });
                if (res.ok) {
                    document.getElementById(`card-${id}`)?.remove();
                    showToast('Projet supprimé définitivement.');
                    checkEmpty('project-list', 'empty-projects');
                } else {
                    showToast('Erreur : ' + res.status, 'error');
                }
            } catch(e) {
                showToast('Erreur de connexion : ' + e.message, 'error');
            }
        }


        async function loadCorbeilleTech() {
            try {
                const res = await fetch('/corbeille-tech');
                if (!res.ok) throw new Error('Erreur serveur');
                const data = await res.json();

                const list  = document.getElementById('tech-list');
                const empty = document.getElementById('empty-tech');
                list.innerHTML = '';

                if (!data || data.length === 0) {
                    empty.style.display = 'block';
                    return;
                }

                empty.style.display = 'none';
                data.forEach(t => {
                    const card = document.createElement('div');
                    card.className = 'item-card';
                    card.id = `tech-card-${t.id}`;
                    card.innerHTML = `
                        <div class="item-info">
                            <div class="titre">${t.nom}</div>
                            <div class="meta">
                                <span>ID : ${t.tech_id}</span>
                                <span>Supprimé le : ${formatDate(t.date_suppression)}</span>
                            </div>
                            ${t.icone ? `<img class="tech-icon" src="${t.icone}" alt="${t.nom}">` : ''}
                        </div>
                        <div class="card-actions">
                            <button class="btn-restore" onclick="restorerTech(${t.id})">↩ Restaurer</button>
                            <button class="btn-del"     onclick="supprimerDefinitiveTech(${t.id})">Supprimer</button>
                        </div>
                    `;
                    list.appendChild(card);
                });
            } catch(e) {
                showToast('Impossible de charger les technologies : ' + e.message, 'error');
            }
        }

        async function restorerTech(id) {
            try {
                const res = await fetch(`/restore-tech?id=${id}`, { method: 'POST' });
                if (res.ok) {
                    document.getElementById(`tech-card-${id}`)?.remove();
                    showToast('Technologie restaurée.');
                    checkEmpty('tech-list', 'empty-tech');
                } else {
                    showToast('Erreur : ' + res.status, 'error');
                }
            } catch(e) {
                showToast('Erreur de connexion : ' + e.message, 'error');
            }
        }

        async function supprimerDefinitiveTech(id) {
            try {
                const res = await fetch(`/delete-tech-definitif?id=${id}`, { method: 'DELETE' });
                if (res.ok) {
                    document.getElementById(`tech-card-${id}`)?.remove();
                    showToast('Technologie supprimée définitivement.');
                    checkEmpty('tech-list', 'empty-tech');
                } else {
                    showToast('Erreur : ' + res.status, 'error');
                }
            } catch(e) {
                showToast('Erreur de connexion : ' + e.message, 'error');
            }
        }

        async function viderCorbeille() {
            try {
                const res = await fetch('/corbeille-vider', { method: 'DELETE' });
                if (res.ok) {
                    document.getElementById('project-list').innerHTML = '';
                    document.getElementById('tech-list').innerHTML = '';
                    document.getElementById('empty-projects').style.display = 'block';
                    document.getElementById('empty-tech').style.display = 'block';
                    showToast('Corbeille vidée.');
                } else {
                    showToast('Erreur : ' + res.status, 'error');
                }
            } catch(e) {
                showToast('Erreur de connexion : ' + e.message, 'error');
            }
        }

        loadCorbeille();
        loadCorbeilleTech();
