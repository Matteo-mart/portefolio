    let selectMode = false;
    let selectedIds = new Set();

    function showSection(id) {
        const sect = document.getElementById(id);
        const isVisible = sect.style.display === 'block';
        document.querySelectorAll('.form-section').forEach(s => s.style.display = 'none');
        sect.style.display = isVisible ? 'none' : 'block';
        document.getElementById('menu').classList.remove('active');
    }

    function toggleSelectMode() {
        selectMode = !selectMode;
        const btn = document.getElementById('toggleSelectBtn');
        btn.classList.toggle('active');
        if(!selectMode) cancelSelection();
    }

    function handleCardClick(card, event) {
        if (!selectMode) return;
        const id = card.dataset.id;
        if (selectedIds.has(id)) {
            selectedIds.delete(id);
            card.classList.remove('selected');
        } else {
            selectedIds.add(id);
            card.classList.add('selected');
        }
        updateBar();
    }

    function updateBar() {
        const bar = document.getElementById('selection-bar');
        document.getElementById('sel-count-num').textContent = selectedIds.size;
        bar.classList.toggle('visible', selectedIds.size > 0);
    }

    function cancelSelection() {
        selectedIds.clear();
        document.querySelectorAll('.project-card').forEach(c => c.classList.remove('selected'));
        updateBar();
    }

    function filterProjects(query) {
        const q = query.trim().toLowerCase();
        const cards = document.querySelectorAll('.project-card[data-id]');
        const clearBtn = document.getElementById('project-search-clear');
        const noResults = document.getElementById('project-no-results');
        let visible = 0;

        clearBtn.classList.toggle('visible', q.length > 0);

        cards.forEach(card => {
            const titre = card.dataset.titre || '';
            const tech  = card.dataset.tech  || '';
            const match = titre.toLowerCase().includes(q) || tech.toLowerCase().includes(q);

            card.classList.toggle('hidden', !match);
            if (match) {
                visible++;
                // Highlight
                card.querySelector('.card-titre').innerHTML = highlight(titre, q);
                card.querySelector('.card-tech').innerHTML  = highlight(tech,  q);
            } else {
                card.querySelector('.card-titre').innerHTML = titre;
                card.querySelector('.card-tech').innerHTML  = tech;
            }
        });

        noResults.classList.toggle('visible', visible === 0 && q.length > 0);
        document.getElementById('project-count').textContent = `${visible} projet(s)`;
    }

    function filterTech(query) {
        const q = query.trim().toLowerCase();
        const cards = document.querySelectorAll('.techno-card[data-id]');
        const clearBtn = document.getElementById('tech-search-clear');
        const noResults = document.getElementById('tech-no-results');
        let visible = 0;

        clearBtn.classList.toggle('visible', q.length > 0);

        cards.forEach(card => {
            const nom = card.dataset.nom || '';
            const match = nom.toLowerCase().includes(q);

            card.classList.toggle('hidden', !match);
            if (match) {
                visible++;
                const link = card.querySelector('.tech-nom');
                if (link) link.innerHTML = highlight(nom, q);
            } else {
                const link = card.querySelector('.tech-nom');
                if (link) link.innerHTML = card.dataset.nom;
            }
        });

        noResults.classList.toggle('visible', visible === 0 && q.length > 0);
        document.getElementById('tech-count').textContent = `${visible} outil(s)`;
    }

    function highlight(text, query) {
        if (!query) return text;
        const escaped = query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
        return text.replace(new RegExp(`(${escaped})`, 'gi'), '<mark>$1</mark>');
    }

    function clearSearch(type) {
        if (type === 'project') {
            const input = document.getElementById('project-search');
            input.value = '';
            filterProjects('');
            input.focus();
        } else {
            const input = document.getElementById('tech-search');
            input.value = '';
            filterTech('');
            input.focus();
        }
    }

    async function sendData() {
        const formData = new FormData();
        formData.append("title", document.getElementById('title').value);
        formData.append("date", document.getElementById('date_creation').value);
        formData.append("description", document.getElementById('desc').value);
        formData.append("technologie", document.getElementById('tech').value);
        formData.append("explication", document.getElementById('expl').value);
        
        const files = document.getElementById('image').files;
        for (let i = 0; i < files.length; i++) formData.append("image", files[i]);

        const res = await fetch('/add-post', { method: 'POST', body: formData });
        if (res.ok) location.reload();
    }

    async function deleteData() {
        const id = document.getElementById('delete-id').value;
        const res = await fetch(`/delete-project?id=${id}`, { method: 'DELETE' });
        if (res.ok) location.reload();
    }

    async function moveSelectedToCorbeille() {
        for (let id of selectedIds) {
            await fetch(`/move-to-corbeille?id=${id}`, { method: 'POST' });
        }
        location.reload();
    }

    let selectTechMode = false;
    let selectedTechIds = new Set();

    function toggleSelectTechMode() {
        selectTechMode = !selectTechMode;
        const btn = document.getElementById('toggleSelectTechBtn');
        btn.classList.toggle('active');
        if (!selectTechMode) cancelTechSelection();
    }

    function handleTechClick(card) {
        if (!selectTechMode) return;
        const id = card.dataset.id;
        if (selectedTechIds.has(id)) {
            selectedTechIds.delete(id);
            card.classList.remove('selected');
        } else {
            selectedTechIds.add(id);
            card.classList.add('selected');
        }
        updateTechBar();
    }

    function updateTechBar() {
        const bar = document.getElementById('selection-bar-tech');
        document.getElementById('sel-tech-count').textContent = selectedTechIds.size;
        bar.classList.toggle('visible', selectedTechIds.size > 0);
    }

    function cancelTechSelection() {
        selectedTechIds.clear();
        document.querySelectorAll('.techno-card').forEach(c => c.classList.remove('selected'));
        updateTechBar();
    }

    async function moveSelectedTechToCorbeille() {
        for (let id of selectedTechIds) {
            await fetch(`/delete-technologie?id=${id}`, { method: 'DELETE' });
        }
        location.reload();
    }

    function sortProjects(criteria) {
        const list = document.getElementById('projectList');
        const cards = Array.from(list.querySelectorAll('.project-card[data-id]'));

        cards.sort((a, b) => {
            switch (criteria) {
                case 'nom-asc':
                    return a.dataset.titre.localeCompare(b.dataset.titre);
                case 'nom-desc':
                    return b.dataset.titre.localeCompare(a.dataset.titre);
                case 'tech-asc':
                    return a.dataset.tech.localeCompare(b.dataset.tech);
                case 'tech-desc':
                    return b.dataset.tech.localeCompare(a.dataset.tech);
                default:
                    return 0;
            }
        });

        cards.forEach(card => list.insertBefore(card, list.querySelector('.no-results')));
    }   