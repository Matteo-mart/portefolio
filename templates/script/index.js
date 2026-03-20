function showToast(msg, type = 'success') {
    const t = document.getElementById('toast');
    if (!t) return;
    t.textContent   = msg;
    t.className     = type;
    t.style.display = 'block';
    setTimeout(() => { t.style.display = 'none'; }, 3500);
}

function formatDate(iso) {
    if (!iso) return '—';
    return new Date(iso).toLocaleDateString('fr-FR', {
        day: '2-digit', month: '2-digit', year: 'numeric',
        hour: '2-digit', minute: '2-digit'
    });
}

function initBurger(menuId, btnId) {
    const menu   = document.getElementById(menuId);
    const burger = document.getElementById(btnId);
    if (!menu || !burger) return;

    burger.addEventListener('click', e => {
        e.stopPropagation();
        menu.classList.toggle('active');
    });
    document.addEventListener('click', e => {
        if (!menu.contains(e.target) && e.target !== burger) {
            menu.classList.remove('active');
        }
    });
}

async function apiFetch(url, options = {}) {
    try {
        const res = await fetch(url, options);
        const text = await res.text();
        let data;
        try { data = JSON.parse(text); } catch { data = text; }
        return { ok: res.ok, status: res.status, data };
    } catch (err) {
        return { ok: false, status: 0, data: err.message };
    }
}