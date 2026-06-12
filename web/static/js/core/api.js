/**
 * API Client - Gestion centralisée des appels HTTP
 */

class APIClient {
    constructor(baseURL = '') {
        this.baseURL = baseURL;
        this.defaultOptions = {
            headers: {
                'Content-Type': 'application/json',
            },
        };
    }

    /**
     * Effectue une requête HTTP avec retry logic
     */
    async request(url, options = {}, retries = 3) {
        const fullURL = `${this.baseURL}${url}`;
        const config = {
            ...this.defaultOptions,
            ...options,
            headers: {
                ...this.defaultOptions.headers,
                ...options.headers,
            },
        };

        for (let attempt = 0; attempt <= retries; attempt++) {
            try {
                const response = await fetch(fullURL, config);

                // Parser la réponse
                let data;
                const contentType = response.headers.get('content-type');
                if (contentType && contentType.includes('application/json')) {
                    data = await response.json();
                } else {
                    data = await response.text();
                }

                if (!response.ok) {
                    throw new APIError(
                        data.error || `HTTP ${response.status}`,
                        response.status,
                        data
                    );
                }

                return { ok: true, status: response.status, data };

            } catch (error) {
                // Si c'est la dernière tentative, throw l'erreur
                if (attempt === retries) {
                    if (error instanceof APIError) {
                        throw error;
                    }
                    throw new APIError('Erreur réseau', 0, error.message);
                }

                // Attendre avant de réessayer (exponential backoff)
                await this.sleep(Math.pow(2, attempt) * 1000);
            }
        }
    }

    /**
     * GET request
     */
    async get(url, params = {}) {
        const queryString = new URLSearchParams(params).toString();
        const fullURL = queryString ? `${url}?${queryString}` : url;
        return this.request(fullURL, { method: 'GET' });
    }

    /**
     * POST request
     */
    async post(url, data, options = {}) {
        const config = {
            method: 'POST',
            ...options,
        };

        // Si c'est du FormData, ne pas définir Content-Type (auto)
        if (data instanceof FormData) {
            delete config.headers;
            config.body = data;
        } else {
            config.body = JSON.stringify(data);
        }

        return this.request(url, config);
    }

    /**
     * PUT request
     */
    async put(url, data) {
        return this.request(url, {
            method: 'PUT',
            body: JSON.stringify(data),
        });
    }

    /**
     * DELETE request
     */
    async delete(url, params = {}) {
        const queryString = new URLSearchParams(params).toString();
        const fullURL = queryString ? `${url}?${queryString}` : url;
        return this.request(fullURL, { method: 'DELETE' });
    }

    /**
     * Utilitaire sleep pour retry logic
     */
    sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
}

/**
 * Erreur API personnalisée
 */
class APIError extends Error {
    constructor(message, status, data) {
        super(message);
        this.name = 'APIError';
        this.status = status;
        this.data = data;
    }
}

// Instance globale du client API
const API = new APIClient();

// Export pour utilisation dans d'autres modules
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { API, APIClient, APIError };
}
