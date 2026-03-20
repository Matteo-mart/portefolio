document.addEventListener('click', function(e) {
    const menu = document.getElementById('menu');
    const burger = document.querySelector('.burger-icon');
    if (menu && burger && !menu.contains(e.target) && !burger.contains(e.target)) {
        menu.classList.remove('active');
    }
});

function showToast(msg, type = 'success') {
    const t = document.getElementById('toast');
    t.textContent = msg;
    t.className = type;
    t.style.display = 'block';
    setTimeout(() => { t.style.display = 'none'; }, 4000);
}