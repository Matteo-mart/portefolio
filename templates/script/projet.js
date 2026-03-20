    document.querySelectorAll('.media-item').forEach(div => {
        const fileName = div.getAttribute('data-file');
        const ext = fileName.split('.').pop().toLowerCase();
        const filePath = "/uploads/" + fileName;
        const images = ['jpg', 'jpeg', 'png', 'gif', 'webp'];
        if (images.includes(ext)) {
            div.innerHTML = `<a href="${filePath}" target="_blank"><img src="${filePath}" class="main-img" alt="Aperçu"></a>`;
        } else if (ext === 'pdf') {
            div.innerHTML = `<a href="${filePath}" target="_blank" class="doc-card"><span class="doc-name">Document PDF</span><span class="btn-back" style="font-size:0.7rem;padding:5px 10px;">Lire le PDF</span></a>`;
        } else {
            div.innerHTML = `<a href="${filePath}" download class="doc-card"></span><span class="doc-name">${fileName}</span><span class="btn-back" style="font-size:0.7rem;padding:5px 10px;">Télécharger</span></a>`;
        }
    });

    function formatEntries(boxId) {
        const box = document.getElementById(boxId);
        if (!box) return;
        const raw = box.innerText.trim();
        if (!raw) return;

        const lines = raw.split('\n');
        const hasLabels = lines.some(l => /^[A-Z]\d+\s*[—–-]|^\d+\.|^[•\-]/.test(l.trim()));

        if (hasLabels) {
            box.innerHTML = lines
                .filter(l => l.trim() !== '')
                .map(l => {
                    const formatted = l.replace(
                        /^([A-Z]\d+\s*[—–-]|\d+\.|[•\-])\s*/,
                        match => `<strong>${match.trim()}</strong> `
                    );
                    return `<span class="line-entry">${formatted}</span>`;
                })
                .join('');
        }
    }

    formatEntries('probleme-box');
    formatEntries('solution-box');

    async function deleteData() {
        const id = window.location.pathname.split('/').pop();
        if (!id) return;
        if (!confirm("Confirmer la suppression du projet n°" + id + " ?")) return;
        try {
            const res = await fetch(`/delete-project?id=${id}`, { method: 'DELETE' });
            if (res.ok) { alert("Supprimé avec succès !"); window.location.href = "/"; }
            else { alert("Erreur lors de la suppression. Status: " + res.status); }
        } catch (e) {
            alert("Erreur de connexion: " + e.message);
        }
    }

    async function MoveToCorbeille() {
        const id = window.location.pathname.split('/').pop();
        try {
            const res = await fetch(`/move-to-corbeille?id=${id}`, { method: 'POST' });
            if (res.ok) { window.location.href = "/"; }
            else { alert("Erreur: " + res.status); }
        } catch (e) {
            alert("Erreur de connexion: " + e.message);
        }
    }