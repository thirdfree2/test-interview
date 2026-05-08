import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class AuthService {

    setToken(token: string) {
        localStorage.setItem(
            'access_token',
            token
        );
    }

    getToken() {
        if (typeof window === 'undefined') {
            return null;
        }

        return localStorage.getItem(
            'access_token'
        );
    }

    clearToken() {
        localStorage.removeItem(
            'access_token'
        );
    }

    isAuthenticated(): boolean {
        return !!this.getToken();
    }
}