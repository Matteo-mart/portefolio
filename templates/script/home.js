let selectMode  = false;
let selectedIds = new Set();

function toggleSelectMode() {
    selectMode = !selectMode;
    document.getElementById('toggleSelectBtn').classList.toggle('active');
    if (!selectMode) cancelSelection();
}

function handleCardClick(card) {
    if (!selectMode) return;
    toggleSelection(card, selectedIds, 'selection-bar', 'sel-count-num');
}

function cancelSelection() {
    clearSelection(selectedIds, '.project-card', 'selection-bar', 'sel-count-num');
}

async function moveSelectedToCorbeille() {
    for (const id of selectedIds) {
        await apiFetch(`/move-to-corbeille?id=${id}`, { method: 'POST' });
    }
    location.reload();
}

let selectTechMode  = false;
let selectedTechIds = new Set();

function toggleSelectTechMode() {
    selectTechMode = !selectTechMode;
    document.getElementById('toggleSelectTechBtn').classList.toggle('active');
    if (!selectTechMode) cancelTechSelection();
}

function handleTechClick(card) {
    if (!selectTechMode) return;
    toggleSelection(card, selectedTechIds, 'selection-bar-tech', 'sel-tech-count');
}

function cancelTechSelection() {
    clearSelection(selectedTechIds, '.techno-card', 'selection-bar-tech', 'sel-tech-count');
}

async function moveSelectedTechToCorbeille() {
    for (const id of selectedTechIds) {
        await apiFetch(`/delete-technologie?id=${id}`, { method: 'DELETE' });
    }
    location.reload();
}

function toggleSelection(card, set, barId, countId) {
    const id = card.dataset.id;
    if (set.has(id)) { set.delete(id); card.classList.remove('selected'); }
    else             { set.add(id);    card.classList.add('selected'); }
    updateBar(set, barId, countId);
}

function clearSelection(set, selector, barId, countId) {
    set.clear();
    document.querySelectorAll(selector).forEach(c => c.classList.remove('selected'));
    updateBar(set, barId, countId);
}

function updateBar(set, barId, countId) {
    document.getElementById(countId).textContent = set.size;
    document.getElementById(barId).classList.toggle('visible', set.size > 0);
}

function highlight(text, query) {
    if (!query) return text;
    const esc = query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
    return text.replace(new RegExp(`(${esc})`, 'gi'), '<mark>$1</mark>');
}

function filterCards({ cards, query, clearBtnId, noResultsId, countId, countLabel, getTitles }) {
    const q         = query.trim().toLowerCase();
    let   visible   = 0;

    document.getElementById(clearBtnId).classList.toggle('visible', q.length > 0);

    cards.forEach(card => {
        const { main, secondary } = getTitles(card);
        const match = main.toLowerCase().includes(q) || (secondary || '').toLowerCase().includes(q);

        card.classList.toggle('hidden', !match);
        if (match) {
            visible++;
            if (card.querySelector('.card-titre'))
                card.querySelector('.card-titre').innerHTML = highlight(main, q);
            if (secondary && card.querySelector('.card-tech'))
                card.querySelector('.card-tech').innerHTML  = highlight(secondary, q);
            if (card.querySelector('.tech-nom'))
                card.querySelector('.tech-nom').innerHTML   = highlight(main, q);
        } else {
            if (card.querySelector('.card-titre')) card.querySelector('.card-titre').innerHTML = main;
            if (card.querySelector('.card-tech'))  card.querySelector('.card-tech').innerHTML  = secondary || '';
            if (card.querySelector('.tech-nom'))   card.querySelector('.tech-nom').innerHTML   = main;
        }
    });

    document.getElementById(noResultsId).classList.toggle('visible', visible === 0 && q.length > 0);
    document.getElementById(countId).textContent = `${visible} ${countLabel}`;
}

function filtrerProjects(query) {
    filterCards({
        cards:       Array.from(document.querySelectorAll('.project-card[data-id]')),
        query,
        clearBtnId:  'project-search-clear',
        noResultsId: 'project-no-results',
        countId:     'project-count',
        countLabel:  'projet(s)',
        getTitles:   c => ({ main: c.dataset.titre || '', secondary: c.dataset.tech || '' }),
    });
}

function filterTech(query) {
    filterCards({
        cards:       Array.from(document.querySelectorAll('.techno-card[data-id]')),
        query,
        clearBtnId:  'tech-search-clear',
        noResultsId: 'tech-no-results',
        countId:     'tech-count',
        countLabel:  'outil(s)',
        getTitles:   c => ({ main: c.dataset.nom || '' }),
    });
}

function clearSearch(type) {
    const id    = type === 'project' ? 'project-search' : 'tech-search';
    const input = document.getElementById(id);
    input.value = '';
    type === 'project' ? filtrerProjects('') : filterTech('');
    input.focus();
}

function sortProjects(criteria) {
    const list  = document.getElementById('projectList');
    const cards = Array.from(list.querySelectorAll('.project-card[data-id]'));
    const noRes = list.querySelector('.no-results');

    const comparators = {
        'nom-asc':  (a, b) => a.dataset.titre.localeCompare(b.dataset.titre),
        'nom-desc': (a, b) => b.dataset.titre.localeCompare(a.dataset.titre),
        'tech-asc': (a, b) => a.dataset.tech.localeCompare(b.dataset.tech),
        'tech-desc':(a, b) => b.dataset.tech.localeCompare(a.dataset.tech),
    };

    if (comparators[criteria]) {
        cards.sort(comparators[criteria]).forEach(c => list.insertBefore(c, noRes));
    }
}


function toggleSearchPanel() {
    const panel = document.getElementById('searchPanel');
    const btn = document.getElementById('searchToggleBtn');
    const isOpen = panel.classList.toggle('open');
    btn.classList.toggle('active', isOpen);
    if (isOpen) {
        setTimeout(() => document.getElementById('project-search').focus(), 100);
    }
}

document.addEventListener('click', function(e) {
    const wrapper = document.querySelector('.search-toggle-wrapper');
    if (wrapper && !wrapper.contains(e.target)) {
        document.getElementById('searchPanel').classList.remove('open');
        document.getElementById('searchToggleBtn').classList.remove('active');
    }
});