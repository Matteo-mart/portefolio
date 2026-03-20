const burgerBtn = document.getElementById('burger-btn');
const menuNav   = document.getElementById('menu-nav');

burgerBtn.addEventListener('click', function (event) {
    event.stopPropagation();
    menuNav.classList.toggle('actif');
});

document.addEventListener('click', function () {
    menuNav.classList.remove('actif');
});

function afficherSection(nomSection, ongletClique) {
    document.querySelectorAll('.section-form').forEach(function (section) {
        section.classList.remove('active');
    });

    document.querySelectorAll('.onglet-btn').forEach(function (onglet) {
        onglet.classList.remove('actif');
    });

    document.getElementById('section-' + nomSection).classList.add('active');
    ongletClique.classList.add('actif');
}

function afficherApercu() {
    const zoneApercu = document.getElementById('zone-apercu');
    const fichiers   = document.getElementById('aj-fichiers').files;

    zoneApercu.innerHTML = '';

    Array.from(fichiers).forEach(function (fichier) {
        const div     = document.createElement('div');
        div.className = 'apercu-item';

        if (fichier.type === 'application/pdf') {
            div.innerHTML = '<span>PDF</span>';
        } else {
            const lecteur    = new FileReader();
            lecteur.onload   = function (e) {
                div.innerHTML = '<img src="' + e.target.result + '" alt="aperçu">';
            };
            lecteur.readAsDataURL(fichier);
        }

        zoneApercu.appendChild(div);
    });
}

function afficherMessage(idDiv, texte, succes) {
    const div         = document.getElementById(idDiv);
    div.textContent   = texte;
    div.className     = 'message ' + (succes ? 'succes' : 'erreur');
    div.style.display = 'block';
}

async function ajouterProjet() {
    const donnees = new FormData();
    donnees.append('title',       document.getElementById('aj-titre').value);
    donnees.append('date',        document.getElementById('aj-date').value);
    donnees.append('description', document.getElementById('aj-description').value);
    donnees.append('technologie', document.getElementById('aj-tech').value);
    donnees.append('explication', document.getElementById('aj-explication').value);
    donnees.append('probleme',    document.getElementById('aj-probleme').value);
    donnees.append('solution',    document.getElementById('aj-solution').value);
    donnees.append('url_source',  document.getElementById('aj-url').value);

    const inputFichiers = document.getElementById('aj-fichiers');
    for (let i = 0; i < inputFichiers.files.length; i++) {
        donnees.append('image', inputFichiers.files[i]);
    }

    try {
        const reponse = await fetch('/add-post', {
            method: 'POST',
            body:   donnees,
        });

        if (reponse.ok) {
            window.location.href = '/';
        } else {
            const msgErreur = await reponse.text();
            afficherMessage('msg-ajouter', 'Erreur : ' + msgErreur, false);
        }
    } catch (erreur) {
        afficherMessage('msg-ajouter', 'Connexion impossible : ' + erreur.message, false);
    }
}

async function modifierProjet() {
    const idProjet = document.getElementById('mod-id').value;

    if (!idProjet) {
        afficherMessage('msg-modifier', 'Saisir un ID', false);
        return;
    }

    const donnees = {
        titre:       document.getElementById('mod-titre').value,
        description: document.getElementById('mod-description').value,
        tech:        document.getElementById('mod-tech').value,
        explication: document.getElementById('mod-explication').value,
        probleme:    document.getElementById('mod-probleme').value,
        solution:    document.getElementById('mod-solution').value,
        url_source:  document.getElementById('mod-url').value,
    };

    try {
        const reponse = await fetch('/update-projet/' + idProjet, {
            method:  'PUT',
            headers: { 'Content-Type': 'application/json' },
            body:    JSON.stringify(donnees),
        });

        if (reponse.ok) {
            window.location.href = '/';
        } else {
            const msgErreur = await reponse.text();
            afficherMessage('msg-modifier', 'Erreur : ' + msgErreur, false);
        }
    } catch (erreur) {
        afficherMessage('msg-modifier', 'Connexion impossible : ' + erreur.message, false);
    }
}

async function supprimerProjet() {
    const idProjet = document.getElementById('sup-id').value;

    if (!idProjet) {
        afficherMessage('msg-supprimer', 'Saisir un ID', false);
        return;
    }

    const confirme = confirm('Voulez-vous vraiment supprimer le projet n°' + idProjet + ' ?');
    if (!confirme) return;

    try {
        const reponse = await fetch('/delete-project?id=' + idProjet, {
            method: 'DELETE',
        });

        if (reponse.ok) {
            document.getElementById('sup-id').value = '';
        } else {
            afficherMessage('msg-supprimer', 'Erreur lors de la suppression', false);
        }
    } catch (erreur) {
        afficherMessage('msg-supprimer', 'Connexion impossible : ' + erreur.message, false);
    }
}

async function modifierContact() {
    const idContact = document.getElementById('ct-id').value;

    if (!idContact) {
        afficherMessage('msg-contact', 'Veuillez saisir un ID', false);
        return;
    }

    const donnees = {
        id:        idContact,
        telephone: document.getElementById('ct-telephone').value,
        email:     document.getElementById('ct-email').value,
        linkedin:  document.getElementById('ct-linkedin').value,
        github:    document.getElementById('ct-github').value,
    };

    try {
        const reponse = await fetch('/update-contact', {
            method:  'PUT',
            headers: { 'Content-Type': 'application/json' },
            body:    JSON.stringify(donnees),
        });

        if (reponse.ok) {
            window.location.href = '/';
        } else {
            const msgErreur = await reponse.text();
            afficherMessage('msg-contact', 'Erreur : ' + msgErreur, false);
        }
    } catch (erreur) {
        afficherMessage('msg-contact', 'Connexion impossible : ' + erreur.message, false);
    }
}

async function modifierTechnologie() {
    const idTechnologies = document.getElementById('tech-id').value;

    if (!idTechnologies) {
        afficherMessage('msg-technologies', 'Veuillez saisir un ID', false);
        return;
    }

    const donnees = {
        id:         parseInt(idTechnologies, 10),
        nom:        document.getElementById('ct-nom').value,
        icone:      document.getElementById('ct-icone').value,
        url_source: document.getElementById('ct-url_source').value,
    };

    try {
        const reponse = await fetch('/update-technologies/' + idTechnologies, {
            method:  'PUT',
            headers: { 'Content-Type': 'application/json' },
            body:    JSON.stringify(donnees),
        });

        if (reponse.ok) {
            window.location.href = '/';
        } else {
            const msgErreur = await reponse.text();
            afficherMessage('msg-technologies', 'Erreur : ' + msgErreur, false);
        }
    } catch (erreur) {
        afficherMessage('msg-technologies', 'Connexion impossible : ' + erreur.message, false);
    }
}

async function ajouterTechnologie() {
    const nom = document.getElementById('add-tech-nom').value.trim();

    if (!nom) {
        afficherMessage('msg-add-tech', 'Le nom est requis', false);
        return;
    }

    const donnees = {
        nom:        nom,
        icone:      document.getElementById('add-tech-icone').value.trim(),
        url_source: document.getElementById('add-tech-url').value.trim(),
    };

    try {
        const reponse = await fetch('/add-technologie', {
            method:  'POST',
            headers: { 'Content-Type': 'application/json' },
            body:    JSON.stringify(donnees),
        });

        if (reponse.ok) {
            afficherMessage('msg-add-tech', 'Technologie ajoutée !', true);
            document.getElementById('add-tech-nom').value   = '';
            document.getElementById('add-tech-icone').value = '';
            document.getElementById('add-tech-url').value   = '';
        } else {
            const err = await reponse.text();
            afficherMessage('msg-add-tech', 'Erreur : ' + err, false);
        }
    } catch (e) {
        afficherMessage('msg-add-tech', 'Connexion impossible : ' + e.message, false);
    }
}