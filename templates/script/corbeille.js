const IMAGE_EXT = ['jpg', 'jpeg', 'png', 'gif', 'webp'];

function renderFiles(images) {
    if (!images?.length) return '';
    const items = images.map(f => {
        const ext = f.split('.').pop().toLowerCase();
        if (IMAGE_EXT.includes(ext))
            return `<a href="/uploads/${f}" target="_blank"><img src="/uploads/${f}" alt="${f}" title="${f}"></a>`;
        return `<a href="/uploads/${f}" ${ext === 'pdf' ? 'target="_blank"' : 'download'} class="doc-badge">${f}</a>`;
    }).join('');
    return `<div class="card-files">${items}</div>`;
}

function checkEmpty(listId, emptyId) {
    const isEmpty = document.getElementById(listId).children.length === 0;
    document.getElementById(emptyId).style.display = isEmpty ? 'block' : 'none';
}

function createCard(id, html) {
    const card     = document.createElement('div');
    card.className = 'item-card';
    card.id        = id;
    card.innerHTML = html;
    return card;
}

async function loadCorbeille() {
    const { ok, data } = await apiFetch('/corbeille-list');
    if (!ok) { showToast('Impossible de charger les projets : ' + data, 'error'); return; }

    const list  = document.getElementById('project-list');
    const empty = document.getElementById('empty-projects');
    list.innerHTML = '';

    if (!data?.length) { empty.style.display = 'block'; return; }
    empty.style.display = 'none';

    data.forEach(p => list.appendChild(createCard(`card-${p.id}`, `
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
    `)));
}

async function restorer(id) {
    const { ok, status } = await apiFetch(`/corbeille-restore?id=${id}`, { method: 'POST' });
    if (ok) { document.getElementById(`card-${id}`)?.remove(); showToast('Projet restauré.'); checkEmpty('project-list', 'empty-projects'); }
    else showToast('Erreur : ' + status, 'error');
}

async function supprimerDefinitif(id) {
    const { ok, status } = await apiFetch(`/corbeille-delete?id=${id}`, { method: 'DELETE' });
    if (ok) { document.getElementById(`card-${id}`)?.remove(); showToast('Projet supprimé définitivement.'); checkEmpty('project-list', 'empty-projects'); }
    else showToast('Erreur : ' + status, 'error');
}

async function loadCorbeilleTech() {
    const { ok, data } = await apiFetch('/corbeille-tech');
    if (!ok) { showToast('Impossible de charger les technologies : ' + data, 'error'); return; }

    const list  = document.getElementById('tech-list');
    const empty = document.getElementById('empty-tech');
    list.innerHTML = '';

    if (!data?.length) { empty.style.display = 'block'; return; }
    empty.style.display = 'none';

    data.forEach(t => list.appendChild(createCard(`tech-card-${t.id}`, `
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
    `)));
}

async function restorerTech(id) {
    const { ok, status } = await apiFetch(`/restore-tech?id=${id}`, { method: 'POST' });
    if (ok) { document.getElementById(`tech-card-${id}`)?.remove(); showToast('Technologie restaurée.'); checkEmpty('tech-list', 'empty-tech'); }
    else showToast('Erreur : ' + status, 'error');
}

async function supprimerDefinitiveTech(id) {
    const { ok, status } = await apiFetch(`/delete-tech-definitif?id=${id}`, { method: 'DELETE' });
    if (ok) { document.getElementById(`tech-card-${id}`)?.remove(); showToast('Technologie supprimée définitivement.'); checkEmpty('tech-list', 'empty-tech'); }
    else showToast('Erreur : ' + status, 'error');
}

async function viderCorbeille() {
    const { ok, status } = await apiFetch('/corbeille-vider', { method: 'DELETE' });
    if (ok) {
        ['project-list', 'tech-list'].forEach(id => document.getElementById(id).innerHTML = '');
        ['empty-projects', 'empty-tech'].forEach(id => document.getElementById(id).style.display = 'block');
        showToast('Corbeille vidée.');
    } else {
        showToast('Erreur : ' + status, 'error');
    }
}

loadCorbeille();
loadCorbeilleTech();