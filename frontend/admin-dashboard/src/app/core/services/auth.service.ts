import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class AuthService {
    private apiUrl = 'http://localhost:8080/auth';
    private tokenKey = 'auth_token';

    constructor(private http: HttpClient) { }

    register(email: string, password: string): Observable<any> {
        return this.http.post(`${this.apiUrl}/register`, { email, password });
    }

    login(email: string, password: string): Observable<any> {
        return this.http.post<{ token: string, user: any }>(`${this.apiUrl}/login`, { email, password }).pipe(
            tap(response => {
                this.setToken(response.token);
                this.setUser(response.user);
            })
        );
    }

    logout(): void {
        localStorage.removeItem(this.tokenKey);
        localStorage.removeItem('user_data');
    }

    getToken(): string | null {
        return localStorage.getItem(this.tokenKey);
    }

    getUser(): any {
        const user = localStorage.getItem('user_data');
        return user ? JSON.parse(user) : null;
    }

    private setToken(token: string): void {
        localStorage.setItem(this.tokenKey, token);
    }

    private setUser(user: any): void {
        localStorage.setItem('user_data', JSON.stringify(user));
    }

    isAuthenticated(): boolean {
        return !!this.getToken();
    }
}
