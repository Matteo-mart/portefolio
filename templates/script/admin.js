function afficherSection(nomSection, ongletClique) {
    document.querySelectorAll('.section-form').forEach(s => s.classList.remove('active'));
    document.querySelectorAll('.onglet-btn').forEach(o => o.classList.remove('actif'));
    document.getElementById('section-' + nomSection).classList.add('active');
    ongletClique.classList.add('actif');
}

function afficherApercu() {
    const zone    = document.getElementById('zone-apercu');
    zone.innerHTML = '';

    Array.from(document.getElementById('aj-fichiers').files).forEach(fichier => {
        const div     = document.createElement('div');
        div.className = 'apercu-item';

        if (fichier.type === 'application/pdf') {
            div.innerHTML = '<span>PDF</span>';
        } else {
            const reader   = new FileReader();
            reader.onload  = e => { div.innerHTML = `<img src="${e.target.result}" alt="aperçu">`; };
            reader.readAsDataURL(fichier);
        }
        zone.appendChild(div);
    });
}

function afficherMessage(idDiv, texte, succes) {
    const div       = document.getElementById(idDiv);
    div.textContent = texte;
    div.className   = 'message ' + (succes ? 'succes' : 'erreur');
    div.style.display = 'block';
}

async function ajouterProjet() {
    const formData = new FormData();
    const champs = ['titre','date','description','tech','explication','probleme','solution','url'];
    const cles   = ['title','date','description','technologie','explication','probleme','solution','url_source'];
    champs.forEach((c, i) => formData.append(cles[i], document.getElementById('aj-' + c).value));

    Array.from(document.getElementById('aj-fichiers').files)
        .forEach(f => formData.append('image', f));

    const { ok, data } = await apiFetch('/add-post', { method: 'POST', body: formData });
    if (ok) window.location.href = '/';
    else afficherMessage('msg-ajouter', 'Erreur : ' + data, false);
}

async function modifierProjet() {
    const id = document.getElementById('mod-id').value;
    if (!id) { afficherMessage('msg-modifier', 'Saisir un ID', false); return; }

    const body = {
        titre:       document.getElementById('mod-titre').value,
        description: document.getElementById('mod-description').value,
        tech:        document.getElementById('mod-tech').value,
        explication: document.getElementById('mod-explication').value,
        probleme:    document.getElementById('mod-probleme').value,
        solution:    document.getElementById('mod-solution').value,
        url_source:  document.getElementById('mod-url').value,
    };

    const { ok, data } = await apiFetch('/update-projet/' + id, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
    });
    if (ok) window.location.href = '/';
    else afficherMessage('msg-modifier', 'Erreur : ' + data, false);
}

async function supprimerProjet() {
    const id = document.getElementById('sup-id').value;
    if (!id) { afficherMessage('msg-supprimer', 'Saisir un ID', false); return; }
    if (!confirm(`Voulez-vous vraiment supprimer le projet n°${id} ?`)) return;

    const { ok } = await apiFetch(`/delete-project?id=${id}`, { method: 'DELETE' });
    if (ok) document.getElementById('sup-id').value = '';
    else afficherMessage('msg-supprimer', 'Erreur lors de la suppression', false);
}

async function modifierContact() {
    const id = document.getElementById('ct-id').value;
    if (!id) { afficherMessage('msg-contact', 'Veuillez saisir un ID', false); return; }

    const body = {
        id,
        telephone: document.getElementById('ct-telephone').value,
        email:     document.getElementById('ct-email').value,
        linkedin:  document.getElementById('ct-linkedin').value,
        github:    document.getElementById('ct-github').value,
    };

    const { ok, data } = await apiFetch('/update-contact', {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
    });
    if (ok) window.location.href = '/';
    else afficherMessage('msg-contact', 'Erreur : ' + data, false);
}

async function modifierTechnologie() {
    const id = document.getElementById('tech-id').value;
    if (!id) { afficherMessage('msg-technologies', 'Veuillez saisir un ID', false); return; }

    const body = {
        id:         parseInt(id, 10),
        nom:        document.getElementById('ct-nom').value,
        icone:      document.getElementById('ct-icone').value,
        url_source: document.getElementById('ct-url_source').value,
    };

    const { ok, data } = await apiFetch('/update-technologies/' + id, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
    });
    if (ok) window.location.href = '/';
    else afficherMessage('msg-technologies', 'Erreur : ' + data, false);
}

async function ajouterTechnologie() {
    const nom = document.getElementById('add-tech-nom').value.trim();
    if (!nom) { afficherMessage('msg-add-tech', 'Le nom est requis', false); return; }

    const body = {
        nom,
        icone:      document.getElementById('add-tech-icone').value.trim(),
        url_source: document.getElementById('add-tech-url').value.trim(),
    };

    const { ok, data } = await apiFetch('/add-technologie', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(body),
    });

    if (ok) {
        afficherMessage('msg-add-tech', 'Technologie ajoutée !', true);
        ['add-tech-nom', 'add-tech-icone', 'add-tech-url'].forEach(id => {
            document.getElementById(id).value = '';
        });
    } else {
        afficherMessage('msg-add-tech', 'Erreur : ' + data, false);
    }
}