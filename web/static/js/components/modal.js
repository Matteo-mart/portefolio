/**
 * Modal Component - Système de modales réutilisable
 */

class ModalManager {
    constructor() {
        this.currentModal = null;
        this.focusTrap = null;
    }

    /**
     * Affiche une modale de confirmation
     * @param {string} title - Titre de la modale
     * @param {string} message - Message de confirmation
     * @param {Function} onConfirm - Callback si confirmé
     * @param {Object} options - Options (confirmText, cancelText, type)
     */
    confirm(title, message, onConfirm, options = {}) {
        const {
            confirmText = 'Confirmer',
            cancelText = 'Annuler',
            type = 'default', // 'default', 'danger'
        } = options;

        const modal = this.create({
            title,
            content: `<p class="modal-message">${this.escapeHTML(message)}</p>`,
            buttons: [
                {
                    text: cancelText,
                    class: 'btn-secondary',
                    onClick: () => this.close(),
                },
                {
                    text: confirmText,
                    class: type === 'danger' ? 'btn-rouge' : 'btn-bleu',
                    onClick: () => {
                        this.close();
                        if (onConfirm) onConfirm();
                    },
                },
            ],
        });

        this.show(modal);
    }

    /**
     * Affiche une modale personnalisée
     * @param {Object} config - Configuration {title, content, buttons, size}
     */
    create(config) {
        const {
            title = '',
            content = '',
            buttons = [],
            size = 'medium', // 'small', 'medium', 'large'
        } = config;

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';

        const modal = document.createElement('div');
        modal.className = `modal modal-${size}`;
        modal.setAttribute('role', 'dialog');
        modal.setAttribute('aria-modal', 'true');
        modal.setAttribute('aria-labelledby', 'modal-title');

        // Header
        const header = document.createElement('div');
        header.className = 'modal-header';
        header.innerHTML = `
            <h3 id="modal-title" class="modal-title">${this.escapeHTML(title)}</h3>
            <button class="modal-close" aria-label="Fermer">
                <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
                    <path d="M18 6L6 18M6 6L18 18" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                </svg>
            </button>
        `;

        // Body
        const body = document.createElement('div');
        body.className = 'modal-body';
        if (typeof content === 'string') {
            body.innerHTML = content;
        } else {
            body.appendChild(content);
        }

        // Footer avec boutons
        const footer = document.createElement('div');
        footer.className = 'modal-footer';
        buttons.forEach(btn => {
            const button = document.createElement('button');
            button.className = `btn ${btn.class || ''}`;
            button.textContent = btn.text;
            button.onclick = btn.onClick;
            footer.appendChild(button);
        });

        modal.appendChild(header);
        modal.appendChild(body);
        if (buttons.length > 0) {
            modal.appendChild(footer);
        }

        overlay.appendChild(modal);

        // Event listeners
        const closeBtn = header.querySelector('.modal-close');
        closeBtn.addEventListener('click', () => this.close());

        overlay.addEventListener('click', (e) => {
            if (e.target === overlay) {
                this.close();
            }
        });

        // ESC key
        const escHandler = (e) => {
            if (e.key === 'Escape') {
                this.close();
            }
        };
        document.addEventListener('keydown', escHandler);
        overlay._escHandler = escHandler;

        return overlay;
    }

    show(modalElement) {
        if (this.currentModal) {
            this.close();
        }

        this.currentModal = modalElement;
        document.body.appendChild(modalElement);
        document.body.style.overflow = 'hidden';

        // Animation d'entrée
        requestAnimationFrame(() => {
            modalElement.classList.add('modal-show');
        });

        // Focus trap
        this.setupFocusTrap(modalElement);
    }

    close() {
        if (!this.currentModal) return;

        const modal = this.currentModal;
        modal.classList.remove('modal-show');

        // Cleanup
        if (modal._escHandler) {
            document.removeEventListener('keydown', modal._escHandler);
        }

        setTimeout(() => {
            if (modal.parentNode) {
                modal.parentNode.removeChild(modal);
            }
            document.body.style.overflow = '';
            this.currentModal = null;
        }, 300);
    }

    setupFocusTrap(modalElement) {
        const focusableElements = modalElement.querySelectorAll(
            'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
        );

        if (focusableElements.length === 0) return;

        const firstElement = focusableElements[0];
        const lastElement = focusableElements[focusableElements.length - 1];

        // Focus le premier élément
        firstElement.focus();

        // Trap focus
        modalElement.addEventListener('keydown', (e) => {
            if (e.key !== 'Tab') return;

            if (e.shiftKey) {
                if (document.activeElement === firstElement) {
                    lastElement.focus();
                    e.preventDefault();
                }
            } else {
                if (document.activeElement === lastElement) {
                    firstElement.focus();
                    e.preventDefault();
                }
            }
        });
    }

    escapeHTML(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    // Méthodes de raccourci
    alert(title, message) {
        return this.confirm(title, message, null, {
            confirmText: 'OK',
            cancelText: null,
        });
    }

    deleteConfirm(message, onConfirm) {
        return this.confirm(
            'Confirmer la suppression',
            message,
            onConfirm,
            {
                confirmText: 'Supprimer',
                cancelText: 'Annuler',
                type: 'danger',
            }
        );
    }
}

// Instance globale
const Modal = new ModalManager();

// Export
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { Modal, ModalManager };
}
